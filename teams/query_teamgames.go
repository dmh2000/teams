package teams

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// QueryTeamGames
func QueryTeamGames(db string, uri string, key string, value string) ([]Team, error) {
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
	coll := client.Database(db).Collection("teams")

	// is there a key and value?
	var cursor *mongo.Cursor
	if (key == "") && (value == "") {
		// get all documents
		cursor, err = coll.Find(ctx, bson.D{})
	} else {
		// get by key/value
		cursor, err = coll.Find(ctx, bson.D{{Key: key, Value: value}})
	}
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
	var teamGames []Team
	for cursor.Next(ctx) {
		var t Team

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
