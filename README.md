# Golang gRPC Examples

## Install Dependencies

Install gRPC...

```bash
$ go get -u google.golang.org/grpc
```

Install Protobuf v3...

```bash
# OSX
$ brew install protobuf

# else
https://github.com/google/protobuf/releases
```

Install protobuf plugin for Go...

```bash
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

Update your path (ie: '~/.bash_profile')...

```
export PATH=$PATH:$GOPATH/bin
```

## Generate Go Code

Generate the Go code from the models...

```bash
$ protoc -I howdy/ howdy/howdy.proto --go_out=plugins=grpc:howdy
```

## Run the client / server

```bash
# server
$ go run howdy/server/main.go

# client
$ go run howdy/client/main.go
RESPONSE: Howdy Mark!
```
