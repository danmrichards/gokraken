build:
	go run cmd/main.go
	go generate .

	cp cmd/validation.html validation.go
	gofmt -w validation.go