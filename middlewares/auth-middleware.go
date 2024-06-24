package middlewares

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/piyush7833/Chat-Api/helpers"
)

// type ContextKey string

// const jwtContextKey = ContextKey("jwtData")

func AuthMiddleware(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := helpers.GetCookie(r, "token")
		// fmt.Println(tokenString, "token from middleware")
		if err != nil {
			// fmt.Println(err, "error from middleware")
			helpers.Error(w, http.StatusUnauthorized, "login required")
			return
		}

		data, err := helpers.VerifyJwt(tokenString)
		if err != nil {
			// fmt.Println(err, "error from middleware")
			helpers.Error(w, http.StatusUnauthorized, "login required")
			return
		}
		// claims, ok := data.(*helpers.Claims)
		// if !ok {
		// 	helpers.Error(w, http.StatusUnauthorized, "Invalid token claims")
		// 	return
		// }
		// fmt.Println(claims.Id, "data from middleware")

		// var cols = []string{"id"}
		// user, error := functions.GetParticularUserById(claims.Id, cols)

		// if error.StatusCode != 0 {
		// 	helpers.Error(w, http.StatusUnauthorized, "Invalid token claims")
		// 	return
		// }

		// if user == nil {
		// 	helpers.Error(w, http.StatusUnauthorized, "Invalid token claims")
		// 	return
		// }
		context.Set(r, "userId", data)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}
