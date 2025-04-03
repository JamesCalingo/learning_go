package exercises

import (
	"fmt"
	"strings"
	"sort"
)

// EXERCISE 1: creating the structs
// NOTE: I added an extra struct for players to hold their first and last names as separate strings - it's probably not necessary in this case but a bit closer to what actual databases likely have
type Player struct {
	FirstName string
	LastName  string
}

type Team struct {
	Name   string
	Roster []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

type kv struct {
	key string
	Value int
}

// EXERCISE 2: methods
func (l *League) MatchResult(t1 Team, t2 Team, p1 int, p2 int, wins map[string]int) {
	if strings.EqualFold(t1.Name, t2.Name) {
		fmt.Println("Is this a scrimmage? Both teams are the same")
		return
	}
	if p1 > p2 {
		wins[t1.Name] = wins[t1.Name] + 1
	} else {
		wins[t2.Name] = wins[t2.Name] + 1
	}
}

func (l League) Ranking() []string {
	var arranged []kv
	for k, v := range l.Wins{
		arranged = append(arranged, kv{k, v})
	}
	sort.Slice(arranged, func(i, j int) bool {
		return arranged[i].Value > arranged[j].Value
	})
	fmt.Println(arranged)
	var standings []string
	for _, team := range arranged{
		standings = append(standings, team.key)
	}
	return standings
}

func Ch7() {
	players := []string{"1", "2", "3", "4", "5"}
	teamA := Team{"Annihilators", players}
	teamB := Team{"Behemoths", players}
	teamC := Team{"Crushers", players}
	teamD := Team{"Destroyers", players}
	teams := []Team{teamA, teamB, teamC, teamD}
	wins := map[string]int{
		"Annihilators": 0,
		"Behemoths":    0,
		"Crushers":     0,
		"Destroyers":   0,
	}
	//SOMETHING TO NOTE: There's an argument that you don't need to initialize your map with the team names, as they'll populate when they win. This does mean, however, that if a team goes winless, they'll never show up in the map...
	theLeague := League{teams, wins}
	theLeague.MatchResult(teamA, teamB, 28, 3, wins)
	theLeague.MatchResult(teamC, teamD, 34, 28, wins)
	theLeague.MatchResult(teamA, teamD, 15, 0, wins)
	theLeague.MatchResult(teamC, teamB, 10, 22, wins)
	theLeague.MatchResult(teamC, teamA, 7, 21, wins)
	theLeague.MatchResult(teamD, teamB, 11, 14, wins)
	theLeague.MatchResult(teamC, teamB, 3, 0, wins)
	theLeague.MatchResult(teamC, teamD, 3, 0, wins)
	fmt.Println(theLeague.Ranking())
}
