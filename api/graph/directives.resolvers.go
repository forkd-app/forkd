package graph

import (
	"context"
	"fmt"
	"forkd/services/auth"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func AuthDirective(authService auth.AuthService) func(context.Context, any, graphql.Resolver, *bool) (any, error) {
	return func(ctx context.Context, _obj any, next graphql.Resolver, required *bool) (any, error) {
		ctx = authService.GetUserSessionAndSetOnContext(ctx)
		user, session := authService.GetUserSessionFromCtx(ctx)
		if (user == nil || session == nil) && required != nil && *required {
			return nil, fmt.Errorf("missing auth")
		}
		if session.Expiry.Time.Before(time.Now()) {
			return nil, fmt.Errorf("session expired")
		}
		return next(ctx)
	}
}
