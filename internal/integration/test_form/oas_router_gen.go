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
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
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
			case 'o': // Prefix: "only"
				if l := len("only"); len(elem) >= l && elem[0:l] == "only" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'F': // Prefix: "Form"
					if l := len("Form"); len(elem) >= l && elem[0:l] == "Form" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleOnlyFormRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
				case 'M': // Prefix: "MultipartF"
					if l := len("MultipartF"); len(elem) >= l && elem[0:l] == "MultipartF" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'i': // Prefix: "ile"
						if l := len("ile"); len(elem) >= l && elem[0:l] == "ile" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleOnlyMultipartFileRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					case 'o': // Prefix: "orm"
						if l := len("orm"); len(elem) >= l && elem[0:l] == "orm" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleOnlyMultipartFormRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					}
				}
			case 't': // Prefix: "test"
				if l := len("test"); len(elem) >= l && elem[0:l] == "test" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'F': // Prefix: "FormURLEncoded"
					if l := len("FormURLEncoded"); len(elem) >= l && elem[0:l] == "FormURLEncoded" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleTestFormURLEncodedRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
				case 'M': // Prefix: "Multipart"
					if l := len("Multipart"); len(elem) >= l && elem[0:l] == "Multipart" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "POST":
							s.handleTestMultipartRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
					switch elem[0] {
					case 'U': // Prefix: "Upload"
						if l := len("Upload"); len(elem) >= l && elem[0:l] == "Upload" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleTestMultipartUploadRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					}
				case 'R': // Prefix: "ReuseForm"
					if l := len("ReuseForm"); len(elem) >= l && elem[0:l] == "ReuseForm" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'O': // Prefix: "OptionalSchema"
						if l := len("OptionalSchema"); len(elem) >= l && elem[0:l] == "OptionalSchema" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleTestReuseFormOptionalSchemaRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					case 'S': // Prefix: "Schema"
						if l := len("Schema"); len(elem) >= l && elem[0:l] == "Schema" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleTestReuseFormSchemaRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					}
				case 'S': // Prefix: "ShareFormSchema"
					if l := len("ShareFormSchema"); len(elem) >= l && elem[0:l] == "ShareFormSchema" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleTestShareFormSchemaRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

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
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [0]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

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
			case 'o': // Prefix: "only"
				if l := len("only"); len(elem) >= l && elem[0:l] == "only" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'F': // Prefix: "Form"
					if l := len("Form"); len(elem) >= l && elem[0:l] == "Form" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: OnlyForm
							r.name = "OnlyForm"
							r.summary = ""
							r.operationID = "onlyForm"
							r.pathPattern = "/onlyForm"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'M': // Prefix: "MultipartF"
					if l := len("MultipartF"); len(elem) >= l && elem[0:l] == "MultipartF" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'i': // Prefix: "ile"
						if l := len("ile"); len(elem) >= l && elem[0:l] == "ile" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: OnlyMultipartFile
								r.name = "OnlyMultipartFile"
								r.summary = ""
								r.operationID = "onlyMultipartFile"
								r.pathPattern = "/onlyMultipartFile"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					case 'o': // Prefix: "orm"
						if l := len("orm"); len(elem) >= l && elem[0:l] == "orm" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: OnlyMultipartForm
								r.name = "OnlyMultipartForm"
								r.summary = ""
								r.operationID = "onlyMultipartForm"
								r.pathPattern = "/onlyMultipartForm"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					}
				}
			case 't': // Prefix: "test"
				if l := len("test"); len(elem) >= l && elem[0:l] == "test" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'F': // Prefix: "FormURLEncoded"
					if l := len("FormURLEncoded"); len(elem) >= l && elem[0:l] == "FormURLEncoded" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: TestFormURLEncoded
							r.name = "TestFormURLEncoded"
							r.summary = ""
							r.operationID = "testFormURLEncoded"
							r.pathPattern = "/testFormURLEncoded"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'M': // Prefix: "Multipart"
					if l := len("Multipart"); len(elem) >= l && elem[0:l] == "Multipart" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							r.name = "TestMultipart"
							r.summary = ""
							r.operationID = "testMultipart"
							r.pathPattern = "/testMultipart"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case 'U': // Prefix: "Upload"
						if l := len("Upload"); len(elem) >= l && elem[0:l] == "Upload" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: TestMultipartUpload
								r.name = "TestMultipartUpload"
								r.summary = ""
								r.operationID = "testMultipartUpload"
								r.pathPattern = "/testMultipartUpload"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					}
				case 'R': // Prefix: "ReuseForm"
					if l := len("ReuseForm"); len(elem) >= l && elem[0:l] == "ReuseForm" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'O': // Prefix: "OptionalSchema"
						if l := len("OptionalSchema"); len(elem) >= l && elem[0:l] == "OptionalSchema" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: TestReuseFormOptionalSchema
								r.name = "TestReuseFormOptionalSchema"
								r.summary = ""
								r.operationID = "testReuseFormOptionalSchema"
								r.pathPattern = "/testReuseFormOptionalSchema"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					case 'S': // Prefix: "Schema"
						if l := len("Schema"); len(elem) >= l && elem[0:l] == "Schema" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: TestReuseFormSchema
								r.name = "TestReuseFormSchema"
								r.summary = ""
								r.operationID = "testReuseFormSchema"
								r.pathPattern = "/testReuseFormSchema"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					}
				case 'S': // Prefix: "ShareFormSchema"
					if l := len("ShareFormSchema"); len(elem) >= l && elem[0:l] == "ShareFormSchema" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: TestShareFormSchema
							r.name = "TestShareFormSchema"
							r.summary = ""
							r.operationID = "testShareFormSchema"
							r.pathPattern = "/testShareFormSchema"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			}
		}
	}
	return r, false
}
