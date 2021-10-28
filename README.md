# round robin tournament

This is a golang module which will provide functions to generate a 2d slice of string containing matches of given teams.
The matches will be generated using the [round-robin tournament](https://en.wikipedia.org/wiki/Round-robin_tournament).
Every team will play against every other team.

## example

The following golang function call will create the 2d slice of strings above.
```go
GenerateRoundRobinTournamentMatchesByNumber(4)
```

```
[
  [Team 1 Team 2]
  [Team 3 Team 4]
  [Team 1 Team 4]
  [Team 2 Team 3]
  [Team 1 Team 3]
  [Team 4 Team 2]
]
```

You will find other examples in [example](example) directory.
