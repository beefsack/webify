FROM golang:1.14

WORKDIR /go/src/script-httpd

COPY . .
RUN go get && go build

# By default, script-httpd listens on all interfaces on port 8080
EXPOSE 80
ENV ADDR :80

# By default, script-httpd executes /script
ENV SCRIPT /script

CMD ["./script-httpd"]