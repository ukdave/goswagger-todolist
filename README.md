# Go-Swagger To Do List API

A To Do List API built with [go-swagger](https://goswagger.io/).

## Dependencies

*   [go](https://golang.org/)
*   [dep](https://github.com/golang/dep)
*   [go-swagger](https://goswagger.io/)

```shell
brew install go
brew install dep

brew tap go-swagger/go-swagger
brew install go-swagger
```

## Code Generation

The server code is built from the [swagger spec](./swagger.yml) file:

```shell
mkdir gen
swagger validate ./swagger.yaml
swagger generate server -t gen -f ./swagger.yaml --exclude-main -A todo
dep init
```

## Running

```shell
go run main.go --port 3000
curl -i http://localhost:3000/api/
```
