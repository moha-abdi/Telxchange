package responses

type BalanceQuertResponse struct {
	BaseResponse
	ServiceInfo struct {
		ResponseAttributes struct {
			ResultCode     string    `json:"resultCode"`
			AccountId      int       `json:"accountId"`
			AccountTitle   string    `json:"accountTitle"`
			CurrencySymbol string    `json:"currencySymbol"`
			CurrentBalance string    `json:"currentBalance"`
			ShadowDebit    int       `json:"shadowDebit"`
			ReplyMessage   string    `json:"replyMessage"`
			Accounts       []Account `json:"accounts"`
		} `json:"responseAttributes"`
	} `json:"serviceInfo"`
}

type Account struct {
	AccountId      int    `json:"accountId"`
	AccountTitle   string `json:"accountTitle"`
	CurrencySymbol string `json:"currencySymbol"`
	CurrentBalance string `json:"currentBalance"`
	ShadowDebit    int    `json:"shadowDebit"`
	ReplyMessage   string `json:"replyMessage,omitempty"`
}
