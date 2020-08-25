#!/bin/bash

# Helper script so that we can pass the script as an environment variable.

SCRIPT_HTTPD_CMD=${SCRIPT_HTTPD_CMD:-/script}
"$(dirname "$0")/script-httpd" "$SCRIPT_HTTPD_CMD"