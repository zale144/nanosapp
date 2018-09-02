package main

import (
	"log"
	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	proto "github.com/zale144/nanosapp/services/account/proto"
	"github.com/zale144/nanosapp/services/account/db"
	"github.com/zale144/nanosapp/services/account/handler"
	"os"
	"fmt"
)

func main() {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbConnString := fmt.Sprintf("postgres://%s:%s@db/%s?sslmode=disable", dbUser, dbPass, dbName)

	db.DBInfo = dbConnString

	err := db.InitDB()
	if err != nil {
		log.Fatalf("cannot initialize db: %v", err)
		return
	}

	serv := k8s.NewService(
		micro.Name("account"),
	)
	serv.Init()

	proto.RegisterAccountHandler(serv.Server(), &handler.Account{})

	if err := serv.Run(); err != nil {
		log.Fatal(err)
	}

}
