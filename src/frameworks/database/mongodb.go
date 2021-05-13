package database

import (
	"context"
	"errors"

	"gitlab.com/altiano/goreen-tea/src/shared"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type (
	mongoDb struct {
		client   *mongo.Client
		database *mongo.Database
	}
	mongoColl struct {
		collection *mongo.Collection
	}

	mongoSingleResult struct {
		singleResult *mongo.SingleResult
	}

	mongoCursor struct {
		cursor *mongo.Cursor
	}
)

/*
	NewMongoDB returns DB
*/

func NewMongoDB(config shared.Config) Db {
	ctx := context.Background()
	uri := config.MongoURI
	name := config.MongoDBName

	if uri == "" {
		panic(errors.New("uri is required"))
	}

	if name == "" {
		panic(errors.New("database name is required"))
	}

	opts := options.Client().
		ApplyURI(uri)

	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	database := client.Database(name)

	return &mongoDb{
		database: database,
		client:   client,
	}
}

func (d *mongoDb) Collection(name string, opts ...*options.CollectionOptions) Coll {
	return &mongoColl{collection: d.database.Collection(name, opts...)}
}

/*
	Collection implements all available operations.
*/

func (c *mongoColl) Drop(ctx context.Context) error {
	return c.collection.Drop(ctx)
}

func (c *mongoColl) Aggregate(ctx context.Context, pipeline interface{},
	opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	return c.collection.Aggregate(ctx, pipeline, opts...)
}

func (c *mongoColl) Find(ctx context.Context, filter interface{},
	opts ...*options.FindOptions) (Cursor, error) {
	return c.collection.Find(ctx, filter, opts...)
}

func (c *mongoColl) FindOne(ctx context.Context, filter interface{},
	opts ...*options.FindOneOptions) SingleResult {
	return &mongoSingleResult{
		singleResult: c.collection.FindOne(ctx, filter, opts...),
	}
}

func (c *mongoColl) BulkWrite(ctx context.Context, models []mongo.WriteModel,
	opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	return c.collection.BulkWrite(ctx, models, opts...)
}

func (c *mongoColl) CountDocuments(ctx context.Context, filter interface{},
	opts ...*options.CountOptions) (int64, error) {
	return c.collection.CountDocuments(ctx, filter, opts...)
}

func (c *mongoColl) DeleteOne(ctx context.Context, filter interface{},
	opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return c.collection.DeleteOne(ctx, filter, opts...)
}

func (c *mongoColl) DeleteMany(ctx context.Context, filter interface{},
	opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return c.collection.DeleteMany(ctx, filter, opts...)
}

func (c *mongoColl) UpdateMany(ctx context.Context, filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return c.collection.UpdateMany(ctx, filter, update, opts...)
}

func (c *mongoColl) UpdateOne(ctx context.Context, filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return c.collection.UpdateOne(ctx, filter, update, opts...)
}

func (c *mongoColl) InsertMany(ctx context.Context, documents []interface{},
	opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return c.collection.InsertMany(ctx, documents, opts...)
}

func (c *mongoColl) InsertOne(ctx context.Context, document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return c.collection.InsertOne(ctx, document, opts...)
}

func (c *mongoColl) FindOneAndUpdate(ctx context.Context, filter interface{},
	update interface{}, opts ...*options.FindOneAndUpdateOptions) SingleResult {
	return &mongoSingleResult{
		singleResult: c.collection.FindOneAndUpdate(ctx, filter, update, opts...),
	}
}

func (c *mongoColl) FindOneAndDelete(ctx context.Context, filter interface{},
	opts ...*options.FindOneAndDeleteOptions) SingleResult {
	return &mongoSingleResult{
		singleResult: c.collection.FindOneAndDelete(ctx, filter, opts...),
	}
}

// Helpers
func (c *mongoColl) FindByID(ctx context.Context, id interface{}, v interface{}, opts ...*options.FindOneOptions) error {
	// primitive.ObjectIDFromHex("asd")
	return c.FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(v)
}

/*
	SingleResult
*/

func (c *mongoSingleResult) Decode(v interface{}) error {
	return c.singleResult.Decode(v)
}

func (c *mongoSingleResult) Err() error {
	return c.singleResult.Err()
}
