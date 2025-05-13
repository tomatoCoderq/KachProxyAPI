package models

type Play struct {
	Id string `json:"id"`
	Scene string `json:"scene"`
	Author string `json:"author"`
	Name   string `json:"name"`
	Genre string `json:"genre"`
	AgeRating string `json:"age_rating"`
	Month int `json:"month"`
	Day   int `json:"day"`
}
