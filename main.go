package main

import (
	"flag"
	"log"

	"github.com/go-openapi/loads"

	"github.com/ukdave/goswagger-todolist/gen/restapi"
	"github.com/ukdave/goswagger-todolist/gen/restapi/operations"
)

func main() {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewTodoAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// parse flags
	var portFlag = flag.Int("port", 3000, "The port to listen on")
	flag.Parse()

	// set the port this service will be run on
	server.Port = *portFlag

	// TODO: Set Handle
	server.ConfigureAPI()

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
