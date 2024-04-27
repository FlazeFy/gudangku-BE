package configs

import (
	"github.com/tkanos/gonfig"
)

type ConfigurationJWT struct {
	JWT_EXP string
}

func GetConfigJWT() ConfigurationJWT {
	conf := ConfigurationJWT{}
	gonfig.GetConf("configs/jwts.json", &conf)
	return conf
}
