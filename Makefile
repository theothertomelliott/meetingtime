all: test

test:
	go test -v -covermode=count -coverprofile=coverage.out
	goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $(COVERALLS_TOKEN)