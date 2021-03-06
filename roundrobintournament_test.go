package roundrobintournament

import (
	"fmt"
	"testing"
)

// testing odd teams number
func TestGenerateRoundRobinTournamentMatchesByTeams3Teams(t *testing.T) {
	teams := []string{
		"TeamA",
		"TeamB",
		"TeamC",
	}

	matches := GenerateRoundRobinTournamentMatchesByTeams(teams)

	expectedLenMatches := expectedMatches(len(teams))

	if len(matches) != expectedLenMatches {
		t.Fatalf("Length of matches (%d) not as expected (%d)", len(matches), expectedLenMatches)
	}

	expectedmatches := [][]string{
		{"TeamA", "TeamB"},
		{"TeamB", "TeamC"},
		{"TeamA", "TeamC"},
	}

	if !slicesEqual(expectedmatches, matches) {
		t.Fatalf("Expected matches for a three teams tournament was not satisfied:\nexpected:\n%s\n\ngot:\n%s\n", expectedmatches, matches)
	}
}

// testing even teams number
func TestGenerateRoundRobinTournamentMatchesByTeams4Teams(t *testing.T) {
	teams := []string{
		"TeamA",
		"TeamB",
		"TeamC",
		"TeamD",
	}

	matches := GenerateRoundRobinTournamentMatchesByTeams(teams)

	expectedLenMatches := expectedMatches(len(teams))

	if len(matches) != expectedLenMatches {
		t.Fatalf("Length of matches (%d) not as expected (%d)", len(matches), expectedLenMatches)
	}

	expectedmatches := [][]string{
		{"TeamA", "TeamB"},
		{"TeamC", "TeamD"},
		{"TeamA", "TeamD"},
		{"TeamB", "TeamC"},
		{"TeamA", "TeamC"},
		{"TeamD", "TeamB"},
	}

	if !slicesEqual(expectedmatches, matches) {
		t.Fatalf("Expected matches for a four teams tournament was not satisfied:\nexpected:\n%s\n\ngot:\n%s\n", expectedmatches, matches)
	}
}

// testing large number of even teams
func TestGenerateRoundRobinTournamentMatchesByTeams10Teams(t *testing.T) {
	teams := make([]string, 0)
	for i := 1; i <= 10; i++ {
		teams = append(teams, fmt.Sprintf("Team %d", i))
	}

	matches := GenerateRoundRobinTournamentMatchesByTeams(teams)

	expectedLenMatches := expectedMatches(len(teams))

	if len(matches) != expectedLenMatches {
		t.Fatalf("Length of matches (%d) not as expected (%d)", len(matches), expectedLenMatches)
	}
}

// testing large number of odd teams
func TestGenerateRoundRobinTournamentMatchesByTeams11Teams(t *testing.T) {
	teams := make([]string, 0)
	for i := 1; i <= 11; i++ {
		teams = append(teams, fmt.Sprintf("Team %d", i))
	}

	matches := GenerateRoundRobinTournamentMatchesByTeams(teams)

	expectedLenMatches := expectedMatches(len(teams))

	if len(matches) != expectedLenMatches {
		t.Fatalf("Length of matches (%d) not as expected (%d)", len(matches), expectedLenMatches)
	}
}

func TestGenerateRoundRobinTournamentMatchesByNumberEven(t *testing.T) {
	teamSize := 14

	matches := GenerateRoundRobinTournamentMatchesByNumber(teamSize)

	expectedLenMatches := expectedMatches(teamSize)

	if len(matches) != expectedLenMatches {
		t.Fatalf("Length of matches (%d) not as expected (%d)", len(matches), expectedLenMatches)
	}
}

func TestGenerateRoundRobinTournamentMatchesByNumberOdd(t *testing.T) {
	teamSize := 15

	matches := GenerateRoundRobinTournamentMatchesByNumber(teamSize)

	expectedLenMatches := expectedMatches(teamSize)

	if len(matches) != expectedLenMatches {
		t.Fatalf("Length of matches (%d) not as expected (%d)", len(matches), expectedLenMatches)
	}

}

// calculates the expected number of matches by a given team length
func expectedMatches(lenTeams int) int {
	return lenTeams * (lenTeams - 1) / 2
}

// compares len and every item of two slices
func slicesEqual(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		for k := range a[i] {
			if a[i][k] != b[i][k] {
				return false
			}
		}
	}
	return true
}
