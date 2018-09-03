package client

import (
	"context"
	acc "github.com/zale144/nanosapp/services/account/proto"
	"github.com/micro/go-micro"
)

var Service micro.Service

type AccountClient struct {}

// Get calls the account microservice and fetches a new account
func (ac AccountClient) Get(username string) (*acc.AccountResponse, error) {
	aClient := acc.NewAccountService("session", Service.Client())
	return aClient.Get(context.TODO(), &acc.AccountGetRequest{
		Username:  username,
	})
}

// Add calls the account microservice and creates a new account
func (ac AccountClient) Add(username, password string) (*acc.AccountResponse, error) {
	aClient := acc.NewAccountService("session", Service.Client())
	return aClient.Add(context.TODO(), &acc.AccountAddRequest{
		Username:  username,
		Password: password,
	})
}