package configs

import (
	"log"

	"github.com/BurntSushi/toml"
)

//Config struct that holds configs
type Config struct {
	Server server
	APIKey apikey
}

type server struct {
	IP   string
	Port string
	TLS  string
}

type apikey struct {
	TransAPIKey string `toml:"transapikey"`
}

//Configs is the variable that holds config information for microservice
var Configs Config

//Initializes configurations
func init() {
	if _, err := toml.DecodeFile("./configs/config.toml", &Configs); err != nil {
		log.Println("Check your configs.")
		log.Fatalln(err.Error())
		return
	}
}
