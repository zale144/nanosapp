package client

import (
	"context"
	"github.com/zale144/nanosapp/services/web/commons"
	ad "github.com/zale144/nanosapp/services/adCampaign/proto"
)

type AdCampaignClient struct {}

// GetAll fetches all ad campaigns from the 'adcampaign' microservice
func (ac AdCampaignClient) GetAll() ([]*ad.AdCampaign, error) {
	adClient := ad.NewAdCampaignService("adcampaign", commons.Service.Client())
	adCampaignRsp, err := adClient.GetAll(context.TODO(), &ad.Request{Token: ""})
	if err != nil {
		return nil, err
	}
	return adCampaignRsp.AdCampaigns, nil
}
