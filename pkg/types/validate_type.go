package types

type ValidateErrorResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Errors  []ValidateFieldError `json:"errors"`
}

type ValidateFieldError struct {
	Code    int    `json:"code"`
	Field   string `json:"field"`
	Message string `json:"message"`
}
