package persistence

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
Persistence represents structure to handle database operations
*/
type Persistence struct {
	URI             string
	client          *mongo.Client
	database        *mongo.Database
	boardCollection *mongo.Collection
}

const database = "tinamar"
const boardCollection = "league_board"

/*
Connect initiates connection to MongoDB instance
*/
func (p *Persistence) Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(p.URI))

	if err != nil {
		return err
	}

	p.client = client

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		return err
	}

	p.database = client.Database(database)
	p.boardCollection = p.database.Collection(boardCollection)

	return nil
}
