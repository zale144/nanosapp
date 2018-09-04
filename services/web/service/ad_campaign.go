package service

import (
	"github.com/labstack/echo"
	"github.com/zale144/nanosapp/services/web/client"
	"net/http"
	"fmt"
	"log"
	"encoding/json"
)

type AdCampaignService struct {}

// GetAll handles requests to get all ad campaigns
func (as AdCampaignService) GetAll(c echo.Context) error {

	adCampaigns, err := client.AdCampaignClient{}.GetAll()
	if err != nil {
		log.Println(err)
		err := fmt.Errorf("error getting ad campaigns")
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}
	j, err := json.Marshal(adCampaigns)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(j))
	return c.JSON(http.StatusOK, adCampaigns)
}