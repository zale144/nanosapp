package handler

import (
	"log"
	"context"

	proto "github.com/zale144/nanosapp/services/account/proto"
	"github.com/zale144/nanosapp/services/account/storage"
)

// Account implements the proto service Account
type Account struct{}

// Get handles the get account request
func (m *Account) Get(ctx context.Context, req *proto.AccountGetRequest, rsp *proto.AccountResponse) error {
	account, err := storage.AccountStorage{}.GetByUsername(req.Username)
	if err != nil {
		log.Println(err)
		return err
	}
	rsp.Account = account.Username
	return nil
}

// Add handles the add account request
func (m *Account) Add(ctx context.Context, req *proto.AccountAddRequest, rsp *proto.AccountResponse) error {
	account, err := storage.AccountStorage{}.GetByUsername(req.Username)
	if err != nil {
		log.Println(err)
		return err
	}
	rsp.Account = account.Username
	return nil
}