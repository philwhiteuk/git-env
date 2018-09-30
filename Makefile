default: test build


build:
	export GOPATH=$$(pwd) && go build -o git-env main.go

clean:
	rm -rf ./pkg ./git-env;

test: 
	export GOPATH=$$(pwd) && go test -timeout 30s commands
