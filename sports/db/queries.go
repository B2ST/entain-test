package db

const (
	eventsList = "list"
)

func getEventsQueries() map[string]string {
	return map[string]string{
		eventsList: `
			SELECT 
				id, 
				sport, 
				team1,
				team2, 
				visible, 
				advertised_start_time,
				CASE 
					WHEN advertised_start_time < CURRENT_TIMESTAMP THEN 'CLOSED' ELSE 'OPEN'
				END AS status
			FROM events
		`,
	}
}
