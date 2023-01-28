package config

import (
	"SpotterBackend/src/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type AppConfig struct {
	IsDebug    bool             `yaml:"is_debug" env_default:"true"`
	ConnConfig ConnectionConfig `yaml:"listen"`
	DbConfig   DataBaseConfig   `yaml:"database"`
}

type ConnectionConfig struct {
	Type        string `yaml:"type" env-default:"port"`
	BindAddress string `yaml:"bind_ip" env-default:"0.0.0.0:8888"`
}

type DataBaseConfig struct {
	Host     string `yaml:"host" env-default:"0.0.0.0"`
	Port     string `yaml:"port" env-default:"5432"`
	DbName   string `yaml:"db_name" env-default:"spotter_db"`
	Username string `yaml:"username" env-default:"nick"`
	Password string `yaml:"password" env-default:"testpaassword"`
}

var config *AppConfig
var once sync.Once

func Init() {
	log := logging.GetLogger()
	config = &AppConfig{}
	err := cleanenv.ReadConfig("config.yml", config)
	if err != nil {
		log.Errorf("Failed to init config: %s", err)
	}
}

func GetConfig() *AppConfig {
	once.Do(Init)
	return config
}

// Example

//type ConfigDatabase struct {
//	Port     string `yaml:"port" env:"PORT" env-default:"5432"`
//	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
//	DbName     string `yaml:"name" env:"NAME" env-default:"postgres"`
//	Username     string `yaml:"user" env:"USER" env-default:"user"`
//	Password string `yaml:"password" env:"PASSWORD"`
//}

//Spotter

//---
//
//is_debug: true
//listen:
//type: port
//bind_address:  0.0.0.0:8888
//database:
//host: localhost
//port: 5432
//name: spotter_db
//user: nick
//password: testpaassword
