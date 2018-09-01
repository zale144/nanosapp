package model

type AdCampaign struct {
	ID          int64       `bson:"id"`
	Name        string    `bson:"name"`
	Goal        string    `bson:"goal"`
	TotalBudget int64       `bson:"total_budget"`
	Status      string	  `bson:"status"`
	Platforms 	*Platforms `bson:"platforms"`
}

type Platforms struct {
	Facebook 	*Platform `bson:"facebook"`
	Instagram 	*Platform `bson:"instagram"`
	Google 		*Platform `bson:"google"`
}

type Platform struct {
	Status          string 		   `bson:"status"`
	TotalBudget     int64    	   `bson:"total_budget"`
	RemainingBudget int64    	   `bson:"remaining_budget"`
	StartDate       int64  		   `bson:"start_date"`
	EndDate         int64  		   `bson:"end_date"`
	TargetAudiance  TargetAudiance `bson:"target_audiance"`
	Creatives 		Creatives	   `bson:"creatives"`
	Insights 		Insights	   `bson:"insights"`
}

type TargetAudiance struct {
	Languages []string `bson:"languages"`
	Genders   []string `bson:"genders"`
	AgeRange  []int64  `bson:"age_range"`
	Locations []string `bson:"locations"`
	KeyWords  []string `bson:"KeyWords"`
	Interests []string `bson:"interests"`
}

type Creatives struct {
	Header      string `bson:"header"`
	Header1     string `bson:"header_1"`
	Header2     string `bson:"header_2"`
	Description string `bson:"description"`
	URL         string `bson:"url"`
	Image       string `bson:"image"`
}

type Insights struct {
	Impressions      int64     `bson:"impressions"`
	Clicks           int64     `bson:"clicks"`
	WebsiteVisits    int64     `bson:"website_visits"`
	CostPerClick     float64 `bson:"cost_per_click"`
	ClickThroughRate float64 `bson:"click_through_rate"`
	AdvancedKpi1     float64 `bson:"advanced_kpi_1"`
	AdvancedKpi2     float64 `bson:"advanced_kpi_2"`
	NanosScore       float64 `bson:"nanos_score"`
}