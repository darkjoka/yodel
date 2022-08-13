package auth

import (
	"context"
	"net/http"

	"github.com/darkjoka/yodel/graph/jwt"
	"github.com/darkjoka/yodel/graph/model"
)

var userCtxKey = "user"

type contextKey struct {
	name string
}

func Middleware(u model.UserScheme) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			id, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user and check if user exists in db
			user, err := u.FindById(id, context.Background())
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) (*model.User, bool) {
	raw, ok := ctx.Value(userCtxKey).(*model.User)
	return raw, ok
}
