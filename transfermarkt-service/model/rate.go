package model

type Rate struct {
	Id        string `json:"id"`
	OfferId   string `json:"offerId"`
	RateValue int    `json:"rateValue"`
}
