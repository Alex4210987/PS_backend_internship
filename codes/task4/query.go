package task4

import (
	"database/sql"
	"log"
)

func QueryRoutes(db *sql.DB) []Route {
	querySQL := "SELECT id, start_point, end_point FROM routes"

	rows, err := db.Query(querySQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var routes []Route

	for rows.Next() {
		var route Route
		err := rows.Scan(&route.ID, &route.StartPoint, &route.EndPoint)
		if err != nil {
			log.Fatal(err)
		}
		routes = append(routes, route)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return routes
}

func QueryPreferences(db *sql.DB) []Preference {
	querySQL := "SELECT id, user_id, mode, outputmode, tactics FROM preferences"

	rows, err := db.Query(querySQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var preferences []Preference

	for rows.Next() {
		var preference Preference
		err := rows.Scan(&preference.ID, &preference.UserID, &preference.mode, &preference.outputmode, &preference.tactics)
		if err != nil {
			log.Fatal(err)
		}
		preferences = append(preferences, preference)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return preferences
}

func QueryAliases(db *sql.DB) []Alias {
	querySQL := "SELECT id, user_id, location, alias FROM aliases"

	rows, err := db.Query(querySQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var aliases []Alias

	for rows.Next() {
		var alias Alias
		err := rows.Scan(&alias.ID, &alias.UserID, &alias.Location, &alias.Alias)
		if err != nil {
			log.Fatal(err)
		}
		aliases = append(aliases, alias)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return aliases
}
