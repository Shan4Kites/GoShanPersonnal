#!/usr/bin/env bash
# Stops the process if something fails
set -xe

# All of the dependencies needed/fetched for the project should be given here.
go get "github.com/sirupsen/logrus"
go get "github.com/onrik/logrus/filename"
go get "github.com/gorilla/mux"
go get "github.com/rs/xid"
go get "golang.org/x/net/context"

# create the application binary that eb uses
go build -o bin/application application.go
