package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

var (
	FilePath = "config.json"

	Conf Configuration
)

type Configuration struct {
	ListenURL string `json:"ListenURL" default:":4444" envconfig:"PORT"`
	PORT      string `envconfig:"PORT"`
	BasePath  string `json:"BasePath" default:"/restik"`

	Postgres struct {
		Host     string `json:"Host"`
		Port     string `json:"Port"`
		DBName   string `json:"DBName" default:"restik"`
		User     string `json:"User" default:"admin"`
		Password string `json:"Password" default:"1488"`
	} `json:"Postgres"`

	HerokuPg string `json:"HerokuPg" envconfig:"DATABASE_URL"`

	UseLogFile bool   `json:"UseLogFile" default:"false"`
	LogFile    string `json:"LogFile" default:"restik.log"`
	LogLevel   string `json:"LogLevel" default:"debug"`
}

func Load() error {
	if err := readFile(&Conf); err != nil {
		return err
	}

	if err := readEnv(&Conf); err != nil {
		return err
	}

	fmt.Printf("%+v", Conf)
	return nil
}

func readFile(cfg *Configuration) error {
	fileContent, err := os.Open(FilePath)
	if err != nil {
		return err
	}

	if err = json.NewDecoder(fileContent).Decode(&Conf); err != nil {
		return err
	}
	return nil
}

func readEnv(cfg *Configuration) error {
	err := envconfig.Process("", cfg)
	if err != nil {
		return err
	}
	return nil
}
