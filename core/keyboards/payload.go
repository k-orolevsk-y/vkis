package keyboards

import "encoding/json"

type Payload struct {
	Command string          `json:"command"`
	Other   json.RawMessage `json:"other,omitempty"`
}
