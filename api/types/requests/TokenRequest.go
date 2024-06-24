package requests

type TokenRequest struct {
	BaseRequest
	ServiceInfo struct {
		RequestAttributes struct {
			DeviceId             string              `json:"deviceId"`
			DeviceIdType         string              `json:"deviceIdType"`
			LocationInfo         LocationInformation `json:"locationInformation"`
			ReceiverLocalionInfo LocationInformation `json:"recieverlocationInformation"`
			Username             string              `json:"username"`
		} `json:"requestAttributes"`
		ServiceCode string `json:"serviceCode"`
		ServiceName string `json:"serviceName"`
	} `json:"serviceInfo"`
}

func NewTokenRequest() *TokenRequest {
	tokenRequest := &TokenRequest{
		BaseRequest: *NewBaseRequest(),
	}
	tokenRequest.ServiceInfo.ServiceCode = "0101"
	tokenRequest.ServiceInfo.ServiceName = "CustomerLogin"
	return tokenRequest
}
