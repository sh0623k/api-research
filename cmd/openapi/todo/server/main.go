package main

import (
	generatedServer "api-research/generated/openapi/server"
	"api-research/pkg/interfaces/openapi/todo/v1/server"
	"fmt"
	"log"
	"net/http"
	"strconv"

	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
)

const Port = 3000

func main() {
	swagger, err := generatedServer.GetSwagger()
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	swagger.Servers = nil
	router := gin.Default()
	router.Use(middleware.OapiRequestValidator(swagger))
	router = generatedServer.RegisterHandlers(router, server.NewTodoServer())
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(Port), router))
}
