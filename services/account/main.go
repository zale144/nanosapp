package main

import (
	"log"
	"flag"

	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	proto "github.com/zale144/nanosapp/services/account/proto"
	"github.com/zale144/nanosapp/services/account/db"
	"github.com/zale144/nanosapp/services/account/handler"
)

var (
	dbInfo = flag.String("db-info", "postgres://postgres:admin@localhost/nanoaccount?sslmode=disable", "database connection string")
)

func main() {

	flag.Parse()

	db.DBInfo = *dbInfo

	log.Println(db.DBInfo)

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
