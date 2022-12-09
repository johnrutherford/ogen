// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
	}
	if prefix := s.cfg.Prefix; len(prefix) > 0 {
		if strings.HasPrefix(elem, prefix) {
			// Cut prefix from the path.
			elem = strings.TrimPrefix(elem, prefix)
		} else {
			// Prefix doesn't match.
			s.notFound(w, r)
			return
		}
	}
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'c': // Prefix: "co"
				if l := len("co"); len(elem) >= l && elem[0:l] == "co" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'm': // Prefix: "mplicatedParameterName"
					if l := len("mplicatedParameterName"); len(elem) >= l && elem[0:l] == "mplicatedParameterName" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleComplicatedParameterNameGetRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'n': // Prefix: "ntentQueryParameter"
					if l := len("ntentQueryParameter"); len(elem) >= l && elem[0:l] == "ntentQueryParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleContentQueryParameterRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
			case 'h': // Prefix: "headerParameter"
				if l := len("headerParameter"); len(elem) >= l && elem[0:l] == "headerParameter" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleHeaderParameterRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			case 'o': // Prefix: "objectQueryParameter"
				if l := len("objectQueryParameter"); len(elem) >= l && elem[0:l] == "objectQueryParameter" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleObjectQueryParameterRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			case 'p': // Prefix: "pathObjectParameter/"
				if l := len("pathObjectParameter/"); len(elem) >= l && elem[0:l] == "pathObjectParameter/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "param"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handlePathObjectParameterRequest([1]string{
							args[0],
						}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			case 's': // Prefix: "same_name/"
				if l := len("same_name/"); len(elem) >= l && elem[0:l] == "same_name/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "path"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleSameNameRequest([1]string{
							args[0],
						}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	operationID string
	count       int
	args        [1]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
func (s *Server) FindRoute(method, path string) (r Route, _ bool) {
	var (
		args = [1]string{}
		elem = path
	)
	r.args = args
	if elem == "" {
		return r, false
	}
	defer func() {
		for i, arg := range r.args[:r.count] {
			if unescaped, err := url.PathUnescape(arg); err == nil {
				r.args[i] = unescaped
			}
		}
	}()

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'c': // Prefix: "co"
				if l := len("co"); len(elem) >= l && elem[0:l] == "co" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'm': // Prefix: "mplicatedParameterName"
					if l := len("mplicatedParameterName"); len(elem) >= l && elem[0:l] == "mplicatedParameterName" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: ComplicatedParameterNameGet
							r.name = "ComplicatedParameterNameGet"
							r.operationID = ""
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'n': // Prefix: "ntentQueryParameter"
					if l := len("ntentQueryParameter"); len(elem) >= l && elem[0:l] == "ntentQueryParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: ContentQueryParameter
							r.name = "ContentQueryParameter"
							r.operationID = "contentQueryParameter"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			case 'h': // Prefix: "headerParameter"
				if l := len("headerParameter"); len(elem) >= l && elem[0:l] == "headerParameter" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: HeaderParameter
						r.name = "HeaderParameter"
						r.operationID = "headerParameter"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'o': // Prefix: "objectQueryParameter"
				if l := len("objectQueryParameter"); len(elem) >= l && elem[0:l] == "objectQueryParameter" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: ObjectQueryParameter
						r.name = "ObjectQueryParameter"
						r.operationID = "objectQueryParameter"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'p': // Prefix: "pathObjectParameter/"
				if l := len("pathObjectParameter/"); len(elem) >= l && elem[0:l] == "pathObjectParameter/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "param"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: PathObjectParameter
						r.name = "PathObjectParameter"
						r.operationID = "pathObjectParameter"
						r.args = args
						r.count = 1
						return r, true
					default:
						return
					}
				}
			case 's': // Prefix: "same_name/"
				if l := len("same_name/"); len(elem) >= l && elem[0:l] == "same_name/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "path"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: SameName
						r.name = "SameName"
						r.operationID = "sameName"
						r.args = args
						r.count = 1
						return r, true
					default:
						return
					}
				}
			}
		}
	}
	return r, false
}
