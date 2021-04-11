package customerOrder

import (
	"context"

	coModels "gitlab.com/altiano/golang-boilerplate/src/domain/customerOrder/models"
	"gitlab.com/altiano/golang-boilerplate/src/domain/customerOrder/repo"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks/email"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks/trace"
)

type (
	Domain interface {
		Create(ctx context.Context, req coModels.CreateReq) (coModels.CustomerOrder, error)
		NotifyRedeemedCode(ctx context.Context, co coModels.CustomerOrder) error
		NotifyWaiterName(ctx context.Context, co coModels.CustomerOrder, waiterName string) error
	}

	domain struct {
		Tracer trace.Tracer
		Email  email.Emailer
		Repo   repo.Repo
	}
)

func NewDomain(trace trace.Tracer, email email.Emailer, repo repo.Repo) Domain {
	return domain{
		Tracer: trace,
		Email:  email,
		Repo:   repo,
	}
}
