package handler

import (
	"log"
	"context"

	"github.com/zale144/nanosapp/services/account/model"
	"github.com/zale144/nanosapp/services/account/storage"
	proto "github.com/zale144/nanosapp/services/account/proto"
)

// Account implements the proto service Account
type Account struct{}

// Get handles the gRPC get account request
func (m *Account) Get(ctx context.Context, req *proto.AccountRequest, rsp *proto.AccountResponse) error {
	account, err := storage.AccountStorage{}.GetByUsername(req.Username)
	if err != nil {
		log.Println(err)
		return err
	}
	rsp.Account = &proto.Account{
		Username: account.Username,
		Password: account.Password,
	}
	return nil
}

// Add handles the add account request
func (m *Account) Add(ctx context.Context, req *proto.Account, rsp *proto.AccountResponse) error {
	err := storage.AccountStorage{}.Insert(model.Account{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	rsp.Account = &proto.Account{
		Username: req.Username,
	}
	return nil
}