package radioid

type Contact struct {
	RadioId  int64  `json:"id"`
	Callsign string `json:"callsign"`
	Name     string `json:"fname"`
	Surname  string `json:"surname"`
	City     string `json:"city"`
	Province string `json:"state"`
	Country  string `json:"country"`
}
