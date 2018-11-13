#!/usr/bin/env bash
# Stops the process if something fails
set -xe

# All of the dependencies needed/fetched for the project should be given here.

#dep ensure will create the required dependency folder(vendor) only inside $GOPATH/src.
#Hence creating a temporary folder under $GOPATH/src where the application is built.
#And then copying it to the $APP_STAGING_DIR

APP_BUILD_DIR="$GOPATH/src/to-be-defined"  # We will build the app here
# Remove the $APP_BUILD_DIR just in case it was left behind in a failed build.
rm -rf $APP_BUILD_DIR
mkdir -p $APP_BUILD_DIR

#move the content of of $APP_STAGING_DIR to APP_BUILD_DIR
mv * .[^.]* $APP_BUILD_DIR
cd $APP_BUILD_DIR

export INSTALL_DIRECTORY="./bin"
sudo mkdir $INSTALL_DIRECTORY
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
./bin/dep ensure -vendor-only

# create the application binary that eb uses
go build -o bin/application application.go
echo "Builded the application successfully"

mv * .[^.]* $APP_STAGING_DIR
cd $APP_STAGING_DIR
rm -rf $APP_BUILD_DIR

