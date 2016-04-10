package main

import (
	"encoding/json"
	"os"
)

type TinkerConfiguration struct {
	Users []UserConfiguration `json:"users"`
}

type UserConfiguration struct {
	Address      string   `json:"address"`
	ContactEmail string   `json:"contact_email"`
	Roadwarrior  bool     `json:"roadwarrior"`
	Subnets      []string `json:"subnets"`
	Username     string   `json:"username"`
}

func (t *TinkerConfiguration) LoadFromFile(filename string) (err error) {
	configFile, err := os.Open(filename)
	if err != nil {
		log.Errorf("error opening configuration file %s: %s", filename, err)
		return err
	}
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(t); err != nil {
		log.Errorf("error loading data from configuration file %s: %s", filename, err)
		return err
	}
	return nil
}
