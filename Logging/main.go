package main

import (
	"bufio"
	"os"
	"strings"

	"example.com/log"
)

func main() {
	file, error := os.Open("myapp.log")
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

		if strings.Contains(str, "ERROR") {
			log.Println(str)
		}
	}
}
