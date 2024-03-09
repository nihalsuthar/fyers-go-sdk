package fyersapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type OrderBookQuery struct {
	Id       string `url:"id,omitempty"`
	OrderTag string `url:"order_tag,omitempty"`
}

type OrderBookResp struct {
	S         string         `json:"s"`
	Code      int            `json:"code"`
	Message   string         `json:"message"`
	OrderBook OrderBookItems `json:"orderBook"`
}
type OrderHistResp struct {
	S       string         `json:"s"`
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    OrderBookItems `json:"data"`
}

type MultiOrderSubResponse struct {
	Statuscode        int       `json:"statusCode"`
	Body              OrderResp `json:"body"`
	StatusDescription string    `json:"statusDescription"`
}

type MultiOrderResponse struct {
	S       string                `json:"s"`
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    MultiOrderSubResponse `json:"data"`
}

type OrderResp struct {
	S       string `json:"s"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	ID      string `json:"id"`
}

type OrderBookItem struct {
	ClientID          string  `json:"clientId"`
	ID                string  `json:"id"`
	ExchOrdID         string  `json:"exchOrdId"`
	Qty               int     `json:"qty"`
	RemainingQuantity int     `json:"remainingQuantity"`
	FilledQty         int     `json:"filledQty"`
	DiscloseQty       int     `json:"discloseQty"`
	LimitPrice        float64 `json:"limitPrice"`
	StopPrice         int     `json:"stopPrice"`
	TradedPrice       float64 `json:"tradedPrice"`
	Type              int     `json:"type"`
	FyToken           string  `json:"fyToken"`
	Exchange          int     `json:"exchange"`
	Segment           int     `json:"segment"`
	Symbol            string  `json:"symbol"`
	Instrument        int     `json:"instrument"`
	Message           string  `json:"message"`
	OfflineOrder      bool    `json:"offlineOrder"`
	OrderDateTime     string  `json:"orderDateTime"`
	OrderValidity     string  `json:"orderValidity"`
	Pan               string  `json:"pan"`
	ProductType       string  `json:"productType"`
	Side              int     `json:"side"`
	Status            int     `json:"status"`
	Source            string  `json:"source"`
	ExSym             string  `json:"ex_sym"`
	Description       string  `json:"description"`
	Ch                float64 `json:"ch"`
	Chp               float64 `json:"chp"`
	Lp                float64 `json:"lp"`
	SlNo              int     `json:"slNo"`
	DqQtyRem          int     `json:"dqQtyRem"`
	OrderNumStatus    string  `json:"orderNumStatus"`
	DisclosedQty      int     `json:"disclosedQty"`
	OrderTag          string  `json:"orderTag"`
}

type OrderBookItems []OrderBookItem

type OrderPlacementBody struct {
	Symbol       string `json:"symbol"`
	Qty          int    `json:"qty"`
	Type         int    `json:"type"`
	Side         int    `json:"side"`
	ProductType  string `json:"productType"`
	LimitPrice   int    `json:"limitPrice"`
	StopPrice    int    `json:"stopPrice"`
	Validity     string `json:"validity"`
	DisclosedQty int    `json:"disclosedQty"`
	OfflineOrder bool   `json:"offlineOrder"`
	OrderTag     string `json:"orderTag,omitempty"`
}

type OrderModificationBody struct {
	Id         string  `json:"id"`
	Qty        int     `json:"qty,omitempty"`
	Type       int     `json:"type,omitempty"`
	LimitPrice float64 `json:"limitPrice,omitempty"`
	StopPrice  float64 `json:"stopPrice,omitempty"`
	Validity   string  `json:"validity,omitempty"`
	TakeProfit float64 `json:"takeProfit,omitempty"`
	StopLoss   float64 `json:"stopLoss,omitempty"`
}

type OrderCancelBody struct {
	ID string `json:"id"`
}

func (f *FyersApi) GetOrderbook(ctx context.Context, Orderquery OrderBookQuery) (OrderBookResp, error) {
	var Resp OrderBookResp
	Url := f.baseURI + URIGetOrders
	qry, err := query.Values(Orderquery)
	if err != nil {
		f.Logger.ErrorLog(ctx, err, "query making failed", Orderquery)
		return Resp, err
	}
	err = f.DoWrapper(ctx, http.MethodGet, Url, qry, f.HeaderBuilder(), nil, &Resp)
	if err != nil {
		return Resp, err
	}
	return Resp, nil
}

func (f *FyersApi) GetOrderHistory(ctx context.Context, OrderId string) (OrderHistResp, error) {
	var Resp OrderHistResp
	Url := f.baseURI + fmt.Sprintf(URIGetOrderHistory, OrderId)

	err := f.DoWrapper(ctx, http.MethodGet, Url, nil, f.HeaderBuilder(), nil, &Resp)
	if err != nil {
		return Resp, err
	}
	return Resp, nil
}

func (f *FyersApi) PlaceOrder(ctx context.Context, Body OrderPlacementBody) (OrderResp, error) {
	var Resp OrderResp
	Url := f.baseURI + URIOrder

	err := f.DoWrapper(ctx, http.MethodPost, Url, nil, f.HeaderBuilder(), Body, &Resp)
	if err != nil {
		return Resp, err
	}
	return Resp, nil
}

func (f *FyersApi) PlaceMultiOrder(ctx context.Context, Body []OrderPlacementBody) (MultiOrderResponse, error) {
	var Resp MultiOrderResponse
	Url := f.baseURI + URIMultiOrder

	err := f.DoWrapper(ctx, http.MethodPost, Url, nil, f.HeaderBuilder(), Body, &Resp)
	if err != nil {
		return Resp, err
	}
	return Resp, nil
}

func (f *FyersApi) ModifyOrder(ctx context.Context, Body OrderModificationBody) (OrderResp, error) {
	var Resp OrderResp
	Url := f.baseURI + URIOrder

	err := f.DoWrapper(ctx, http.MethodPatch, Url, nil, f.HeaderBuilder(), Body, &Resp)
	if err != nil {
		return Resp, err
	}
	return Resp, nil
}

func (f *FyersApi) ModifyMultiOrder(ctx context.Context, Body []OrderPlacementBody) (MultiOrderResponse, error) {
	var Resp MultiOrderResponse
	Url := f.baseURI + URIMultiOrder

	err := f.DoWrapper(ctx, http.MethodPatch, Url, nil, f.HeaderBuilder(), Body, &Resp)
	if err != nil {
		return Resp, err
	}
	return Resp, nil
}

func (f *FyersApi) CancelOrder(ctx context.Context, Body OrderCancelBody) (OrderResp, error) {
	var Resp OrderResp
	Url := f.baseURI + URIOrder

	err := f.DoWrapper(ctx, http.MethodDelete, Url, nil, f.HeaderBuilder(), Body, &Resp)
	if err != nil {
		return Resp, err
	}
	return Resp, nil
}

func (f *FyersApi) DeleteMultiOrder(ctx context.Context, Body []OrderPlacementBody) (MultiOrderResponse, error) {
	var Resp MultiOrderResponse
	Url := f.baseURI + URIMultiOrder

	err := f.DoWrapper(ctx, http.MethodDelete, Url, nil, f.HeaderBuilder(), Body, &Resp)
	if err != nil {
		return Resp, err
	}
	return Resp, nil
}
