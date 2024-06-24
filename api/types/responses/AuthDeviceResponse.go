package responses

type AuthDeviceResponse struct {
	BaseResponse
	ServiceInfo struct {
		ResponseAttributes struct {
			ResultCode   string `json:"resultCode"`
			ReplyMessage string `json:"replyMessage"`
		} `json:"responseAttributes"`
	} `json:"serviceInfo"`
}
