# prototool

> This repository is a fork [uber/prototool](https://github.com/uber/prototool).

## Add new plugins

| Name                                                                            | Version |
|---------------------------------------------------------------------------------|---------|
| [protoc-gen-go](google.golang.org/protobuf/cmd/protoc-gen-go)                   | 1.27.1  |
| [protoc-gen-go-grpc](google.golang.org/grpc/cmd/protoc-gen-go-grpc)             | 1.2.0   |
| [protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate)        | 0.6.13  |
| [grpc-tools](https://github.com/grpc/grpc-node/tree/master/packages/grpc-tools) | 1.11.2  |
| [ts-protoc-gen](https://github.com/improbable-eng/ts-protoc-gen)                | 0.15.0  |

## Configuration

Example `prototool.yaml` wich generates both Golang code:

```yaml
protoc:
  version: 3.19.4
lint:
  group: uber2
generate:
  go_options:
    import_path: github.com/citilink/prototool/example/proto
  plugins:
    - name: go
      type: go
      output: ../gen
    - name: go-grpc
      type: go
      output: ../gen
    - name: validate
      type: go
      flags: lang=go
      output: ../gen
```

Example `prototool.yaml` wich generates both NodeJS and browser code:

```yaml
protoc:
  version: 3.19.4
lint:
  group: uber2
generate:
  plugins:
    # Shared
    - name: js
      flags: import_style=commonjs,binary
      output: ../gen
    # Node
    - name: grpc
      flags: grpc_js
      output: ../gen
      path: grpc_node_plugin
    - name: ts
      flags: service=grpc-node,mode=grpc-js
      output: ../gen
      path: protoc-gen-ts
    # Browser
    - name: grpc-web
      flags: import_style=commonjs+dts,mode=grpcwebtext
      output: ../gen
      path: protoc-gen-grpc-web
```

## Docker build

**x86_64**
```bash
docker build -t "citilink/prototool:latest" .
```

**aarch64**
```bash
docker build --platform=linux/x86_64 -t "citilink/prototool:latest" .
```

## How to use

**x86_64**
```bash
docker run --rm -v "$(pwd):/work" citilink/prototool:latest prototool generate
```

**aarch64**
```bash
docker run --rm --platform=linux/x86_64 -v "$(pwd):/work" citilink/prototool:latest prototool generate
```

## Documentation

For more information please refer to the [original documentation](https://github.com/uber/prototool/blob/dev/docs/docker.md).
