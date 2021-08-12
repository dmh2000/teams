package main

import (
	"fmt"

	teams "github.com/dmh2000/teams/teams"
)

// set this to the address clients will connect to
const serveradr = "172.17.59.222:8080"

// set this to the connection string for your mongodb instance
const mongouri = "mongodb://172.17.48.1:27017"

// set this to the database name that will be created on your mongodb instance
const database = "bb-team-games"

// set this to thenstance ofname  teams
func printTeam(t teams.Team) {
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
	// make sure the database name above is ok to drop!
	teams.DropDatabase(mongouri, database)

	// ======================================
	// 1. Read Json File
	// ======================================
	tg, err := teams.ReadTeams("teams.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// ======================================
	// 2. Load it into Mongodb
	// ======================================
	err = teams.LoadTeams(mongouri, database, tg)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ======================================
	// 3. Make sure the database works
	// ======================================
	// query the database for all documents
	r, err := teams.QueryTeams(mongouri, database, "", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	// should return 150
	fmt.Println(len(r))

	// query mongodb for a specific document
	// note when querying in mongodb the field names are all lower case
	// In the Go context they are according to the struct definition which is title case
	r, err = teams.QueryTeams(mongouri, database, "id", "WAS")
	if err != nil {
		fmt.Println(err)
		return
	}
	printTeam(r[0])

	// query mongodb for a specific document
	// note when querying in mongodb the field names are all lower case
	// In the Go context they are according to the struct definition which is title case
	r, err = teams.QueryTeams(mongouri, database, "name", "Nationals")
	if err != nil {
		fmt.Println(err)
		return
	}
	printTeam(r[0])

	// ======================================
	// GraphQL Server
	// ======================================
	teams.Serve(serveradr, mongouri, database)
}
