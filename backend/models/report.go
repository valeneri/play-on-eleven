package models

type Report struct {
	Id    string `json:"id" bson:"_id,omitempty"`
	Title string `json:"title" bson:"title"`
}
