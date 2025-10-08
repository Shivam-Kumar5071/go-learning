package config

import (
	"flag"
	"log"
	"os"
)

type HTTPServer struct{
	Address string
}

type Config struct{
	Env string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer `yaml:"http_server "`
}


func MustLoad(){

	var cfgPath string
	
	cfgPath = os.Getenv("CONFIG_PATH")

	if cfgPath == ""{
		flags := flag.String("config","","path to the cofiguration file")
		flag.Parse()
		cfgPath = *flags

		if cfgPath == ""{
			log.Fatal()
		}

	}


}