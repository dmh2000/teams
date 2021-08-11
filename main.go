package main

import (
	"fmt"

	teams "github.com/dmh2000/teams/teams"
)

const mongdbUrl = "mongodb://172.17.48.1:27017"

// WARNING : THIS CODE WILL DELETE ANY DATABASE WITH THIS NAME
// CHANGE IT AS NEEDED SO THERE IS NO CONFLICT WITH ANY EXISTING DATABASE
const database = "bb-team-games"

// print one instance of teams
func printTeam(t teams.Teams) {
	fmt.Println("ID    :", t.ID)
	fmt.Println("Name  :", t.Name)
	fmt.Println("Wins  :", t.Wins)
	fmt.Println("Losses:", t.Losses)
	fmt.Println("Ties  :", t.Ties)
	fmt.Println("Other :", t.Other)
}

func main() {
	fmt.Println("Populating Database")

	// drop the existing database so it starts clean
	teams.DropDatabase(mongdbUrl, database)

	// read the JSON file, returns a slice of teams
	tg, err := teams.ReadTeams("team-games.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	// load into the database (be sure to drop it first or you will get duplicates	)
	err = teams.LoadTeams(mongdbUrl, database, tg)
	if err != nil {
		fmt.Println(err)
		return
	}

	// note when querying : in mongodb the field names are all lower case
	// in the Go context they are according to the struct definition
	// which is title case

	// query the database for all documents
	r, err := teams.QueryTeams(database, mongdbUrl, "", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	// should return 151
	fmt.Println(len(r))

	// query for a specific document
	r, err = teams.QueryTeams(database, mongdbUrl, "id", "WAS")
	if err != nil {
		fmt.Println(err)
		return
	}
	printTeam(r[0])
}
