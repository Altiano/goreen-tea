package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/altiano/goreen-tea/src/app"
	"gitlab.com/altiano/goreen-tea/src/frameworks"
	"gitlab.com/altiano/goreen-tea/src/frameworks/trace"
	"gitlab.com/altiano/goreen-tea/src/shared"
)

type (
	echoRest struct {
		Config shared.Config
		App    app.App
		Tracer trace.Tracer
	}
)

func NewEcho(config shared.Config, tracer trace.Tracer, app app.App) frameworks.Server {
	return echoRest{
		Config: config,
		Tracer: tracer,
		App:    app,
	}
}

func (i echoRest) Run() {
	e := echo.New()

	customerOrder := e.Group("customerOrder")
	customerOrder.POST("", i.createCustomerOrder)

	addr := fmt.Sprintf(":%v", i.Config.AppPort)
	e.Logger.Fatal(e.Start(addr))
}

func (i echoRest) createCustomerOrder(c echo.Context) error {
	var req app.OnsiteOrderReq
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	res, err := i.App.OnsiteOrder(context.Background(), req)
	if err != nil {
		code := http.StatusInternalServerError

		if errors.Is(err, shared.ErrBadRequest) {
			code = http.StatusBadRequest
		}

		return c.String(code, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
