package requests

import (
	"github.com/moha-abdi/telxchange/config"
)

type ExchangeRateRequest struct {
	BaseRequest
	ServiceInfo struct {
		RequestAttributes struct {
			DeviceId             string              `json:"deviceId"`
			DeviceIdType         string              `json:"deviceIdType"`
			LocationInfo         LocationInformation `json:"locationInformation"`
			ReceiverLocalionInfo LocationInformation `json:"recieverlocationInformation"`
			MobileNo             string              `json:"mobileNo"`
			PartnerUID           string              `json:"partnerUID"`
			SourceCurrencyCode   string              `json:"sourceCurrencyCode"`
			TargetCurrencyCode   string              `json:"targetCurrencyCode"`
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

func NewExchangeRateRequest() *ExchangeRateRequest {
	exchangeRequest := &ExchangeRateRequest{
		BaseRequest: *NewBaseRequest(),
	}
	exchangeRequest.ServiceInfo.RequestAttributes.DeviceId = "TLX"
	exchangeRequest.ServiceInfo.ServiceCode = "0261"
	exchangeRequest.ServiceInfo.ServiceName = "ExchangeFromMerchantService"
	exchangeRequest.ServiceInfo.RequestAttributes.ServiceType = "MMT"
	exchangeRequest.ServiceInfo.RequestAttributes.UserType = "CUSTOMER"
	exchangeRequest.ServiceInfo.RequestAttributes.ChannelName = string(config.DefaultChannel)
	if config.DefaultChannel == config.MobileApp {
		exchangeRequest.ServiceInfo.RequestAttributes.AppVersion = "8.2.2"
		exchangeRequest.ServiceInfo.RequestAttributes.DeviceOS = "iOS"
	}

	return exchangeRequest
}
