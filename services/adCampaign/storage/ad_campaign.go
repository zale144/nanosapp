package storage

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"log"
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

// Establish a connection to database
func (m *AdCampaignStorage) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	DB = session.DB(m.Database)
}

// Find list of ad campaigns
func (m *AdCampaignStorage) GetAll() ([]model.AdCampaign, error) {
	var movies []model.AdCampaign
	err :=DB.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

// Insert an ad campaign into database
func (m *AdCampaignStorage) Insert(adCampaign model.AdCampaign) error {
	err := DB.C(COLLECTION).Insert(&adCampaign)
	return err
}