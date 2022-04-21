// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
)

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	s.cfg.NotFound(w, r)
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	// Static code generated router with unwrapped path search.
	switch r.Method {
	case "GET":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/healthz"
			if l := len("/healthz"); len(elem) >= l && elem[0:l] == "/healthz" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				// Leaf: ProbeLiveness
				s.handleProbeLivenessRequest([0]string{}, w, r)

				return
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name  string
	count int
	args  [0]string
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.name
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
func (s *Server) FindRoute(method, path string) (r Route, _ bool) {
	var (
		args = [0]string{}
		elem = path
	)
	r.args = args
	if elem == "" {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch method {
	case "GET":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/healthz"
			if l := len("/healthz"); len(elem) >= l && elem[0:l] == "/healthz" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				// Leaf: ProbeLiveness
				r.name = "ProbeLiveness"
				r.args = args
				r.count = 0
				return r, true
			}
		}
	}
	return r, false
}
