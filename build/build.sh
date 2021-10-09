BASEDIR=$(dirname $(dirname $(readlink -f $0)))

export APPENV=production
export GOARCH=amd64
export GOOS=linux
export CGO_ENABLED=0
go build -ldflags "-w" -o $BASEDIR/bin/svrmain $BASEDIR/main.go
