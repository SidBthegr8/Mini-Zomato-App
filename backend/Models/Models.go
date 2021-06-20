package Models
type Restaurant struct {
	Name        string `json:"name"`
	Rating      string `json:"rating"`
	Cuisine     string `json:"cuisine"`
	Address     string `json:"address"`
	OpeningTime string `json:"openingTime"`
	ClosingTime string `json:"closingTime"`
	CostForTwo  string `json:"costForTwo"`

}