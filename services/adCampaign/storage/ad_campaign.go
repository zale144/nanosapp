package storage

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"log"
	proto "github.com/zale144/nanosapp/services/adCampaign/proto"
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

// GetAll fetches a list of ad campaigns
func (m *AdCampaignStorage) GetAll() ([]*proto.AdCampaign, error) {
	var adCampaigns []*proto.AdCampaign
	err := DB.C(COLLECTION).Find(bson.M{}).All(&adCampaigns)
	return adCampaigns, err
}

// Insert adds an ad campaign into the database
func (m *AdCampaignStorage) Insert(adCampaign proto.AdCampaign) error {
	err := DB.C(COLLECTION).Insert(&adCampaign)
	return err
}

// DeleteAll removes all ad campaigns
func (m *AdCampaignStorage) DeleteAll() error {
	return DB.C(COLLECTION).DropCollection()
}