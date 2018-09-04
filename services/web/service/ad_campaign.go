package service

import (
	"github.com/labstack/echo"
	"github.com/zale144/nanosapp/services/web/client"
	"net/http"
	"fmt"
)

type AdCampaignService struct {}

func (as AdCampaignService) GetAll(c echo.Context) error {

	//adCampaigns := []ad.AdCampaign{}
	adCampaigns, err := client.AdCampaignClient{}.GetAll()
	if err != nil {
		err := fmt.Errorf("error getting ad campaigns")
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}
	/*for _, a := range adCampaignsP {
		adCampaigns = append(adCampaigns, *a)
	}*/

	return c.JSON(http.StatusOK, adCampaigns)
}