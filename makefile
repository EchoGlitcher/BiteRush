DOCKER_IMAGE := biterush
DOCKER_TAG := latest
GO_BIN_OUT := .builds/biterush
JET_OUT := internal/db/jet

all: build

dockerpush: dockerbuild
	docker push $(DOCKER_IMAGE):$(DOCKER_TAG)

dockerbuild: fmt
	docker build -f ./Dockerfile . -t $(DOCKER_IMAGE):$(DOCKER_TAG)

codegen: clean fmt

	go run ./cmd/jetgen
	go mod tidy

fmt:
	go fmt ./...

clean:
	rm -rf $(GO_BIN_OUT) $(JET_OUT)


build: fmt
	mkdir -p .builds
	go build -o $(GO_BIN_OUT) ./cmd/server