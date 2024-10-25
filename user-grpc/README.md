## Install gRPC Tools

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Set Path

```
export PATH=$PATH:$(go env GOPATH)/bin
```

## Source Bash
```
source ~/.bashrc
```

## Check Version

```
protoc-gen-go --version
protoc-gen-go-grpc --version
```

## Generate Code Proto
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/user.proto
```
