# script-httpd

script-httpd is a simple tool to turn a command line script into a web service.

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

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)