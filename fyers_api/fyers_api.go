package fyersapi

import (
	"fmt"
	"fyers-go-sdk/common"
	"net/http"
	"time"
)

type FyersApi struct {
	AppId       string
	AppSecret   string
	RedirectUrl string
	accessToken string
	version     string
	baseURI     string
	baseURIData string
	debug       bool
	Logger      common.LoggerInterface
	NetworkInterface
}

type FyersApiOptions struct {
	Version     string
	Debug       bool
	BaseURI     string
	BaseURIData string
	LogFilePath string
	WriteToFile *bool
	Logger      common.LoggerInterface
	NetworkInterface
}

// API endpoints
const (
	URIGetOrders       string = "/orders"
	URIGetTrades       string = "/trades"
	URIGetPositionss   string = "/positions"
	URIGetOrderHistory string = "/orders/%s" // "/orders/{id}"
	URIOrder           string = "/orders/sync"
	URIMultiOrder      string = "/multi-order/sync"
	URIGetHistory      string = "/history"
	URIGetQuotes       string = "/quotes"
	URIGetDepth        string = "/depth"
)

// BASE URI
const (
	DefaultBaseURI     string = "https://api-t1.fyers.in/api/v3"
	DefaultBaseURIData string = "https://api-t1.fyers.in/data"
)

func (f *FyersApi) HeaderBuilder() http.Header {
	Headers := make(http.Header)
	Headers.Add("Authorization", fmt.Sprintf("%s:%s", f.AppId, f.accessToken))
	Headers.Add("Version", f.version)
	return Headers
}

func (f *FyersApi) SetAccessToken(accessToken string) {
	f.accessToken = accessToken
}

func (f *FyersApi) SetAppID(appId string) {
	f.AppId = appId
}

func (f *FyersApi) SetRedirectURL(redirectURL string) {
	f.RedirectUrl = redirectURL
}

func (f *FyersApi) SetAppSecret(appSecret string) {
	f.AppSecret = appSecret
}

func (f *FyersApi) GetAccessToken() string {
	return f.accessToken
}

func NewFyersApi(Opt ...FyersApiOptions) *FyersApi {
	if len(Opt) != 0 {
		var fp string
		var writeToFile bool = true
		opt := Opt[0]
		obj := FyersApi{
			version:     "2.1",
			debug:       opt.Debug,
			baseURI:     DefaultBaseURI,
			baseURIData: DefaultBaseURIData,
		}
		if opt.Version != "" {
			obj.version = opt.Version
		}
		if opt.BaseURI != "" {
			obj.baseURI = opt.BaseURI
		}
		if opt.BaseURIData != "" {
			obj.baseURIData = opt.BaseURIData
		}
		if opt.LogFilePath != "" {
			fp = opt.LogFilePath
		} else {
			// Get the current time in the system's local timezone
			currentTime := time.Now()

			// Format the time as a string in the desired layout
			dateString := currentTime.Format("2006-01-02")

			fp = "./" + dateString + ".log"
		}
		if opt.WriteToFile != nil {
			writeToFile = *opt.WriteToFile
		}
		if opt.Logger != nil {
			obj.Logger = opt.Logger
		} else {
			obj.Logger = common.NewZLogger(writeToFile, fp)
		}
		if opt.NetworkInterface != nil {
			obj.NetworkInterface = opt.NetworkInterface
		} else {
			obj.NetworkInterface = newHttpLayer(obj.Logger, obj.debug)
		}
		return &obj
	}
	// Get the current time in the system's local timezone
	currentTime := time.Now()

	// Format the time as a string in the desired layout
	dateString := currentTime.Format("2006-01-02")

	filePath := "./" + dateString + ".log"

	logger := common.NewZLogger(true, filePath)

	return &FyersApi{
		version:          "2.1",
		debug:            false,
		baseURI:          DefaultBaseURI,
		baseURIData:      DefaultBaseURIData,
		Logger:           logger,
		NetworkInterface: newHttpLayer(logger, false),
	}
}
