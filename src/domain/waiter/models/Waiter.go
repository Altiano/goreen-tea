package waiterModels

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Waiter struct {
		ID         primitive.ObjectID `bson:"_id,omitempty"`
		Name       string             `bson:"Name"`
		Rated      float64            `bson:"Rated"` // 1-10
		TotalServe int                `bson:"TotalServe"`
	}
)
