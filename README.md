# Simple Go : Json -> Mongodb -> GraphQL -> Client

## Overview

The point of this repo is to provide, in _simple, minimal_ code, an example in Go, of reading a JSON file, uploading it to a Mongodb database, and serving the data using GraphQL.

There's a bit of learning curve if you are starting something like this from scratch, expecially with Mongodb and the GraphQL server. There are other examples but there isn't a single example in one place that has the pieces.

## Main Program

The main program runs through all the steps to get this going. At each step you can just look at the function being called, find its implementation in the 'teams' directory, and see what it does.

### Setup

- have an instance of Mongodb running somewhere you can connect to, and get its connection string.
- at the top of _main.go_
  - set the mongodb connection string (mongouri)
  - set the ip address and port that the GraphQL server will use. This is the address that clients will connect to for GrapQL Queries
  - set the database name. It defaults to _bb-team-games_, but make sure it unique to your system so you don't step on something important.
- this required Go 1.16 or later due to changes in ioutil

## 1. Read Json File

- _teams/read_teams.go_
  - read the JSON file
  - using encoding/json, unmarshal it into Go objects
  - at this point you have a slice of Team objects

The data used here is derived from the data products provided by [Retrosheet.org](retrosheet.org). Retrosheet does a lot of work to provide historical baseball data going all the way back to 1876. The data used here is a very small subset of what they provide.

Thanks to Retrosheet for making this data available.

<pre>
     The information used here was obtained free of
     charge from and is copyrighted by Retrosheet.  Interested
     parties may contact Retrosheet at "www.retrosheet.org". 
</pre>

## 2. Load it into Mongodb

- _teams/load_teams.go_
  - drops the existing database. it does this so the example starts clean every time its run
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

- the example runs some test queries
  - that mongodb is accessible
  - that the data has been uploaded
  - that the queries work properly

## 4. Run a GraphQL server

The example uses [graphql-go from graph-gophers](https://github.com/graph-gophers/graphql-go). There are several options if you search for golang graphql. I found this one to be the most straightforward to get running. It has examples in their repo if you need something more complex that what is here.

To implement the server, there are a few steps, all in _server_graphql.go_.

- compose a schema. In this case is a raw string with the standard GraphQL schema language.
- set up a Resolver for the objects that are served, in this case one or more _Team_ objects.
  - _teamResolver_
  - add a method for each of the fields in the Object
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

- use Postman (easiest to get going)
- use a client program, for example using node.js

### Here are the queries you can execute

```GraphQL
{
    teamsByName(name: "Orioles") {
        id
        name
        wins
        losses
        other
        uuid
    }

    teamsByID(id: "WAS") {
        id
        name
        wins
        losses
        other
        uuid
    }

    teamsAll {
        name
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

### Using Postman

- Install [Postman](www.postman.com)
- Create POST requests with GraphQL Body
- Try one or more of these queries
