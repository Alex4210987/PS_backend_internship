package task4

import (
	"database/sql"
	"log"
)

func UpdateRoute(db *sql.DB, id int, route Route) {
	updateSQL := "UPDATE routes SET start_point = ?, end_point = ? WHERE id = ?"

	_, err := db.Exec(updateSQL, route.StartPoint, route.EndPoint, id)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdatePreference(db *sql.DB, id int, preference Preference) {
	updateSQL := "UPDATE preferences SET user_id = ?, mode = ?, outputmode = ?, tactics = ? WHERE id = ?"

	_, err := db.Exec(updateSQL, preference.UserID, preference.mode, preference.outputmode, preference.tactics, id)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateAlias(db *sql.DB, id int, alias Alias) {
	updateSQL := "UPDATE aliases SET user_id = ?, location = ?, alias = ? WHERE id = ?"

	_, err := db.Exec(updateSQL, alias.UserID, alias.Location, alias.Alias, id)
	if err != nil {
		log.Fatal(err)
	}
}
