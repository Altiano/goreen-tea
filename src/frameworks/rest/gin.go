package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/altiano/goreen-tea/src/app"
	"gitlab.com/altiano/goreen-tea/src/frameworks"
	"gitlab.com/altiano/goreen-tea/src/frameworks/trace"
	"gitlab.com/altiano/goreen-tea/src/shared"
)

type (
	ginRest struct {
		Config shared.Config
		App    app.App
		Tracer trace.Tracer
	}
)

func NewGin(config shared.Config, tracer trace.Tracer, app app.App) frameworks.Server {
	return ginRest{
		Config: config,
		Tracer: tracer,
		App:    app,
	}
}

func (i ginRest) Run() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//
	v1 := r.Group("/customerOrder")
	{
		v1.POST("/", i.createCustomerOrder)
	}

	//
	addr := fmt.Sprintf(":%v", i.Config.AppPort)
	r.Run(addr)
}

func (i ginRest) createCustomerOrder(c *gin.Context) {
	var req app.OnsiteOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := i.App.OnsiteOrder(context.Background(), req)

	if err != nil {
		code := http.StatusInternalServerError

		if errors.Is(err, shared.ErrBadRequest) {
			code = http.StatusBadRequest
		}

		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
