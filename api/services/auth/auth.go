package auth

import (
	"context"
	"fmt"
	"forkd/db"
	"forkd/util"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type key int

const (
	USER_KEY key = iota
	SESSION_KEY
	TOKEN_KEY
)

type MagicLinkToken struct {
	Token pgtype.UUID
}

type MagicLinkCode struct {
	ID pgtype.UUID
}

type SessionToken struct {
	ID pgtype.UUID
}

type UserWithSessionToken struct {
	db.User
	Token string
}

type AuthService interface {
	CreateMagicLink(context.Context, pgtype.UUID) (string, error)
	CreateSession(context.Context, pgtype.UUID) (UserWithSessionToken, error)
	DeleteSession(context.Context, pgtype.UUID) error
	GetUserSessionAndSetOnContext(context.Context) context.Context
	GetUserSessionFromCtx(context.Context) (*db.User, *db.Session)
	SessionWrapper(http.HandlerFunc) http.HandlerFunc
	SetTokenOnCtx(context.Context, string) context.Context
	UpsertUser(context.Context, string) (db.User, error)
	ValidateMagicLink(context.Context, string, string) (pgtype.UUID, error)
	getTokenFromCtx(context.Context) *string
	getUserFromCtx(context.Context) *db.User
	getSessionFromCtx(context.Context) *db.Session
	setUserOnCtx(context.Context, db.User) context.Context
	setSessionOnCtx(context.Context, db.Session) context.Context
}

type authService struct {
	queries *db.Queries
}

func (a authService) DeleteSession(ctx context.Context, sessionId pgtype.UUID) error {
	return a.queries.DeleteSession(ctx, sessionId)
}

// GetUserSessionAndSetOnContext implements AuthService.
func (a authService) GetUserSessionAndSetOnContext(ctx context.Context) context.Context {
	sessionId := a.GetSessionIdFromCtxToken(ctx)
	if !sessionId.Valid {
		return ctx
	}
	result, err := a.queries.GetUserBySessionId(ctx, sessionId)
	if err != nil {
		return ctx
	}
	return a.setSessionOnCtx(a.setUserOnCtx(ctx, result.User), result.Session)
}

func (a authService) ValidateMagicLink(ctx context.Context, code string, token string) (pgtype.UUID, error) {
	var id pgtype.UUID
	var magicLinkCode MagicLinkCode
	var magicLinkToken MagicLinkToken
	err := util.DecodeBase64StringToStruct(code, &magicLinkCode)
	if err != nil {
		return id, err
	}
	err = util.DecodeBase64StringToStruct(code, &magicLinkToken)
	if err != nil {
		return id, err
	}
	magicLink, err := a.queries.GetMagicLink(ctx, db.GetMagicLinkParams{
		ID:    magicLinkCode.ID,
		Token: magicLinkToken.Token,
	})
	if err != nil {
		return id, err
	}
	if magicLink.Expiry.Time.Before(time.Now()) {
		return id, fmt.Errorf("otp/magic link expired")
	}
	return magicLink.UserID, nil
}

func (a authService) CreateMagicLink(ctx context.Context, userId pgtype.UUID) (string, error) {
	// Set the expiry for 10 minutes
	expiry := time.Now().Add(10 * time.Minute)
	token := pgtype.UUID{
		Bytes: uuid.New(),
		Valid: true,
	}
	magicLink, err := a.queries.CreateMagicLink(ctx, db.CreateMagicLinkParams{
		UserID: userId,
		Token:  token,
		Expiry: pgtype.Timestamp{
			Time:  expiry,
			Valid: true,
		},
	})
	if err != nil {
		return "", err
	}
	magicLinkCodeStruct := MagicLinkCode{
		ID: magicLink.ID,
	}
	magicLinkCode, err := util.EncodeStructToBase64String(magicLinkCodeStruct)
	if err != nil {
		return "", err
	}
	fmt.Printf("MAGIC LINK CODE: %s\n", magicLinkCode)
	magicLinkTokenStruct := MagicLinkToken{
		Token: magicLink.Token,
	}
	magicLinkToken, err := util.EncodeStructToBase64String(magicLinkTokenStruct)
	if err != nil {
		return "", err
	}
	return magicLinkToken, nil
}

func (a authService) CreateSession(ctx context.Context, userId pgtype.UUID) (UserWithSessionToken, error) {
	// Set expiry to 2 weeks, this will refresh everyime we access the session
	expiry := time.Now().AddDate(0, 0, 14)
	result, err := a.queries.CreateSession(ctx, db.CreateSessionParams{
		UserID: userId,
		Expiry: pgtype.Timestamp{
			Time:  expiry,
			Valid: true,
		},
	})
	if err != nil {
		return UserWithSessionToken{}, nil
	}
	session := SessionToken{
		ID: result.ID,
	}
	sessionToken, err := util.EncodeStructToBase64String(session)
	if err != nil {
		return UserWithSessionToken{}, nil
	}
	return newUserWithToken(result.User, sessionToken), nil
}

func (a authService) GetUserSessionFromCtx(ctx context.Context) (*db.User, *db.Session) {
	var userAddr *db.User
	var sessionAddr *db.Session
	user, ok := ctx.Value(USER_KEY).(db.User)
	if ok {
		userAddr = &user
	}
	session, ok := ctx.Value(SESSION_KEY).(db.Session)
	if ok {
		sessionAddr = &session
	}
	return userAddr, sessionAddr
}

func (a authService) UpsertUser(ctx context.Context, email string) (db.User, error) {
	displayName := strings.Split(email, "@")[0]
	upsert, err := a.queries.UpsertUser(ctx, db.UpsertUserParams{
		Email:       email,
		DisplayName: displayName,
	})
	if err != nil {
		return db.User{}, err
	}
	return db.User(upsert), nil
}

func (a authService) GetSessionIdFromCtxToken(ctx context.Context) pgtype.UUID {
	id := pgtype.UUID{}
	token := a.getTokenFromCtx(ctx)
	if token == nil {
		return id
	}
	var session SessionToken
	err := util.DecodeBase64StringToStruct(*token, &session)
	if err != nil {
		return id
	}
	return session.ID
}

func (a authService) SessionWrapper(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val, ok := r.Header[http.CanonicalHeaderKey("authorization")]
		if !ok {
			handler(w, r)
			return
		}
		req := r.WithContext(a.SetTokenOnCtx(r.Context(), val[0]))
		handler(w, req)
	}
}

func (a authService) setUserOnCtx(ctx context.Context, user db.User) context.Context {
	return context.WithValue(ctx, USER_KEY, user)
}

func (a authService) getUserFromCtx(ctx context.Context) *db.User {
	user, ok := ctx.Value(USER_KEY).(db.User)
	if !ok {
		return nil
	}
	return &user
}

func (a authService) setSessionOnCtx(ctx context.Context, session db.Session) context.Context {
	return context.WithValue(ctx, SESSION_KEY, session)
}

func (a authService) getSessionFromCtx(ctx context.Context) *db.Session {
	user, ok := ctx.Value(SESSION_KEY).(db.Session)
	if !ok {
		return nil
	}
	return &user
}

func (a authService) SetTokenOnCtx(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, TOKEN_KEY, token)
}

func (a authService) getTokenFromCtx(ctx context.Context) *string {
	token, ok := ctx.Value(TOKEN_KEY).(string)
	if !ok {
		return nil
	}
	return &token
}

func newUserWithToken(user db.User, token string) UserWithSessionToken {
	return UserWithSessionToken{
		user,
		token,
	}
}

func New(queries *db.Queries) AuthService {
	return authService{
		queries,
	}
}
