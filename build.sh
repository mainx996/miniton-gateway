#!/bin/sh
RM=`command -v rm`
GO=`command -v go`
MKDIR=`command -v mkdir`
CP=`command -v cp`
MV=`command -v mv`
TAR=`command -v tar`
$RM -rf log
$RM -rf bin
$MKDIR bin
$CP -rf config bin/config

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $GO build -o miniton main.go

$MV miniton bin/miniton
$TAR -czvf miniton.tar bin
$RM -rf bin