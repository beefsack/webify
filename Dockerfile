FROM golang:1.14

WORKDIR /go/src/script-httpd

COPY main.go .
RUN go build

COPY docker/run.sh .

# By default, script-httpd listens on all interfaces on port 8080
EXPOSE 80
ENV SCRIPT_HTTPD_ADDR :80

# By default, script-httpd executes /script
ENV SCRIPT_HTTPD_CMD /script

CMD ["./run.sh"]