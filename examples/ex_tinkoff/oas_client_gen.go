// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
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
	"go.opentelemetry.io/otel/codes"
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
	_ = bits.LeadingZeros64
	_ = big.Rat{}
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
	_ = codes.Unset
)

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	cfg       config
	requests  metric.Int64Counter
	errors    metric.Int64Counter
	duration  metric.Int64Histogram
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...Option) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	c := &Client{
		cfg:       newConfig(opts...),
		serverURL: u,
	}
	if c.requests, err = c.cfg.Meter.NewInt64Counter(otelogen.ClientRequestCount); err != nil {
		return nil, err
	}
	if c.errors, err = c.cfg.Meter.NewInt64Counter(otelogen.ClientErrorsCount); err != nil {
		return nil, err
	}
	if c.duration, err = c.cfg.Meter.NewInt64Histogram(otelogen.ClientDuration); err != nil {
		return nil, err
	}
	return c, nil
}

// MarketBondsGet invokes  operation.
//
// GET /market/bonds
func (c *Client) MarketBondsGet(ctx context.Context) (res MarketBondsGetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "MarketBondsGet",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/market/bonds"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMarketBondsGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MarketCandlesGet invokes  operation.
//
// GET /market/candles
func (c *Client) MarketCandlesGet(ctx context.Context, params MarketCandlesGetParams) (res MarketCandlesGetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "MarketCandlesGet",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/market/candles"

	q := u.Query()
	{
		// Encode "figi" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Figi))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["figi"] = e.Result()
	}
	{
		// Encode "from" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.TimeToString(params.From))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["from"] = e.Result()
	}
	{
		// Encode "to" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.TimeToString(params.To))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["to"] = e.Result()
	}
	{
		// Encode "interval" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(string(params.Interval)))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["interval"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMarketCandlesGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MarketCurrenciesGet invokes  operation.
//
// GET /market/currencies
func (c *Client) MarketCurrenciesGet(ctx context.Context) (res MarketCurrenciesGetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "MarketCurrenciesGet",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/market/currencies"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMarketCurrenciesGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MarketEtfsGet invokes  operation.
//
// GET /market/etfs
func (c *Client) MarketEtfsGet(ctx context.Context) (res MarketEtfsGetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "MarketEtfsGet",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/market/etfs"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMarketEtfsGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MarketOrderbookGet invokes  operation.
//
// GET /market/orderbook
func (c *Client) MarketOrderbookGet(ctx context.Context, params MarketOrderbookGetParams) (res MarketOrderbookGetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "MarketOrderbookGet",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/market/orderbook"

	q := u.Query()
	{
		// Encode "figi" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Figi))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["figi"] = e.Result()
	}
	{
		// Encode "depth" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int32ToString(params.Depth))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["depth"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMarketOrderbookGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MarketSearchByFigiGet invokes  operation.
//
// GET /market/search/by-figi
func (c *Client) MarketSearchByFigiGet(ctx context.Context, params MarketSearchByFigiGetParams) (res MarketSearchByFigiGetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "MarketSearchByFigiGet",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/market/search/by-figi"

	q := u.Query()
	{
		// Encode "figi" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Figi))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["figi"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMarketSearchByFigiGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MarketSearchByTickerGet invokes  operation.
//
// GET /market/search/by-ticker
func (c *Client) MarketSearchByTickerGet(ctx context.Context, params MarketSearchByTickerGetParams) (res MarketSearchByTickerGetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "MarketSearchByTickerGet",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/market/search/by-ticker"

	q := u.Query()
	{
		// Encode "ticker" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Ticker))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["ticker"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMarketSearchByTickerGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MarketStocksGet invokes  operation.
//
// GET /market/stocks
func (c *Client) MarketStocksGet(ctx context.Context) (res MarketStocksGetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "MarketStocksGet",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/market/stocks"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMarketStocksGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OperationsGet invokes  operation.
//
// GET /operations
func (c *Client) OperationsGet(ctx context.Context, params OperationsGetParams) (res OperationsGetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "OperationsGet",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/operations"

	q := u.Query()
	{
		// Encode "from" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.TimeToString(params.From))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["from"] = e.Result()
	}
	{
		// Encode "to" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.TimeToString(params.To))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["to"] = e.Result()
	}
	{
		// Encode "figi" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.Figi.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["figi"] = e.Result()
	}
	{
		// Encode "brokerAccountId" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["brokerAccountId"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeOperationsGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OrdersCancelPost invokes  operation.
//
// POST /orders/cancel
func (c *Client) OrdersCancelPost(ctx context.Context, params OrdersCancelPostParams) (res OrdersCancelPostRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "OrdersCancelPost",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/orders/cancel"

	q := u.Query()
	{
		// Encode "orderId" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.OrderId))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["orderId"] = e.Result()
	}
	{
		// Encode "brokerAccountId" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["brokerAccountId"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "POST", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeOrdersCancelPostResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OrdersGet invokes  operation.
//
// GET /orders
func (c *Client) OrdersGet(ctx context.Context, params OrdersGetParams) (res OrdersGetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "OrdersGet",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/orders"

	q := u.Query()
	{
		// Encode "brokerAccountId" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["brokerAccountId"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeOrdersGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OrdersLimitOrderPost invokes  operation.
//
// POST /orders/limit-order
func (c *Client) OrdersLimitOrderPost(ctx context.Context, request LimitOrderRequest, params OrdersLimitOrderPostParams) (res OrdersLimitOrderPostRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "OrdersLimitOrderPost",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeOrdersLimitOrderPostRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutWriter(buf)
	reqBody = bytes.NewReader(buf.Buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/orders/limit-order"

	q := u.Query()
	{
		// Encode "figi" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Figi))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["figi"] = e.Result()
	}
	{
		// Encode "brokerAccountId" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["brokerAccountId"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeOrdersLimitOrderPostResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OrdersMarketOrderPost invokes  operation.
//
// POST /orders/market-order
func (c *Client) OrdersMarketOrderPost(ctx context.Context, request MarketOrderRequest, params OrdersMarketOrderPostParams) (res OrdersMarketOrderPostRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "OrdersMarketOrderPost",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeOrdersMarketOrderPostRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutWriter(buf)
	reqBody = bytes.NewReader(buf.Buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/orders/market-order"

	q := u.Query()
	{
		// Encode "figi" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Figi))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["figi"] = e.Result()
	}
	{
		// Encode "brokerAccountId" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["brokerAccountId"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeOrdersMarketOrderPostResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PortfolioCurrenciesGet invokes  operation.
//
// GET /portfolio/currencies
func (c *Client) PortfolioCurrenciesGet(ctx context.Context, params PortfolioCurrenciesGetParams) (res PortfolioCurrenciesGetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "PortfolioCurrenciesGet",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/portfolio/currencies"

	q := u.Query()
	{
		// Encode "brokerAccountId" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["brokerAccountId"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePortfolioCurrenciesGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PortfolioGet invokes  operation.
//
// GET /portfolio
func (c *Client) PortfolioGet(ctx context.Context, params PortfolioGetParams) (res PortfolioGetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "PortfolioGet",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/portfolio"

	q := u.Query()
	{
		// Encode "brokerAccountId" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["brokerAccountId"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePortfolioGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SandboxClearPost invokes  operation.
//
// Удаление всех позиций клиента.
//
// POST /sandbox/clear
func (c *Client) SandboxClearPost(ctx context.Context, params SandboxClearPostParams) (res SandboxClearPostRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "SandboxClearPost",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/sandbox/clear"

	q := u.Query()
	{
		// Encode "brokerAccountId" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["brokerAccountId"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "POST", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeSandboxClearPostResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SandboxCurrenciesBalancePost invokes  operation.
//
// POST /sandbox/currencies/balance
func (c *Client) SandboxCurrenciesBalancePost(ctx context.Context, request SandboxSetCurrencyBalanceRequest, params SandboxCurrenciesBalancePostParams) (res SandboxCurrenciesBalancePostRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "SandboxCurrenciesBalancePost",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeSandboxCurrenciesBalancePostRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutWriter(buf)
	reqBody = bytes.NewReader(buf.Buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/sandbox/currencies/balance"

	q := u.Query()
	{
		// Encode "brokerAccountId" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["brokerAccountId"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeSandboxCurrenciesBalancePostResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SandboxPositionsBalancePost invokes  operation.
//
// POST /sandbox/positions/balance
func (c *Client) SandboxPositionsBalancePost(ctx context.Context, request SandboxSetPositionBalanceRequest, params SandboxPositionsBalancePostParams) (res SandboxPositionsBalancePostRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "SandboxPositionsBalancePost",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeSandboxPositionsBalancePostRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutWriter(buf)
	reqBody = bytes.NewReader(buf.Buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/sandbox/positions/balance"

	q := u.Query()
	{
		// Encode "brokerAccountId" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["brokerAccountId"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeSandboxPositionsBalancePostResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SandboxRegisterPost invokes  operation.
//
// Создание счета и валютных позиций для клиента.
//
// POST /sandbox/register
func (c *Client) SandboxRegisterPost(ctx context.Context, request OptSandboxRegisterRequest) (res SandboxRegisterPostRes, err error) {
	if err := func() error {
		if request.Set {
			if err := func() error {
				if err := request.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "SandboxRegisterPost",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeSandboxRegisterPostRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutWriter(buf)
	reqBody = bytes.NewReader(buf.Buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/sandbox/register"

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeSandboxRegisterPostResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SandboxRemovePost invokes  operation.
//
// Удаление счета клиента.
//
// POST /sandbox/remove
func (c *Client) SandboxRemovePost(ctx context.Context, params SandboxRemovePostParams) (res SandboxRemovePostRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "SandboxRemovePost",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/sandbox/remove"

	q := u.Query()
	{
		// Encode "brokerAccountId" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["brokerAccountId"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "POST", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeSandboxRemovePostResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UserAccountsGet invokes  operation.
//
// GET /user/accounts
func (c *Client) UserAccountsGet(ctx context.Context) (res UserAccountsGetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, "UserAccountsGet",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/user/accounts"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeUserAccountsGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
