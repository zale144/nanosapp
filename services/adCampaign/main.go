
package main

import (
	"log"

	k8s "github.com/micro/kubernetes/go/micro"
	proto "github.com/zale144/nanosapp/services/adCampaign/proto"
	"github.com/micro/go-micro"
	"github.com/zale144/nanosapp/services/adCampaign/service"
	"github.com/zale144/nanosapp/services/adCampaign/storage"
	"flag"
)

var (
	dao = storage.AdCampaignStorage{}
	dbServer = flag.String("db-server", "localhost", "database server")
	db = flag.String("db", "AdCampaigns", "database name")
)

func main() {

	dao.Server = *dbServer
	dao.Database = *db
	dao.Connect()

	srvc := k8s.NewService(
		micro.Name("adcampaign"),
		micro.Version("latest"),
	)

	srvc .Init()
	proto.RegisterAdCampaignServiceHandler(srvc .Server(), &service.AdCampaignService{})

	if err := srvc.Run(); err != nil {
		log.Fatal(err)
	}

}
