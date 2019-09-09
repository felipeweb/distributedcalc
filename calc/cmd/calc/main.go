package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/caarlos0/httperr"
	"github.com/felipeweb/distributedcalc/calc"
	"github.com/gorilla/handlers"
)

func main() {
	http.Handle("/", handlers.LoggingHandler(os.Stdout, httperr.NewF(calc.Handler)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}
