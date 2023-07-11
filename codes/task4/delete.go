package task4

import (
	"database/sql"
	"log"
)

func DeleteRoute(db *sql.DB, id int) {
	deleteSQL := "DELETE FROM routes WHERE id = ?"

	_, err := db.Exec(deleteSQL, id)
	if err != nil {
		log.Fatal(err)
	}
}

func DeletePreference(db *sql.DB, id int) {
	deleteSQL := "DELETE FROM preferences WHERE id = ?"

	_, err := db.Exec(deleteSQL, id)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteAlias(db *sql.DB, id int) {
	deleteSQL := "DELETE FROM aliases WHERE id = ?"

	_, err := db.Exec(deleteSQL, id)
	if err != nil {
		log.Fatal(err)
	}
}
