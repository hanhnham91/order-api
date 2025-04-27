package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"gitlab.com/pak-server/pkg/paklogger"
)

var config *Config

type Config struct {
	Stage            StageType `envconfig:"STAGE"`
	TokenSecretKey   string    `envconfig:""`
	Port             string    `envconfig:"PROJECT_PORT"`
	FirebaseAdminSDK string    `envconfig:"FIREBASE_ADMIN_SDK"`
}

func init() {
	config = &Config{}

	if err := godotenv.Load(); err != nil {
		paklogger.Warnln(errors.Wrap(err, "godotenv"))
	}

	if err := envconfig.Process("", config); err != nil {
		paklogger.Fatal(err)
	}

	if len(config.Port) == 0 {
		config.Port = "8080"
	}
}

func GetConfig() *Config {
	return config
}

type StageType string

const (
	StageTypeProd    StageType = "PROD"
	StageTypeStaging StageType = "STG"
	StageTypeDev     StageType = "DEV"
	StageTypeLocal   StageType = "LOCAL"
)
