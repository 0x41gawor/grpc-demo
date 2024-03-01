# gRPC Demo
The business case is a server that generates invoices. 

Just look at the [`.protobuf`](invoicer.proto) file how it describes the remote procedure calls. More theory [here](docs/main.md).


## protobuf file

```protobuf
service Invoicer {
    rpc Create(CreateRequest) returns (CreateResponse);
}
```
The service which we are implementing on a business level is a invoice generator. Client passes some information to fill up the invoice and server returns byte files. 

Here on protobuf code you can see that service `Invoicer` has a method/procedure `Create`. This method takes some args and return some values. Here both of these we call messages as this is all remote and we define The Protocol.
```protobuf
message Amount {
    int64 amount = 1;
    string currency = 2;
}

message CreateRequest {
    Amount amount = 1;
    string from = 2;
    string to = 3;
}

message CreateResponse {
    bytes pdf = 1;
    bytes docx = 2;
}
```

## Generating code from protobuf file

```sh
# installation of protobuf complier (protoc) 
# complier creates .go files from .proto files
# provides programming interface in go for .proto defined data structs
sudo apt install protobuf-complier
# interface between protoc a golang
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
export PATH=$PATH:$(go env GOPATH)/bin
# to create go code in current dir
protoc greet.proto --go_out=$(pwd)
```