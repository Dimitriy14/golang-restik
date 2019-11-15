package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	FilePath = "../../config.json"

	Conf Configuration
)

type Configuration struct {
	ListenURL string `json:"ListenURL" default:"4444"`
	BasePath  string `json:"BasePath" default:"/restik"`

	Postgres struct {
		Host     string `json:"Host"`
		Port     string `json:"Port"`
		DBName   string `json:"DBName" default:"restik"`
		User     string `json:"User" default:"admin"`
		Password string `json:"Password" default:"1488"`
	} `json:"Postgres"`

	UseLogFile bool   `json:"UseLogFile" default:"false"`
	LogFile    string `json:"LogFile" default:"restik.log"`
	LogLevel   string `json:"LogLevel" default:"debug"`
}

func Load() error {
	fileContent, err := ioutil.ReadFile(FilePath)
	if err != nil {
		return fmt.Errorf("cannot read config file (file: %s): %s", FilePath, err)
	}

	return json.Unmarshal(fileContent, &Conf)
}
