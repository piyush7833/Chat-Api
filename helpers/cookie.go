package helpers

import (
	"net/http"
	"time"
)

func GetCookie(r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func SetCookie(w http.ResponseWriter, name string, value string, expires time.Time) error {
	http.SetCookie(w, &http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expires,
		Path:    "/",
	})
	return nil
}
