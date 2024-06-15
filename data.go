package main

import (
	"encoding/json"
	"os"
)

type Data struct {
	Users []User `json:"users"`
}

func loadData(filename string) (Data, error) {
	booksBytes, err := os.ReadFile(filename)
	checkError(err)

	var data Data
	err = json.Unmarshal(booksBytes, &data)
	checkError(err)

	return data, err
}

func saveData(filename string, data Data) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	checkError(err)

	err = os.WriteFile(filename, bytes, 0644)
	return err
}