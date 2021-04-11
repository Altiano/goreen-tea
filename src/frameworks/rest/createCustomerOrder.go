package rest

import (
	"context"
	"net/http"

	"github.com/kataras/iris/v12"
	"gitlab.com/altiano/golang-boilerplate/src/app"
)

func (i irisRest) createCustomerOrder(irisCtx iris.Context) {
	var req app.OnsiteOrderReq
	err := irisCtx.ReadJSON(&req)

	//
	if err != nil {
		irisCtx.StatusCode(http.StatusBadRequest)
		mapResponse(irisCtx, nil, err)
		return
	}

	//
	ctx, span := i.Tracer.Start(context.Background(), "Rest.CreateCustomerOrder")
	defer span.End()
	res, err := i.App.OnsiteOrder(ctx, req)

	//
	mapStatusCode(irisCtx, http.StatusCreated, err)
	mapResponse(irisCtx, res, err)

}
