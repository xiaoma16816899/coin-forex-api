package types

type PaginationMeta struct {
	Page    uint64 `json:"page"`
	Limit   uint64 `json:"limit"`
	Total   uint64 `json:"total"`
	HasNext bool   `json:"hasNext"`
}

type PaginationResponse[T any] struct {
	Data []T             `json:"data"`
	Meta *PaginationMeta `json:"meta"`
}

type ErrorResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
}

type ListResponse[T any] struct {
	Data []T `json:"data"`
}
type DataResponse[T any] struct {
	Data T `json:"data"`
}
