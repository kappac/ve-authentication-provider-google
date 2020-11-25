FROM golang:1.15

WORKDIR /service

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -race -o service.exec ./cmd/service" -command="./service.exec"