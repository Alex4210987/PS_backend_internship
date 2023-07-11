package task4

import (
	"database/sql"
	"log"
)

func InsertRoute(db *sql.DB, route Route) {
	insertSQL := "INSERT INTO routes (start_point, end_point) VALUES (?, ?)"

	_, err := db.Exec(insertSQL, route.StartPoint, route.EndPoint)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertPreference(db *sql.DB, preference Preference) {
	insertSQL := "INSERT INTO preferences (user_id, mode, outputmode, tactics) VALUES (?, ?, ?, ?)"

	_, err := db.Exec(insertSQL, preference.UserID, preference.mode, preference.outputmode, preference.tactics)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertAlias(db *sql.DB, alias Alias) {
	insertSQL := "INSERT INTO aliases (user_id, location, alias) VALUES (?, ?, ?)"

	_, err := db.Exec(insertSQL, alias.UserID, alias.Location, alias.Alias)
	if err != nil {
		log.Fatal(err)
	}
}
