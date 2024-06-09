package helpers

import "net/http"

func CheckNilErr(err error, message string, w http.ResponseWriter) bool {
	if err != nil {
		Error(w, 500, message)
		return true
	}
	return false
}
