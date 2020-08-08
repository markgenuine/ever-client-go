.PHONY: clean build

export GOBIN := ${PWD}/bin
export GO111MODULE=on #work or not
export CGO_ENABLED=1

NAME := goton
#VERSION :=
#REVISION := 

default: build

clean: 
	@echo Cleaning...
	@rm bin/* -rf
	@echo Done.

build: 
	@echo build
	@env go build -ldflags="-r /lib/darwin" -v -o ./bin/$(NAME)
	@echo Done.
	
build-linux:
	@echo build to linux
	@env GOOS=linux GOARCH=amd64 go build -ldflags="-r /lib/linux" -v -o ../bin/$(NAME)
	@echo Done.

build-windows:
	@echo build to windows
	@env GOOS=windows GOARCH=amd64 go build -ldflags="-r /lib/windows" -v -o ../bin/$(NAME)
	@echo Done.

 