package config

import (
	"flag"
	"github.com/NubeIO/configor"
	"path"
)

type Configuration struct {
	Server struct {
		Address string
		Port    int `default:"1883"`
	}
	Credential struct {
		Auth     bool   `default:"false"`
		Username string `default:"admin"`
		Password string `default:"test_admin"`
	}
	Storage struct {
		EnablePersistence bool   `default:"true"`
		DB                string `default:"mqtt.db"`
	}
	Location struct {
		GlobalDir string `default:"./"`
		ConfigDir string `default:"config"`
		DataDir   string `default:"data"`
	}
	Prod bool `default:"false"`
}

var config *Configuration = nil

func Get() *Configuration {
	return config
}

func CreateApp() *Configuration {
	config = new(Configuration)
	config = config.Parse()
	err := configor.New(&configor.Config{EnvironmentPrefix: "RUBIX_PIIO"}).Load(config, path.Join(config.GetAbsConfigDir(), "config.yml"))
	if err != nil {
		panic(err)
	}
	return config
}

func (conf *Configuration) Parse() *Configuration {
	port := flag.Int("p", 1883, "Port")
	globalDir := flag.String("g", "./", "Global Directory")
	dataDir := flag.String("d", "data", "Data Directory")
	configDir := flag.String("c", "config", "Config Directory")
	prod := flag.Bool("prod", false, "Deployment Mode")
	auth := flag.Bool("auth", false, "enable mqtt auth")
	password := flag.String("password", "test_password", "auth password")
	flag.Parse()
	conf.Server.Port = *port
	conf.Location.GlobalDir = *globalDir
	conf.Location.DataDir = *dataDir
	conf.Location.ConfigDir = *configDir
	conf.Prod = *prod
	conf.Credential.Auth = *auth
	conf.Credential.Password = *password
	return conf
}

func (conf *Configuration) GetAbsDataDir() string {
	return path.Join(conf.Location.GlobalDir, conf.Location.DataDir)
}

func (conf *Configuration) GetAbsConfigDir() string {
	return path.Join(conf.Location.GlobalDir, conf.Location.ConfigDir)
}
