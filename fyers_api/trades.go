package fyersapi

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
)

type TradeBookQuery struct {
	OrderTag string `url:"order_tag,omitempty"`
}

type TradeBookResp struct {
	S         string         `json:"s"`
	Code      int            `json:"code"`
	Message   string         `json:"message"`
	OrderBook TradeBookItems `json:"tradeBook"`
}

type TradeBookItem struct {
	ClientID        string  `json:"clientId"`
	OrderDateTime   string  `json:"orderDateTime"`
	OrderNumber     string  `json:"orderNumber"`
	ExchangeOrderNo string  `json:"exchangeOrderNo"`
	Exchange        int     `json:"exchange"`
	Side            int     `json:"side"`
	Segment         int     `json:"segment"`
	OrderType       int     `json:"orderType"`
	FyToken         string  `json:"fyToken"`
	ProductType     string  `json:"productType"`
	TradedQty       int     `json:"tradedQty"`
	TradePrice      float64 `json:"tradePrice"`
	TradeValue      float64 `json:"tradeValue"`
	TradeNumber     string  `json:"tradeNumber"`
	Row             int     `json:"row"`
	Symbol          string  `json:"symbol"`
	OrderTag        string  `json:"orderTag"`
}

type TradeBookItems []TradeBookItem

func (f *FyersApi) GetTradebook(ctx context.Context, Tradequery TradeBookQuery) (TradeBookResp, error) {
	var Resp TradeBookResp
	Url := f.baseURI + URIGetTrades
	qry, err := query.Values(Tradequery)
	if err != nil {
		f.Logger.ErrorLog(ctx, err, "query making failed", Tradequery)
		return Resp, err
	}
	err = f.DoWrapper(ctx, http.MethodGet, Url, qry, f.HeaderBuilder(), nil, &Resp)
	if err != nil {
		return Resp, err
	}
	return Resp, nil
}
