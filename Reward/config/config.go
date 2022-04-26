package config

import (
	"fmt"

	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/literals"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type WebServerConfig struct {
	Port         string `required:"true" split_words:"true"`
	RoutePrefix  string `required:"false" split_words:"true" default:"/reward"`
	Db           string `required:"false" split_words:"true"`
	DbCollection string `required:"false" split_words:"true"`
}

func FromEnv() (*WebServerConfig, error) {
	cfgFilename := "../../etc/localhost.env"
	if err := godotenv.Load(cfgFilename); err != nil {
		fmt.Println("No config files found to load to env.")
	}

	webServerConfig := &WebServerConfig{}
	err := envconfig.Process(literals.AppPrefix, webServerConfig)
	if err != nil {
		return nil, err
	}

	return webServerConfig, nil
}
