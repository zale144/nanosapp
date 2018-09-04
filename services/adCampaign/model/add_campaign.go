package model

type AdCampaign struct {
	ID          int64      `json:"id" bson:"id"`
	Name        string     `json:"name" bson:"name"`
	Goal        string     `json:"goal" bson:"goal"`
	TotalBudget int64      `json:"total_budget" bson:"total_budget"`
	Status      string	   `json:"status" bson:"status"`
	Platforms 	*Platforms `json:"platforms" bson:"platforms"`
}

type Platforms struct {
	Facebook 	*Platform `json:"facebook" bson:"facebook"`
	Instagram 	*Platform `json:"instagram" bson:"instagram"`
	Google 		*Platform `json:"google" bson:"google"`
}

type Platform struct {
	Status          string 		    `json:"status" bson:"status"`
	TotalBudget     int64    	    `json:"total_budget" bson:"total_budget"`
	RemainingBudget int64    	    `json:"remaining_budget" bson:"remaining_budget"`
	StartDate       int64  		    `json:"start_date" bson:"start_date"`
	EndDate         int64  		    `json:"end_date" bson:"end_date"`
	TargetAudiance  *TargetAudiance `json:"target_audiance" bson:"target_audiance"`
	Creatives 		*Creatives	    `json:"creatives" bson:"creatives"`
	Insights 		*Insights	    `json:"insights" bson:"insights"`
}

type TargetAudiance struct {
	Languages []string `json:"languages" bson:"languages"`
	Genders   []string `json:"genders" bson:"genders"`
	AgeRange  []int64  `json:"age_range" bson:"age_range"`
	Locations []string `json:"locations" bson:"locations"`
	KeyWords  []string `json:"keyWords" bson:"keyWords"`
	Interests []string `json:"interests" bson:"interests"`
}

type Creatives struct {
	Header      string `json:"header" bson:"header"`
	Header1     string `json:"header_1" bson:"header_1"`
	Header2     string `json:"header_2" bson:"header_2"`
	Description string `json:"description" bson:"description"`
	URL         string `json:"url" bson:"url"`
	Image       string `json:"image" bson:"image"`
}

type Insights struct {
	Impressions      int64   `json:"impressions" bson:"impressions"`
	Clicks           int64   `json:"clicks" bson:"clicks"`
	WebsiteVisits    int64   `json:"website_visits" bson:"website_visits"`
	CostPerClick     float64 `json:"cost_per_click" bson:"cost_per_click"`
	ClickThroughRate float64 `json:"click_through_rate" bson:"click_through_rate"`
	AdvancedKpi1     float64 `json:"advanced_kpi_1" bson:"advanced_kpi_1"`
	AdvancedKpi2     float64 `json:"advanced_kpi_2" bson:"advanced_kpi_2"`
	NanosScore       float64 `json:"nanos_score" bson:"nanos_score"`
}