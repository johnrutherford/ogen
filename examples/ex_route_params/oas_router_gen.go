// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"strings"
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
	args := [2]string{}
	// Static code generated router with unwrapped path search.
	switch r.Method {
	case "GET":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/name"
			if l := len("/name"); len(elem) >= l && elem[0:l] == "/name" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				s.handleDataGetAnyRequest([0]string{}, w, r)

				return
			}
			switch elem[0] {
			case '/': // Prefix: "/"
				if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					s.handleDataGetIDRequest([1]string{
						args[0],
					}, w, r)

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "key"
					// Leaf parameter
					args[1] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf: DataGet
						s.handleDataGetRequest([2]string{
							args[0],
							args[1],
						}, w, r)

						return
					}
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name  string
	count int
	args  [2]string
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
		args = [2]string{}
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
		case '/': // Prefix: "/name"
			if l := len("/name"); len(elem) >= l && elem[0:l] == "/name" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				r.name = "DataGetAny"
				r.args = args
				r.count = 0
				return r, true
			}
			switch elem[0] {
			case '/': // Prefix: "/"
				if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					r.name = "DataGetID"
					r.args = args
					r.count = 1
					return r, true
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "key"
					// Leaf parameter
					args[1] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf: DataGet
						r.name = "DataGet"
						r.args = args
						r.count = 2
						return r, true
					}
				}
			}
		}
	}
	return r, false
}
