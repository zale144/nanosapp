package client

import (
	ad "github.com/zale144/nanosapp/services/adCampaign/proto"
	"github.com/zale144/nanosapp/services/web/commons"
	"context"
)

type AdCampaignClient struct {}

func (ac AdCampaignClient) GetAll() ([]*ad.AdCampaign, error) {

	adClient := ad.NewAdCampaignService("adcampaign", commons.Service.Client())
	adCampaignRsp, err := adClient.GetAll(context.TODO(), &ad.Request{Token: ""})
	if err != nil {
		return nil, err
	}

	return adCampaignRsp.AdCampaigns, nil
}
