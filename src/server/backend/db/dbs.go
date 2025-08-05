package db

type DBSession struct {
	UUID string `json:"uuid"`
	IP string `json:"ip"`
	Time int `json:"time"`
	Inputs []string `json:"inputs"`
}

type DBClient struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Sessions []DBSession `json:"sessions"`
	URL string `json:"url"`
}

var DBClients []DBClient