swagger:
	mkdir -p gen
	go run vendor/github.com/go-swagger/go-swagger/cmd/swagger/swagger.go generate server -t gen -f ./swagger.yaml --exclude-main -A todo

.PHONY: swagger
