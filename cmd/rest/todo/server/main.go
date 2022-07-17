package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	generatedServer "web-service/generated/rest/server"
	"web-service/pkg/interfaces/rest/todo/v1/server"

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
