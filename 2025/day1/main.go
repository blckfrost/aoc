package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// starting position
	pos := 50
	// keep count of zero
	atZero := 0

	// open the file
	inputFile, err := os.Open("day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	// read file lines
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()

		// first char is either L/R
		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case 'L':
			pos = (pos - distance + 100) % 100

			if pos == 0 {
				atZero++
			}

		case 'R':
			pos = (pos + distance) % 100
			if pos == 0 {
				atZero++
			}
		}
	}
	fmt.Println("final position:", pos)
	fmt.Println("password:", atZero)

}
