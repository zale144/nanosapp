package storage

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/zale144/nanosapp/services/adCampaign/model"
)

type AdCampaignStorage struct {
	Server   string
	Database string
}

var DB *mgo.Database

const (
	COLLECTION = "ad_campaigns"
)

// Connect establishes a connection to the database
func (m *AdCampaignStorage) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	DB = session.DB(m.Database)
}

// GetAll fetches a list of ad campaigns from the 'ad_campaign' collection
func (m *AdCampaignStorage) GetAll() ([]model.AdCampaign, error) {
	var adCampaigns []model.AdCampaign
	err := DB.C(COLLECTION).Find(bson.M{}).All(&adCampaigns)
	return adCampaigns, err
}

// Insert adds an ad campaign into the 'ad_campaign' collection
func (m *AdCampaignStorage) Insert(adCampaign model.AdCampaign) error {
	err := DB.C(COLLECTION).Insert(&adCampaign)
	return err
}

// DeleteAll removes the 'ad_campaign' collection
func (m *AdCampaignStorage) DeleteAll() error {
	return DB.C(COLLECTION).DropCollection()
}