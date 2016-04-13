
GOBIN := $(GOPATH)/bin

build: prebuild
	CGO_ENABLED=0 $(GOBIN)/godep go build -v -a --installsuffix cgo -ldflags '-s' -o ./bin/goapi

docker_run:
	CGO_ENABLED=0 GOOS=linux $(GOBIN)/godep go build -v -a --installsuffix cgo -ldflags '-s' -o goapi
	docker build -t goapi -f Dockerfile .
	docker run --rm -p 4200:4200 goapi

prebuild: generate fmt vet


run: build
	./bin/goapi

fmt:
	go fmt ./...

vet:
	go vet ./...

generate:
	go generate ./...

clean:
	- rm -fr goapi
	- rm -fr ./bin/goapi

builder:
	go fmt ./...


# delete any changes to the Godeps folder
godeps_reset:
	git reset -- Godeps
	git checkout -- Godeps
	git clean -df Godeps

