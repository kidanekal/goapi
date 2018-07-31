build: prebuild
	CGO_ENABLED=0 go build -v -a --installsuffix cgo -ldflags '-s' -o ./bin/goapi

docker_run:
	docker build -t goapi -f Dockerfile .
	docker run --rm -p 4200:4200 goapi

prebuild: generate fmt


run: build
	./bin/goapi

fmt:
	go fmt ./...

generate:
	go generate ./...

clean:
	- rm -fr goapi
	- rm -fr ./bin/goapi

builder:
	go fmt ./...
