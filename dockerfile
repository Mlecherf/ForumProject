FROM golang:1.16.5
RUN mkdir /app
WORKDIR /app

COPY . .
# RUN go get -d -v ./...
# RUN go install -v ./...
RUN go mod download
# RUN go get github.com/mattn/go-sqlite3

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build serveur.go" --command=./serveur

EXPOSE 8080