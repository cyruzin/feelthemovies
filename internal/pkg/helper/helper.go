package helper

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// Attach receives a map of int/[]int and attach the IDs on the given pivot table.
func Attach(s map[int64][]int, pivot string, db *sql.DB) (int64, error) {
	for index, ids := range s {
		for _, values := range ids {
			stmt, err := db.Prepare("INSERT INTO " + pivot + " VALUES (?,?)")
			if err != nil {
				return 0, err
			}
			defer stmt.Close()
			_, err = stmt.Exec(index, values)
			// Error handler for duplicate entries
			if mysqlError, ok := err.(*mysql.MySQLError); ok {
				if mysqlError.Number == 1062 {
					return 0, err
				}
			}
		}
	}
	return 1, nil
}

// Detach receives a map of int/[]int and Detach the IDs on the given pivot table.
func Detach(s map[int64][]int, pivot, field string, db *sql.DB) (int64, error) {
	for index := range s {
		stmt, err := db.Prepare("DELETE FROM " + pivot + " WHERE " + field + " = ?")
		if err != nil {
			return 0, err
		}
		defer stmt.Close()
		_, err = stmt.Exec(index)
		if err != nil {
			return 0, err
		}
	}
	return 1, nil
}

// Sync receives a map of int/[]int and sync the IDs on the given pivot table.
func Sync(s map[int64][]int, pivot, field string, db *sql.DB) (int64, error) {
	empty := IsEmpty(s)
	if !empty {
		_, err := Detach(s, pivot, field, db)
		if err != nil {
			return 0, err
		}
		_, err = Attach(s, pivot, db)
		if err != nil {
			return 0, err
		}
	} else {
		_, err := Detach(s, pivot, field, db)
		if err != nil {
			return 0, err
		}
	}
	return 1, nil
}

// IsEmpty checks if a given map of int/[]int is empty.
func IsEmpty(s map[int64][]int) bool {
	empty := true
	for _, ids := range s {
		if len(ids) > 0 {
			empty = false
		}
	}
	return empty
}

// ToJSON receives an interface as argument and returns a JSON string.
func ToJSON(j interface{}) (string, error) {
	data, err := json.Marshal(j)
	if err != nil {
		return "", err
	}
	res := fmt.Sprintf("%s", data)
	return res, nil
}

// ToJSONIndent receives an interface as argument and returns a JSON string indented.
func ToJSONIndent(j interface{}) (string, error) {
	data, err := json.MarshalIndent(j, "", "\t")
	if err != nil {
		return "", err
	}
	res := fmt.Sprintf("%s", data)
	return res, nil
}

// HashPassword encrypts a given password using bcrypt algorithm.
func HashPassword(password string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPasswordHash checks if the given passwords matches.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
