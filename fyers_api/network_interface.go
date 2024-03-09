package fyersapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"fyers-go-sdk/common"
	"io"
	"net/http"
	"net/url"
	"time"
)

type NetworkInterface interface {
	DoWrapper(ctx context.Context, httpMethod, URL string, query url.Values, header http.Header, requestBody, input any) error
}

type HttpLayer struct {
	Logger common.LoggerInterface
	Debug  bool
}

func newHttpLayer(logger common.LoggerInterface, debug bool) NetworkInterface {
	return &HttpLayer{
		Logger: logger,
		Debug:  debug,
	}
}

func (h *HttpLayer) DoWrapper(ctx context.Context, httpMethod, URL string, query url.Values, headers http.Header, requestBody, input any) error {

	if h.Debug {
		h.Logger.DebugLog(ctx, "wrapping https request", map[string]interface{}{"URL": URL, "query": query, "headers": headers, "requestbody": requestBody})
	}
	if query != nil {
		URL = fmt.Sprintf("%s?%s", URL, query.Encode())
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		h.Logger.ErrorLog(ctx, err, "failed to marshal request body", requestBody)
		return err
	}

	req, err := http.NewRequest(httpMethod, URL, bytes.NewBuffer(body))
	if err != nil {
		h.Logger.ErrorLog(ctx, err, "failed to make new request", nil)
		return err
	}

	// Set headers
	req.Header = headers

	// Perform the request
	client := http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		h.Logger.ErrorLog(ctx, err, "https request failed", nil)
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		h.Logger.ErrorLog(ctx, err, "error reading response body", nil)
		return err
	}
	err = json.Unmarshal(respBody, input)
	if err != nil {
		h.Logger.ErrorLog(ctx, err, "error parsing response body", nil)
	}
	return nil
}
