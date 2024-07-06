package responses

type LoginResponse struct {
	BaseResponse
	ServiceInfo struct {
		ResponseAttributes struct {
			Username           string        `json:"username"`
			UserType           string        `json:"userType"`
			ResultCode         string        `json:"resultCode"`
			ReplyMessage       string        `json:"replyMessage"`
			SessionId          string        `json:"sessionId"`
			DefaultAccount     string        `json:"defaultAccount"`
			AccountInformation []AccountInfo `json:"accountInformation"`
		} `json:"responseAttributes"`
	} `json:"serviceInfo"`
}

type AccountInfo struct {
	AccountId        string `json:"accountId"`
	AccountTitle     string `json:"accountTitle"`
	AccountNumber    string `json:"accountNumber"`
	AccountType      string `json:"accountType"`
	AccountCurrency  string `json:"accountCurrency"`
	CurrencyName     string `json:"currencyName"`
	CurrencySymbol   string `json:"currencySymbol"`
	CurrencyId       string `json:"currencyId"`
	IsDefaultAccount bool   `json:"isDefaultAccount"`
}
