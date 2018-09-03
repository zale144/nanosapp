package main

import (
	"os"
	"log"

	k8s "github.com/micro/kubernetes/go/micro"
	proto "github.com/zale144/nanosapp/services/adCampaign/proto"
	"github.com/micro/go-micro"
	"github.com/zale144/nanosapp/services/adCampaign/service"
	"github.com/zale144/nanosapp/services/adCampaign/storage"
)

var (
	dao = storage.AdCampaignStorage{}
)

func main() {

	dao.Server = os.Getenv("DB_SERVER")
	dao.Database = os.Getenv("DB_NAME")
	dao.Connect()

	// import data
	err := (&service.AdCampaignService{}).DataImport()
	if err != nil {
		log.Fatal(err)
	}

	srvc := k8s.NewService(
		micro.Name("adcampaign"),
		micro.Version("latest"),
	)

	srvc.Init()
	proto.RegisterAdCampaignServiceHandler(srvc.Server(), &service.AdCampaignService{})

	if err := srvc.Run(); err != nil {
		log.Fatal(err)
	}

}
