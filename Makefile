
build: deps
	@echo "Build..."
	@go build .

deps:
	@echo "Fetch deps..."
	@go get github.com/codegangsta/cli
	@go get github.com/parnurzeal/gorequest
	@go get github.com/fatih/color
	@go get github.com/nightlyone/lockfile


