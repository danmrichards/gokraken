build:
	go run cmd/main.go
	go generate .

	cp cmd/validation.tpl validation.go
	gofmt -w validation.go