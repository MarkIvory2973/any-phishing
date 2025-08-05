package db

import (
	"encoding/json"
	"os"
)

func UpdateDBClients() error {
	rawClients, err :=json.Marshal(DBClients)
	if err != nil {
		return err
	}

	err = os.WriteFile("./dbs/clients.json", rawClients, 0775)
	if err != nil {
		return err
	}

	return nil
}