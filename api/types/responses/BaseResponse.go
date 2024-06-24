package responses

type BaseResponse struct {
	SchemaVersion string `json:"schemaVersion"`
	RequestID     string `json:"requestId"`
	Timestamp     int64  `json:"timestamp"`
	Channel       string `json:"channel"`
	SystemInfo    struct {
		SystemID     string `json:"systemId"`
		SystemSecret string `json:"systemSecret"`
	} `json:"systemInfo"`
}
