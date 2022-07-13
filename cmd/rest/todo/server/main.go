package main

import (
	"fmt"
	"net/http"
	generatedServer "web-service/generated/rest/server"
	"web-service/pkg/interfaces/rest/todo/v1/server"

	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	todoServer := server.NewTodoServer()
	swagger, err := generatedServer.GetSwagger()
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	swagger.Servers = nil
	router := gin.Default()
	router.Use(middleware.OapiRequestValidator(swagger))
	router = generatedServer.RegisterHandlers(router, todoServer)
	httpServer := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("localhost:%d", 8080),
	}
	err = httpServer.ListenAndServe()
	if err != nil {
		return
	}
}
