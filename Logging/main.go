package main

import (
	"bufio"
	"os"
	"strings"
	"flag"

	"github.com/michaeliyke/Golang/log"
)

func main() {
	path := flag.String("path", "myapp.log", "Path to log file")
	level := flag.String(
		"level", 
		"ERROR", 
		"Log level to search for.Options are: ERROR, DEBUG, INFO, and CRITICAL",
	)
	flag.Parse()

	file, error := os.Open(*path)
	if error != nil {
		log.Fatal("ERROR OCCURED: ", error)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		if strings.Contains(str, *level) {
			log.Println(str)
		}
	}
}
