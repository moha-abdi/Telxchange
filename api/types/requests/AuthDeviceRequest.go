package requests

import "github.com/moha-abdi/telxchange/config"

type AuthDeviceRequest struct {
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
			ActivationTypeId     string              `json:"activationTypeId"`
			OtpCode              string              `json:"otpCode"`
		} `json:"requestAttributes"`
		ServiceCode string `json:"serviceCode"`
		ServiceName string `json:"serviceName"`
	} `json:"serviceInfo"`
}

func NewAuthDeviceRequest() *AuthDeviceRequest {
	authDeviceRequest := &AuthDeviceRequest{
		BaseRequest: *NewBaseRequest(),
	}
	authDeviceRequest.ServiceInfo.RequestAttributes.VerificationCodeType = "18"
	authDeviceRequest.ServiceInfo.ServiceName = "authenticateDevice"
	authDeviceRequest.ServiceInfo.RequestAttributes.UserType = "CUSTOMER"
	authDeviceRequest.ServiceInfo.RequestAttributes.ActivationTypeId = "18"
	if config.DefaultChannel == config.MobileApp {
		authDeviceRequest.ServiceInfo.RequestAttributes.AppVersion = "8.2.2"
		authDeviceRequest.ServiceInfo.RequestAttributes.DeviceOS = "iOS"
	}

	return authDeviceRequest
}
