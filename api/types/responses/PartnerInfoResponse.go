package responses

type PartnerInfo struct {
	Name             string `json:"NAME"`
	ID               int
	SubscriptionID   string `json:"SUBSCRIPTIONID"`
	PartnerID        int    `json:"PARTNERID"`
	PartnerSegmentID int    `json:"PARTNERSEGMENTID"`
	SegmentName      string `json:"SEGMENTNAME"`
	Status           string `json:"STATUS"`
}

type PartnerInfoResponse struct {
	BaseResponse
	ServiceInfo struct {
		ResponseAttributes struct {
			PartnerInfo  PartnerInfo `json:"partnerInfo"`
			SessionId    string      `json:"sessionId"`
			ResultCode   string      `json:"resultCode"`
			ReplyMessage string      `json:"replyMessage"`
		} `json:"responseAttributes"`
	} `json:"serviceInfo"`
}
