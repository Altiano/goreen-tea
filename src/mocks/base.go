package mocks

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	NilErr        error = nil
	ErrDummy            = errors.New("an_error")
	ObjectIDDummy       = primitive.NewObjectIDFromTimestamp(time.Date(2010, 1, 1, 1, 1, 1, 1, time.UTC))
)

type (
	Base struct {
		Ctx context.Context
	}
)

func setupBase() Base {
	return Base{
		Ctx: context.Background(),
	}
}
