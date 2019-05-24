package data

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"testing"
)

//TestWriteJSONToAFile ... Write random info of the stucts team* to a file to understand how it works
func TestWriteJSONToAFile(t *testing.T) {
	csvFile, err := os.Open("test/SP1_1819.csv")
	if err != nil {
		t.Errorf("Error: %#v", err)
	}
	reader := csv.NewReader(csvFile)
	teams := make(map[string]Team)
	matchs := []Match{}
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error(err)
		}
		log.Print(teams, matchs, line)
	}
}
