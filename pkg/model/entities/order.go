package entities

type Order struct {
	Order            string `json:"Order"`
	ExternalID       string `json:"ExternalID"`
	Login            string `json:"Login"`
	Dealer           string `json:"Dealer"`
	Symbol           string `json:"Symbol"`
	Digits           string `json:"Digits"`
	DigitsCurrency   string `json:"DigitsCurrency"`
	ContractSize     string `json:"ContractSize"`
	State            string `json:"State"`
	Reason           string `json:"Reason"`
	TimeSetup        string `json:"TimeSetup"`
	TimeExpiration   string `json:"TimeExpiration"`
	TimeDone         string `json:"TimeDone"`
	TimeSetupMsc     string `json:"TimeSetupMsc"`
	TimeDoneMsc      string `json:"TimeDoneMsc"`
	ModifyFlags      string `json:"ModifyFlags"`
	Type             string `json:"Type"`
	TypeFill         string `json:"TypeFill"`
	TypeTime         string `json:"TypeTime"`
	PriceOrder       string `json:"PriceOrder"`
	PriceTrigger     string `json:"PriceTrigger"`
	PriceCurrent     string `json:"PriceCurrent"`
	PriceSL          string `json:"PriceSL"`
	PriceTP          string `json:"PriceTP"`
	VolumeInitial    string `json:"VolumeInitial"`
	VolumeInitialExt string `json:"VolumeInitialExt"`
	VolumeCurrent    string `json:"VolumeCurrent"`
	VolumeCurrentExt string `json:"VolumeCurrentExt"`
	ExpertID         string `json:"ExpertID"`
	PositionID       string `json:"PositionID"`
	PositionByID     string `json:"PositionByID"`
	Comment          string `json:"Comment"`
	RateMargin       string `json:"RateMargin"`
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
	OrigTimeSetup      string `json:"OrigTimeSetup"`
	OrigTimeExpiration string `json:"OrigTimeExpiration"`
	OrigTimeDone       string `json:"OrigTimeDone"`
}
