package models

type RideKaki struct {
	StartingLatitude  string `json:"starting_latitude" bson:"starting_latitude" example:"1.2962727"`
	StartingLongitude string `json:"starting_longitude" bson:"starting_longitude" example:"103.8501578"`
	EndingLatitude    string `json:"ending_latitude" bson:"ending_latitude" example:"1.3005317"`
	EndingLongitude   string `json:"ending_longitude" bson:"ending_longitude" example:"103.8452356"`
}
