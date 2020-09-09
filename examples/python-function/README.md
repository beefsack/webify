# Example Python function

This example shows how we can turn a Python function into a web service using
`webify`.

Our example here will provide an A* pathfinding service. A client will send a
world, a start location, and an end location. The service will execute our
pathfinding function and return the result.

## Pathfinding function

Our base pathfinding function is in [`pathfind.py`](pathfind.py):

```python
def pathfind(world, start_x, start_y, end_x, end_y):
    """Find a path from start to end in a world"""
    # ...
```

This function returns a dictionary containing the path, the number of runs, and
a rendered result.

## Connect function to stdin and stdout

Now we need a main function that takes stdin as input, and returns output to
stdout. We will use JSON as the format for both input and output.

Our main function is in [`main.py`](main.py):

```python
def main():
    # Read input from stdin
    input = json.load(sys.stdin)

    # Validate our input against our JSON schema
    validate(instance=input, schema=SCHEMA)

    # Find a path from start to end in our world
    result = pathfind(**input)

    # Print result to stdout
    json.dump(result, sys.stdout)
```

It is **very important** to validate input, particularly on a service exposed
on a network. We do this using a JSON schema which is defined in
[`input.py`](input.py).

## Running in shell

We can test our python script directly at the shell using the example
[`input.json`](input.json) file:

```bash
$ # First we need to install the dependencies
$ pip install -r requirements.txt
...
Successfully installed

$ # Get the raw JSON result
$ python main.py < input.json
{"path": [[0, 0], [1, 0], [2, 0], [2, 1], [2, 2], [2, 3], [3, 3]], "runs": 9, "render": "+----+\n|sxx#|\n| #x |\n|##x#|\n|  xe|\n+----+"}

$ # Just show the rendered world from the result
$ python main.py < input.json | jq -r .render
+----+
|sxx#|
| #x |
|##x#|
|  xe|
+----+
```

## Use `webify` to turn it into a service

Now all we need to do is run `webify` to expose our function to the network over
HTTP:

```bash
$ webify python main.py
2020/09/09 13:47:34 listening on :8080, proxying to python main.py
...

$ # Get the raw JSON result
$ curl -d @input.json http://localhost:8080
{"path": [[0, 0], [1, 0], [2, 0], [2, 1], [2, 2], [2, 3], [3, 3]], "runs": 9, "render": "+----+\n|sxx#|\n| #x |\n|##x#|\n|  xe|\n+----+"}

$ # Just show the rendered world from the result
$ curl -d @input.json http://localhost:8080 | jq -r .render
+----+
|sxx#|
| #x |
|##x#|
|  xe|
+----+
```