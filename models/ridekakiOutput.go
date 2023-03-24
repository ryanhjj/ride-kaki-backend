package models

type RideKakiOutput struct {
	Gojek4EconomyPrice float64 `json:"gojek_4_economy_price" bson:"gojek_4_economy_price" example:"20.50"`
	Gojek6EconomyPrice float64 `json:"gojek_6_economy_price" bson:"gojek_6_economy_price" example:"20.50"`
	Gojek4PremiumPrice float64 `json:"gojek_4_premium_price" bson:"gojek_4_premium_price" example:"20.50"`
	Tada4EconomyPrice  float64 `json:"tada_4_economy_price" bson:"tada_4_economy_price" example:"20.50"`
	Tada6EconomyPrice  float64 `json:"tada_6_economy_price" bson:"tada_6_economy_price" example:"20.50"`
	Tada4PremiumPrice  float64 `json:"tada_4_premium_price" bson:"tada_4_premium_price" example:"20.50"`
}
