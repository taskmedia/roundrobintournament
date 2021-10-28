package roundrobintournament

import (
	"github.com/google/uuid"
)

// GenerateRoundRobinTournamentMatches generates a 2d slice of matches of a single round robin tournament.
// Each team will play one time against all other teams.
func GenerateRoundRobinTournamentMatches(teams []string) [][]string {
	matches := make([][]string, 0)

	dummy := "even"
	if len(teams)%2 != 0 {
		dummy = uuid.New().String()
		teams = append(teams, dummy)
	}

	rotation := make([]string, len(teams))
	copy(rotation, teams)

	for i := 0; i < (len(teams) - 1); i++ {
		rotationLen := len(rotation)
		for i := 0; i < len(rotation); i = i + 2 {
			matches = append(matches, []string{rotation[i], rotation[i+1]})
		}

		// rotate teams for next round
		rotationHelper := append([]string{}, rotation[0])                     // append first team
		rotationHelper = append(rotationHelper, rotation[rotationLen-1])      // append last team
		rotationHelper = append(rotationHelper, rotation[1:rotationLen-1]...) // append remaining teams

		rotation = rotationHelper
	}

	// remove dummy matches
	if dummy != "even" {
		i := 0
		for _, x := range matches {
			if !stringSlicecontains(x, dummy) {
				matches[i] = x
				i++
			}
		}
		matches = matches[:i]
	}

	return matches
}

// Iterate over slice of string  to check whether it an element or not
func stringSlicecontains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
