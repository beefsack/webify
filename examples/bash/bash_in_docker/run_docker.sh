#!/bin/bash

docker run -it --rm -p 8080:80 -v $(realpath logic.sh):/script beefsack/webify:latest
