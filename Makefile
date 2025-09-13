app_name := hexagonal-demo

build-app:
  go build -o ./bin/$(app_name) ./cmd/http/main.go
