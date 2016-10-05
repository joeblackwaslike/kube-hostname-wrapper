USER = joeblackwaslike
PROJECT = kube-hostname-wrapper
TAG = $(shell git tag | sort -n | tail -1)

BIN_FILES = $(shell ls bin)


build:
	@docker run --rm -it -v $$GOPATH:/go -w /go/src/github.com/$(USER)/$(PROJECT) golang:cross sh -c 'export GOARCH=amd64; for GOOS in darwin linux; do echo "Building $$GOOS-$$GOARCH"; export GOOS=$$GOOS; go build -v -o bin/$(PROJECT)-$$GOOS; done'

clean:
	@rm -rf bin/*

bump-tag:
	@git tag -a $(shell echo $(TAG) | awk -F. '1{$$NF+=1; OFS="."; print $$0}') -m "New Release"

commit-all:
	@git add .
	@git commit

push:
	@git push origin master

release:
	@-git push origin $(TAG)
	@github-release release --user $(USER) --repo $(PROJECT) --tag $(TAG)

upload-release:
	@for f in $(BIN_FILES); do github-release upload --user $(USER) --repo $(PROJECT) --tag $(TAG) --name "$$f" --file "bin/$$f"; done

default: build
