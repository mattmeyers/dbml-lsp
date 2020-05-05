package lsp

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// Request holds the information for an LSP request.
type Request struct {
	Headers map[string]string `json:"-"`
	RequestMessage
}

// RequestMessage represents a request. If the ID is not set, then the request
// is considered a notification.
type RequestMessage struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      *ID             `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}

func newRequest(r io.Reader) (*Request, error) {
	reader := bufio.NewReader(r)

	headers, err := parseHeaders(reader)
	if err != nil {
		return nil, err
	}

	bLen, err := strconv.Atoi(headers["content-length"])
	if err != nil {
		return nil, errors.New("invalid Content-Length header value")
	}

	buf := make([]byte, bLen)
	_, err = reader.Read(buf)
	if err != nil {
		return nil, err
	}

	var req Request
	err = json.Unmarshal(buf, &req)
	if err != nil {
		return nil, err
	} else if req.JSONRPC != JSONRPCVersion {
		return nil, errors.New("invalid jsonrpc version")
	}
	req.Headers = headers

	return &req, nil
}
