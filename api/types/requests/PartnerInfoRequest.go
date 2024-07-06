package requests

import (
	"github.com/moha-abdi/telxchange/config"
)

type PartnerInfoRequest struct {
	BaseRequest
	ServiceInfo struct {
		RequestAttributes struct {
			DeviceId             string              `json:"deviceId"`
			DeviceIdType         string              `json:"deviceIdType"`
			LocationInfo         LocationInformation `json:"locationInformation"`
			ReceiverLocalionInfo LocationInformation `json:"recieverlocationInformation"`
			MobileNo             string              `json:"mobileNo"`
			SubPartnerUID        string              `json:"subPartnerUID"`
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

func NewPartnerInfoRequest() *PartnerInfoRequest {
	partnerRequest := &PartnerInfoRequest{
		BaseRequest: *NewBaseRequest(),
	}
	partnerRequest.ServiceInfo.RequestAttributes.DeviceId = "TLX"
	partnerRequest.ServiceInfo.ServiceCode = "999"
	partnerRequest.ServiceInfo.ServiceName = "GetPartnerInfoByUID"
	partnerRequest.ServiceInfo.RequestAttributes.ServiceType = "MMT"
	partnerRequest.ServiceInfo.RequestAttributes.UserType = "CUSTOMER"
	partnerRequest.ServiceInfo.RequestAttributes.ChannelName = string(config.DefaultChannel)
	if config.DefaultChannel == config.MobileApp {
		partnerRequest.ServiceInfo.RequestAttributes.AppVersion = "8.2.2"
		partnerRequest.ServiceInfo.RequestAttributes.DeviceOS = "iOS"
	}

	return partnerRequest
}
