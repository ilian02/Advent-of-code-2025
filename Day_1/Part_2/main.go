package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Move struct {
	Dir   string
	Steps int
}

func main() {
	filename := "input.txt"

	var f, err = os.Open(filename)

	if err != nil {
		fmt.Println("Error opening input.txt")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)
	re := regexp.MustCompile(`^\s*([RL])\s*[, ]?\s*(\d+)\s*$`)

	var moves []Move

	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		m := re.FindStringSubmatch(line)
		if m == nil {
			fmt.Fprintf(os.Stderr, "Could not parse line")
			continue
		}
		steps, _ := strconv.Atoi(m[2])
		moves = append(moves, Move{Dir: m[1], Steps: steps})
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Read errpr")
		os.Exit(1)
	}

	clicks := 0
	currentPlace := 50

	for _, mv := range moves {
		if mv.Dir == "R" {
			for i := 0; i < mv.Steps; i++ {
				currentPlace = (currentPlace + 1) % 100
				if currentPlace == 0 {
					clicks++
				}
			}
		} else { // "L"
			for i := 0; i < mv.Steps; i++ {
				currentPlace = (currentPlace - 1 + 100) % 100
				if currentPlace == 0 {
					clicks++
				}
			}
		}
	}

	fmt.Printf("%d", clicks)
}
