package responses

type VerificationCodeResponse struct {
	BaseResponse
	ServiceInfo struct {
		ResponseAttributes struct {
			ResultCode   string `json:"resultCode"`
			ReplyMessage string `json:"replyMessage"`
			SessionId    string `json:"sessionId"`
		} `json:"responseAttributes"`
	} `json:"serviceInfo"`
	SubscriberId string `json:"subscriberId"`
}
