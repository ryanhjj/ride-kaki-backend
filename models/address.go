package models

type Address struct {
	Address       string  `json:"address" bson:"address" example:"81 Victoria St, Singapore 188065"`
	GooglePlaceId string  `json:"google_place_id" bson:"google_place_id" example:"ChIJGddBg6MZ2jERACsxW7Ovm_4"`
	Latitude      float64 `json:"latitude" bson:"latitude" example:"1.2962727"`
	Longitude     float64 `json:"longitude" bson:"longitude" example:"103.8501578"`
	Name          string  `json:"name" bson:"name" example:"SMU"`
}
