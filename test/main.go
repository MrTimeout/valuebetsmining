package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	praticseArgs()
}

func praticseArgs() {
	a := flag.Bool("t", false, "something about boolean")
	b := flag.Int("i", 0, "something about int")
	c := flag.String("s", "default", "something about string")

	flag.Parse()

	log.Println(*a, *b, *c)

	for _, value := range flag.Args() {
		log.Println(value)
	}

}

func simpelArgs() {
	times := flag.Int("n", 10, "Represents how many rows will be prompted of wildcard(by default)")
	char := flag.String("char", "*", "Represents the char to be prompetd (wildcard by default)")
	loop := func(n int, char string) string {
		var s strings.Builder
		if n <= 0 {
			n = 10
		}
		if len(strings.TrimSpace(char)) == 0 {
			char = "-"
		} else {
			char = char[:1]
		}
		for i := 0; i <= n; i++ {
			for j := 0; j <= i; j++ {
				s.WriteString(char)
			}
			s.WriteString("\n")
		}
		return s.String()
	}
	flag.Parse()

	log.Println("\n" + loop(*times, *char))

}

func parseArgs() {
	words := flag.Bool("verbose", false, "If it set to true, then it will print verbose output. False, otherwise")
	net := flag.Bool("net", false, "If it sets to true, then it dowload the data. False otherwise")

	flag.Parse()

	if flag.Parsed() {
		for _, value := range os.Args[1:] {
			switch value {
			case "--verbose":
				if *words {
					log.Println("Verbose")
				}
			case "--net":
				if *net {
					log.Println("Connecting to the web page to download content")
				}
			default:
				os.Exit(1)
			}
		}
		log.Println("Program finished: ", *words, *net)
	}

}
