package responses

type ExchangeRateResponse struct {
	BaseResponse
	ServiceInfo struct {
		ResponseAttributes struct {
			ExchangeRates []struct {
				Rate             int    `json:"RATE"`
				IsLocked         string `json:"ISLOCKED"`
				MinExchangeLimit int64  `json:"MINEXCHANGELIMIT"`
				MaxExchangeLimit int64  `json:"MAXEXCHANGELIMIT"`
				Name             string `json:"NAME"`
			} `json:"exchangeRates"`
			SessionId    string `json:"sessionId"`
			ResultCode   string `json:"resultCode"`
			ReplyMessage string `json:"replyMessage"`
		} `json:"responseAttributes"`
	} `json:"serviceInfo"`
}
