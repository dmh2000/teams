package main

import (
	"fmt"

	"github.com/dmh2000/teamgames/teamgames"
)

const mongdbUrl = "mongodb://172.17.48.1:27017"
const database  = "baseball"

func printTeam(t teamgames.TeamGames) {
	fmt.Println("ID    :",t.ID)
	fmt.Println("Name  :",t.Name)
	fmt.Println("Wins  :",t.Wins)
	fmt.Println("Losses:",t.Losses)
	fmt.Println("Ties  :",t.Ties)
	fmt.Println("Other :",t.Other)

}

func main() {
	// // read the JSON file, returns a slice of TeamGames
	// tg, err := teamgames.ReadTeamGames("team-games.json")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // load into the database (be sure to drop it first or you will get duplicates	)
	// err = teamgames.LoadTeamGames(database,mongdbUrl,tg)
	// if err != nil {
	// 	fmt.Println(err);
	// 	return;
	// }

	// note when querying : in mongodb the field names are all lower case
	// in the Go context they are according to the struct definition
	// which is title case

	// query the database for all documents
	r, err := teamgames.QueryTeamGames(database,mongdbUrl,"","")
	if err != nil {
		fmt.Println(err)
		return;
	}
	// should return 151
	fmt.Println(len(r))

	// query for a specific document
	r, err = teamgames.QueryTeamGames(database,mongdbUrl,"id","WAS")
	if err != nil {
		fmt.Println(err)
		return;
	}
	printTeam(r[0])
}
