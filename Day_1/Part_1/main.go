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
	currentPlace := FIFTY

	for _, mv := range moves {

		if mv.Dir == "R" {
			currentPlace = Add(currentPlace, Number(mv.Steps))
		} else {
			currentPlace = Subtract(currentPlace, Number(mv.Steps))
		}

		if currentPlace == ZERO {
			clicks++
		}
	}

	fmt.Printf("%d", clicks)
}

func Add(a, b Number) Number {
	return Number(int(a)+int(b)) % 100
}

func Subtract(a, b Number) Number {
	return Number(int(a)-int(b)) % 100
}

type Number int

const (
	ZERO Number = iota
	ONE
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	ELEVEN
	TWELVE
	THIRTEEN
	FOURTEEN
	FIFTEEN
	SIXTEEN
	SEVENTEEN
	EIGHTEEN
	NINETEEN
	TWENTY
	TWENTY_ONE
	TWENTY_TWO
	TWENTY_THREE
	TWENTY_FOUR
	TWENTY_FIVE
	TWENTY_SIX
	TWENTY_SEVEN
	TWENTY_EIGHT
	TWENTY_NINE
	THIRTY
	THIRTY_ONE
	THIRTY_TWO
	THIRTY_THREE
	THIRTY_FOUR
	THIRTY_FIVE
	THIRTY_SIX
	THIRTY_SEVEN
	THIRTY_EIGHT
	THIRTY_NINE
	FORTY
	FORTY_ONE
	FORTY_TWO
	FORTY_THREE
	FORTY_FOUR
	FORTY_FIVE
	FORTY_SIX
	FORTY_SEVEN
	FORTY_EIGHT
	FORTY_NINE
	FIFTY
	FIFTY_ONE
	FIFTY_TWO
	FIFTY_THREE
	FIFTY_FOUR
	FIFTY_FIVE
	FIFTY_SIX
	FIFTY_SEVEN
	FIFTY_EIGHT
	FIFTY_NINE
	SIXTY
	SIXTY_ONE
	SIXTY_TWO
	SIXTY_THREE
	SIXTY_FOUR
	SIXTY_FIVE
	SIXTY_SIX
	SIXTY_SEVEN
	SIXTY_EIGHT
	SIXTY_NINE
	SEVENTY
	SEVENTY_ONE
	SEVENTY_TWO
	SEVENTY_THREE
	SEVENTY_FOUR
	SEVENTY_FIVE
	SEVENTY_SIX
	SEVENTY_SEVEN
	SEVENTY_EIGHT
	SEVENTY_NINE
	EIGHTY
	EIGHTY_ONE
	EIGHTY_TWO
	EIGHTY_THREE
	EIGHTY_FOUR
	EIGHTY_FIVE
	EIGHTY_SIX
	EIGHTY_SEVEN
	EIGHTY_EIGHT
	EIGHTY_NINE
	NINETY
	NINETY_ONE
	NINETY_TWO
	NINETY_THREE
	NINETY_FOUR
	NINETY_FIVE
	NINETY_SIX
	NINETY_SEVEN
	NINETY_EIGHT
	NINETY_NINE
)
