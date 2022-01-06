// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
)

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}

	args := map[string]string{}
	// Static code generated router with unwrapped path search.
	switch r.Method {
	case "GET":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if prefix := "/"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
				elem = elem[len(prefix):]
			} else {
				break
			}

			if len(elem) == 0 {
				s.handleOperationsGetRequest(args, w, r)
				return
			}
			switch elem[0] {
			case 'm': // Prefix: "market/"
				if prefix := "market/"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleMarketCandlesGetRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'b': // Prefix: "bonds"
					if prefix := "bonds"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: MarketBondsGet
					s.handleMarketBondsGetRequest(args, w, r)
					return
				case 'c': // Prefix: "c"
					if prefix := "c"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleMarketCurrenciesGetRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'a': // Prefix: "andles"
						if prefix := "andles"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: MarketCandlesGet
						s.handleMarketCandlesGetRequest(args, w, r)
						return
					case 'u': // Prefix: "urrencies"
						if prefix := "urrencies"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: MarketCurrenciesGet
						s.handleMarketCurrenciesGetRequest(args, w, r)
						return
					}
				case 'e': // Prefix: "etfs"
					if prefix := "etfs"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: MarketEtfsGet
					s.handleMarketEtfsGetRequest(args, w, r)
					return
				case 'o': // Prefix: "orderbook"
					if prefix := "orderbook"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: MarketOrderbookGet
					s.handleMarketOrderbookGetRequest(args, w, r)
					return
				case 's': // Prefix: "s"
					if prefix := "s"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleMarketStocksGetRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'e': // Prefix: "earch/by-"
						if prefix := "earch/by-"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						if len(elem) == 0 {
							s.handleMarketSearchByTickerGetRequest(args, w, r)
							return
						}
						switch elem[0] {
						case 'f': // Prefix: "figi"
							if prefix := "figi"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: MarketSearchByFigiGet
							s.handleMarketSearchByFigiGetRequest(args, w, r)
							return
						case 't': // Prefix: "ticker"
							if prefix := "ticker"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: MarketSearchByTickerGet
							s.handleMarketSearchByTickerGetRequest(args, w, r)
							return
						}
					case 't': // Prefix: "tocks"
						if prefix := "tocks"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: MarketStocksGet
						s.handleMarketStocksGetRequest(args, w, r)
						return
					}
				}
			case 'o': // Prefix: "o"
				if prefix := "o"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleOrdersGetRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'p': // Prefix: "perations"
					if prefix := "perations"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: OperationsGet
					s.handleOperationsGetRequest(args, w, r)
					return
				case 'r': // Prefix: "rders"
					if prefix := "rders"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: OrdersGet
					s.handleOrdersGetRequest(args, w, r)
					return
				}
			case 'p': // Prefix: "portfolio"
				if prefix := "portfolio"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handlePortfolioGetRequest(args, w, r)
					return
				}
				switch elem[0] {
				case '/': // Prefix: "/currencies"
					if prefix := "/currencies"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: PortfolioCurrenciesGet
					s.handlePortfolioCurrenciesGetRequest(args, w, r)
					return
				}
			case 'u': // Prefix: "user/accounts"
				if prefix := "user/accounts"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				// Leaf: UserAccountsGet
				s.handleUserAccountsGetRequest(args, w, r)
				return
			}
		}
	case "POST":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if prefix := "/"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
				elem = elem[len(prefix):]
			} else {
				break
			}

			if len(elem) == 0 {
				s.handleSandboxClearPostRequest(args, w, r)
				return
			}
			switch elem[0] {
			case 'o': // Prefix: "orders/"
				if prefix := "orders/"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleOrdersLimitOrderPostRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'c': // Prefix: "cancel"
					if prefix := "cancel"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: OrdersCancelPost
					s.handleOrdersCancelPostRequest(args, w, r)
					return
				case 'l': // Prefix: "limit-order"
					if prefix := "limit-order"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: OrdersLimitOrderPost
					s.handleOrdersLimitOrderPostRequest(args, w, r)
					return
				case 'm': // Prefix: "market-order"
					if prefix := "market-order"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: OrdersMarketOrderPost
					s.handleOrdersMarketOrderPostRequest(args, w, r)
					return
				}
			case 's': // Prefix: "sandbox/"
				if prefix := "sandbox/"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleSandboxPositionsBalancePostRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'c': // Prefix: "c"
					if prefix := "c"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleSandboxCurrenciesBalancePostRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'l': // Prefix: "lear"
						if prefix := "lear"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: SandboxClearPost
						s.handleSandboxClearPostRequest(args, w, r)
						return
					case 'u': // Prefix: "urrencies/balance"
						if prefix := "urrencies/balance"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: SandboxCurrenciesBalancePost
						s.handleSandboxCurrenciesBalancePostRequest(args, w, r)
						return
					}
				case 'p': // Prefix: "positions/balance"
					if prefix := "positions/balance"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: SandboxPositionsBalancePost
					s.handleSandboxPositionsBalancePostRequest(args, w, r)
					return
				case 'r': // Prefix: "re"
					if prefix := "re"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleSandboxRemovePostRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'g': // Prefix: "gister"
						if prefix := "gister"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: SandboxRegisterPost
						s.handleSandboxRegisterPostRequest(args, w, r)
						return
					case 'm': // Prefix: "move"
						if prefix := "move"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: SandboxRemovePost
						s.handleSandboxRemovePostRequest(args, w, r)
						return
					}
				}
			}
		}
	}
	s.notFound(w, r)
}
