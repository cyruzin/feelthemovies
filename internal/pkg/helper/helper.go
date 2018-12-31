package helper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

// Attach receives a map of int/[]int and attach
// the IDs on the given pivot table.
func Attach(s map[int64][]int, pivot string, db *sql.DB) (int64, error) {

	for index, ids := range s {
		for _, values := range ids {

			query := fmt.Sprintf("INSERT INTO %s VALUES (?,?)", pivot)

			stmt, err := db.Prepare(query)

			if err != nil {
				log.Fatal(err)
			}

			defer stmt.Close()

			_, err = stmt.Exec(index, values)

			if err != nil {
				log.Fatal(err)
			}

		}
	}

	return 1, nil
}

// Detach receives a map of int/[]int and Detach
// the IDs on the given pivot table.
func Detach(s map[int64][]int, pivot, field string, db *sql.DB) (int64, error) {
	for index := range s {

		query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", pivot, field)

		stmt, err := db.Prepare(query)

		if err != nil {
			log.Fatal(err)
		}

		defer stmt.Close()

		_, err = stmt.Exec(index)

		if err != nil {
			log.Fatal(err)
		}
	}
	return 1, nil
}

// Sync receives a map of int/[]int and sync
// the IDs on the given pivot table.
func Sync(s map[int64][]int, pivot, field string, db *sql.DB) (int64, error) {

	empty, err := IsEmpty(s)

	if err != nil {
		log.Fatal(err)
	}

	if !empty {

		_, err := Detach(s, pivot, field, db)

		if err != nil {
			log.Fatal(err)
		}

		_, err = Attach(s, pivot, db)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err := Detach(s, pivot, field, db)

		if err != nil {
			log.Fatal(err)
		}
	}

	return 1, nil
}

// IsEmpty checks if a given map of int/[]int is empty.
func IsEmpty(s map[int64][]int) (bool, error) {
	empty := true

	for _, ids := range s {
		if len(ids) > 0 {
			empty = false
		}
	}

	return empty, nil
}

// ToJSON receives an interface as argument and
// returns a JSON string.
func ToJSON(j interface{}) (string, error) {
	data, err := json.Marshal(j)

	if err != nil {
		log.Println(err)
	}

	res := fmt.Sprintf("%s", data)

	return res, nil
}

// ToJSONIndent receives an interface as argument
// and returns a JSON string indented.
func ToJSONIndent(j interface{}) (string, error) {
	data, err := json.MarshalIndent(j, "", "\t")

	if err != nil {
		log.Println(err)
	}

	res := fmt.Sprintf("%s", data)

	return res, nil
}
