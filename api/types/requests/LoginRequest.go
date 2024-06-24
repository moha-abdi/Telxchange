package requests

import "github.com/moha-abdi/telxchange/config"

type LoginRequest struct {
	BaseRequest
	ServiceInfo struct {
		RequestAttributes struct {
			DeviceId             string              `json:"deviceId"`
			DeviceIdType         string              `json:"deviceIdType"`
			LocationInfo         LocationInformation `json:"locationInformation"`
			ReceiverLocalionInfo LocationInformation `json:"recieverlocationInformation"`
			Username             string              `json:"username"`
			UserPassword         string              `json:"userPassword"`
			UserType             string              `json:"userType"`
			ChannelName          string              `json:"channelName"`
			ServiceType          string              `json:"serviceType"`
			AppVersion           string              `json:"appVersion,omitempty"`
			DeviceOS             string              `json:"deviceOS,omitempty"`
			AuthMode             string              `json:"authMode,omitempty"`
		} `json:"requestAttributes"`
		ServiceCode string `json:"serviceCode"`
		ServiceName string `json:"serviceName"`
	} `json:"serviceInfo"`
}

func NewLoginRequest() *LoginRequest {
	loginRequest := &LoginRequest{
		BaseRequest: *NewBaseRequest(),
	}
	loginRequest.ServiceInfo.RequestAttributes.DeviceId = "TLX"
	loginRequest.ServiceInfo.ServiceCode = "0101"
	loginRequest.ServiceInfo.ServiceName = "CustomerLogin"
	loginRequest.ServiceInfo.RequestAttributes.UserType = "CUSTOMER"
	loginRequest.ServiceInfo.RequestAttributes.ChannelName = string(config.DefaultChannel)
	if config.DefaultChannel == config.MobileApp {
		loginRequest.ServiceInfo.RequestAttributes.AppVersion = "8.2.2"
		loginRequest.ServiceInfo.RequestAttributes.DeviceOS = "iOS"
		loginRequest.ServiceInfo.RequestAttributes.AuthMode = string(config.DefaultAuthMode)
	}

	return loginRequest
}
