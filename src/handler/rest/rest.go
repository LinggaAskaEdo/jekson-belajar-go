package rest

import (
	"belajar-go/src/middleware"
	"sync"

	"belajar-go/src/service"

	"github.com/gin-gonic/gin"
)

type rest struct {
	gin  *gin.Engine
	svc  *service.Service
	port string
}

// this singleton pattern to make sure InitRestHandlerr is called only during execution
var onceRestHandler = &sync.Once{}

func InitRestHandler(svc *service.Service, port string) {
	onceRestHandler.Do(func() {
		e := &rest{
			gin:  gin.Default(),
			svc:  svc,
			port: port,
		}
		e.Serve()
	})
}

func (e *rest) Serve() {
	e.gin.GET("/users", e.ListUsers)
	e.gin.POST("/user", middleware.LimitBodySize(2<<20), e.CreateUser)

	if err := e.gin.Run(":" + e.port); err != nil {
		panic(err)
	}
}
