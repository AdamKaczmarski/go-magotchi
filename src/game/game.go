package game

import (
	"fmt"
	"math/rand"
)

const higher, lower, maxRounds = "H", "L", 5

func Game() (score int) {
	pick1, pick2 := -1, -1
	var userGuess string
	fmt.Printf("Please answer either \"%v\" for higher or \"%v\" for lower\n", higher, lower)
	for i := 0; i < maxRounds; i++ {
		pick1, pick2 = rand.Intn(10), rand.Intn(10)
		for pick1 == pick2 {
			pick1, pick2 = rand.Intn(10), rand.Intn(10)
		}
		getUsersGuess(pick1, &userGuess)
		fmt.Printf("The second number is: %d\n", pick2)
		if validateGuess(pick1, pick2, userGuess) {
			score++
			fmt.Println(":)")
		} else {
			fmt.Println(":(")
		}
	}
	fmt.Printf("Score: %d/%d\n", score, maxRounds)
	return score
}

func validateGuess(pick1, pick2 int, userGuess string) bool {
	return (userGuess == higher && pick2 > pick1) || (userGuess == lower && pick2 < pick1)
}

func getUsersGuess(pick1 int, userGuess *string) {
	fmt.Printf("Higher or lower than %d?\n", pick1)
	for {
		fmt.Scan(&*userGuess)
		if *userGuess == higher || *userGuess == lower {
			break
		}
		fmt.Printf("Please answer either \"%v\" for higher or \"%v\" for lower\n", higher, lower)
	}
}
