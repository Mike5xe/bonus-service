package config

import "github.com/kelseyhightower/envconfig"

type Service struct {
	Port int `envconfig:"BONUS_SERVICE_PORT" default:"8080"`
}

func LoadServiceConfig() (cfg Service, err error) {
	if err = envconfig.Process("", &cfg); err != nil {
		return Service{}, err
	}

	return cfg, nil
}
