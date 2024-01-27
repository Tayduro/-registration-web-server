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

func ReadConfig(path string) (*Config, error) {
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

func (c *Config) DBURL() string {

	fmt.Println(fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.UserName, c.Password, c.Host, c.Port, c.DBname, c.Sslmode))
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.UserName, c.Password, c.Host, c.Port, c.DBname, c.Sslmode)
}
