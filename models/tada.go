package models

type Tada struct {
	IsStudentFleet bool      `json:"is_student_fleet" bson:"is_student_fleet" example:"false"`
	Locations      []Address `json:"locations" bson:"locations""`
}
