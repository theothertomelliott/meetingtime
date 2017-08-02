all: test

ifndef TRAVIS
test:
	go test -cover -race ./...
else
test:
	go test -v -covermode=count -coverprofile=coverage.out
	goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $(COVERALLS_TOKEN)
endif