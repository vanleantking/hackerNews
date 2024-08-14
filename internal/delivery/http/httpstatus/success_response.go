package httpstatus

type Response[T any] struct {
	Data       T       `json:"data,omitempty"`
	Message    string  `json:"message"`
	Errors     *string `json:"errors,omitempty"`
	StatusCode int     `json:"status_code"`
}

const (
	Success = "success"
	Error   = "error"
)
