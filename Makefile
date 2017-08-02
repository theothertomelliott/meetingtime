all: test

ifndef TRAVIS
test:
	go test -cover -race ./...
else
test:
	./cover_packages.sh
	goveralls -coverprofile=profile.cov -service=travis-ci -repotoken $(COVERALLS_TOKEN)
endif