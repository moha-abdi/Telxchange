package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

// NOTE: If more timestamp-related functionality is added in the future,
// consider moving FlexibleTimestamp to a separate package for better
// separation of concerns and maintainability. Good luck :)

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
