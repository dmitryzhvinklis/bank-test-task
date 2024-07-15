package main

import (
	"log"
	"net/http"

	"github.com/dmitryzhvinklis/bank-test-task/routers"
)

func main() {
	router := routers.SetupRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
