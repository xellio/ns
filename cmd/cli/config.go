package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/jinzhu/gorm"
	yaml "gopkg.in/yaml.v2"
)

type config struct {
	db    *gorm.DB
	DBCfg *dbConfig `yaml:"database"`
}

type dbConfig struct {
	Login          string        `yaml:"login"`
	Password       string        `yaml:"password"`
	Host           string        `yaml:"host"`
	Port           string        `yaml:"port"`
	Name           string        `yaml:"name"`
	ConMaxLifetime time.Duration `yaml:"conMaxLifetime"`
	MaxIdleConns   int           `yaml:"maxIdleConns"`
	MaxOpenConns   int           `yaml:"maxOpenConns"`
}

//
// Configuration creates and returns a config struct for the given config yaml file
//
func configuration() (*config, error) {
	filepath := "./cmd/cli/config.yml"
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	config := &config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open("mysql",
		fmt.Sprintf(
			"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
			config.DBCfg.Login,
			config.DBCfg.Password,
			config.DBCfg.Name,
		),
	)
	if err != nil {
		return nil, err
	}

	db.DB().SetConnMaxLifetime(config.DBCfg.ConMaxLifetime * time.Second)
	db.DB().SetMaxIdleConns(config.DBCfg.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.DBCfg.MaxOpenConns)

	db.LogMode(false)

	config.db = db
	return config, nil
}

//
// DB returns the *gorm.DB database connection
//
func (cfg *config) DB() *gorm.DB {
	return cfg.db
}
