// lib-instance-gen-go: File auto generated -- DO NOT EDIT!!!
.DEFAULT_GOAL=build

include Makefile.db

build:
	go fmt ./...
	go vet ./...
	go build -o bin/tv-tracker app/*.go

install:
	cp bin/tv-tracker /usr/local/sbin/tv-tracker

golib-latest:
	go get -u github.com/go-playground/validator/v10@latest
	go get -u github.com/google/uuid@latest
	go get -u github.com/labstack/echo/v4@latest
	go get -u github.com/mattn/go-sqlite3@latest
	go get -u github.com/patrickmn/go-cache@latest
	go get -u golang.org/x/crypto@latest
	go get -u github.com/skeletonkey/lib-core-go@latest
	go get -u github.com/skeletonkey/lib-instance-gen-go@latest

	go mod tidy

app-init:
	go generate
