package domain

type CustomError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
