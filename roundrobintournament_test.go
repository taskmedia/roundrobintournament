package roundrobintournament

import (
	"testing"
)

// testing odd teams number
func TestGenerateRoundRobinTournamentMatches3Teams(t *testing.T) {
	teams := []string{
		"TeamA",
		"TeamB",
		"TeamC",
	}

	matches := GenerateRoundRobinTournamentMatches(teams)

	expectedLenMatches := len(teams) * (len(teams) - 1) / 2

	if len(matches) != expectedLenMatches {
		t.Fatalf("Length of matches (%d) not as expected (%d)", len(matches), expectedLenMatches)
	}

	expectedmatches := [][]string{
		{"TeamA", "TeamB"},
		{"TeamA", "TeamC"},
		{"TeamC", "TeamB"},
	}

	if !slicesEqual(expectedmatches, matches) {
		t.Fatalf("Expected matches for a three teams tournament was not satisfied:\nexpected:\n%s\n\ngot:\n%s\n", expectedmatches, matches)
	}
}

// testing even teams number
func TestGenerateRoundRobinTournamentMatches4Teams(t *testing.T) {
	teams := []string{
		"TeamA",
		"TeamB",
		"TeamC",
		"TeamD",
	}

	matches := GenerateRoundRobinTournamentMatches(teams)

	expectedLenMatches := len(teams) * (len(teams) - 1) / 2

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
