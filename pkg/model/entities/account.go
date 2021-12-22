package entities

type Account struct {
	Login             string `json:"Login"`
	CurrencyDigits    string `json:"CurrencyDigits"`
	Balance           string `json:"Balance"`
	Credit            string `json:"Credit"`
	Margin            string `json:"Margin"`
	MarginFree        string `json:"MarginFree"`
	MarginLevel       string `json:"MarginLevel"`
	MarginLeverage    string `json:"MarginLeverage"`
	Profit            string `json:"Profit"`
	Storage           string `json:"Storage"`
	Commission        string `json:"Commission"`
	Floating          string `json:"Floating"`
	Equity            string `json:"Equity"`
	SOActivation      string `json:"SOActivation"`
	SOTime            string `json:"SOTime"`
	SOLevel           string `json:"SOLevel"`
	SOEquity          string `json:"SOEquity"`
	SOMargin          string `json:"SOMargin"`
	Assets            string `json:"Assets"`
	Liabilities       string `json:"Liabilities"`
	BlockedCommission string `json:"BlockedCommission"`
	BlockedProfit     string `json:"BlockedProfit"`
	MarginInitial     string `json:"MarginInitial"`
	MarginMaintenance string `json:"MarginMaintenance"`
}
