package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Load()
	cookies.Config()
	utils.LoadTemplates()
	r := router.Generate()
	port := fmt.Sprintf(":%d", config.Port)
	fmt.Printf("Running WebApp on port %s ...\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
