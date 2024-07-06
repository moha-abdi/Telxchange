package responses

import (
	"encoding/json"
	"fmt"
	"time"
)

type FlexibleTimestamp time.Time

func (ft *FlexibleTimestamp) UnmarshalJSON(data []byte) error {
	var rawValue interface{}
	if err := json.Unmarshal(data, &rawValue); err != nil {
		return err
	}

	switch v := rawValue.(type) {
	case string:
		t, err := time.Parse(time.RFC3339, v)
		if err != nil {
			return err
		}
		*ft = FlexibleTimestamp(t)
	case float64:
		*ft = FlexibleTimestamp(time.Unix(int64(v), 0))
	default:
		return fmt.Errorf("invalid type for timestamp: %T", v)
	}

	return nil
}

type BaseResponse struct {
	SchemaVersion string            `json:"schemaVersion"`
	RequestID     string            `json:"requestId"`
	Timestamp     FlexibleTimestamp `json:"timestamp"`
	Channel       string            `json:"channel"`
	SystemInfo    struct {
		SystemID     string `json:"systemId"`
		SystemSecret string `json:"systemSecret"`
	} `json:"systemInfo"`
}

type BaseResponseInt struct {
	SchemaVersion string `json:"schemaVersion"`
	RequestID     string `json:"requestId"`
	Timestamp     int64  `json:"timestamp"`
	Channel       string `json:"channel"`
	SystemInfo    struct {
		SystemID     string `json:"systemId"`
		SystemSecret string `json:"systemSecret"`
	} `json:"systemInfo"`
}
