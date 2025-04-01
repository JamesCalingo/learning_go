package exercises

import "fmt"

//EXERCISE 1: creating the structs
//NOTE: I added an extra struct for players to hold their first and last names as separate strings - it's probably not necessary in this case but a bit closer to what actual databases likely have
type Player struct {
	FirstName string
	LastName string
}

type Team struct {
	Name string
	Roster []string
}

type League struct {
	Teams []Team
	Wins map[string]int
}

//EXERCISE 2: methods
func (l *League) MatchResult (t1 string, t2 string, p1 int, p2 int) {
	if p1 > p2 {
		fmt.Println("TEAM ONE WINS")
	}else {
		fmt.Println("TEAM TWO WINS")
	}
}

func (l League) Ranking() []string {
	var standings []string
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
	"Behemoths": 0,
	"Crushers":0,
	"Destroyers":0,
}
theLeague := League{teams, wins}
theLeague.MatchResult("Annihilators", "Behemoths", 28, 3)
}