#!/bin/bash

go get github.com/joho/godotenv

cd src
go build -o ../gctl main.go utils.go colour.go genny.go token.go cache.go entity.go search.go rules.go blacklist.go
cd ..
