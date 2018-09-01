package client

import (
	"context"
	"github.com/zale144/nanosapp/services/web/commons"
)


type AccountClient struct {}

// Get calls the account microservice and fetches a new session
func (ac AccountClient) Get(account, password string) (string, error) {
	aClient := acc.NewAccountService("session", commons.Service.Client())
	rsp, err := aClient.Get(context.TODO(), &acc.SessionRequest{
		Account:  account,
		Password: password,
	})
	if err != nil {
		return "", err
	}
	return rsp.Account, nil
}
