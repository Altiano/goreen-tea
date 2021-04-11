package repo

import (
	"context"

	waiterModels "gitlab.com/altiano/golang-boilerplate/src/domain/waiter/models"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks/database"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks/trace"
)

var (
	CollName = "account"
)

type (
	Repo interface {
		GetByFreetime(ctx context.Context) ([]waiterModels.Waiter, error)
		IncreaseTotalServe(ctx context.Context, waiter waiterModels.Waiter) error
	}

	repo struct {
		Tracer trace.Tracer
		Coll   database.Coll
	}
)

func NewRepo(db database.Db, trace trace.Tracer) Repo {
	return repo{
		Coll:   db.Collection("waiter"),
		Tracer: trace,
	}
}
