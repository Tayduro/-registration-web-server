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
	Key      string `yaml:"key"`
}

func NewConfig() *Config {
	return &Config{
	}
}

func ReadConfig(path string) (*Config, error) {
	//config := NewConfig()

	config := &Config{}

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Error %v", err)
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Error %v", err)
		return nil, err
	}

	return config, nil
}

func (c *Config) DBURL() string  {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.UserName, c.Password, c.Host, c.Port, c.DBname, c.Sslmode)
}
// "postgres://postgres:12345@localhoct:5432/users?sslmode=disable"
//func ConfigServer() (string, error) {
//	config, err := LoadConfig("./cmd/signup-server/config.yaml")
//	if err != nil {
//		return "", err
//	}
//}

//func GetKey(path string) (string, error) {
//	yfile, err := ioutil.ReadFile(path)
//
//	if err != nil {
//
//		//log.Fatal(err)
//		return "", err
//	}
//
//	//conf := &Config{}
//	conf := NewConfig()
//
//	err = yaml.Unmarshal(yfile, conf)
//
//	if err != nil {
//
//		//log.Fatal(err)
//		return "", err
//	}
//	return conf.Key , nil
//}
