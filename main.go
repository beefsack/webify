package main

import (
	"github.com/beefsack/webify/lib"

	"log"
	"net/http"
	"strings"
)

func main() {
	opts := lib.ParseConfig()

	log.Printf("listening on %s, proxying to %s", opts.Addr, strings.Join(opts.Script, " "))
	log.Fatal(http.ListenAndServe(opts.Addr, &lib.Server{Opts: opts}))
}
