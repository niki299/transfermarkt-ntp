package model

import "time"

type Offer struct {
	Id       string    `json:"id"`
	Player   Player    `json:"player"`
	Date     time.Time `json:"date"` // publish date
	Rates    []Rate    `json:"rates"`
	Comments []Comment `json:"comments"`
}
