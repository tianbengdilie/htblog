BASEDIR=$(dirname $(dirname $(readlink -f $0)))
CURRENTDIR=$(pwd)
export GOOS=linux
export GIT_TERMINAL_PROMPT=1
export GOSUMDB=off
export GOINSECURE=oa.com
export GOFLAGS=-insecure
export GOPRIVATE=git.code.oa.com
export ROOT_PATH=$BASEDIR
unset TESTMOD
cd $BASEDIR
air
cd $CURRENTDIR