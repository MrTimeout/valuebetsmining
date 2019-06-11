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
	"valuebetsmining/data/logic"
)

const (
	//DefaultDir ... dir of the project
	DefaultDir = "data"
)

func main() {
	if t, err := whereIAm(); err != nil {
		panic(err)
	} else if !t {
		panic("Please, execute me in my directory")
	}
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "output-file":
			parsingOutputFile()
		case "connection-config":
			parsingConnection()
		case "run":
			parsingRun()
		default:
			fmt.Println("Commands expected: output-file or connection-file or run")
			os.Exit(2)
		}
	} else {
		panic("Write a command")
	}
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

func parsingRun() {
	run := flag.NewFlagSet("run", flag.ExitOnError)
	outputFile := run.String("output-file", entities.DefaultFileConfigOutput, fmt.Sprintf("If you use this option, use it alone and pass a correct path to the param. By default is the %s file. Review it to know and understand how to work with it", entities.DefaultFileConfigOutputSchema))
	connectionFile := run.String("connection-file", entities.DefaultFileConfigConnection, fmt.Sprintf("If you use this option, use it alone and pass a correct path to the param. By default is the %s file. Review it to know and understand how to work with it", entities.DefaultFileConfigConnectionSchema))
	output := run.Bool("run-output", true, "Run output script to create the files. By default is true")
	connection := run.Bool("run-connection", true, "Run connections to get the data. By default is true")

	if len(os.Args) >= 2 {
		run.Parse(os.Args[2:])

		if len(strings.TrimSpace(*outputFile)) != 0 {
			c, err := entities.ReadFileTakingArgs(*outputFile)
			if err != nil {
				panic(err)
			}
			file, err := ioutil.ReadFile(*outputFile)
			if err != nil {
				panic(err)
			}
			config := entities.ConfigFile{}
			err = json.Unmarshal(file, &config)
			if err != nil {
				panic(err)
			}
			entities.OutputConfigFile = *outputFile
			entities.MaxLines = c.Lines
			entities.ExtFile = c.Extension
			entities.DefaultLenGoal = c.Lines
			entities.DefaultLenResult = c.Lines
		} else {
			panic("Error parsing nil string")
		}

		if len(strings.TrimSpace(*connectionFile)) != 0 {
			_, err := entities.ReadFile(*connectionFile)
			if err != nil {
				panic(err)
			}
			file, err := ioutil.ReadFile(*connectionFile)
			if err != nil {
				panic(err)
			}
			config := entities.Config{}
			err = json.Unmarshal(file, &config)
			if err != nil {
				panic(err)
			}
			entities.ConnectionConfigFile = *connectionFile
		} else {
			panic("Error parsing nil string")
		}
		err := logic.ProcessData(*connection, *output)
		if err != nil {
			panic(err)
		}
	}

}

func parsingOutputFile() {
	outputFile := flag.NewFlagSet("output-file", flag.ExitOnError)
	numberMatchs := outputFile.Int("number-matchs", 10, "Number of matchs to calculate properties")
	typeFile := outputFile.String("type-file", "csv", "Type of the extension of the output file")
	skeleton := outputFile.Bool("skeleton-file", false, "Return the squeleton of the json file")
	configFile := outputFile.String("config-file", "", fmt.Sprintf("If you use this option, use it alone and pass a correct path to the param. By default is the %s file. Review it to know and understand how to work with it", entities.DefaultFileConfigOutputSchema))

	outputFile.Parse(os.Args[2:])

	if *skeleton {
		str, err := entities.SkeletonConfigFile()
		if err != nil {
			panic(err)
		}
		fmt.Println(str)
	} else if *configFile != "" {
		if *configFile == entities.DefaultFileConfigOutput {
			panic(entities.ErrInvalidConfigFile)
		}
		file, err := ioutil.ReadFile(*configFile)
		if err != nil {
			panic(err)
		}
		c := entities.ConfigFile{}
		err = json.Unmarshal(file, &c)
		if err != nil {
			panic(err)
		}
	} else {
		oFile, err := entities.NewConfigFile(*numberMatchs, *typeFile)
		if err != nil {
			panic(err)
		}
		d, err := json.MarshalIndent(oFile, "", "	")
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(entities.DefaultFileConfigOutputSchema, d, 0666)
		if err != nil {
			panic(err)
		}
	}

}

func parsingConnection() {
	conConfig := flag.NewFlagSet("connection-config", flag.ExitOnError)
	//fileType := conConfig.String("type-file", "csv", "You have to choose the file extension to parse data. Possible values: 'json' or 'csv'")
	yearFrom := conConfig.Int("year-from", 10, "Year from to search about teams and divisions. By default is 10 an is also the minimun, representing 2010(is inclusive)")
	yearTo := conConfig.Int("year-to", 19, "Year to to seach about teams and divisions. By default is 19 an is also the maximun, representing 2019(inclusive)")
	country := conConfig.String("country", "all", "Country to search in the machine. By default is Spain and it neccesary to use nomenclature of the file params.md to use it well. If you dont write any country, it will cacth all countries")
	division := conConfig.String("division", "all", "Divisions to search about a country. Possible values are all|other division. See params.md to more info")
	skeleton := conConfig.Bool("skeleton-json", false, "It prints the skeleton of the json to show the structure of the file to pass like a param.")
	configFile := conConfig.String("config-file", "", fmt.Sprintf("If you use this option, use it alone and pass a correct path to the param. By default is the %s file. Review it to know and understand how to work with it", entities.DefaultFileConfigConnection))

	conConfig.Parse(os.Args[2:])

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
		if *configFile == entities.DefaultFileConfigConnection {
			panic(entities.ErrInvalidConfigFile)
		}
		file, err := ioutil.ReadFile(*configFile)
		if err != nil {
			panic(err)
		}
		c := entities.Config{}
		err = json.Unmarshal(file, &c)
		if err != nil {
			panic(err)
		}
	} else {
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
		fmt.Println(string(json))
		err = ioutil.WriteFile(entities.DefaultFileConfigConnectionSchema, json, 0666)
		if err != nil {
			panic(err)
		}
	}
}
