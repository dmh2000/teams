package teamgames

import (
	"context"
	"errors"

	"github.com/dmh2000/retrosheet/src/jsontypes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetTeamGamesByID full inline
func GetTeamGamesByID(db string, uri string, id string) ([]jsontypes.TeamGames, error) {
	var err error
	var client *mongo.Client

	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	// get the personnel collection
	coll := client.Database(db).Collection("teamgames")

	// find documents that have the key "abbr" with the specified value
	cursor, err := coll.Find(ctx, bson.D{{Key: "ID", Value: id}})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	// insertion failed
	if err != nil {
		return nil, err
	}

	// some other fail, auth error can do this
	if cursor == nil {
		return nil, errors.New("no result : authorization?")
	}

	// decode all matching records
	var teamGames []jsontypes.TeamGames
	for cursor.Next(ctx) {
		var t jsontypes.TeamGames

		// decode into a team
		err = cursor.Decode(&t)
		if err != nil {
			return nil, err
		}

		// add to list
		teamGames = append(teamGames, t)
	}
	return teamGames, nil
}
