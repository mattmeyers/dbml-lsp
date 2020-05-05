package lsp

import (
	"encoding/json"
	"log"
	"net"
)

// JSONRPCVersion is the required JSON RPC version for all requests and
// responses.
const JSONRPCVersion = "2.0"

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

// Handler represents a function that can handle an LSP request.
type Handler func(*Request, *ResponseMessage)

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

	// If the ID is not set, the request is considered a notification. We must
	// not return a response according to the specification.
	if req.ID == nil {
		h(req, nil)
		return
	}

	res := &ResponseMessage{
		JSONRPC: JSONRPCVersion,
		ID:      req.ID,
	}

	h(req, res)

	// The LSP specification requires the result field to be missing if an
	// error is present. We force the result to be nil just in case the user
	// set it.
	if res.Error != nil {
		res.Result = nil
	}

	out, err := json.Marshal(res)
	if err != nil {
		log.Println("[ERROR]: " + err.Error())
		return
	}

	_, err = conn.Write(out)
	if err != nil {
		log.Println("[ERROR]: " + err.Error())
	}
}
