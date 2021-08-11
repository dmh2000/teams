package main

import (
	"fmt"

	"github.com/dmh2000/teamgames/teamgames"
)

func main() {
	err := teamgames.MergeTeams("teams.json","team-games.json")

	if err != nil {
		fmt.Println(err)
		return
	}
}
