package db

import (
	"encoding/json"
	"os"
)

func LoadDBClients() error {
	rawClients, err := os.ReadFile("./dbs/clients.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawClients, &DBClients)
	if err != nil {
		return err
	}

	return nil
}