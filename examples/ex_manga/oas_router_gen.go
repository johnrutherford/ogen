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
	args := [3]string{}

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
			case 'a': // Prefix: "api/galler"
				if l := len("api/galler"); len(elem) >= l && elem[0:l] == "api/galler" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'i': // Prefix: "ies/"
					if l := len("ies/"); len(elem) >= l && elem[0:l] == "ies/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 's': // Prefix: "search"
						if l := len("search"); len(elem) >= l && elem[0:l] == "search" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleSearchRequest([0]string{}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					case 't': // Prefix: "tagged"
						if l := len("tagged"); len(elem) >= l && elem[0:l] == "tagged" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleSearchByTagIDRequest([0]string{}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				case 'y': // Prefix: "y/"
					if l := len("y/"); len(elem) >= l && elem[0:l] == "y/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "book_id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleGetBookRequest([1]string{
								args[0],
							}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
			case 'g': // Prefix: "galleries/"
				if l := len("galleries/"); len(elem) >= l && elem[0:l] == "galleries/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "media_id"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

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
					case 'c': // Prefix: "cover."
						if l := len("cover."); len(elem) >= l && elem[0:l] == "cover." {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "format"
						// Leaf parameter
						args[1] = elem
						elem = ""

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleGetPageCoverImageRequest([2]string{
									args[0],
									args[1],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
					// Param: "page"
					// Match until one of ".t"
					idx := strings.IndexAny(elem, ".t")
					if idx < 0 {
						idx = len(elem)
					}
					args[1] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case '.': // Prefix: "."
						if l := len("."); len(elem) >= l && elem[0:l] == "." {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "format"
						// Leaf parameter
						args[2] = elem
						elem = ""

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleGetPageImageRequest([3]string{
									args[0],
									args[1],
									args[2],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					case 't': // Prefix: "t."
						if l := len("t."); len(elem) >= l && elem[0:l] == "t." {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "format"
						// Leaf parameter
						args[2] = elem
						elem = ""

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleGetPageThumbnailImageRequest([3]string{
									args[0],
									args[1],
									args[2],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
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
	args        [3]string
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
		args = [3]string{}
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
			case 'a': // Prefix: "api/galler"
				if l := len("api/galler"); len(elem) >= l && elem[0:l] == "api/galler" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'i': // Prefix: "ies/"
					if l := len("ies/"); len(elem) >= l && elem[0:l] == "ies/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 's': // Prefix: "search"
						if l := len("search"); len(elem) >= l && elem[0:l] == "search" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: Search
								r.name = "Search"
								r.operationID = "search"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					case 't': // Prefix: "tagged"
						if l := len("tagged"); len(elem) >= l && elem[0:l] == "tagged" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: SearchByTagID
								r.name = "SearchByTagID"
								r.operationID = "searchByTagID"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					}
				case 'y': // Prefix: "y/"
					if l := len("y/"); len(elem) >= l && elem[0:l] == "y/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "book_id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: GetBook
							r.name = "GetBook"
							r.operationID = "getBook"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
				}
			case 'g': // Prefix: "galleries/"
				if l := len("galleries/"); len(elem) >= l && elem[0:l] == "galleries/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "media_id"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

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
					case 'c': // Prefix: "cover."
						if l := len("cover."); len(elem) >= l && elem[0:l] == "cover." {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "format"
						// Leaf parameter
						args[1] = elem
						elem = ""

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: GetPageCoverImage
								r.name = "GetPageCoverImage"
								r.operationID = "getPageCoverImage"
								r.args = args
								r.count = 2
								return r, true
							default:
								return
							}
						}
					}
					// Param: "page"
					// Match until one of ".t"
					idx := strings.IndexAny(elem, ".t")
					if idx < 0 {
						idx = len(elem)
					}
					args[1] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case '.': // Prefix: "."
						if l := len("."); len(elem) >= l && elem[0:l] == "." {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "format"
						// Leaf parameter
						args[2] = elem
						elem = ""

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: GetPageImage
								r.name = "GetPageImage"
								r.operationID = "getPageImage"
								r.args = args
								r.count = 3
								return r, true
							default:
								return
							}
						}
					case 't': // Prefix: "t."
						if l := len("t."); len(elem) >= l && elem[0:l] == "t." {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "format"
						// Leaf parameter
						args[2] = elem
						elem = ""

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: GetPageThumbnailImage
								r.name = "GetPageThumbnailImage"
								r.operationID = "getPageThumbnailImage"
								r.args = args
								r.count = 3
								return r, true
							default:
								return
							}
						}
					}
				}
			}
		}
	}
	return r, false
}
