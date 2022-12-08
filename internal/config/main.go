package config

type MainConfig struct {
	Service Service
}

func LoadMainConfig() (cfg MainConfig, err error) {
	if cfg.Service, err = LoadServiceConfig(); err != nil {
		return MainConfig{}, err
	}

	return cfg, nil
}
