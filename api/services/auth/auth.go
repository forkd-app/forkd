package auth

import (
	"context"
	"errors"
	"fmt"
	"forkd/db"
	"forkd/services/email"
	"forkd/util"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
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

type MagicLinkLookup struct {
	Token string
	Code  string
}

type AuthService interface {
	CreateSession(ctx context.Context, userId pgtype.UUID, code *string) (UserWithSessionToken, error)
	DeleteSession(ctx context.Context, sessionId pgtype.UUID) error
	GetUserSessionAndSetOnCtx(ctx context.Context) context.Context
	GetUserSessionFromCtx(ctx context.Context) (*db.User, *db.Session)
	SessionWrapper(http.HandlerFunc) http.HandlerFunc
	SetTokenOnCtx(ctx context.Context, token string) context.Context
	ValidateMagicLink(ctx context.Context, code string, token string) (pgtype.UUID, error)
	Signup(ctx context.Context, email string, displayName string) (*string, error)
	RequestMagicLink(ctx context.Context, email string) (*string, error)
	createMagicLink(ctx context.Context, user db.User) (*MagicLinkLookup, error)
	getTokenFromCtx(ctx context.Context) *string
	getUserFromCtx(ctx context.Context) *db.User
	getSessionFromCtx(ctx context.Context) *db.Session
	setUserOnCtx(ctx context.Context, user db.User) context.Context
	setSessionOnCtx(ctx context.Context, session db.Session) context.Context
}

type authService struct {
	conn         *pgxpool.Pool
	queries      *db.Queries
	emailService email.EmailService
}

// TODO: Refactor to use a transaction.
func (a authService) RequestMagicLink(ctx context.Context, email string) (*string, error) {
	user, err := a.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("email not registered: %s", email)
		}

		return nil, err
	}

	lookup, err := a.createMagicLink(ctx, user)
	if err != nil {
		return nil, err
	}

	return &lookup.Token, nil
}

// TODO: Refactor to use a transaction.
func (a authService) Signup(ctx context.Context, email string, displayName string) (*string, error) {
	user, err := a.queries.CreateUser(ctx, db.CreateUserParams{
		Email:       email,
		DisplayName: displayName,
	})
	if err != nil {
		fmt.Println(err)
		// If the error returned is a postgres error with the code for a unique violation, we just return null
		// You can find a list of the error codes here: https://www.postgresql.org/docs/current/errcodes-appendix.html
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, nil
		}

		return nil, err
	}

	lookup, err := a.createMagicLink(ctx, user)
	if err != nil {
		return nil, err
	}

	return &lookup.Token, nil
}

func (a authService) DeleteSession(ctx context.Context, sessionId pgtype.UUID) error {
	return a.queries.DeleteSession(ctx, sessionId)
}

func (a authService) GetUserSessionAndSetOnCtx(ctx context.Context) context.Context {
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

	err = util.DecodeBase64StringToStruct(token, &magicLinkToken)
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
		return id, fmt.Errorf("magic link expired")
	}

	return magicLink.UserID, nil
}

func (a authService) createMagicLink(ctx context.Context, user db.User) (*MagicLinkLookup, error) {
	// Set the expiry for 10 minutes
	expiry := time.Now().Add(10 * time.Minute)
	token := pgtype.UUID{
		Bytes: uuid.New(),
		Valid: true,
	}

	magicLink, err := a.queries.CreateMagicLink(ctx, db.CreateMagicLinkParams{
		UserID: user.ID,
		Token:  token,
		Expiry: pgtype.Timestamp{
			Time:  expiry,
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	magicLinkCodeStruct := MagicLinkCode{
		ID: magicLink.ID,
	}

	magicLinkCode, err := util.EncodeStructToBase64String(magicLinkCodeStruct)
	if err != nil {
		return nil, err
	}

	magicLinkTokenStruct := MagicLinkToken{
		Token: magicLink.Token,
	}

	magicLinkToken, err := util.EncodeStructToBase64String(magicLinkTokenStruct)
	if err != nil {
		return nil, err
	}

	lookup := MagicLinkLookup{
		Token: magicLinkToken,
		Code:  magicLinkCode,
	}

	// FIXME: Remove this before we push this to the public internet.
	// THIS SHOULD ONLY BE USED FOR LOCAL TESTING
	if util.GetEnv().GetSendMagicLinkEmail() {
		// TODO: This should probably be done async, either just using a goroutine or a task queue
		emailData, err := a.emailService.SendMagicLink(ctx, lookup.Code, user.Email)
		if err != nil {
			return nil, err
		} else if emailData.Data.Failed > 0 || emailData.Data.Succeeded < 1 {
			return nil, fmt.Errorf("error sending auth email: %+v", emailData.Data.Failures)
		}
	} else {
		fmt.Printf("MAGIC LINK CODE: %s\n", lookup.Code)
	}

	return &lookup, nil
}

func (a authService) CreateSession(ctx context.Context, userId pgtype.UUID, code *string) (UserWithSessionToken, error) {
	tx, err := a.conn.Begin(ctx)
	if err != nil {
		return UserWithSessionToken{}, err
	}
	// I'm ignoring the linter here because I feel we can safely ignore the error here
	defer tx.Rollback(ctx) //nolint:errcheck
	qtx := a.queries.WithTx(tx)
	// Set expiry to 2 weeks, this will refresh everyime we access the session
	expiry := time.Now().AddDate(0, 0, 14)

	result, err := qtx.CreateSession(ctx, db.CreateSessionParams{
		UserID: userId,
		Expiry: pgtype.Timestamp{
			Time:  expiry,
			Valid: true,
		},
	})
	if err != nil {
		return UserWithSessionToken{}, err
	}

	session := SessionToken{
		ID: result.ID,
	}

	sessionToken, err := util.EncodeStructToBase64String(session)
	if err != nil {
		return UserWithSessionToken{}, err
	}

	if code != nil {
		var magicLinkCode MagicLinkCode

		err = util.DecodeBase64StringToStruct(*code, &magicLinkCode)
		if err != nil {
			return UserWithSessionToken{}, err
		}

		err = qtx.DeleteMagicLinkById(ctx, magicLinkCode.ID)
		if err != nil {
			return UserWithSessionToken{}, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return UserWithSessionToken{}, err
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

func New(queries *db.Queries, conn *pgxpool.Pool, emailService email.EmailService) AuthService {
	return authService{
		queries:      queries,
		conn:         conn,
		emailService: emailService,
	}
}
