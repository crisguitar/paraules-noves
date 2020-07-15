package common

type AppError struct {
	HttpCode int    `json:"code"`
	Message  string `json:"error"`
}

func (e AppError) Error() string {
	return e.Message
}
