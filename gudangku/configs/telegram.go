package configs

import (
	"github.com/tkanos/gonfig"
)

type ConfigurationTele struct {
	TELE_TOKEN string
}

func GetConfigTele() ConfigurationTele {
	conf := ConfigurationTele{}
	gonfig.GetConf("configs/telegram.json", &conf)
	return conf
}
