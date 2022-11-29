// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// MarketBondsGet implements GET /market/bonds operation.
	//
	// Получение списка облигаций.
	//
	// GET /market/bonds
	MarketBondsGet(ctx context.Context) (MarketBondsGetRes, error)
	// MarketCandlesGet implements GET /market/candles operation.
	//
	// Получение исторических свечей по FIGI.
	//
	// GET /market/candles
	MarketCandlesGet(ctx context.Context, params MarketCandlesGetParams) (MarketCandlesGetRes, error)
	// MarketCurrenciesGet implements GET /market/currencies operation.
	//
	// Получение списка валютных пар.
	//
	// GET /market/currencies
	MarketCurrenciesGet(ctx context.Context) (MarketCurrenciesGetRes, error)
	// MarketEtfsGet implements GET /market/etfs operation.
	//
	// Получение списка ETF.
	//
	// GET /market/etfs
	MarketEtfsGet(ctx context.Context) (MarketEtfsGetRes, error)
	// MarketOrderbookGet implements GET /market/orderbook operation.
	//
	// Получение стакана по FIGI.
	//
	// GET /market/orderbook
	MarketOrderbookGet(ctx context.Context, params MarketOrderbookGetParams) (MarketOrderbookGetRes, error)
	// MarketSearchByFigiGet implements GET /market/search/by-figi operation.
	//
	// Получение инструмента по FIGI.
	//
	// GET /market/search/by-figi
	MarketSearchByFigiGet(ctx context.Context, params MarketSearchByFigiGetParams) (MarketSearchByFigiGetRes, error)
	// MarketSearchByTickerGet implements GET /market/search/by-ticker operation.
	//
	// Получение инструмента по тикеру.
	//
	// GET /market/search/by-ticker
	MarketSearchByTickerGet(ctx context.Context, params MarketSearchByTickerGetParams) (MarketSearchByTickerGetRes, error)
	// MarketStocksGet implements GET /market/stocks operation.
	//
	// Получение списка акций.
	//
	// GET /market/stocks
	MarketStocksGet(ctx context.Context) (MarketStocksGetRes, error)
	// OperationsGet implements GET /operations operation.
	//
	// Получение списка операций.
	//
	// GET /operations
	OperationsGet(ctx context.Context, params OperationsGetParams) (OperationsGetRes, error)
	// OrdersCancelPost implements POST /orders/cancel operation.
	//
	// Отмена заявки.
	//
	// POST /orders/cancel
	OrdersCancelPost(ctx context.Context, params OrdersCancelPostParams) (OrdersCancelPostRes, error)
	// OrdersGet implements GET /orders operation.
	//
	// Получение списка активных заявок.
	//
	// GET /orders
	OrdersGet(ctx context.Context, params OrdersGetParams) (OrdersGetRes, error)
	// OrdersLimitOrderPost implements POST /orders/limit-order operation.
	//
	// Создание лимитной заявки.
	//
	// POST /orders/limit-order
	OrdersLimitOrderPost(ctx context.Context, req *LimitOrderRequest, params OrdersLimitOrderPostParams) (OrdersLimitOrderPostRes, error)
	// OrdersMarketOrderPost implements POST /orders/market-order operation.
	//
	// Создание рыночной заявки.
	//
	// POST /orders/market-order
	OrdersMarketOrderPost(ctx context.Context, req *MarketOrderRequest, params OrdersMarketOrderPostParams) (OrdersMarketOrderPostRes, error)
	// PortfolioCurrenciesGet implements GET /portfolio/currencies operation.
	//
	// Получение валютных активов клиента.
	//
	// GET /portfolio/currencies
	PortfolioCurrenciesGet(ctx context.Context, params PortfolioCurrenciesGetParams) (PortfolioCurrenciesGetRes, error)
	// PortfolioGet implements GET /portfolio operation.
	//
	// Получение портфеля клиента.
	//
	// GET /portfolio
	PortfolioGet(ctx context.Context, params PortfolioGetParams) (PortfolioGetRes, error)
	// SandboxClearPost implements POST /sandbox/clear operation.
	//
	// Удаление всех позиций клиента.
	//
	// POST /sandbox/clear
	SandboxClearPost(ctx context.Context, params SandboxClearPostParams) (SandboxClearPostRes, error)
	// SandboxCurrenciesBalancePost implements POST /sandbox/currencies/balance operation.
	//
	// Выставление баланса по валютным позициям.
	//
	// POST /sandbox/currencies/balance
	SandboxCurrenciesBalancePost(ctx context.Context, req *SandboxSetCurrencyBalanceRequest, params SandboxCurrenciesBalancePostParams) (SandboxCurrenciesBalancePostRes, error)
	// SandboxPositionsBalancePost implements POST /sandbox/positions/balance operation.
	//
	// Выставление баланса по инструментным позициям.
	//
	// POST /sandbox/positions/balance
	SandboxPositionsBalancePost(ctx context.Context, req *SandboxSetPositionBalanceRequest, params SandboxPositionsBalancePostParams) (SandboxPositionsBalancePostRes, error)
	// SandboxRegisterPost implements POST /sandbox/register operation.
	//
	// Создание счета и валютных позиций для клиента.
	//
	// POST /sandbox/register
	SandboxRegisterPost(ctx context.Context, req OptSandboxRegisterRequest) (SandboxRegisterPostRes, error)
	// SandboxRemovePost implements POST /sandbox/remove operation.
	//
	// Удаление счета клиента.
	//
	// POST /sandbox/remove
	SandboxRemovePost(ctx context.Context, params SandboxRemovePostParams) (SandboxRemovePostRes, error)
	// UserAccountsGet implements GET /user/accounts operation.
	//
	// Получение брокерских счетов клиента.
	//
	// GET /user/accounts
	UserAccountsGet(ctx context.Context) (UserAccountsGetRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
