# Define variables
CI_COMMIT_SHA := $(shell git rev-parse --short HEAD)
IMAGE_NAME := goapi
RELEASE_NAME := goapi-dev
DEPLOYMENT_PATH := ./deployment
VALUES_FILE := $(DEPLOYMENT_PATH)/values.yaml
BINARY_PATH := ./bin/goapi

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

# Build Docker image
.PHONY: docker_build
docker_build: build
	@echo "Building new Docker image with tag $(CI_COMMIT_SHA)..."
	docker build -t $(IMAGE_NAME):$(CI_COMMIT_SHA) .

# Deploy using Helm
.PHONY: deploy
deploy: docker_build
	@echo "Deploying with Helm..."
	helm upgrade --install $(RELEASE_NAME) $(DEPLOYMENT_PATH) -f $(VALUES_FILE) --set image.repository=$(IMAGE_NAME) --set image.tag=$(CI_COMMIT_SHA)
	@echo "Deployment completed successfully."

# Show help
.PHONY: help
help:
	@echo "Makefile targets:"
	@echo "  all      - Clean, build, and deploy the Docker image"
	@echo "  prebuild - Run prebuild steps (generate, fmt)"
	@echo "  build    - Build the Go binary"
	@echo "  run      - Run the Go binary"
	@echo "  fmt      - Format Go code"
	@echo "  generate - Generate Go code"
	@echo "  clean    - Clean up build artifacts"
	@echo "  docker_build - Build the Docker image"
	@echo "  docker_run   - Run the Docker container"
	@echo "  deploy   - Deploy the Docker image using Helm"
	@echo "  test     - Run Go tests"
	@echo "  help     - Show this help message"