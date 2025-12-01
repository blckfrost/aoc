package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func partOne() {
	// starting position
	pos := 50
	// keep count of zero
	password := 0

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
			fmt.Printf("error converting string to int: %v\n", err)
		}

		switch direction {
		case 'L':
			pos = (pos - distance + 100) % 100

			if pos == 0 {
				password++
			}

		case 'R':
			pos = (pos + distance) % 100
			if pos == 0 {
				password++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("final position: %v password: %v\n", pos, password)

}

func parTwo() {
	pos := 50
	password := 0

	file, err := os.Open("day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		direction := line[0]
		distance, err := strconv.Atoi(line[1:])

		if err != nil {
			fmt.Printf("error converting string to int: %v\n", err)
		}

		for i := 0; i < distance; i++ {
			switch direction {
			case 'L':
				pos = (pos - 1 + 100) % 100
			case 'R':
				pos = (pos + 1) % 100
			}

			if pos == 0 {
				password++
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("final position: %v password: %v\n", pos, password)

}

func main() {
	partOne()
	parTwo()
}
