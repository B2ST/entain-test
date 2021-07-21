package db

import (
	"time"

	"syreclabs.com/go/faker"
)

func (r *sportsRepo) seed() error {
	statement, err := r.db.Prepare(`CREATE TABLE IF NOT EXISTS events (id INTEGER PRIMARY KEY, sport TEXT, team1 TEXT, team2 TEXT, visible INTEGER, advertised_start_time DATETIME)`)
	if err == nil {
		_, err = statement.Exec()
	}

	for i := 1; i <= 100; i++ {
		statement, err = r.db.Prepare(`INSERT OR IGNORE INTO events(id, sport, team1, team2, visible, advertised_start_time) VALUES (?,?,?,?,?,?)`)
		if err == nil {
			_, err = statement.Exec(
				i,
				faker.Team().Name(),
				faker.Team().Creature(),
				faker.Team().Creature(),
				faker.Number().Between(0, 1),
				faker.Time().Between(time.Now().AddDate(0, 0, -1), time.Now().AddDate(0, 0, 2)).Format(time.RFC3339),
			)
		}
	}

	return err
}
