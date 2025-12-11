package file

import (
	"encoding/json"
	"os"
)

func LoadVerification[T any]() (T, error) {
	var rec T

	f, err := os.Open("users.json")
	if err != nil {
		return rec, err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&rec)
	return rec, err
}
