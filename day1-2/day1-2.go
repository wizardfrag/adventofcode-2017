package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"log"
)

func main() {
	var err error
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var digits []int
		var half int
		var length int
		var sum int

		length = len(strings.Trim(s.Text(), "\r\n"))

		digits = make([]int, length)
		half = length / 2

		for i, r := range s.Text() {
			digits[i], err = strconv.Atoi(string(r))
			if err != nil {
				log.Panic("invalid rune", r)
			}
		}

		log.Printf("length [%+v] half [%+v]", length, half)

		for i, digit := range digits {
			j := i + half

			if j >= length {
				j -= length
			}
			if digits[j] == digit {
				sum += digit
			}
		}

		log.Printf("sum [%+v]", sum)
	}
}
