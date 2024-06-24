package requests

import "github.com/moha-abdi/telxchange/config"

type VerificationCodeRequest struct {
	BaseRequest
	ServiceInfo struct {
		RequestAttributes struct {
			DeviceId             string              `json:"deviceId"`
			DeviceIdType         string              `json:"deviceIdType"`
			LocationInfo         LocationInformation `json:"locationInformation"`
			ReceiverLocalionInfo LocationInformation `json:"recieverlocationInformation"`
			MobileNumber         string              `json:"mobileNumber"`
			VerificationCodeType string              `json:"verificationCodeType"`
			UserType             string              `json:"userType"`
			AppVersion           string              `json:"appVersion,omitempty"`
			DeviceOS             string              `json:"deviceOS,omitempty"`
		} `json:"requestAttributes"`
		ServiceCode string `json:"serviceCode"`
		ServiceName string `json:"serviceName"`
	} `json:"serviceInfo"`
}

func NewVerificationCodeRequest() *VerificationCodeRequest {
	verificationCodeRequest := &VerificationCodeRequest{
		BaseRequest: *NewBaseRequest(),
	}
	verificationCodeRequest.ServiceInfo.RequestAttributes.VerificationCodeType = "18"
	verificationCodeRequest.BaseRequest.Channel = string(config.DefaultChannel)
	verificationCodeRequest.ServiceInfo.ServiceCode = "0102"
	verificationCodeRequest.ServiceInfo.ServiceName = "requestVerificationCode"
	verificationCodeRequest.ServiceInfo.RequestAttributes.UserType = "CUSTOMER"
	if config.DefaultChannel == config.MobileApp {
		verificationCodeRequest.ServiceInfo.RequestAttributes.AppVersion = "8.2.2"
		verificationCodeRequest.ServiceInfo.RequestAttributes.DeviceOS = "iOS"
	}

	return verificationCodeRequest
}
