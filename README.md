# Foreign At Home

Foreign At Home is an application allowing foreign language learners to find
language speaking and language studying groups in their vicinity.

## Building

### Prerequisites

* [Go SDK installed](https://go.dev/doc/install)
* [protoc installed](https://grpc.io/docs/protoc-installation/)
* Go plugins for the protocol compiler:
  1. Install the protocol compiler plugins for Go using the following commands:

  ```shell
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
  ```
  
  2. Update your PATH so that the protoc compiler can find the plugins:

  ```shell
  export PATH="$PATH:$(go env GOPATH)/bin"
  ```

### Now just do as I say

1. Clone repository.
1. Use `go mod download` to download all backend dependencies
1. Use yarn to install frontend dependencies:
`cd front & yarn init`
