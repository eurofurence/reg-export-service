package main

import (
	"fmt"
	"log"

	"github.com/eurofurence/reg-export-service/internal/repository/config"
	"github.com/eurofurence/reg-export-service/internal/repository/logging/consolelogging/logformat"
	"github.com/eurofurence/reg-export-service/web"
)

func main() {
	err := config.LoadConfiguration("config.yaml")
	if err != nil {
		log.Fatal(logformat.Logformat("ERROR", "00000000", fmt.Sprintf("Error while loading configuration: %v", err)))
	}
	log.Println(logformat.Logformat("INFO", "00000000", "Initializing..."))
	server := web.Create()
	web.Serve(server)
}
