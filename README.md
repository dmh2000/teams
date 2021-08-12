# Simple Go : Json -> Mongodb -> GraphQL -> Client

## Overview

The point of this repo is to provide, in _simple, minimal_ code, an example in Go, of reading a JSON file, uploading it to a Mongodb database, and serving the data using GraphQL.

There's a bit of learning curve if you are starting something like this from scratch, expecially with Mongodb and the GraphQL server. There are other examples but I haven't seen an example in one place that has all the pieces.

The data used here is derived from the data products provided by [Retrosheet.org](retrosheet.org). Retrosheet does a lot of work to provide historical baseball data going all the way back to 1871. The data used here is a very small subset of what they provide.

Thanks to Retrosheet for making this data available.

<pre>
     The information used here was obtained free of
     charge from and is copyrighted by Retrosheet.  Interested
     parties may contact Retrosheet at "www.retrosheet.org". 
</pre>

## Main Program

The main program runs through all the steps to get this going. At each step you can just look at the function being called, find its implementation in the 'teams' directory, and see what it does.

### Setup

- have an instance of Mongodb running somewhere you can connect to, and get its connection string.
- at the top of _main.go_
  - set the mongodb connection string (_mongouri_)
  - set the ip address and port (_serveradr_) that the GraphQL server will use. This is the address that clients will connect to for GrapQL Queries
  - set the database name (_database_). It defaults to _bb-team-games_, but make sure it unique to your system so you don't step on something important.
  - if you are running mongodb and this app on the same system, you can probably use _localhost_ for the IP addresses.
- this required Go 1.16 or later due to changes in ioutil

## 1. Read Json File

- a JSON file _teams.json_ is included
- _teams/read_teams.go_
  - function _read_teams_
  - read the JSON file
  - using encoding/json, unmarshal it into Go objects

### The Data

- at this point you have a slice of Team objects

  - there can and will be multiple teams for most Names because the data goes back to 1871 and it differentiates teams when they move cities
  - there is a whole lot of data in Retrosheet, but I boiled this down to a few fields to use as an example.

  - type definition of _Team_

```Go
type Team struct {
	ID       string     // Abbreviation of team name from Retrosheet data
	Name     string     // Name of the team
	Location string     // city or state
	Year     string     // year the franchise started
	Wins     int        // total wins
	Losses   int        // total losses
	Ties     int        // total ties (yes there can be ties!)
	Other    int        // total rainouts and other suspensions
	Games    int        // sum of wins,losses,ties and other
	UUID     string     // a unique ID I added as a key
}
```

## 2. Load it into Mongodb

- _teams/load_teams.go_
  - function _DropDatabase_
    - drops the existing database. it does this so the example starts clean every time its run
  - function _LoadTeams_
    - uploads the []Team slice to the Mongodb server
  - most of this is [boilerplate Golang Mongodb](https://docs.mongodb.com/drivers/go/)
    - create a client
    - connect to mongodb
    - marshal one or more Go object into BSON documents
    - specifiy a database and collection
    - insert the documents into the collection
    - disconnect from mongod
  - once you get the boilerplate, the trickiest part is marshalling the BSON data

## 3. Make sure the database works

- _query_teams_
  - function _QueryTeams_
    - runs some test queries
    - tests that mongodb is accessible
    - tests that the data has been uploaded
    - test that the queries work properly

## 4. Run a GraphQL server

The example uses [graphql-go from graph-gophers](https://github.com/graph-gophers/graphql-go). There are several options if you search for golang graphql. I found this one to be the most straightforward to get running. It has examples in their repo if you need something more complex that what is here.

To implement the server, there are a few steps, all in _server_graphql.go_.

- compose a schema. In this case is a raw string with the standard GraphQL schema language.
- set up a Resolver for the objects that are served, in this case one or more _Team_ objects.
  - _teamResolver_
  - add a method for each of the fields in the object
- create a Root Resolver
  - add methods to the rootResolver type for the supported Queries
  - these methods access the database queries as needed
- create a Root Schema
  - parses the schema string and complains if you make a mistake
- create an HTTP handler using graphql-relay
- start the server
  - a normal golang _http_ server

When you run the program with "go run main.go" run through all the steps and end up spawning a GraphQL server that clients can connect to.

## Make some client side queries

Once the server is running you can execute some client side queries against it. There are a couple of options. There are two here:

- use a client program, for example using node.js
- use Postman (easiest to get going)

### Here are the GraphqQL syntax queries you can execute

```GraphQL
{
    teamsByName(name: "Orioles") {
        id
        name
        location
        year
        wins
        losses
        ties
        other
        games
        uuid
    }

    teamsByID(id: "WAS") {
        id
        name
        location
        year
        wins
        losses
        ties
        other
        games
        uuid
    }

    teamsAll {
        id
        name
        location
        year
    }

}
```

### Typescript/Node Client

The _client_ directory implements a simple command line client that is written in TypeScript and runs with node.js.

- go run main.go
- cd _client_
- edit index.ts and set the server IP:PORT for your system
- npm install
- npm start
  - uses tsc and nodemon watchers so you can change the typescript code and the compilation and execution will run automatically

A tiny bit of the Retrosheet data is incomplete so there could be an 'unknown' show up here or there.

### Using Postman

- Install [Postman](www.postman.com)
- Create POST requests with GraphQL Body
- Try one or more of these queries
