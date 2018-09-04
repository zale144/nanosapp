package client

import (
	"context"
	acc "github.com/zale144/nanosapp/services/account/proto"
	"github.com/zale144/nanosapp/services/web/commons"
)

type AccountClient struct {}

// Get calls the account microservice and fetches a new account
func (ac AccountClient) Get(username string) (*acc.Account, error) {
	aClient := acc.NewAccountService("account", commons.Service.Client())
	accountResponse, err := aClient.Get(context.TODO(), &acc.AccountRequest{
		Username: username,
	})
	if err != nil {
		return nil, err
	}
	return accountResponse.Account, nil
}

// Add calls the account microservice and creates a new account
func (ac AccountClient) Add(username, password string) (*acc.Account, error) {
	aClient := acc.NewAccountService("account", commons.Service.Client())

	accountResponse, err := aClient.Add(context.TODO(), &acc.Account{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return accountResponse.Account, nil
}