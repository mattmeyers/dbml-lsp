package lsp

// Error represents a JSON RPC error. Error codes in the range of -32768 to
// -32000 are reserved by the JSON RPC 2.0 specification.
type Error int

// Defined error codes
const (
	// JSON RPC 2.0 Errors
	ParseError           Error = -32700
	InvalidRequest       Error = -32600
	MethodNotFound       Error = -32601
	InvalidParams        Error = -32602
	InternalError        Error = -32603
	serverErrorStart     Error = -32099
	serverErrorEnd       Error = -32000
	ServerNotInitialized Error = -32002
	UnknownErrorCode     Error = -32001

	// LSP Specific Errors
	RequestCancelled Error = -32800
	ContentModified  Error = -32801
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
