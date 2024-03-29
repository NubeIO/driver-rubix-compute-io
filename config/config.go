package config

import (
	"flag"
	"github.com/NubeIO/configor"
	"path"
)

type Configuration struct {
	Server struct {
		Address string
		Port    int `default:"5001"`
	}

	Location struct {
		GlobalDir string `default:"./"`
		ConfigDir string `default:"config"`
		DataDir   string `default:"data"`
	}
	Prod  bool `default:"false"`
	Debug bool `default:"false"`
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
	port := flag.Int("p", 5001, "Port")
	rootDir := flag.String("r", "./", "Root Directory")
	appDir := flag.String("a", "./", "App Directory")
	dataDir := flag.String("d", "data", "Data Directory")
	configDir := flag.String("c", "config", "Config Directory")
	prod := flag.Bool("prod", false, "Deployment Mode")
	debug := flag.Bool("debug", false, "test mode")
	flag.Parse()
	config.Server.Port = *port
	conf.Location.GlobalDir = path.Join(*rootDir, *appDir)
	config.Location.DataDir = *dataDir
	config.Location.ConfigDir = *configDir
	config.Prod = *prod
	config.Debug = *debug
	return config
}

func (conf *Configuration) GetAbsDataDir() string {
	return path.Join(conf.Location.GlobalDir, conf.Location.DataDir)
}

func (conf *Configuration) GetAbsConfigDir() string {
	return path.Join(conf.Location.GlobalDir, conf.Location.ConfigDir)
}
