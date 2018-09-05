package main

import (
	"os"
	"log"
	"fmt"

	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	"github.com/zale144/nanosapp/services/account/db"
	"github.com/zale144/nanosapp/services/account/handler"
	proto "github.com/zale144/nanosapp/services/account/proto"
)

func main() {

	// get environment variables passed from the deployment descriptor
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	db.DBInfo = fmt.Sprintf("postgres://%s:%s@db/%s?sslmode=disable", dbUser, dbPass, dbName)

	// start and init the db connection
	err := db.InitDB()
	if err != nil {
		log.Fatalf("cannot initialize db: %v", err)
		return
	}

	// register the 'account' microservice on the Kubernetes cluster
	serv := k8s.NewService(
		micro.Name("account"),
	)
	serv.Init()

	proto.RegisterAccountServiceHandler(serv.Server(), &handler.Account{})

	if err := serv.Run(); err != nil {
		log.Fatal(err)
	}

}
