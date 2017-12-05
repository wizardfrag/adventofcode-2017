package main

import (
	"bufio"
	"os"
	"strconv"
	"log"
)

func main() {
	var err error
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var sum int
		var first int = -1
		var current int
		var prev int
		for _, rune := range s.Text() {
			if current, err = strconv.Atoi(string(rune)); err == nil {
				if first == -1 {
					first = current
				}
				log.Printf("current [%+v] rune [%+v]", current, rune)
				if current == prev {
					sum += prev
				}
				prev = current
			} else {
				log.Fatalf("invalid rune", rune)
			}
		}
		if current == first {
			sum += current
		}

		log.Printf("SUM: %+v", sum)
	}
}