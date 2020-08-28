FROM golang:1.14

WORKDIR /go/src/webify

COPY . .
RUN go get && go build

# By default, webify listens on all interfaces on port 8080
EXPOSE 80
ENV ADDR :80

# By default, webify executes /script
ENV SCRIPT /script

CMD ["./webify"]