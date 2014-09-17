#!/bin/bash

sudo npm install grunt -g
sudo npm install grunt-cli -g
go get github.com/gin-gonic/gin

echo ""
echo "OS :"
uname -a
lsb_release -a
echo ""


# Check installation
echo "Postgresql version: "
psql --version
echo ""
echo "Node version: "
node --version
echo ""
echo "NPM version: "
npm -version
echo ""
echo "Grunt version: "
grunt --version
echo ""
echo "Go version: "
go version
