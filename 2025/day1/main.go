package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/* Advent of Code 2025 - Day 1
L -> toward lower numbers
R -> toward higher numbers
dstance -> how many clicks dial should be rotated in that direction
numbers are 0 - 99
dial is a full circle so after 99 comes 0 again
 ---> dial starts at 50 <---
actual password -> number of times the dial reaches zero after 0 after any
rotation sequence
*/

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

		if direction == 'L' {
			pos = (pos - distance + 100) % 100

			if pos == 0 {
				atZero++
			}

		} else {
			pos = (pos + distance) % 100
			if pos == 0 {
				atZero++
			}
		}
	}
	fmt.Println("final position:", pos)
	fmt.Println("number of times we reach zero:", atZero)

}
