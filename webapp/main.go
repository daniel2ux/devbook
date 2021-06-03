package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
)

func main() {
	fmt.Println("Running WebApp...")

	r := router.Generate()
	log.Fatal(http.ListenAndServe(":3000", r))
}
