package teamgames

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// LoadTeamGames ...
func LoadTeamGames(db string, uri string, teamGames []TeamGames) error {

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	// interface slice to contain bson marshalled records
	dox := make([]interface{}, 0)

	// marshall all the documents into 'dox'
	for _, v := range teamGames {
		// marshall the element into a byte slice
		b, err := bson.Marshal(v)
		if err != nil {
			return err
		}
		// append to array of documents
		dox = append(dox, b)
	}

	if len(dox) == 0 {
		// no data
		return errors.New("no team-games data found, check path name")
	}

	// get the collection
	coll := client.Database(db).Collection("teamgames")

	// insert the documents
	res, err := coll.InsertMany(ctx, dox)

	// insertion failed
	if err != nil {
		return err
	}

	// some other fail, auth error can do this
	if res == nil {
		return errors.New("no result : authorization?")
	}

	defer client.Disconnect(ctx)

	return nil
}
