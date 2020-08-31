<p align="center"><img src="https://i.imgur.com/I2HvDxv.png"></p>
<h1 align="center">webify</h1>
<p align="center"><b>Turn shell commands into web services</b></p>
<p align="center">
  <a href="https://github.com/beefsack/webify/actions"><img src="https://github.com/beefsack/webify/workflows/build/badge.svg" alt="Build Status"></a>
  <a href="https://goreportcard.com/report/github.com/beefsack/webify"><img src="https://goreportcard.com/badge/github.com/beefsack/webify" alt="Go Report Card"></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="MIT License"></a>
</p>
<p align="center"><img src="https://i.imgur.com/OipBB3v.gif"></p>

webify is a very basic CGI server which forwards all requests to a single
script. A design goal is to be as zero-config as possible.

webify invokes your script and writes the request body to your process'
stdin. Stdout is then passed back to the client as the HTTP response body.

If your script returns a non-zero exit code, the HTTP response status code will
be 500.

## Installation

webify is available from the [project's releases page](https://github.com/beefsack/webify/releases).

## Usage

```bash
# Make a web service out of `wc` to count the characters in the request body.
$ webify wc -c
2020/08/25 12:42:32 listening on :8080, proxying to wc -c

...

$ curl -d 'This is a really long sentence' http://localhost:8080
30
```

### Official Docker image

The official Docker image is [beefsack/webify](https://hub.docker.com/r/beefsack/webify).

It can be configured using the following environment variables:

* `ADDR` - the address to listen on inside the container, defaults to `:80`
* `SCRIPT` - the command to execute, defaults to `/script`

#### Mounting script and running official image

```
$ docker run -it --rm -p 8080:80 -v /path/to/my/script:/script beefsack/webify:latest
2020/08/25 04:27:46 listening on :80, proxying to /script

...

$ curl -d 'Some data' http://localhost:8080
```

#### Building a new image using official image as base

Create a `Dockerfile` like the following:

```
FROM beefsack/webify:latest
COPY myscript /script
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
