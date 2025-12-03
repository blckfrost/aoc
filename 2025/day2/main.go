package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	password := 0
	password2 := 0

	file, err := os.ReadFile("day2_input.txt")
	checkErr(err)

	text := string(file)
	ids := strings.Split(strings.TrimSpace(text), ",")

	for _, id := range ids {
		parts := strings.Split(id, "-")

		start, err := strconv.Atoi(parts[0])
		checkErr(err)

		end, err := strconv.Atoi(parts[1])
		checkErr(err)

		for n := start; n <= end; n++ {
			if isInvalid(n) {
				password += n
			}
			if isInvalid2(n) {
				password2 += n
			}

		}
	}
	fmt.Println("Password:", password)
	fmt.Println("Password 2:", password2)
}

func isInvalid(n int) bool {
	s := strconv.Itoa(n)
	l := len(s)
	if l%2 != 0 {
		return false
	}
	half := l / 2
	return s[:half] == s[half:]
}

/*
Repeating 11111
Repeating 1010
Repeating 123123123
Repeating 12341234
*/

func isInvalid2(n int) bool {
	s := strconv.Itoa(n)
	l := len(s)

	// Try every possible chunk size
	for p := 1; p <= l/2; p++ {
		if l%p != 0 {
			continue // the chunk must evenly divide the number
		}

		chunk := s[:p]
		valid := true

		for i := 0; i < l; i += p {
			if s[i:i+p] != chunk {
				valid = false
				break
			}
		}

		if valid {
			return true
		}
	}
	return false
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}

}
