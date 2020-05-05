package lsp

// Error represents a JSON RPC error. Error codes in the range of -32768 to
// -32000 are reserved by the JSON RPC 2.0 specification.
type Error int

// Defined error codes
const (
	// JSON RPC 2.0 Errors
	ErrParseError           Error = -32700
	ErrInvalidRequest       Error = -32600
	ErrMethodNotFound       Error = -32601
	ErrInvalidParams        Error = -32602
	ErrInternalError        Error = -32603
	ErrserverErrorStart     Error = -32099
	ErrserverErrorEnd       Error = -32000
	ErrServerNotInitialized Error = -32002
	ErrUnknownErrorCode     Error = -32001

	// LSP Specific Errors
	ErrRequestCancelled Error = -32800
	ErrContentModified  Error = -32801
)

var errorStrings = map[int]string{
	-32700: "parse error",
	-32600: "invalid request",
	-32601: "method not found",
	-32602: "invalid params",
	-32603: "internal error",
	-32099: "server error start",
	-32000: "server error end",
	-32002: "server not initialized",
	-32001: "unknown error code",
	-32800: "request cancelled",
	-32801: "content modified",
}

func (e Error) Error() string { return errorStrings[int(e)] }

// ErrorResponse represents an error to be returned to the client.
type ErrorResponse struct {
	Code    Error       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewErrorResponse creates a new error response. If the empty string is
// passed for the msg parameter, then the Error's message is used.
func NewErrorResponse(err Error, msg string, data interface{}) *ErrorResponse {
	if msg == "" {
		msg = err.Error()
	}

	return &ErrorResponse{
		Code:    err,
		Message: msg,
		Data:    data,
	}
}
