package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//Check for the first and last digit
		actual := []rune(scanner.Text())
		var first int
		var last int

		for i := 0; i < len(actual); i++ {
			if
		}		
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
