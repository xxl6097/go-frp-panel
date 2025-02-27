set CGO_ENABLED=0 set GOOS=linux GOARCH=amd64 go build *.go
CGO_ENABLED=0 GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -ldflags "$ldflags -s -w -linkmode internal" -o ${distDir} ${appdir}

set CGO_ENABLED=0
set GOOS=linux
set GOARCH=arm64
go build -ldflags "-s -w -linkmode internal" -o z service.go

打版0.1.1

## client config
```
serverAddr = "127.0.0.1"
serverPort = 7000
user = "user1"
metadatas.token = "123"
```
