package db

const (
	racesList = "list"
)

func getRaceQueries() map[string]string {
	return map[string]string{
		racesList: `
			SELECT 
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				advertised_start_time,
				CASE 
					WHEN advertised_start_time < CURRENT_TIMESTAMP THEN 'CLOSED' ELSE 'OPEN'
				END AS status
			FROM races
		`,
	}
}
