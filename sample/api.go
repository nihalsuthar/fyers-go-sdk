package main

import (
	"context"
	"fmt"
	fyersapi "github.com/nihalsuthar/fyers-go-sdk/fyers_api"
)

func main() {
	obj := fyersapi.NewFyersApi()
	obj.SetAppID("Qxxxx5-1xx")
	obj.SetAppSecret("4xxxxxC")
	obj.SetRedirectURL("https://www.google.com")
	obj.SetAccessToken("xxxxx")

	resp, err := obj.GetOrderbook(context.Background(), fyersapi.OrderBookQuery{})
	fmt.Println(resp, err)

}
