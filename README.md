# gRPC-Go-Web-TODO

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
- **[Node][]**: [latest releases][node-release].

## Development

- [**mkcert** install][mkcert-install]

```sh
# on macOS
brew install mkcert

# root CA install
mkcert -install

# install package & tool
mage install

# generate api & code
mage gen

# livereload development
mage dev

# build & embedding for a single binary
mage build
```

[Go]: https://golang.org
[Node]: https://nodejs.org
[go-releases]: https://golang.org/doc/devel/release.html
[node-release]: https://nodejs.org/en/blog
[mkcert-install]: https://github.com/FiloSottile/mkcert#installation
