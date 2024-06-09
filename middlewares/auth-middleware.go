package middlewares

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/piyush7833/Chat-Api/functions"
	"github.com/piyush7833/Chat-Api/helpers"
)

// type ContextKey string

// const jwtContextKey = ContextKey("jwtData")

func AuthMiddleware(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := helpers.GetCookie(r, "token")
		if err != nil {
			helpers.Error(w, http.StatusUnauthorized, err.Error())
			return
		}

		data, err := helpers.VerifyJwt(tokenString)
		if err != nil {
			helpers.Error(w, http.StatusUnauthorized, err.Error())
			return
		}

		claims, ok := data.(*helpers.Claims)
		if !ok {
			helpers.Error(w, http.StatusUnauthorized, "Invalid token claims")
			return
		}

		// user, err := helpers.Prisma.User.FindUnique(db.User.ID.Equals(claims.Id)).Exec(r.Context())
		var cols = []string{"id"}
		user, error := functions.GetParticularUserById(claims.Id, cols)
		if error.StatusCode != 0 {
			helpers.Error(w, http.StatusUnauthorized, "Invalid token claims")
			return
		}

		if user == nil {
			helpers.Error(w, http.StatusUnauthorized, "Invalid token claims")
			return
		}
		context.Set(r, "userId", data)
		// ctx := context.WithValue(r.Context(), jwtContextKey, data)
		// fmt.Println(ctx.Value(jwtContextKey).(*helpers.Claims), "data from middleware")
		// context
		// r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}
