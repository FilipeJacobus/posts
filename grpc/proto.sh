#!/usr/bin/env sh

protoc --go-grpc_out=. hello.proto --go_out=.