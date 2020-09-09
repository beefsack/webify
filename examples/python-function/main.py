from jsonschema import validate

import json
import sys

# We have a JSON schema to validate our input
from input import SCHEMA
# We have an A* pathfinder defined in pathfind.py
from pathfind import pathfind


def main():
    # Read input from stdin
    input = json.load(sys.stdin)

    # Validate our input against our JSON schema
    validate(instance=input, schema=SCHEMA)

    # Find a path from start to end in our world
    result = pathfind(**input)

    # Print result to stdout
    json.dump(result, sys.stdout)


if __name__ == '__main__':
    main()
