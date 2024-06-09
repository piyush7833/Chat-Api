package types

type ErrorType struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}
