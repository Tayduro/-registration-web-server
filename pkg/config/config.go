package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port     int    `yaml:"port"`
	UserName string `yaml:"username"`
	Host     string `yaml:"host"`
	DBname   string `yaml:"dbname"`
	Password string `yaml:"password"`
	Sslmode  string `yaml:"sslmode"`
	Key string `yaml:"key"`
}

func NewConfig() *Config  {
	return &Config{
<<<<<<< HEAD
		//Port: 0,
		//UserName: "",
		//Host: "",
		//DBname:"",
		//Password:"",
		//Sslmode: "",
=======
		Port: 0,
		UserName: "",
		Host: "",
		DBname:"",
		Password:"",
		Sslmode: "",
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3
	}
}

func LoadConfig() *Config {
	config := NewConfig()

	yamlFile, err :=ioutil.ReadFile("./cmd/signup-server/config.yaml")
	if err != nil{
		log.Fatalf("Error %v",err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil{
		log.Fatalf("Error %v",err)
	}
	return config
}

func ConfigServer() string {
	config := LoadConfig()
	stringOfData := fmt.Sprintf( "postgres://%s:%s@%s:%d/%s?sslmode=%s",config.UserName, config.Password, config.Host, config.Port, config.DBname ,config.Sslmode)
	return stringOfData
}

<<<<<<< HEAD
func GetKey() string {
	yfile, err := ioutil.ReadFile("./cmd/signup-server/config.yaml")

	if err != nil {

		log.Fatal(err)
	}

	conf := *&Config{}

	err = yaml.Unmarshal(yfile, &conf)

	if err != nil {

		log.Fatal(err)
	}
	return conf.Key
}

=======
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3
