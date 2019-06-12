package main

import (
	"log"
	"os"
)

func main() {
	IsDir("leagues/CSV")
	// IsFile("leagues/CSV")
	IsDir("leagues/CSV/Spain")
	// IsFile("leagues/CSV/Spain")
	IsDir("leagues/CSV")
}

//IsDir ...
func IsDir(name string) {
	f, err := os.Stat(name)
	if os.IsNotExist(err) {
		err = os.MkdirAll(name, 0755)
		if err != nil {
			log.Println(err)
		}
	} else {
		err = os.RemoveAll("leagues/CSV")
		if err != nil {
			log.Println(err)
		}
		err = os.MkdirAll(name, 0755)
		if err != nil {
			log.Println(err)
		}
	}
	log.Println(f)
}

//IsFile ...
func IsFile(name string) {
	_, err := os.Create(name + "/file.csv")
	if err != nil {
		log.Println(err)
	}
}
