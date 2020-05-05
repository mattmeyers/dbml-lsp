package lsp

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"net"
	"strconv"
)

// Server represents a Language Server Protocol server.
type Server struct {
	Addr     string
	Handlers map[string]Handler
}

// NewServer initializes a new Server struct.
func NewServer(addr string) *Server {
	return &Server{
		Addr:     addr,
		Handlers: make(map[string]Handler),
	}
}

// RegisterHandler adds the provided handler to the Server's handler map. If
// a handler is already provided for the given method, it will be overwritten
// with the new handler.
func (s *Server) RegisterHandler(method string, h Handler) {
	s.Handlers[method] = h
}

// ListenAndServe resolves the Server's configured address and begins listening
// for connections. Any error returned from this method will always be non nil.
func (s *Server) ListenAndServe() error {
	lAddr, err := net.ResolveTCPAddr("tcp", s.Addr)
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP("tcp", lAddr)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	req, err := newRequest(conn)
	if err != nil {
		conn.Write([]byte(err.Error()))
		return
	}

	h, ok := s.Handlers[req.Method]
	if !ok {
		conn.Write([]byte("invalid method"))
		return
	}

	h(req, conn)
}

// Handler represents a function that can handle an LSP request.
type Handler func(*Request, ResponseWriter)

// Request holds the information for an LSP request.
type Request struct {
	Headers map[string]string `json:"-"`
	JSONRPC string            `json:"jsonrpc"`
	ID      ID                `json:"id"`
	Method  string            `json:"method"`
	Params  json.RawMessage   `json:"params"`
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
	}
	req.Headers = headers

	return &req, nil
}

// ResponseWriter is the interface for writing the response to a request.
type ResponseWriter interface {
	io.Writer
}
