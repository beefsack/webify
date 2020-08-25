package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
)

/// Server is a simple proxy server to pipe HTTP requests to a subprocess' stdin
/// and the subprocess' stdout to the HTTP response.
type Server struct {
	script []string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Start subprocess
	cmd := exec.Command(s.script[0], s.script[1:]...)

	// Get handles to subprocess stdin, stdout and stderr
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		log.Printf("error accessing subprocess stdin: %v", err)
		respError(w)
		return
	}
	defer stdinPipe.Close()
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		log.Printf("error accessing subprocess stderr: %v", err)
		respError(w)
		return
	}
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("error accessing subprocess stdout: %v", err)
		respError(w)
		return
	}

	// Start the subprocess
	err = cmd.Start()
	if err != nil {
		log.Printf("error starting subprocess: %v", err)
		respError(w)
		return
	}

	// We use a WaitGroup to wait for all goroutines to finish
	wg := sync.WaitGroup{}

	// Write request body to subprocess stdin
	wg.Add(1)
	go func() {
		defer func() {
			stdinPipe.Close()
			wg.Done()
		}()
		_, err = io.Copy(stdinPipe, r.Body)
		if err != nil {
			log.Printf("error writing request body to subprocess stdin: %v", err)
			respError(w)
			return
		}
	}()

	// Read all stderr and write to parent stderr if not empty
	wg.Add(1)
	go func() {
		defer wg.Done()
		stderr, err := ioutil.ReadAll(stderrPipe)
		if err != nil {
			log.Printf("error reading subprocess stderr: %v", err)
			respError(w)
			return
		}
		if len(stderr) > 0 {
			log.Print(string(stderr))
		}
	}()

	// Read all stdout, but don't write to the response as we need the exit
	// status of the subcommand to know our HTTP response code
	wg.Add(1)
	var stdout []byte
	go func() {
		defer wg.Done()
		so, err := ioutil.ReadAll(stdoutPipe)
		stdout = so
		if err != nil {
			log.Printf("error reading subprocess stdout: %v", err)
			respError(w)
			return
		}
	}()

	// Wait for the subprocess to complete
	cmdErr := cmd.Wait()
	if cmdErr != nil {
		// We don't return here because we also want to try to write stdout if
		// there was some output
		log.Printf("error running subprocess: %v", err)
		respError(w)
	}
	// Also wait for all of our goroutines to finish, in case there was
	// buffering
	wg.Wait()

	// Write stdout as the response body
	_, err = w.Write(stdout)
	if err != nil {
		log.Printf("error writing response body: %v", err)
	}
}

/// respError sends an error response back to the client. Currently this is just
/// a 500 status code.
func respError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("script-httpd requires a script to execute")
	}
	script := os.Args[1:]

	addr := os.Getenv("SCRIPT_HTTPD_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	log.Printf("listening on %s, proxying to %s", addr, strings.Join(script, " "))
	log.Fatal(http.ListenAndServe(addr, &Server{
		script: script,
	}))
}
