package test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"testing"
)

const (
	scheme  = "http"
	apiHost = "127.0.0.1"
	port    = "8002"
	request = "/hello"
)

func TestHello(t *testing.T) {
	url := fmt.Sprintf("%v://%v:%v%v", scheme, apiHost, port, request)

	resp, err := resty.New().R().Get(url)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("resp: %v\n", resp)
}
