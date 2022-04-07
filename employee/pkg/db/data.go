package db

import (
	"encoding/json"
	"github.com/asqat/mgid/employee/pkg/model"
	"io/ioutil"
)

const dbPath = "dataset.json"

var database []model.Employee

func GetDatabase() []model.Employee {
	return database
}

func InitDummyDatabase() error {
	blob, err := ioutil.ReadFile(dbPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(blob, &database)
}
