package entities

type Position struct {
	Position         string `json:"Position"`
	ExternalID       string `json:"ExternalID"`
	Login            string `json:"Login"`
	Dealer           string `json:"Dealer"`
	Symbol           string `json:"Symbol"`
	Action           string `json:"Action"`
	Digits           string `json:"Digits"`
	DigitsCurrency   string `json:"DigitsCurrency"`
	Reason           string `json:"Reason"`
	ContractSize     string `json:"ContractSize"`
	TimeCreate       int    `json:"TimeCreate"`
	TimeUpdate       int    `json:"TimeUpdate"`
	TimeCreateMsc    string `json:"TimeCreateMsc"`
	TimeUpdateMsc    string `json:"TimeUpdateMsc"`
	ModifyFlags      string `json:"ModifyFlags"`
	PriceOpen        string `json:"PriceOpen"`
	PriceCurrent     string `json:"PriceCurrent"`
	PriceSL          string `json:"PriceSL"`
	PriceTP          string `json:"PriceTP"`
	Volume           string `json:"Volume"`
	VolumeExt        string `json:"VolumeExt"`
	Profit           string `json:"Profit"`
	Storage          string `json:"Storage"`
	RateProfit       string `json:"RateProfit"`
	RateMargin       string `json:"RateMargin"`
	ExpertID         string `json:"ExpertID"`
	ExpertPositionID string `json:"ExpertPositionID"`
	Comment          string `json:"Comment"`
	ActivationMode   string `json:"ActivationMode"`
	ActivationTime   string `json:"ActivationTime"`
	ActivationPrice  string `json:"ActivationPrice"`
	ActivationFlags  string `json:"ActivationFlags"`
	ApiData          []struct {
		AppID       string `json:"AppID"`
		ID          string `json:"ID"`
		ValueInt    string `json:"ValueInt"`
		ValueUInt   string `json:"ValueUInt"`
		ValueDouble string `json:"ValueDouble"`
	} `json:"ApiData"`
	OrigTimeCreate string `json:"OrigTimeCreate"`
	OrigTimeUpdate string `json:"OrigTimeUpdate"`
}
