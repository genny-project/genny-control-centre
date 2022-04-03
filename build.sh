#!/bin/bash

go get github.com/joho/godotenv

go build -o gctl src/*.go
