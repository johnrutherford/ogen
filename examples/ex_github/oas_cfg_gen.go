// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"regexp"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

var regexMap = map[string]*regexp.Regexp{
	"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4})$": regexp.MustCompile("^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4})$"),
	"^(?:first|last|after:\\d+)$":                             regexp.MustCompile("^(?:first|last|after:\\d+)$"),
	"^(?:top|bottom|after:\\d+)$":                             regexp.MustCompile("^(?:top|bottom|after:\\d+)$"),
	"^[0-9a-fA-F]+$":                                          regexp.MustCompile("^[0-9a-fA-F]+$"),
	"^ssh-(rsa|dss|ed25519) |^ecdsa-sha2-nistp(256|384|521) ": regexp.MustCompile("^ssh-(rsa|dss|ed25519) |^ecdsa-sha2-nistp(256|384|521) "),
}

// ErrorHandler is error handler.
type ErrorHandler = ogenerrors.ErrorHandler

type config struct {
	TracerProvider     trace.TracerProvider
	Tracer             trace.Tracer
	MeterProvider      metric.MeterProvider
	Meter              metric.Meter
	Client             ht.Client
	NotFound           http.HandlerFunc
	MethodNotAllowed   func(w http.ResponseWriter, r *http.Request, allowed string)
	ErrorHandler       ErrorHandler
	MaxMultipartMemory int64
}

func newConfig(opts ...Option) config {
	cfg := config{
		TracerProvider: otel.GetTracerProvider(),
		MeterProvider:  metric.NewNoopMeterProvider(),
		Client:         http.DefaultClient,
		NotFound:       http.NotFound,
		MethodNotAllowed: func(w http.ResponseWriter, r *http.Request, allowed string) {
			w.Header().Set("Allow", allowed)
			w.WriteHeader(http.StatusMethodNotAllowed)
		},
		ErrorHandler:       ogenerrors.DefaultErrorHandler,
		MaxMultipartMemory: 32 << 20, // 32 MB
	}
	for _, opt := range opts {
		opt.apply(&cfg)
	}
	cfg.Tracer = cfg.TracerProvider.Tracer(otelogen.Name,
		trace.WithInstrumentationVersion(otelogen.SemVersion()),
	)
	cfg.Meter = cfg.MeterProvider.Meter(otelogen.Name)
	return cfg
}

type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (o optionFunc) apply(c *config) {
	o(c)
}

// WithTracerProvider specifies a tracer provider to use for creating a tracer.
//
// If none is specified, the global provider is used.
func WithTracerProvider(provider trace.TracerProvider) Option {
	return optionFunc(func(cfg *config) {
		if provider != nil {
			cfg.TracerProvider = provider
		}
	})
}

// WithMeterProvider specifies a meter provider to use for creating a meter.
//
// If none is specified, the metric.NewNoopMeterProvider is used.
func WithMeterProvider(provider metric.MeterProvider) Option {
	return optionFunc(func(cfg *config) {
		if provider != nil {
			cfg.MeterProvider = provider
		}
	})
}

// WithClient specifies http client to use.
func WithClient(client ht.Client) Option {
	return optionFunc(func(cfg *config) {
		if client != nil {
			cfg.Client = client
		}
	})
}

// WithNotFound specifies Not Found handler to use.
func WithNotFound(notFound http.HandlerFunc) Option {
	return optionFunc(func(cfg *config) {
		if notFound != nil {
			cfg.NotFound = notFound
		}
	})
}

// WithMethodNotAllowed specifies Method Not Allowed handler to use.
func WithMethodNotAllowed(methodNotAllowed func(w http.ResponseWriter, r *http.Request, allowed string)) Option {
	return optionFunc(func(cfg *config) {
		if methodNotAllowed != nil {
			cfg.MethodNotAllowed = methodNotAllowed
		}
	})
}

// WithErrorHandler specifies error handler to use.
func WithErrorHandler(h ErrorHandler) Option {
	return optionFunc(func(cfg *config) {
		if h != nil {
			cfg.ErrorHandler = h
		}
	})
}

// WithMaxMultipartMemory specifies limit of memory for storing file parts.
// File parts which can't be stored in memory will be stored on disk in temporary files.
func WithMaxMultipartMemory(max int64) Option {
	return optionFunc(func(cfg *config) {
		if max > 0 {
			cfg.MaxMultipartMemory = max
		}
	})
}
