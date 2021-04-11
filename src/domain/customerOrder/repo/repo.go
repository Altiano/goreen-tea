package repo

import (
	"context"

	coModels "gitlab.com/altiano/golang-boilerplate/src/domain/customerOrder/models"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks/database"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks/trace"
)

type (
	Repo interface {
		GetLastCode(ctx context.Context) (int, error)
		Save(ctx context.Context, co coModels.CustomerOrder) (coModels.CustomerOrder, error)
		UpdateWaiter(ctx context.Context, co coModels.CustomerOrder, waiterName string) error
		GetReferralCodeCreator(ctx context.Context, refererCode int) (coModels.CustomerOrder, error)
		GetReferralCodeRedeemer(ctx context.Context, refererCode int) (coModels.CustomerOrder, error)
	}

	repo struct {
		Tracer trace.Tracer
		Coll   database.Coll
	}
)

func NewRepo(tracer trace.Tracer, db database.Db) Repo {
	return &repo{
		Tracer: tracer,
		Coll:   db.Collection("customerOrder"),
	}
}
