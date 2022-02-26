package response

type ErrorResponse struct {
	Message    string `json:"message"`
	HttpStatus int    `json:"-"`
}
