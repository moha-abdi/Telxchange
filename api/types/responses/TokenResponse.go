package responses

type TokenResponse struct {
	BaseResponse
	ServiceInfo struct {
		ResponseAttributes struct {
			Token                   string `json:"token"`
			EnablePayloadEncryption bool   `json:"enablePayloadEncryption"`
			ResultCode              int    `json:"resultCode"`
			ReplyMessage            string `json:"replyMessage"`
		} `json:"responseAttributes"`
	} `json:"serviceInfo"`
}
