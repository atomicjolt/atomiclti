package main

import (
	"fmt"
	"github.com/atomicjolt/atomiclti/config"
	"github.com/atomicjolt/atomiclti/controllers"
	"github.com/atomicjolt/atomiclti/resources"
	"log"
	"net/http"
)

func main() {
	localConfig := config.GetServerConfig()
	port := localConfig.ServerPort
	controllerResources, cancelResourcesContext := resources.NewResources()

	defer cancelResourcesContext()

	fmt.Printf("Running in %s mode...\n", config.DetermineEnv())
	fmt.Printf("Listening on port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, controllers.NewRouter(controllerResources)))
}
