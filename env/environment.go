package env

import (
	"encoding/json"
	"io/ioutil"
)

type Record struct {
	ErrorMode    int    `json:"errorMode"`
	ErrorMessage string `json:"errorMessage"`
}

func GetEnvironmentVariable() Record {

	raw, err := ioutil.ReadFile("./dev.env.json")
	if err != nil {
		panic(err)
	}

	var record Record
	json.Unmarshal(raw, &record)

	return record
}
