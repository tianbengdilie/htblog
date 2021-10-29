#!/usr/bin/env bash
SystemName=$(uname)

if [[ $SystemName -eq "Darwin" ]] 
then
    BASEDIR=$(dirname $(dirname $(greadlink -f $0)))
elif [[$SystemName -eq "Linux" ]] 
then
    BASEDIR=$(dirname $(dirname $(readlink -f $0)))
fi

CURRENTDIR=$(pwd)
 
if [[ $SystemName -eq "Darwin" ]] 
then
export GOOS=darwin
elif [[$SystemName -eq "Linux" ]] 
then
export GOOS=linux
fi

export GIT_TERMINAL_PROMPT=1
export GOSUMDB=off
export GOINSECURE=oa.com
export GOFLAGS=-insecure
export GOPRIVATE=git.code.oa.com
export ROOT_PATH=$BASEDIR
# unset TESTMOD
cd $BASEDIR
air
cd $CURRENTDIR