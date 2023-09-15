package test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"testing"
	"tracing/cmd/registry"
)

func TestGenTokenRequest(t *testing.T) {
	resp, err := resty.New().R().
		SetBody(map[string]interface{}{"accountID": "leebai", "password": "123456"}).
		Post(fmt.Sprintf("%v%v", registry.AuthServiceURL(), "/genToken"))

	fmt.Printf("resp: %v, err: %v\n", resp, err)
}

func TestDeleteTokenRequest(t *testing.T) {
	resp, err := resty.New().R().
		SetBody(map[string]interface{}{"accountID": "leebai"}).
		Post(fmt.Sprintf("%v%v", registry.AuthServiceURL(), "/deleteToken"))
	fmt.Printf("resp: %v, err: %v\n", resp, err)
}

func TestVerifyTokenHandler(t *testing.T) {
	resp, err := resty.New().R().
		SetBody(map[string]interface{}{"accountID": "leebai", "token": "token-leebai"}).
		Post(fmt.Sprintf("%v%v", registry.AuthServiceURL(), "/verifyToken"))

	fmt.Printf("resp: %v, err: %v\n", resp, err)
}
