package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/xellio/ns"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var cfg *config

func init() {
	initializeDatabase()
}

func main() {

	defer func() {
		e := cfg.DB().Close()
		if e != nil {
			log.Println(e)
		}
	}()

	scanner := &ns.Scanner{true}

	arguments := os.Args[1:]
	if len(arguments) == 0 {
		AllJSON()
		return
	}

	result, err := scanner.Query(arguments...)
	if err != nil {
		panic(err)
	}

	for _, host := range result.Hosts {
		free := cfg.DB().NewRecord(host)
		if !free {
			log.Println(fmt.Errorf("Unable to save"))
		}
		cfg.DB().Create(&host)
	}
}

//
// AllJSON outputs all stored hosts in JSON format
func AllJSON() {
	var hosts []ns.Host
	cfg.DB().Find(&hosts)

	for i := range hosts {
		cfg.DB().Model(hosts[i]).Related(&hosts[i].Status)
		cfg.DB().Model(hosts[i]).Related(&hosts[i].Addresses)
		cfg.DB().Model(hosts[i]).Related(&hosts[i].Hostnames)

		cfg.DB().Model(hosts[i]).Related(&hosts[i].Ports)
		for p := range hosts[i].Ports {
			cfg.DB().Model(hosts[i].Ports[p]).Related(&hosts[i].Ports[p].State)
			cfg.DB().Model(hosts[i].Ports[p]).Related(&hosts[i].Ports[p].Service)

			cfg.DB().Model(hosts[i].Ports[p].Service).Related(&hosts[i].Ports[p].Service.CPE)

			cfg.DB().Model(hosts[i].Ports[p]).Related(&hosts[i].Ports[p].Scripts)
			for s := range hosts[i].Ports[p].Scripts {
				cfg.DB().Model(hosts[i].Ports[p].Scripts[s]).Related(&hosts[i].Ports[p].Scripts[s].Elements)
			}
		}

		cfg.DB().Model(hosts[i]).Related(&hosts[i].OS)
		cfg.DB().Model(hosts[i].OS).Related(&hosts[i].OS.Portused)
		cfg.DB().Model(hosts[i].OS).Related(&hosts[i].OS.OSFingerprint)
		cfg.DB().Model(hosts[i].OS).Related(&hosts[i].OS.OSMatch)
		for m := range hosts[i].OS.OSMatch {
			cfg.DB().Model(hosts[i].OS.OSMatch[m]).Related(&hosts[i].OS.OSMatch[m].OSClass)
			for c := range hosts[i].OS.OSMatch[m].OSClass {
				cfg.DB().Model(hosts[i].OS.OSMatch[m].OSClass[c]).Related(&hosts[i].OS.OSMatch[m].OSClass[c].CPE)
			}
		}

		cfg.DB().Model(hosts[i]).Related(&hosts[i].Times)
		cfg.DB().Model(hosts[i]).Related(&hosts[i].Trace)
	}

	out, err := json.Marshal(hosts)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}

func initializeDatabase() (err error) {
	cfg, err = configuration()
	if err != nil {
		return
	}

	cfg.DB().Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1").AutoMigrate(
		&ns.Host{},
		&ns.Status{},
		&ns.Times{},
		&ns.Address{},
		&ns.Hostname{},
		&ns.Port{},
		&ns.State{},
		&ns.Service{},
		&ns.Script{},
		&ns.Element{},
		&ns.CPE{},
		&ns.OS{},
		&ns.Portused{},
		&ns.OSMatch{},
		&ns.OSClass{},
		&ns.OSFingerprint{},
		&ns.Trace{},
		&ns.Hop{},
	)

	return
}
