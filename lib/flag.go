package lib

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mattn/go-shellwords"
	"github.com/namsral/flag"
)

const EnvScript = "SCRIPT"

const helpHeader = `
webify is a simple helper to turn a command line script into an HTTP
service.

Homepage: http://github.com/beefsack/webify

webify functions by starting the script for each request and piping the
HTTP request body into stdin of the subprocess. stdout is captured and returned
as the body of the HTTP response.

stderr is not sent to the client, but is logged to the webify process
stderr. stderr can be sent to the client using redirection if required.

The exit status of the script determines the HTTP status code for the response:
200 when the exit status is 0, otherwise 500. Because of this, the response
isn't sent until the script completes.

Example server that responds with the number of lines in the request body:

  webify wc -l

Piping and redirection are supported by calling a shell directly:

  webify bash -c 'date && wc -l'

All options are also exposed as environment variables, entirely in uppercase.
Eg. -addr can also be specified using the environment variable ADDR. The script
arguments can be passed in the SCRIPT environment variable instead.

Available options:
`

type Opts struct {
	Script []string
	Addr   string
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, strings.TrimLeft(helpHeader, "\n"))
		flag.PrintDefaults()
	}
}

func ParseConfig() Opts {
	opts := Opts{}

	flag.StringVar(&opts.Addr, "addr", ":8080", "the TCP network address to listen on, eg. ':80'")

	flag.Parse()

	opts.Script = flag.Args()
	if len(opts.Script) == 0 {
		envScript := os.Getenv(EnvScript)
		if envScript != "" {
			args, err := shellwords.Parse(envScript)
			if err != nil {
				log.Fatalf("error parsing SCRIPT environment variable: %v", err)
			}
			opts.Script = args
		} else {
			// No script was passed via args or env, print usage and exit.
			flag.Usage()
			os.Exit(2)
		}
	}

	return opts
}
