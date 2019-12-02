BIN ?= gocve
VERSION=`cat VERSION`
.PHONY : test

docker-build:
	@docker build -t gocve:${VERSION} .

docker-nw:
	@docker network create gocve

docker-shell:
	@docker run -it --rm --name gocve --network gocve -v `pwd`:/home/gouser gocve:${VERSION} /bin/bash

go-build-linux-amd64:
	@env GOOS=linux GOARCH=amd64 go build -ldflags "-X github.com/jimmyislive/gocve/cmd/gocve/cli.buildVersion=${VERSION}" -o ./_output/bin/linux/amd64/gocve-linux-amd64 github.com/jimmyislive/gocve/cmd/gocve

go-build:
	@make go-build-linux-amd64

test:
	@cd cmd/gocve/cli && go test . -v
