BIN ?= gocve

docker-build:
	@docker build -t gocve:1.0.0 .

docker-nw:
	@docker network create gocve

docker-shell:
	@docker run -it --rm --name gocve --network gocve -v `pwd`:/home/gouser gocve:1.0.0 /bin/bash

go-build:
	@go build -ldflags "-X github.com/jimmyislive/gocve/cmd/gocve/cli.buildVersion=1.0.0" -o ./_output/bin github.com/jimmyislive/gocve/cmd/gocve

