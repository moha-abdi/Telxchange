package requests

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/moha-abdi/telxchange/config"
)

type BaseRequest struct {
	SchemaVersion string `json:"schemaVersion"`
	RequestID     string `json:"requestId"`
	Timestamp     string `json:"timestamp"`
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
		Timestamp:     strconv.FormatInt(time.Now().UnixMilli(), 10),
		Channel:       string(config.DefaultChannel),
	}
}
