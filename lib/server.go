package lib

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"sync"
)

// Server is a simple proxy server to pipe HTTP requests to a subprocess' stdin
// and the subprocess' stdout to the HTTP response.
type Server struct {
	Opts
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Start subprocess
	cmd := exec.Command(s.Script[0], s.Script[1:]...)

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

	// We must consume stdout and stderr before `cmd.Wait()` as per
	// doc and example at https://golang.org/pkg/os/exec/#Cmd.StdoutPipe
	wg.Wait()

	// Wait for the subprocess to complete
	cmdErr := cmd.Wait()
	if cmdErr != nil {
		// We don't return here because we also want to try to write stdout if
		// there was some output
		log.Printf("error running subprocess: %v", err)
		respError(w)
	}

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
