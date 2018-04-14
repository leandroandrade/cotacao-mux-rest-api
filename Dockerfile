# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/leandroandrade/cotacao-mux-rest-api

RUN go get -u github.com/gorilla/mux
RUN go install github.com/leandroandrade/cotacao-mux-rest-api

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/cotacao-mux-rest-api

# Document that the service listens on port 8080.
EXPOSE 3000
