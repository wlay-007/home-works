package file

import (
	"encoding/json"
	"os"
)

func SaveVerification(rec any) error {
	f, err := os.Create("users.json")
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(rec)
}
