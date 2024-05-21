.PHONY : test run docker-build docker-run build

cur := $(shell pwd)

run:
	@go run main.go server -l true -c configs/dev.yaml

build:
	@go build -v -o bin/match-dating ./main.go

test: 
	@go clean -testcache & go test -v ./test/...

docker-build:
	@docker build --tag n960321/match-dating:latest --file build/dockerfile .

docker-run:
	@docker run --name match-dating \
	-p 8080:8080 \
	--volume $(cur)/configs:/app/configs \
	n960321/match-dating:latest
