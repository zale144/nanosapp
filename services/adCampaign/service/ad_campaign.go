package service

import (
	"log"
	"context"
	"github.com/zale144/nanosapp/services/adCampaign/model"
	proto "github.com/zale144/nanosapp/services/adCampaign/proto"
	"github.com/zale144/nanosapp/services/adCampaign/storage"
	"io/ioutil"
	"encoding/json"
	"fmt"
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
	rsp.AdCampaigns = srv.mapAdCampaignsToProto(adCampaigns...)

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
	// convert the proto type into the model type
	modelAdCampaigns := srv.mapAdCampaignsToModel(adCampaigns...)
	fmt.Println(len(modelAdCampaigns))
	// insert the new data into the database
	for _, v := range modelAdCampaigns {
		err = srv.Storage.Insert(*v)
		if err != nil {
			log.Println(err)
			// no return here
		}
	}

	return nil
}

// mapAdCampaignsToProto converts the model.AdCampaign to proto.AdCampaign type
func (srv *AdCampaignService) mapAdCampaignsToProto(adCampaigns ...model.AdCampaign) []*proto.AdCampaign {
	grpcAdCampaigns := []*proto.AdCampaign{}
	for _, c := range adCampaigns {
		ac := &proto.AdCampaign{
			ID:       c.ID,
			Name:    c.Name,
			Goal:      c.Goal,
			TotalBudget: c.TotalBudget,
			Status: c.Status,
		}
		fb := &proto.Platform{
			Status: c.Platforms.Facebook.Status,
			TotalBudget: c.Platforms.Facebook.TotalBudget,
			RemainingBudget: c.Platforms.Facebook.RemainingBudget,
			StartDate: c.Platforms.Facebook.StartDate,
			EndDate: c.Platforms.Facebook.EndDate,

		}
		fta := &proto.TargetAudiance{
			Languages: c.Platforms.Facebook.TargetAudiance.Languages,
			Genders: c.Platforms.Facebook.TargetAudiance.Genders,
			AgeRange: c.Platforms.Facebook.TargetAudiance.AgeRange,
			Locations: c.Platforms.Facebook.TargetAudiance.Locations,
			KeyWords: c.Platforms.Facebook.TargetAudiance.KeyWords,
			Interests: c.Platforms.Facebook.TargetAudiance.Interests,
		}
		fct := &proto.Creatives{
			Header: c.Platforms.Facebook.Creatives.Header,
			Header1:c.Platforms.Facebook.Creatives.Header1,
			Header2:c.Platforms.Facebook.Creatives.Header2,
			Description:c.Platforms.Facebook.Creatives.Description,
			URL:c.Platforms.Facebook.Creatives.URL,
			Image:c.Platforms.Facebook.Creatives.Image,
		}
		fin := &proto.Insights{
			Impressions: c.Platforms.Facebook.Insights.Impressions,
			Clicks: c.Platforms.Facebook.Insights.Clicks,
			WebsiteVisits: c.Platforms.Facebook.Insights.WebsiteVisits,
			CostPerClick: c.Platforms.Facebook.Insights.CostPerClick,
			ClickThroughRate: c.Platforms.Facebook.Insights.ClickThroughRate,
			AdvancedKpi1: c.Platforms.Facebook.Insights.AdvancedKpi1,
			AdvancedKpi2: c.Platforms.Facebook.Insights.AdvancedKpi2,
			NanosScore: c.Platforms.Facebook.Insights.NanosScore,
		}
		fb.TargetAudiance = fta
		fb.Creatives = fct
		fb.Insights = fin

		ins := &proto.Platform{
			Status: c.Platforms.Instagram.Status,
			TotalBudget: c.Platforms.Instagram.TotalBudget,
			RemainingBudget: c.Platforms.Instagram.RemainingBudget,
			StartDate: c.Platforms.Instagram.StartDate,
			EndDate: c.Platforms.Instagram.EndDate,

		}
		insta := &proto.TargetAudiance{
			Languages: c.Platforms.Instagram.TargetAudiance.Languages,
			Genders: c.Platforms.Instagram.TargetAudiance.Genders,
			AgeRange: c.Platforms.Instagram.TargetAudiance.AgeRange,
			Locations: c.Platforms.Instagram.TargetAudiance.Locations,
			KeyWords: c.Platforms.Instagram.TargetAudiance.KeyWords,
			Interests: c.Platforms.Instagram.TargetAudiance.Interests,
		}
		insct := &proto.Creatives{
			Header: c.Platforms.Instagram.Creatives.Header,
			Header1:c.Platforms.Instagram.Creatives.Header1,
			Header2:c.Platforms.Instagram.Creatives.Header2,
			Description:c.Platforms.Instagram.Creatives.Description,
			URL:c.Platforms.Instagram.Creatives.URL,
			Image:c.Platforms.Instagram.Creatives.Image,
		}
		insin := &proto.Insights{
			Impressions: c.Platforms.Instagram.Insights.Impressions,
			Clicks: c.Platforms.Instagram.Insights.Clicks,
			WebsiteVisits: c.Platforms.Instagram.Insights.WebsiteVisits,
			CostPerClick: c.Platforms.Instagram.Insights.CostPerClick,
			ClickThroughRate: c.Platforms.Instagram.Insights.ClickThroughRate,
			AdvancedKpi1: c.Platforms.Instagram.Insights.AdvancedKpi1,
			AdvancedKpi2: c.Platforms.Instagram.Insights.AdvancedKpi2,
			NanosScore: c.Platforms.Instagram.Insights.NanosScore,
		}
		ins.TargetAudiance = insta
		ins.Creatives = insct
		ins.Insights = insin

		gog := &proto.Platform{
			Status: c.Platforms.Google.Status,
			TotalBudget: c.Platforms.Google.TotalBudget,
			RemainingBudget: c.Platforms.Google.RemainingBudget,
			StartDate: c.Platforms.Google.StartDate,
			EndDate: c.Platforms.Google.EndDate,

		}
		gogta := &proto.TargetAudiance{
			Languages: c.Platforms.Google.TargetAudiance.Languages,
			Genders: c.Platforms.Google.TargetAudiance.Genders,
			AgeRange: c.Platforms.Google.TargetAudiance.AgeRange,
			Locations: c.Platforms.Google.TargetAudiance.Locations,
			KeyWords: c.Platforms.Google.TargetAudiance.KeyWords,
			Interests: c.Platforms.Google.TargetAudiance.Interests,
		}
		gogct := &proto.Creatives{
			Header: c.Platforms.Google.Creatives.Header,
			Header1:c.Platforms.Google.Creatives.Header1,
			Header2:c.Platforms.Google.Creatives.Header2,
			Description:c.Platforms.Google.Creatives.Description,
			URL:c.Platforms.Google.Creatives.URL,
			Image:c.Platforms.Google.Creatives.Image,
		}
		gogin := &proto.Insights{
			Impressions: c.Platforms.Google.Insights.Impressions,
			Clicks: c.Platforms.Google.Insights.Clicks,
			WebsiteVisits: c.Platforms.Google.Insights.WebsiteVisits,
			CostPerClick: c.Platforms.Google.Insights.CostPerClick,
			ClickThroughRate: c.Platforms.Google.Insights.ClickThroughRate,
			AdvancedKpi1: c.Platforms.Google.Insights.AdvancedKpi1,
			AdvancedKpi2: c.Platforms.Google.Insights.AdvancedKpi2,
			NanosScore: c.Platforms.Google.Insights.NanosScore,
		}
		gog.TargetAudiance = gogta
		gog.Creatives = gogct
		gog.Insights = gogin

		ac.Platforms = &proto.Platforms{
			Facebook: fb,
			Instagram: ins,
			Google: gog,
		}

		grpcAdCampaigns = append(grpcAdCampaigns, ac)

	}
	return grpcAdCampaigns
}

// mapAdCampaignsToModel converts the proto.AdCampaign to model.AdCampaign type
func (srv *AdCampaignService) mapAdCampaignsToModel(adCampaigns ...proto.AdCampaign) []*model.AdCampaign {
	grpcAdCampaigns := []*model.AdCampaign{}
	for _, c := range adCampaigns {
		ac := &model.AdCampaign{
			ID:       c.ID,
			Name:    c.Name,
			Goal:      c.Goal,
			TotalBudget: c.TotalBudget,
			Status: c.Status,
		}
		fb := &model.Platform{
			Status: c.Platforms.Facebook.Status,
			TotalBudget: c.Platforms.Facebook.TotalBudget,
			RemainingBudget: c.Platforms.Facebook.RemainingBudget,
			StartDate: c.Platforms.Facebook.StartDate,
			EndDate: c.Platforms.Facebook.EndDate,

		}
		fta := &model.TargetAudiance{
			Languages: c.Platforms.Facebook.TargetAudiance.Languages,
			Genders: c.Platforms.Facebook.TargetAudiance.Genders,
			AgeRange: c.Platforms.Facebook.TargetAudiance.AgeRange,
			Locations: c.Platforms.Facebook.TargetAudiance.Locations,
			KeyWords: c.Platforms.Facebook.TargetAudiance.KeyWords,
			Interests: c.Platforms.Facebook.TargetAudiance.Interests,
		}
		fct := &model.Creatives{
			Header: c.Platforms.Facebook.Creatives.Header,
			Header1:c.Platforms.Facebook.Creatives.Header1,
			Header2:c.Platforms.Facebook.Creatives.Header2,
			Description:c.Platforms.Facebook.Creatives.Description,
			URL:c.Platforms.Facebook.Creatives.URL,
			Image:c.Platforms.Facebook.Creatives.Image,
		}
		fin := &model.Insights{
			Impressions: c.Platforms.Facebook.Insights.Impressions,
			Clicks: c.Platforms.Facebook.Insights.Clicks,
			WebsiteVisits: c.Platforms.Facebook.Insights.WebsiteVisits,
			CostPerClick: c.Platforms.Facebook.Insights.CostPerClick,
			ClickThroughRate: c.Platforms.Facebook.Insights.ClickThroughRate,
			AdvancedKpi1: c.Platforms.Facebook.Insights.AdvancedKpi1,
			AdvancedKpi2: c.Platforms.Facebook.Insights.AdvancedKpi2,
			NanosScore: c.Platforms.Facebook.Insights.NanosScore,
		}
		fb.TargetAudiance = fta
		fb.Creatives = fct
		fb.Insights = fin

		ins := &model.Platform{
			Status: c.Platforms.Instagram.Status,
			TotalBudget: c.Platforms.Instagram.TotalBudget,
			RemainingBudget: c.Platforms.Instagram.RemainingBudget,
			StartDate: c.Platforms.Instagram.StartDate,
			EndDate: c.Platforms.Instagram.EndDate,

		}
		insta := &model.TargetAudiance{
			Languages: c.Platforms.Instagram.TargetAudiance.Languages,
			Genders: c.Platforms.Instagram.TargetAudiance.Genders,
			AgeRange: c.Platforms.Instagram.TargetAudiance.AgeRange,
			Locations: c.Platforms.Instagram.TargetAudiance.Locations,
			KeyWords: c.Platforms.Instagram.TargetAudiance.KeyWords,
			Interests: c.Platforms.Instagram.TargetAudiance.Interests,
		}
		insct := &model.Creatives{
			Header: c.Platforms.Instagram.Creatives.Header,
			Header1:c.Platforms.Instagram.Creatives.Header1,
			Header2:c.Platforms.Instagram.Creatives.Header2,
			Description:c.Platforms.Instagram.Creatives.Description,
			URL:c.Platforms.Instagram.Creatives.URL,
			Image:c.Platforms.Instagram.Creatives.Image,
		}
		insin := &model.Insights{
			Impressions: c.Platforms.Instagram.Insights.Impressions,
			Clicks: c.Platforms.Instagram.Insights.Clicks,
			WebsiteVisits: c.Platforms.Instagram.Insights.WebsiteVisits,
			CostPerClick: c.Platforms.Instagram.Insights.CostPerClick,
			ClickThroughRate: c.Platforms.Instagram.Insights.ClickThroughRate,
			AdvancedKpi1: c.Platforms.Instagram.Insights.AdvancedKpi1,
			AdvancedKpi2: c.Platforms.Instagram.Insights.AdvancedKpi2,
			NanosScore: c.Platforms.Instagram.Insights.NanosScore,
		}
		ins.TargetAudiance = insta
		ins.Creatives = insct
		ins.Insights = insin

		gog := &model.Platform{
			Status: c.Platforms.Google.Status,
			TotalBudget: c.Platforms.Google.TotalBudget,
			RemainingBudget: c.Platforms.Google.RemainingBudget,
			StartDate: c.Platforms.Google.StartDate,
			EndDate: c.Platforms.Google.EndDate,

		}
		gogta := &model.TargetAudiance{
			Languages: c.Platforms.Google.TargetAudiance.Languages,
			Genders: c.Platforms.Google.TargetAudiance.Genders,
			AgeRange: c.Platforms.Google.TargetAudiance.AgeRange,
			Locations: c.Platforms.Google.TargetAudiance.Locations,
			KeyWords: c.Platforms.Google.TargetAudiance.KeyWords,
			Interests: c.Platforms.Google.TargetAudiance.Interests,
		}
		gogct := &model.Creatives{
			Header: c.Platforms.Google.Creatives.Header,
			Header1:c.Platforms.Google.Creatives.Header1,
			Header2:c.Platforms.Google.Creatives.Header2,
			Description:c.Platforms.Google.Creatives.Description,
			URL:c.Platforms.Google.Creatives.URL,
			Image:c.Platforms.Google.Creatives.Image,
		}
		gogin := &model.Insights{
			Impressions: c.Platforms.Google.Insights.Impressions,
			Clicks: c.Platforms.Google.Insights.Clicks,
			WebsiteVisits: c.Platforms.Google.Insights.WebsiteVisits,
			CostPerClick: c.Platforms.Google.Insights.CostPerClick,
			ClickThroughRate: c.Platforms.Google.Insights.ClickThroughRate,
			AdvancedKpi1: c.Platforms.Google.Insights.AdvancedKpi1,
			AdvancedKpi2: c.Platforms.Google.Insights.AdvancedKpi2,
			NanosScore: c.Platforms.Google.Insights.NanosScore,
		}
		gog.TargetAudiance = gogta
		gog.Creatives = gogct
		gog.Insights = gogin

		ac.Platforms = &model.Platforms{
			Facebook: fb,
			Instagram: ins,
			Google: gog,
		}

		grpcAdCampaigns = append(grpcAdCampaigns, ac)

	}
	return grpcAdCampaigns
}
