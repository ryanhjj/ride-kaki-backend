package models

type Tada struct {
	IsStudentFleet bool        `json:"is_student_fleet"`
	Locations      []Locations `json:"locations"`
}

type Locations struct {
	Address       string  `json:"address"`
	GooglePlaceId string  `json:"google_place_id"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Name          string  `json:"name"`
}
