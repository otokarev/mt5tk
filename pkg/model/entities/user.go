package entities

type User struct {
	Login                  string `json:"Login"`
	Group                  string `json:"Group"`
	CertSerialNumber       string `json:"CertSerialNumber"`
	Rights                 string `json:"Rights"`
	MQID                   string `json:"MQID"`
	Registration           string `json:"Registration"`
	LastAccess             string `json:"LastAccess"`
	LastPassChange         string `json:"LastPassChange"`
	LastIP                 string `json:"LastIP"`
	Name                   string `json:"Name"`
	FirstName              string `json:"FirstName"`
	LastName               string `json:"LastName"`
	MiddleName             string `json:"MiddleName"`
	Company                string `json:"Company"`
	Account                string `json:"Account"`
	Country                string `json:"Country"`
	Language               string `json:"Language"`
	ClientID               string `json:"ClientID"`
	City                   string `json:"City"`
	State                  string `json:"State"`
	ZipCode                string `json:"ZipCode"`
	Address                string `json:"Address"`
	Phone                  string `json:"Phone"`
	Email                  string `json:"Email"`
	ID                     string `json:"ID"`
	Status                 string `json:"Status"`
	Comment                string `json:"Comment"`
	Color                  string `json:"Color"`
	PhonePassword          string `json:"PhonePassword"`
	Leverage               string `json:"Leverage"`
	Agent                  string `json:"Agent"`
	CurrencyDigits         string `json:"CurrencyDigits"`
	Balance                string `json:"Balance"`
	Credit                 string `json:"Credit"`
	InterestRate           string `json:"InterestRate"`
	CommissionDaily        string `json:"CommissionDaily"`
	CommissionMonthly      string `json:"CommissionMonthly"`
	CommissionAgentDaily   string `json:"CommissionAgentDaily"`
	CommissionAgentMonthly string `json:"CommissionAgentMonthly"`
	BalancePrevDay         string `json:"BalancePrevDay"`
	BalancePrevMonth       string `json:"BalancePrevMonth"`
	EquityPrevDay          string `json:"EquityPrevDay"`
	EquityPrevMonth        string `json:"EquityPrevMonth"`
	TradeAccounts          string `json:"TradeAccounts"`
	ApiData                []struct {
		AppID       string `json:"AppID"`
		ID          string `json:"ID"`
		ValueInt    string `json:"ValueInt"`
		ValueUInt   string `json:"ValueUInt"`
		ValueDouble string `json:"ValueDouble"`
	} `json:"ApiData"`
	LeadCampaign string `json:"LeadCampaign"`
	LeadSource   string `json:"LeadSource"`
}
