package main

import (
	"fmt"

	teams "github.com/dmh2000/teams/teams"
)

// set this to the address clients will connect to
const serveradr = "172.31.189.50:8080"

// set this to the connection string for your mongodb instance
const mongouri = "mongodb://172.31.176.1:27017"

// set this to the database name that will be created on your mongodb instance
const database = "bb-teams"

// set this to thenstance ofname  teams
func printTeam(t teams.Team) {
	fmt.Println("    ID       : ", t.ID)
	fmt.Println("    Name     : ", t.Name)
	fmt.Println("    Location : ", t.Location)
	fmt.Println("    Year     : ",t.Year)
	fmt.Println("    Wins     : ", t.Wins)
	fmt.Println("    Losses   : ", t.Losses)
	fmt.Println("    Ties     : ", t.Ties)
	fmt.Println("    Other    : ", t.Other)
	fmt.Println("    Games    : ", t.Games)

	fmt.Println();
}    

func main() {
	// drop the existing database so it starts clean
	// make sure the database name above is ok to drop!
	fmt.Println("Drop Database")
	teams.DropDatabase(mongouri, database)

	// ======================================
	// 1. Read Json File
	// ======================================
	fmt.Println("Read JSON")
	tg, err := teams.ReadTeams("teams.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// ======================================
	// 2. Load it into Mongodb
	// ======================================
	fmt.Println("Upload to Mongodb")
	err = teams.LoadTeams(mongouri, database, tg)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ======================================
	// 3. Make sure the database works
	// ======================================
	fmt.Println("Test Queries")
	// query the database for all documents
	r, err := teams.QueryTeams(mongouri, database, "", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	// print len, should be 150 entries
	fmt.Print("    Teams    : ",len(r),"\n\n")

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
	fmt.Printf("Database : %s\n", database)
	fmt.Printf("Mongodb  : %s\n", mongouri)
	fmt.Printf("GraphQL  : %s\n", serveradr)
	fmt.Println("Listening...")

	teams.Serve(serveradr, mongouri, database)
}
