default: build

clean:
	@rm ./git-env

build:
	@go build -o git-env src/*.go
