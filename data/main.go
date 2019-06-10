package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"valuebetsmining/data/entities"
)

const (
	//DefaultDir ... dir of the project
	DefaultDir = "src"
)

func main() {
	if t, err := whereIAm(); err != nil {
		panic(err)
	} else if !t {
		panic("Please, execute me in my directory")
	}
	parsingFlags()
}

func whereIAm() (bool, error) {
	dir := DefaultDir
	actualDir, err := os.Getwd()
	if err != nil {
		return false, err
	}
	if dir != strings.Split(actualDir, "/")[len(strings.Split(actualDir, "/"))-1] {
		return false, nil
	}
	return true, nil
}

func parsingFlags() {
	fileType := flag.String("type-file", "csv", "You have to choose the file extension to parse data. Possible values: 'json' or 'csv'")
	yearFrom := flag.Int("year-from", 10, "Year from to search about teams and divisions. By default is 10 an is also the minimun, representing 2010(is inclusive)")
	yearTo := flag.Int("year-to", 19, "Year to to seach about teams and divisions. By default is 19 an is also the maximun, representing 2019(inclusive)")
	country := flag.String("country", "all", "Country to search in the machine. By default is Spain and it neccesary to use nomenclature of the file params.md to use it well. If you dont write any country, it will cacth all countries")
	division := flag.String("division", "all", "Divisions to search about a country. Possible values are all|other division. See params.md to more info")
	skeleton := flag.Bool("skeleton-json", false, "It prints the skeleton of the json to show the structure of the file to pass like a param.")
	configFile := flag.String("config-file", "", "If you use this option, use it alone and pass a correct path to the param. By default is the config/config.json file. Review it to know and understand how to work with it")

	flag.Parse()

	shell, err := entities.NewShell()
	if err != nil {
		log.Panic(err)
	}
	if *skeleton {
		str, err := entities.Skeleton()
		if err != nil {
			panic(err)
		}
		fmt.Println(str)
	} else if *configFile != "" {
		if *configFile == entities.DefaultFileConfig {
			panic(entities.ErrInvalidConfigFile)
		}
	} else {
		log.Println(shell.IsFileType(*fileType))
		log.Println(shell.IsYearFrom(*yearFrom))
		log.Println(shell.IsYearTo(*yearTo))
		log.Println(shell.IsCountry(*country))
		log.Println(shell.IsDivision(*division))
		c, err := entities.NewConfig(*country, *division, *yearFrom, *yearTo)
		if err != nil {
			panic(err)
		}
		json, err := json.MarshalIndent(c, "", "	")
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile("c.json", json, 0666)
		if err != nil {
			panic(err)
		}
	}
}
