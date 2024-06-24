package responses

type LoginResponse struct {
	BaseResponse
	ServiceInfo struct {
		ResponseAttributes struct {
			Username     string `json:"username"`
			UserType     string `json:"userType"`
			ResultCode   string `json:"resultCode"`
			ReplyMessage string `json:"replyMessage"`
		} `json:"responseAttributes"`
	} `json:"serviceInfo"`
}
