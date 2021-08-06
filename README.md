# gRPC-Go-Web-TODO

[![build](https://github.com/iwaltgen/grpc-go-web-todo/actions/workflows/build.yml/badge.svg)](https://github.com/iwaltgen/grpc-go-web-todo/actions/workflows/build.yml)

Interactive TODO App powered by the gRPC-WEB server stream.

```
    +----------Browser----------+                 +----------Backend----------+
    |                           |                 |                           |
    |                           |                 |    static file serving    |
    |                           |                 |                           |
    |         gRPC-Web          | <-- HTTP/2 -->  | gRPC-Web in-process-proxy |
    |                           |                 |                           |
    |                           |                 |          gRPC-go          |
    |                           |                 |                           |
    +---------------------------+                 +---------------------------+
```

## Prerequisites

- **[Go][]**: [latest releases][go-releases].
- **[Node][]**: [latest LTS releases][node-release].

## Development

- [**mkcert** install][mkcert-install]

```sh
# on macOS
brew install mkcert

# root CA install
mkcert -install

# install package & tool
mage install

# dev mode
mage dev

# prod build
mage build
```

[go]: https://golang.org
[node]: https://nodejs.org
[go-releases]: https://golang.org/doc/devel/release.html
[node-release]: https://nodejs.org/en/blog
[mkcert-install]: https://github.com/FiloSottile/mkcert#installation
