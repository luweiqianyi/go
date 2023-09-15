package test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"testing"
	"tracing/cmd/registry"
)

func TestLogin(t *testing.T) {
	resp, err := resty.New().R().
		SetBody(map[string]interface{}{"accountID": "leebai", "password": "123456"}).
		Post(fmt.Sprintf("%v%v", registry.UserServiceURL(), "/login"))

	fmt.Printf("resp: %v, err: %v\n", resp, err)
}

func TestLogout(t *testing.T) {
	resp, err := resty.New().R().
		SetBody(map[string]interface{}{"accountID": "leebai"}).
		Post(fmt.Sprintf("%v%v", registry.UserServiceURL(), "/logout"))

	fmt.Printf("resp: %v, err: %v\n", resp, err)
}
