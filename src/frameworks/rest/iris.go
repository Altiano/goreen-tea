package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/kataras/iris/v12"
	"gitlab.com/altiano/golang-boilerplate/src/app"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks/trace"
	"gitlab.com/altiano/golang-boilerplate/src/shared"
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
