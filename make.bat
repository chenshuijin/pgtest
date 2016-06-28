@echo off

echo set OLDGOPATH=%GOPATH%
set GOPATH=%CD%

go fmt webapi
go build -o ./bin/webapi.exe webapi
