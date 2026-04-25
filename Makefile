.PHONY: build run test docker clean

BINARY=sre-agent
DOCKER_IMAGE=sravanthigorantla/sre-agent

build:
	go build -o bin/$(BINARY) ./cmd/sre-agent

run:
	OPENAI_API_KEY=$(OPENAI_API_KEY) go run ./cmd/sre-agent

test:
	go test -v ./...

docker:
	docker build -t $(DOCKER_IMAGE):latest .

docker-push:
	docker push $(DOCKER_IMAGE):latest

clean:
	rm -rf bin/
