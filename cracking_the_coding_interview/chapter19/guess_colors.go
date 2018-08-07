package main

/*
The Game of Master Mind is played as follows:
    The computer has four slots containing balls that are red(R), yellow(Y), green(G) or blue(B).
    For example, the computer might have RGGB.

    You, the user, are trying to guess the solution. You might, for example, guess YRGB.
    When you guess the correct color for the correct slot, you get a "hit".
    If you guess a color that exists but is in the wrong slot, you get a "pseudo-hit".
    For example, the guess "YRGB" has 2 hits and 1 pseudo hit.

    For each guess, you are told the number of hits and pseudo-hits.
    Write a method that, given a guess and a solution, returns the number of hits and pseudo-hits.
*/
import (
	"fmt"
)

type Result struct {
	Hits       int
	PseudoHits int
}

func Estimate(guess, solution string) *Result {
	r := Result{Hits: 0, PseudoHits: 0}
	var solutionMask = 0
	for i := 0; i < 4; i++ {
		solutionMask |= 1 << (1 + uint(solution[i]) - uint('A'))
	}

	fmt.Println(solutionMask)

	for i := 0; i < 4; i++ {
		if guess[i] == solution[i] {
			r.Hits++
		} else if solutionMask&(1<<(1+uint(guess[i])-uint('A'))) >= 1 {
			r.PseudoHits++
		}
	}
	return &r
}

func main() {
	guess, solution := "YRGB", "RGGB"
	fmt.Println(Estimate(guess, solution))
}
