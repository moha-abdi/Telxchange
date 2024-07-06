package requests

import "github.com/moha-abdi/telxchange/config"

type BalanceQueryRequest struct {
	BaseRequest
	ServiceInfo struct {
		RequestAttributes struct {
			DeviceId             string              `json:"deviceId"`
			DeviceIdType         string              `json:"deviceIdType"`
			LocationInfo         LocationInformation `json:"locationInformation"`
			ReceiverLocalionInfo LocationInformation `json:"recieverlocationInformation"`
			MobileNo             string              `json:"mobileNo"`
			AccountId            string              `json:"accountId"`
			UserType             string              `json:"userType"`
			SessionId            string              `json:"sessionId"`
			ChannelName          string              `json:"channelName"`
			ServiceType          string              `json:"serviceType"`
			AppVersion           string              `json:"appVersion,omitempty"`
			DeviceOS             string              `json:"deviceOS,omitempty"`
		} `json:"requestAttributes"`
		ServiceCode string `json:"serviceCode"`
		ServiceName string `json:"serviceName"`
	} `json:"serviceInfo"`
}

func NewBalanceQueryRequest() *BalanceQueryRequest {
	balanceRequest := &BalanceQueryRequest{
		BaseRequest: *NewBaseRequest(),
	}
	balanceRequest.ServiceInfo.RequestAttributes.DeviceId = "TLX"
	balanceRequest.ServiceInfo.ServiceCode = "0004"
	balanceRequest.ServiceInfo.ServiceName = "BalanceQuery"
	balanceRequest.ServiceInfo.RequestAttributes.ServiceType = "MMT"
	balanceRequest.ServiceInfo.RequestAttributes.UserType = "CUSTOMER"
	balanceRequest.ServiceInfo.RequestAttributes.ChannelName = string(config.DefaultChannel)
	if config.DefaultChannel == config.MobileApp {
		balanceRequest.ServiceInfo.RequestAttributes.AppVersion = "8.2.2"
		balanceRequest.ServiceInfo.RequestAttributes.DeviceOS = "iOS"
	}

	return balanceRequest
}
