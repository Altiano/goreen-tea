package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	Db interface {
		Collection(name string, opts ...*options.CollectionOptions) Coll
	}

	Coll interface {
		Drop(ctx context.Context) error
		Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error)
		Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (Cursor, error)
		FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) SingleResult
		BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error)
		CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
		DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
		DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
		UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
		UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
		InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
		InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
		FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) SingleResult
		FindOneAndDelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) SingleResult

		// Helpers
		FindByID(ctx context.Context, id interface{}, v interface{}, opts ...*options.FindOneOptions) error
	}
	Client interface {
		Database(string) Db
		Connect() error
		StartSession() (mongo.Session, error)
	}
	SingleResult interface {
		Decode(v interface{}) error
		Err() error
	}

	Cursor interface {
		All(ctx context.Context, results interface{}) error
	}
)
