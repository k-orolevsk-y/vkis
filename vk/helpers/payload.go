package helpers

import "encoding/json"

func GetPayload(payload json.RawMessage, object any) error {
	if err := json.Unmarshal(payload, &object); err != nil {
		return err
	}

	return nil
}
