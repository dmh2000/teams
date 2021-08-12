package teams

import (
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/rs/cors"
)

// ===============================
// Schema
// ===============================
const schemaDef = `
schema {
	query:Query
}

type Query {
	"""
	Get data about Teams
	There can be multiple teams that satisfy these queries
	The team data stores multiple instances of particular teams
	because they moved or changed on certain dates
	"""
	teamsByID(id:String!):[Team]
	teamsByName(name:String!):[Team]
	teamsAll():[Team]
}

"""
This must match the Team type except inu all lower case
"""
type Team {
	id: String!
	name : String!,
	location: String!,
	year: String!,
	wins: Int!,
	losses: Int!,
	ties: Int!,
	other: Int!,
	games: Int!,
	uuid: String!
}
`

// ===============================
// Team Resolver
// ===============================
/*
There must be one Resolver function for each Field in the Team Struct
type Team struct {
	ID     string `json:"ID"`
	Name   string `json:"name"`
	Location   string `json:"location"`
	Year   string `json:"year"`
	Wins   int    `json:"wins"`
	Losses int    `json:"losses"`
	Ties   int    `json:"ties"`
	Other  int    `json:"other"`
	Games  int    `json:"games"`
	UUID   string `json:"uuid"`
}
*/
type teamResolver struct {
	t *Team
}

func (r *teamResolver) ID() string {
	return r.t.ID
}

func (r *teamResolver) Name() string {
	return r.t.Name
}

func (r *teamResolver) Location() string {
	return r.t.Location
}

func (r *teamResolver) Year() string {
	return r.t.Year
}

func (r *teamResolver) Wins() int32 {
	return int32(r.t.Wins)
}

func (r *teamResolver) Losses() int32 {
	return int32(r.t.Losses)
}

func (r *teamResolver) Ties() int32 {
	return int32(r.t.Ties)
}

func (r *teamResolver) Other() int32 {
	return int32(r.t.Other)
}

func (r *teamResolver) Games() int32 {
	return int32(r.t.Games)
}

func (r *teamResolver) UUID() string {
	return r.t.UUID
}

// ===============================
// ROOT RESOLVER
// ===============================

type rootResolver struct{}

// get team by abbreviation
func (*rootResolver) TeamsByID(args struct{ ID string }) *[]*teamResolver {
	// query the db for the team by the abbreviation
	// remember, in mongodb the keys are lower case
	teams, err := QueryTeams(mongodb, database, "id", args.ID)
	if err != nil {
		teams = []Team{}
	}

	// resolve the results
	var tr []*teamResolver
	for i := range teams {
		tr = append(tr, &teamResolver{&teams[i]})
	}

	return &tr
}

func (*rootResolver) TeamsByName(args struct{ Name string }) *[]*teamResolver {
	// query the db for the team by the abbreviation
	// remember, in mongodb the keys are lower case
	teams, err := QueryTeams(mongodb, database, "name", args.Name)
	if err != nil {
		teams = []Team{}
	}

	// this is ok because it gives a pointer to elements of the slice
	var tr []*teamResolver = make([]*teamResolver, len(teams))
	for i := range teams {
		tr[i] = &teamResolver{&teams[i]}
	}

	return &tr
}

func (*rootResolver) TeamsAll() *[]*teamResolver {
	// query the db for the team by the abbreviation
	teams, err := QueryTeams(mongodb, database, "", "")
	if err != nil {
		teams = []Team{}
	}

	// resolve the results
	var tr []*teamResolver
	for i := range teams {
		tr = append(tr, &teamResolver{&teams[i]})
	}

	return &tr
}

// RootSchema return the parsed schema
func RootSchema() *graphql.Schema {
	return graphql.MustParseSchema(schemaDef, &rootResolver{})
}

// the main program must set these before running the server
var mongodb string
var database string

// Run : start the GraphQL server
func Serve(portip string, mongo string, db string) {
	// init the references the resolver uses
	mongodb = mongo
	database = db

	// use cors middleware in case queries come from different domain
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	handler := c.Handler(&relay.Handler{Schema: RootSchema()})
	http.Handle("/query", handler)
	log.Fatal(http.ListenAndServe(portip, nil))
}
