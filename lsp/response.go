package lsp

import (
	"encoding/json"
	"strings"
)

// ResponseMessage represents a response. JSONRPC must always be set to "2.0".
// If the request was notification, the the ID will not be set. Finally, if
// an error occurred and Error is set, then Result must be nil.
type ResponseMessage struct {
	JSONRPC string         `json:"jsonrpc"`
	ID      *ID            `json:"id"`
	Result  interface{}    `json:"result"`
	Error   *ErrorResponse `json:"error,omitempty"`
}

// MarshalJSON implements json.Marshaler for ResponseMessage. This method
// handles properly marshalling the Result and Error. If the Error field is
// set, then the Result fields must not be marshalled. Otherwise, the Result
// field must be present.
func (r ResponseMessage) MarshalJSON() ([]byte, error) {
	var sb strings.Builder
	sb.WriteString(`{`)

	sb.WriteString(`"jsonrpc": `)
	s, err := json.Marshal(r.JSONRPC)
	if err != nil {
		return nil, err
	}
	sb.WriteString(string(s))
	sb.WriteString(`,`)

	sb.WriteString(`"id": `)
	s, err = json.Marshal(r.ID)
	if err != nil {
		return nil, err
	}
	sb.WriteString(string(s))
	sb.WriteString(`,`)

	if r.Error == nil {
		sb.WriteString(`"result": `)
		s, err = json.Marshal(r.Result)
		if err != nil {
			return nil, err
		}
		sb.WriteString(string(s))
	} else {
		sb.WriteString(`"error": `)
		s, err = json.Marshal(r.Error)
		if err != nil {
			return nil, err
		}
		sb.WriteString(string(s))
	}

	sb.WriteString(`}`)
	return []byte(sb.String()), nil
}
