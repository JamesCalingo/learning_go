package exercises

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
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
		wins[t1.Name]++
	} else {
		wins[t2.Name]++
	}
}

//NOTE: you'll note that this does not use a pointer unlike MatchResult - that's because we're not actually modifying theLeague; we just need its current data
func (l League) Ranking() []string {
	var arranged []kv
	for k, v := range l.Wins{
		arranged = append(arranged, kv{k, v})
	}
	sort.Slice(arranged, func(i, j int) bool {
		return arranged[i].Value > arranged[j].Value
	})
	var standings []string
	for _, team := range arranged{
		standings = append(standings, team.key)
	}
	return standings
}

//EXERCISE 3: Ranker
type Ranker interface {
	Ranking() []string
}

func RankPrinter (r Ranker, w io.Writer) {
	for _, value := range r.Ranking(){
		io.WriteString(w, value)
		w.Write([]byte("\n"))
	}
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
	theLeague.MatchResult(teamA, teamC, 28, 24, wins)
	fmt.Println(theLeague)
	RankPrinter(theLeague, os.Stdout)
	/* I feel the need to comment on this, because WHAT THE HELL IS GOING ON HERE? Why are we running RankPrinter on theLeague, and how?

	The thing with interfaces is that they hold METHODS and that ANYTHING THAT USES THE METHODS ON AN INTERFACE MATCHES TYPE. Since we have a "Ranking" method on our interface AND a "Ranking" method on our League struct, League meets the Ranker interface on that Ranking method and thus fits (remember, you'll get an error if the type doesn't match).

	When we run RankPrinter, it uses theLeague.Ranking to execute without us having to explicitly type theLeague.Ranking. This is useful as it allows us to "reuse" this method for other sports that don't use a simple binary win-loss method for ranking participants (for example, point based standings for sports with draws/ties or things like racing).
	*/
}
