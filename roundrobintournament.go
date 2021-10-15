package roundrobintournament

// GenerateRoundRobinTournamentMatches generates a 2d slice of matches of a single round robin tournament.
// Each team will play one time against all other teams.
func GenerateRoundRobinTournamentMatches(teams []string) [][]string {
	matches := make([][]string, 0)

	rotation := make([]string, len(teams))
	copy(rotation, teams)

	var cachedTeam string

	for i := 0; i < (len(teams) - 1); i++ {
		rotationLen := len(rotation)

		if len(teams)%2 == 0 {
			// number of teams is even - cachedTeam is not required
			for i := 0; i < len(rotation); i = i + 2 {
				matches = append(matches, []string{rotation[i], rotation[i+1]})
			}
		} else {
			// number of teams is odd - cachedTeam is required
			for i := 0; i < len(rotation)-1; i = i + 2 {
				matches = append(matches, []string{rotation[i], rotation[i+1]})
				if cachedTeam != "" {
					matches = append(matches, []string{cachedTeam, rotation[i+2]})
					cachedTeam = ""
				} else {
					cachedTeam = rotation[i+2]
				}
			}
		}

		// rotate teams for next round
		rotationHelper := append([]string{}, rotation[0])                     // append first team
		rotationHelper = append(rotationHelper, rotation[rotationLen-1])      // append last team
		rotationHelper = append(rotationHelper, rotation[1:rotationLen-1]...) // append remaining teams

		rotation = rotationHelper
	}

	return matches
}
