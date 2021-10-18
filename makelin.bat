@echo off

set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -o poster_linux main.go

echo make exe files is ok!
