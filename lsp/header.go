package lsp

import (
	"bufio"
	"errors"
	"strings"
)

// Headers contains a map of headers. All headers are converted to lowercase
// for consistency. The LSP specification defines two headers: Content-Type and
// Content-Length. The only required header is Content-Length. Any additional
// headers will still be placed in the map, but their usage is undefined.
type Headers map[string]string

// parseHeaders reads through the provided buffer parsing each header into the
// header map. Every header name must be separated from its value with a ": ".
// Every header must end with "\r\n". The header section of the request must
// end with "\r\n". The Content-Length header is the only required header. An
// error will be returned if this header is missing.
func parseHeaders(r *bufio.Reader) (Headers, error) {
	headers := make(Headers)

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			return nil, err
		}

		if len(line) == 0 {
			break
		}

		parts := strings.Split(string(line), ": ")
		if len(parts) != 2 {
			return nil, errors.New("malformed header")
		}
		headers[strings.ToLower(parts[0])] = parts[1]
	}

	if _, ok := headers["content-length"]; !ok {
		return nil, errors.New("missing Content-Length header")
	}

	return headers, nil
}
