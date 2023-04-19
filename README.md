# Advance golang course: Protobuf and gRPC

---

This is the documentation for my Golang gRPC project. 
This project is part of my learning process of golang and gRPC. 
To understand `microservices`.

## What is a protocol buffer?

Protocol buffers are Google's language-neutral, platform-neutral, extensible mechanism for 
serializing structured data – think XML, but smaller, faster, and simpler. 
You define how you want your data to be structured once, then you can use special 
generated source code to easily write and read your structured data to and from a 
variety of data streams and using a variety of languages.

Official documentation: https://developers.google.com/protocol-buffers

## What is gRPC?

gRPC is a modern, open source, high-performance remote procedure call (RPC) framework

Official documentation: https://grpc.io/docs/

---

## Before staring

### Install protoc compiler (version 3.13.0)

```bash
# Linux
apt install -y protobuf-compiler
protoc --version 
```
```bash
# Or MacOS
brew install protobuf
protoc --version 
```

### Install protoc-gen-go and protoc-gen-go-grpc

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest  
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
### Install protobuf  and gRPC dependencies

```bash
go get google.golang.org/protobuf
go get google.golang.org/grpc
```

### Install postgres driver

```bash
go get github.com/lib/pq
```

---

## To start

### Compile the proto file to golang code using protoc

This command will generate the golang code for the proto file.
example: `student.pb.go` in the proto folder

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative studentpb/student.studentpb
```

+ `--go_out` is the output path for the generated golang code
+ `--go_opt` is the option for the generated golang code
+ `--go-grpc_out` is the output path for the generated golang code
+ `--go-grpc_opt` is the option for the generated golang code
+ `proto/student.proto` is the path to the proto file

### Run DockerFile to build the postgres db

```bash
cd database
docker build -t github-thrashy-grpc .
```

### Run the postgres db

```bash
docker run -d -p 54321:5432 --name postgres-grpc github-thrashy-grpc
```