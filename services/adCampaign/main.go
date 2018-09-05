package main

import (
	"os"
	"log"

	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	proto "github.com/zale144/nanosapp/services/adCampaign/proto"
	"github.com/zale144/nanosapp/services/adCampaign/service"
	"github.com/zale144/nanosapp/services/adCampaign/storage"
)

var (
	dao = storage.AdCampaignStorage{}
)

func main() {

	// get environment variables passed from the deployment descriptor
	// and connect to the MongoDB database
	dao.Server = os.Getenv("DB_SERVER")
	dao.Database = os.Getenv("DB_NAME")
	dao.Connect()

	// import the data from the 'data.json' file
	err := (&service.AdCampaignService{}).DataImport()
	if err != nil {
		log.Fatal(err)
	}

	// register the 'adcampaign' microservice on the Kubernetes cluster
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
