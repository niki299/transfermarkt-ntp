package model

type SearchData struct {
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Club      string `json:"club"`
	Position  string `json:"position"`
	PriceFrom int    `json:"priceFrom"`
	PriceTo   int    `json:"priceTo"`
	//eventualno prosirenje sa jos search parametara
	// za pocetak cena i pozicija
}
