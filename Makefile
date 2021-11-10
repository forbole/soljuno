VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT  := $(shell git log -1 --format='%H')

export GO111MODULE = on

all: lint test-unit install

###############################################################################
###                                Build flags                              ###
###############################################################################

LD_FLAGS = -X github.com/forbole/soljuno.Version=$(VERSION) \
	-X github.com/forbole/soljuno.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(LD_FLAGS)'

###############################################################################
###                                  Build                                  ###
###############################################################################

build: go.sum
ifeq ($(OS),Windows_NT)
	@echo "building soljuno binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o build/soljuno.exe ./cmd/soljuno
else
	@echo "building soljuno binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o build/soljuno ./cmd/soljuno
endif
.PHONY: build

###############################################################################
###                                 Install                                 ###
###############################################################################

install: go.sum
	@echo "installing soljuno binary..."
	@go install -mod=readonly $(BUILD_FLAGS) ./cmd/soljuno

###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

stop-docker-test:
	@echo "Stopping Docker container..."
	@docker stop soljuno-test-db || true && docker rm soljuno-test-db || true
.PHONY: stop-docker-test

start-docker-test: stop-docker-test
	@echo "Starting Docker container..."
	@docker run --name soljuno-test-db -e POSTGRES_USER=soljuno -e POSTGRES_PASSWORD=password -e POSTGRES_DB=soljuno -d -p 5433:5432 postgres
.PHONY: start-docker-test

coverage:
	@echo "viewing test coverage..."
	@go tool cover --html=coverage.out

test-unit: start-docker-test
	@echo "executing unit tests..."
	@go test -mod=readonly -v -coverprofile coverage.txt ./...

lint:
	golangci-lint run --out-format=tab

lint-fix:
	golangci-lint run --fix --out-format=tab --issues-exit-code=0
.PHONY: lint lint-fix


format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs goimports -w -local github.com/desmos-labs/juno
.PHONY: format

clean:
	rm -f tools-stamp ./build/**

.PHONY: install build ci-test ci-lint coverage clean
