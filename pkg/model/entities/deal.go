package entities

type Deal struct {
	Deal            string `json:"Deal"`
	ExternalID      string `json:"ExternalID"`
	Login           string `json:"Login"`
	Dealer          string `json:"Dealer"`
	Order           string `json:"Order"`
	Action          string `json:"Action"`
	Entry           string `json:"Entry"`
	Reason          string `json:"Reason"`
	Digits          string `json:"Digits"`
	DigitsCurrency  string `json:"DigitsCurrency"`
	ContractSize    string `json:"ContractSize"`
	Time            string `json:"Time"`
	TimeMsc         string `json:"TimeMsc"`
	Symbol          string `json:"Symbol"`
	Price           string `json:"Price"`
	Volume          string `json:"Volume"`
	VolumeExt       string `json:"VolumeExt"`
	Profit          string `json:"Profit"`
	Storage         string `json:"Storage"`
	Commission      string `json:"Commission"`
	Fee             string `json:"Fee"`
	RateProfit      string `json:"RateProfit"`
	RateMargin      string `json:"RateMargin"`
	ExpertID        string `json:"ExpertID"`
	PositionID      string `json:"PositionID"`
	Comment         string `json:"Comment"`
	ProfitRaw       string `json:"ProfitRaw"`
	PricePosition   string `json:"PricePosition"`
	PriceSL         string `json:"PriceSL"`
	PriceTP         string `json:"PriceTP"`
	VolumeClosed    string `json:"VolumeClosed"`
	VolumeClosedExt string `json:"VolumeClosedExt"`
	TickValue       string `json:"TickValue"`
	TickSize        string `json:"TickSize"`
	Flags           string `json:"Flags"`
	Gateway         string `json:"Gateway"`
	PriceGateway    string `json:"PriceGateway"`
	ModifyFlags     string `json:"ModifyFlags"`
	Value           string `json:"Value"`
	ApiData         []struct {
		AppID       string `json:"AppID"`
		ID          string `json:"ID"`
		ValueInt    string `json:"ValueInt"`
		ValueUInt   string `json:"ValueUInt"`
		ValueDouble string `json:"ValueDouble"`
	} `json:"ApiData"`
	MarketBid  string `json:"MarketBid"`
	MarketAsk  string `json:"MarketAsk"`
	MarketLast string `json:"MarketLast"`
	OrigTime   string `json:"OrigTime"`
}
