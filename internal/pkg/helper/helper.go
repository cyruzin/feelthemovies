package helper

import (
	"encoding/json"
	"fmt"
	"log"
)

// ToJSON receives an interface as argument and returns a JSON string.
func ToJSON(j interface{}) (string, error) {
	data, err := json.Marshal(j)

	if err != nil {
		log.Println(err)
	}

	res := fmt.Sprintf("%s", data)

	return res, nil
}

// ToJSONIndent receives an interface as argument and returns a JSON string indented.
func ToJSONIndent(j interface{}) (string, error) {
	data, err := json.MarshalIndent(j, "", "\t")

	if err != nil {
		log.Println(err)
	}

	res := fmt.Sprintf("%s", data)

	return res, nil
}
