package fyersapi

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

type HistoryQuery struct {
	Symbol     string `url:"symbol"`
	Resolution string `url:"resolution"`
	DateFormat int    `url:"date_format"`
	RangeFrom  string `url:"range_from"`
	RangeTo    string `url:"range_to"`
	ContFlag   string `url:"cont_flag"`
}

type HistoryResp struct {
	Candles [][]any `json:"candles"`
	Code    int     `json:"code"`
	Message string  `json:"message"`
	S       string  `json:"s"`
}

type QuotesQuery struct {
	Symbols []string
}

type quotesQuery struct {
	Symbols string `url:"symbols"`
}

type QuotesResp struct {
	Code    int           `json:"code"`
	D       []QuotesDResp `json:"d"`
	Message string        `json:"message"`
	S       string        `json:"s"`
}

type QuotesDResp struct {
	N string      `json:"n"`
	S string      `json:"s"`
	V QuotesVResp `json:"v"`
}

type QuotesVResp struct {
	Ask            int           `json:"ask"`
	Bid            float64       `json:"bid"`
	Ch             float64       `json:"ch"`
	Chp            float64       `json:"chp"`
	Cmd            QuotesCmdResp `json:"cmd"`
	Description    string        `json:"description"`
	Exchange       string        `json:"exchange"`
	FyToken        string        `json:"fyToken"`
	HighPrice      float64       `json:"high_price"`
	LowPrice       float64       `json:"low_price"`
	Lp             float64       `json:"lp"`
	OpenPrice      float64       `json:"open_price"`
	OriginalName   string        `json:"original_name"`
	PrevClosePrice float64       `json:"prev_close_price"`
	ShortName      string        `json:"short_name"`
	Spread         float64       `json:"spread"`
	Symbol         string        `json:"symbol"`
	Tt             string        `json:"tt"`
	Volume         int           `json:"volume"`
}

type QuotesCmdResp struct {
	C  float64 `json:"c"`
	H  float64 `json:"h"`
	L  float64 `json:"l"`
	O  float64 `json:"o"`
	T  int     `json:"t"`
	Tf string  `json:"tf"`
	V  int     `json:"v"`
}

type DepthQuery struct {
	Symbol    string `url:"symbol"`
	OhlcvFlag int    `url:"ohlcv_flag"`
}

type DepthResp struct {
	S       string                `json:"s"`
	D       map[string]DepthDResp `json:"d"`
	Message string                `json:"message"`
}

type DepthDResp struct {
	Totalbuyqty  int               `json:"totalbuyqty"`
	Totalsellqty int               `json:"totalsellqty"`
	Bids         []DepthBidAskResp `json:"bids"`
	Ask          []DepthBidAskResp `json:"ask"`
	O            float64           `json:"o"`
	H            float64           `json:"h"`
	L            float64           `json:"l"`
	C            float64           `json:"c"`
	Chp          float64           `json:"chp"`
	TickSize     float64           `json:"tick_Size"`
	Ch           float64           `json:"ch"`
	Ltq          int               `json:"ltq"`
	Ltt          int               `json:"ltt"`
	Ltp          float64           `json:"ltp"`
	V            int               `json:"v"`
	Atp          float64           `json:"atp"`
	LowerCkt     float64           `json:"lower_ckt"`
	UpperCkt     float64           `json:"upper_ckt"`
	Expiry       string            `json:"expiry"`
	Oi           int               `json:"oi"`
	Oiflag       bool              `json:"oiflag"`
	Pdoi         int               `json:"pdoi"`
	Oipercent    float64           `json:"oipercent"`
}

type DepthBidAskResp struct {
	Price  float64 `json:"price"`
	Volume int     `json:"volume"`
	Ord    int     `json:"ord"`
}

func (q QuotesQuery) makeString() quotesQuery {
	if q.Symbols == nil {
		return quotesQuery{Symbols: ""}
	}
	var builder strings.Builder

	for index, symbol := range q.Symbols {
		builder.WriteString(symbol)
		if index != len(q.Symbols)-1 {
			builder.WriteString(",")
		}
	}
	return quotesQuery{Symbols: builder.String()}
}

func (f *FyersApi) GetHistory(ctx context.Context, HistQuery HistoryQuery) (HistoryResp, error) {
	var Resp HistoryResp
	Url := f.baseURIData + URIGetHistory
	qry, err := query.Values(HistQuery)
	if err != nil {
		f.Logger.ErrorLog(ctx, err, "query making failed", HistQuery)
		return Resp, err
	}
	err = f.DoWrapper(ctx, http.MethodGet, Url, qry, f.HeaderBuilder(), nil, &Resp)
	if err != nil {
		return Resp, err
	}
	return Resp, nil
}

func (f *FyersApi) GetQuotes(ctx context.Context, QotQuery QuotesQuery) (QuotesResp, error) {
	var Resp QuotesResp
	Url := f.baseURIData + URIGetQuotes
	qotQuery := QotQuery.makeString()
	qry, err := query.Values(qotQuery)
	if err != nil {
		f.Logger.ErrorLog(ctx, err, "query making failed", QotQuery)
		return Resp, err
	}
	err = f.DoWrapper(ctx, http.MethodGet, Url, qry, f.HeaderBuilder(), nil, &Resp)
	if err != nil {
		return Resp, err
	}
	return Resp, nil
}

func (f *FyersApi) GetDepth(ctx context.Context, DepthQuery HistoryQuery) (DepthResp, error) {
	var Resp DepthResp
	Url := f.baseURIData + URIGetDepth
	qry, err := query.Values(DepthQuery)
	if err != nil {
		f.Logger.ErrorLog(ctx, err, "query making failed", DepthQuery)
		return Resp, err
	}
	err = f.DoWrapper(ctx, http.MethodGet, Url, qry, f.HeaderBuilder(), nil, &Resp)
	if err != nil {
		return Resp, err
	}
	return Resp, nil
}
