.DEFAULT_GOAL := build
.PHONY: build install docker dockerpush

REPO=linkpoolio/alphavantage-adapter
LDFLAGS=-ldflags "-X github.com/linkpoolio/alphavantage-adapter/store.Sha=`git rev-parse HEAD`"

build:
	@go build $(LDFLAGS) -o alphavantage-adapter

install:
	@go install $(LDFLAGS)

docker:
	@docker build . -t $(REPO)

dockerpush:
	@docker push $(REPO)