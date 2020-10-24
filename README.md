# Protobuffers

- Respositório: https://github.com/protocolbuffers/protobuf
- Instalação:   https://github.com/protocolbuffers/protobuf/blob/master/src/README.md

# Golang GRPC and Protobuffers:

- Repositório: https://github.com/golang/protobuf
- Compilação: https://developers.google.com/protocol-buffers/docs/reference/go-generated
- GRPC: https://godoc.org/google.golang.org/grpc

```bash
    # Download packages
    go get -u -v github.com/golang/protobuf/...
    go get -u -v google.golang.org/grpc/...
```

```cmd
    # Windows install
    go install github.com\golang\protobuf\protoc-gen-go
```

```bash

    cd golang-grpc-examples

    # Generate Code
    protoc -I=./defs --go_out=plugins=grpc:. ./defs/*.proto

    # -I=./defs       -> indica o diretorio onde estão os arquivos .proto
    # --go_out=.      -> indica o diretorio onde serão gerados os arquivos .go, neste caso o diretório está especificado no arquivo defs/message.proto: option go_package = "messages/messenger";
    # ./defs/*.proto  -> indica quais os arquivos que vão ser gerados, neste caso será todos.
```