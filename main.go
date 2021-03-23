package main

import (
	"log"
	"net/http"

	"github.com/dedeTr/text-spam-detector.git/router"
)

func main() {

	r := router.Router()
	log.Fatal(http.ListenAndServe(":8080", r))

}
