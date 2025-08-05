package db

func QueryDBSessions(filterURL string, filterUUID string) (int, int, DBSession, bool) {
	indexDBClient, dbClient, ok := QueryDBClients(filterURL)
	if !ok {
		return 0, 0, DBSession{}, false
	}

	for indexDBSession, dbSession := range dbClient.Sessions {
		if dbSession.UUID == filterUUID {
			return indexDBClient, indexDBSession, dbSession, true
		}
	}

	return 0, 0, DBSession{}, false
}

func QueryDBClients(filterURL string) (int, DBClient, bool) {
	for indexDBClient, dbClient := range DBClients {
		if dbClient.URL == filterURL {
			return indexDBClient, dbClient, true
		}
	}

	return 0, DBClient{}, false
}