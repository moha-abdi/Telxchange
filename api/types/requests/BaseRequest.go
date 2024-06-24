package requests

import (
	"time"

	"github.com/google/uuid"
	"github.com/moha-abdi/telxchange/config"
)

type BaseRequest struct {
	SchemaVersion string `json:"schemaVersion"`
	RequestID     string `json:"requestId"`
	Timestamp     int64  `json:"timestamp"`
	Channel       string `json:"channel"`
	SystemInfo    struct {
		SystemID     string `json:"systemId"`
		SystemSecret string `json:"systemSecret"`
	} `json:"systemInfo"`
}

func NewBaseRequest() *BaseRequest {
	return &BaseRequest{
		SchemaVersion: "1.0",
		RequestID:     uuid.NewString(),
		Timestamp:     time.Now().Unix(),
		Channel:       string(config.DefaultChannel),
	}
}
