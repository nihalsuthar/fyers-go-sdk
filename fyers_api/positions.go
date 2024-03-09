package fyersapi

import (
	"context"
	"net/http"
)

type PositionBookResp struct {
	S            string                  `json:"s"`
	Code         int                     `json:"code"`
	Message      string                  `json:"message"`
	NetPositions PositionBookNetResps    `json:"netPositions"`
	Overall      PositionBookOverallResp `json:"overall"`
}

type PositionBookNetResp struct {
	NetQty           int    `json:"netQty"`
	Qty              int    `json:"qty"`
	AvgPrice         int    `json:"avgPrice"`
	NetAvg           int    `json:"netAvg"`
	Side             int    `json:"side"`
	ProductType      string `json:"productType"`
	RealizedProfit   int    `json:"realized_profit"`
	UnrealizedProfit int    `json:"unrealized_profit"`
	Pl               int    `json:"pl"`
	Ltp              int    `json:"ltp"`
	BuyQty           int    `json:"buyQty"`
	BuyAvg           int    `json:"buyAvg"`
	BuyVal           int    `json:"buyVal"`
	SellQty          int    `json:"sellQty"`
	SellAvg          int    `json:"sellAvg"`
	SellVal          int    `json:"sellVal"`
	SlNo             int    `json:"slNo"`
	FyToken          string `json:"fyToken"`
	CrossCurrency    string `json:"crossCurrency"`
	RbiRefRate       int    `json:"rbiRefRate"`
	QtyMultiCom      int    `json:"qtyMulti_com"`
	Segment          int    `json:"segment"`
	Symbol           string `json:"symbol"`
	ID               string `json:"id"`
	CfBuyQty         int    `json:"cfBuyQty"`
	CfSellQty        int    `json:"cfSellQty"`
	DayBuyQty        int    `json:"dayBuyQty"`
	DaySellQty       int    `json:"daySellQty"`
	Exchange         int    `json:"exchange"`
}

type PositionBookNetResps []PositionBookNetResp

type PositionBookOverallResp struct {
	CountTotal   int `json:"count_total"`
	CountOpen    int `json:"count_open"`
	PlTotal      int `json:"pl_total"`
	PlRealized   int `json:"pl_realized"`
	PlUnrealized int `json:"pl_unrealized"`
}

type ExitPositionBody struct {
	Id          []string `json:"id,omitempty"`
	ExitAll     int      `json:"exit_all,omitempty"`
	Segment     []int    `json:"segment,omitempty"`
	Side        []int    `json:"side,omitempty"`
	ProductType []string `json:"productType,omitempty"`
}

type ExitPositionResp struct {
	S       string `json:"s"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PositionConversionBody struct {
	Symbol      string `json:"symbol"`
	Side        int    `json:"positionSide"`
	Qty         int    `json:"convertQty"`
	FromProduct string `json:"convertFrom"`
	ToProduct   string `json:"convertTo"`
	Overnight   int    `json:"overnight"`
}

type PositionConversionResp struct {
	S               string `json:"s"`
	Code            int    `json:"code"`
	Message         string `json:"message"`
	PositionDetails int    `json:"positionDetails"`
}

func (f *FyersApi) GetPositionbook(ctx context.Context) (PositionBookResp, error) {
	var Resp PositionBookResp
	Url := f.baseURI + URIGetPositionss

	err := f.DoWrapper(ctx, http.MethodGet, Url, nil, f.HeaderBuilder(), nil, &Resp)
	if err != nil {
		return Resp, err
	}
	return Resp, nil
}

func (f *FyersApi) ExitPosition(ctx context.Context, Body ExitPositionBody) (ExitPositionResp, error) {
	var Resp ExitPositionResp
	Url := f.baseURI + URIGetPositionss

	err := f.DoWrapper(ctx, http.MethodDelete, Url, nil, f.HeaderBuilder(), Body, &Resp)
	if err != nil {
		return Resp, err
	}
	return Resp, nil
}

func (f *FyersApi) ConvertPosition(ctx context.Context, Body PositionConversionBody) (PositionConversionResp, error) {
	var Resp PositionConversionResp
	Url := f.baseURI + URIGetPositionss

	err := f.DoWrapper(ctx, http.MethodPost, Url, nil, f.HeaderBuilder(), Body, &Resp)
	if err != nil {
		return Resp, err
	}
	return Resp, nil
}
