# script-httpd

script-httpd is a simple tool to turn a command line script into a web service.

<p align="center">
  <img src="https://i.imgur.com/1wL9m5C.gif">
</p>

script-httpd is essentially a very basic CGI server which forwards all requests
to a single script. A design goal is to be as close to zero-config as possible.

script-httpd invokes your script and writes the request body to your process'
stdin. Stdout is then passed back to the client as the HTTP response body.

If your script returns a non-zero exit code, the HTTP response status code will
be 500.

## Installation

script-httpd is available from the [project's releases page](https://github.com/beefsack/script-httpd/releases).

## Usage

```bash
# Make a web service out of `wc` to count the characters in the request body.
$ script-httpd wc -c
2020/08/25 12:42:32 listening on :8080, proxying to wc -c

...

$ curl -d 'This is a really long sentence' http://localhost:8080
30
```

### Official Docker image

The official Docker image is [beefsack/script-httpd](https://hub.docker.com/r/beefsack/script-httpd).

It can be configured using the following environment variables:

* `SCRIPT_HTTPD_ADDR` - the address to listen on inside the container, defaults to `:80`
* `SCRIPT_HTTPD_CMD` - the command to execute, defaults to `/script`

#### Mounting script and running official image

```
$ docker run -it --rm -p 8080:80 -v /path/to/my/script:/script beefsack/script-httpd:latest
2020/08/25 04:27:46 listening on :80, proxying to /script

...

$ curl -d 'Some data' http://localhost:8080
```

#### Building a new image using official image as base

Create a `Dockerfile` like the following:

```
FROM beefsack/script-httpd:latest
COPY myscript /script
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
