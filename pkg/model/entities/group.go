package entities

type Group struct {
	Group                string        `json:"Group"`
	Server               string        `json:"Server"`
	PermissionsFlags     string        `json:"PermissionsFlags"`
	AuthMode             string        `json:"AuthMode"`
	AuthPasswordMin      string        `json:"AuthPasswordMin"`
	AuthOTPMode          string        `json:"AuthOTPMode"`
	Company              string        `json:"Company"`
	CompanyPage          string        `json:"CompanyPage"`
	CompanyEmail         string        `json:"CompanyEmail"`
	CompanySupportPage   string        `json:"CompanySupportPage"`
	CompanySupportEmail  string        `json:"CompanySupportEmail"`
	CompanyCatalog       string        `json:"CompanyCatalog"`
	CompanyDepositURL    string        `json:"CompanyDepositURL"`
	CompanyWithdrawalURL string        `json:"CompanyWithdrawalURL"`
	Currency             string        `json:"Currency"`
	CurrencyDigits       string        `json:"CurrencyDigits"`
	ReportsMode          string        `json:"ReportsMode"`
	ReportsFlags         string        `json:"ReportsFlags"`
	ReportsEmail         string        `json:"ReportsEmail"`
	NewsMode             string        `json:"NewsMode"`
	NewsCategory         string        `json:"NewsCategory"`
	NewsLangs            []interface{} `json:"NewsLangs"`
	MailMode             string        `json:"MailMode"`
	TradeFlags           string        `json:"TradeFlags"`
	TradeTransferMode    string        `json:"TradeTransferMode"`
	TradeInterestrate    string        `json:"TradeInterestrate"`
	TradeVirtualCredit   string        `json:"TradeVirtualCredit"`
	MarginMode           string        `json:"MarginMode"`
	MarginFlags          string        `json:"MarginFlags"`
	MarginSOMode         string        `json:"MarginSOMode"`
	MarginFreeMode       string        `json:"MarginFreeMode"`
	MarginCall           string        `json:"MarginCall"`
	MarginStopOut        string        `json:"MarginStopOut"`
	MarginFreeProfitMode string        `json:"MarginFreeProfitMode"`
	DemoLeverage         string        `json:"DemoLeverage"`
	DemoDeposit          string        `json:"DemoDeposit"`
	DemoTradesClean      string        `json:"DemoTradesClean"`
	LimitHistory         string        `json:"LimitHistory"`
	LimitOrders          string        `json:"LimitOrders"`
	LimitSymbols         string        `json:"LimitSymbols"`
	LimitPositions       string        `json:"LimitPositions"`
	Commissions          []struct {
		Name             string `json:"Name"`
		Description      string `json:"Description"`
		Path             string `json:"Path"`
		Mode             string `json:"Mode"`
		RangeMode        string `json:"RangeMode"`
		ChargeMode       string `json:"ChargeMode"`
		TurnoverCurrency string `json:"TurnoverCurrency"`
		EntryMode        string `json:"EntryMode"`
		ActionMode       string `json:"ActionMode"`
		Tiers            []struct {
			Mode      string `json:"Mode"`
			Type      string `json:"Type"`
			Value     string `json:"Value"`
			Minimal   string `json:"Minimal"`
			Maximal   string `json:"Maximal"`
			RangeFrom string `json:"RangeFrom"`
			RangeTo   string `json:"RangeTo"`
			Currency  string `json:"Currency"`
		} `json:"Tiers"`
	} `json:"Commissions"`
	Symbols []GroupSymbol ` json:"Symbols"`
}
