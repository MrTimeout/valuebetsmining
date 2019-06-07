package main

import "log"

func main() {
	err := ProcessData()
	if err != nil {
		log.Panic(err)
	}
}
