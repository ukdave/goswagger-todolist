# Go-Swagger To Do List API

A To Do List API built with [go-swagger](https://goswagger.io/).

## Dependencies

*   [go](https://golang.org/)
*   [dep](https://github.com/golang/dep)

```shell
brew install go dep
```

## Code Generation

The server code is built from the [swagger spec](./swagger.yml) file using
[go-swagger](https://goswagger.io/) and all generated code is put in a folder
`gen`.

```shell
dep ensure
make swagger
```

## Running

```shell
go run main.go --port 3000
http :3000/api/
http :3000/api/ description="Do some stuff" completed:=true
http :3000/api/ description="Do some more stuff" completed:=false
```
