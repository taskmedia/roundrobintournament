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

---

[![releases](https://img.shields.io/github/v/release/taskmedia/roundrobintournament?style=flat-square)](https://github.com/taskmedia/roundrobintournament/releases/latest)
[![docs](https://img.shields.io/badge/docs-pkg.go.dev-blue?style=flat-square)](https://pkg.go.dev/github.com/taskmedia/roundrobintournament)
[![golang version](https://img.shields.io/github/go-mod/go-version/taskmedia/roundrobintournament?style=flat-square)](https://golang.org/dl/#stable)
<br />
[![MIT license](https://img.shields.io/github/license/taskmedia/roundrobintournament?style=flat-square)](https://github.com/taskmedia/roundrobintournament/blob/main/LICENSE)
[![codecoverage](https://img.shields.io/codecov/c/github/taskmedia/roundrobintournament?style=flat-square)](https://app.codecov.io/gh/taskmedia/roundrobintournament)
![code size](https://img.shields.io/github/languages/code-size/taskmedia/roundrobintournament?style=flat-square)
<br />
[![issues](https://img.shields.io/github/issues/taskmedia/roundrobintournament?style=flat-square)](https://github.com/taskmedia/roundrobintournament/issues)
[![pull requests](https://img.shields.io/github/issues-pr/taskmedia/roundrobintournament?style=flat-square)](https://github.com/taskmedia/roundrobintournament/pulls)
<br />
[![twitter](https://img.shields.io/twitter/follow/taskmediaDE?style=social)](https://twitter.com/taskmediaDE)
