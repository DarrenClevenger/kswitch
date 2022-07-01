build:
	go build -o bin/kswitch cmd/main.go

clean:
	rm -r bin/

test:
	go test ./... -v

run: build
	go run cmd/main.go
