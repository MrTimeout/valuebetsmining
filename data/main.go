package main

import (
	"log"
	"valuebetsmining/src/data/logic"
)

func main() {
	err := logic.ProcessData()
	if err != nil {
		log.Panic(err)
	}
}
