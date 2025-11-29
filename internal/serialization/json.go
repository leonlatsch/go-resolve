package serialization

import (
	"encoding/json"
)

func ToJson(data any) (string, error) {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func FromJson[T any](in string, out *T) error {
	if err := json.Unmarshal([]byte(in), &out); err != nil {
		return err
	}

	return nil
}
