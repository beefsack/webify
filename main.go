package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/beefsack/script-httpd/scripthttpd"
)

func main() {
	opts := scripthttpd.ParseConfig()

	log.Printf("listening on %s, proxying to %s", opts.Addr, strings.Join(opts.Script, " "))
	log.Fatal(http.ListenAndServe(opts.Addr, &scripthttpd.Server{Opts: opts}))
}
