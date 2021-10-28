package main

import (
	"fmt"

	"github.com/taskmedia/roundrobintournament"
)

func main() {
	matches := roundrobintournament.GenerateRoundRobinTournamentMatchesByNumber(4)

	for i, match := range matches {
		fmt.Printf("match #%d: %s : %s\n", i+1, match[0], match[1])
	}
}
