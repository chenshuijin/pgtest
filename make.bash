if [ ! -f make.bash ]; then
    echo 'make.bash must be run under its container folder' 1>&2
    exit 1
fi
CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

go fmt webapi
go build -o ./bin/webapi webapi
export GOPATH="$OLDPATH"