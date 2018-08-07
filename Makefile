.DEFAULT_GOAL := build
.PHONY: dep build install docker dockerpush

REPO=linkpoolio/asset-price-cl-ea
LDFLAGS=-ldflags "-X github.com/linkpoolio/alpha-vantage-cl-ea/store.Sha=`git rev-parse HEAD`"

dep:
	@dep ensure

build: dep
	@go build $(LDFLAGS) -o alpha-vantage-cl-ea

install: dep
	@go install $(LDFLAGS)

docker:
	@docker build . -t $(REPO)

dockerpush:
	@docker push $(REPO)