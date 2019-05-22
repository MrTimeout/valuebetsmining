package data

import "log"

func main() {
	config, err := ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	connection := &Connection{config}
	res, err := connection.Get()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res)
}

func processData() {

}
