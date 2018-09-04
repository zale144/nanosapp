package service

import (
	"log"
	"context"
	proto "github.com/zale144/nanosapp/services/adCampaign/proto"
	"github.com/zale144/nanosapp/services/adCampaign/storage"
	"io/ioutil"
	"encoding/json"
)

// AdCampaignService ...
type AdCampaignService struct {
	Storage *storage.AdCampaignStorage
}

// GetAll handles requests to get all Ad Campaigns from the database
func (srv *AdCampaignService) GetAll(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	adCampaigns, err := srv.Storage.GetAll()
	if err != nil {
		log.Println(err)
		return err
	}
	rsp.AdCampaigns = adCampaigns

	return nil
}

// DataImport wipes the ad_campaign collection,
// loads the data.json file and stores it to the database
func (srv *AdCampaignService) DataImport() error {
	// read the data.json file
	data, err := ioutil.ReadFile("data/data.json")
	if err != nil {
		log.Println(err)
		return err
	}
	// unmarshal the json data into a list of AdCampaign structs
	adCampaigns := []proto.AdCampaign{}
	err = json.Unmarshal(data, &adCampaigns)
	if err != nil {
		log.Println(err)
		return err
	}
	// wipe the existing collection
	err = srv.Storage.DeleteAll()
	if err != nil {
		log.Println(err)
		// no return here
	}
	// insert the new data into the database
	for _, v := range adCampaigns {
		err = srv.Storage.Insert(v)
		if err != nil {
			log.Println(err)
			// no return here
		}
	}

	return nil
}
