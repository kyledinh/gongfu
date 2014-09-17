#!/bin/bash
go get github.com/gin-gonic/gin
go get github.com/gorilla/securecookie
go get github.com/julienschmidt/httprouter
go get github.com/lib/pq
go get github.com/russross/blackfriday
go get code.google.com/p/go.crypto/bcrypt

echo ""
echo "Go environments: "
go env
echo ""
echo "Go packages: "
go list 
