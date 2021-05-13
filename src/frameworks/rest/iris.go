package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/kataras/iris/v12"
	"gitlab.com/altiano/goreen-tea/src/app"
	"gitlab.com/altiano/goreen-tea/src/frameworks"
	"gitlab.com/altiano/goreen-tea/src/frameworks/trace"
	"gitlab.com/altiano/goreen-tea/src/shared"
)

type (
	irisRest struct {
		Config shared.Config
		App    app.App
		Tracer trace.Tracer
	}
)

func NewIris(config shared.Config, tracer trace.Tracer, app app.App) frameworks.Server {
	return irisRest{
		Config: config,
		Tracer: tracer,
		App:    app,
	}
}

func (i irisRest) Run() {
	//
	app := iris.New()
	app.Use(iris.Compression)

	//
	customerOrderAPI := app.Party("/customerOrder")
	{
		customerOrderAPI.Post("/", i.createCustomerOrder)
	}

	//
	addr := fmt.Sprintf(":%v", i.Config.AppPort)
	fmt.Println("Iris running at", addr)
	app.Listen(addr)
}

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

/*
	Utils
*/
func mapResponse(ctx iris.Context, res interface{}, err error) {
	if err != nil {
		ctx.JSON(struct {
			Msg string
		}{
			Msg: err.Error(),
		})
		return
	}

	ctx.JSON(res)
}

func mapStatusCode(ctx iris.Context, defaultStatus int, err error) {
	if errors.Is(err, shared.ErrBadRequest) {
		ctx.StatusCode(http.StatusBadRequest)
	} else if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
	} else {
		ctx.StatusCode(defaultStatus)
	}
}
