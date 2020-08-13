# gRPC-Go-Web-TODO

Svelte TODO App with gRPC-WEB, gRPC-Go

## Prerequisites

- **[Go][]**: [latest releases][go-releases].
- **[Node][]**: [latest releases][node-release].

## Development

- [**mkcert** install](mkcert-install)
- [**github-dl** install](github-dl-install)

```bash
# root CA install
mkcert -install

# install package & tool
mage install

# generate api & code
mage gen

# livereload dev
mage dev

# build & embedding
mage build
```

[Go]: https://golang.org
[Node]: https://nodejs.org
[go-releases]: https://golang.org/doc/devel/release.html
[node-release]: https://nodejs.org/en/blog
[mkcert-install]: https://github.com/FiloSottile/mkcert#installation
[github-dl-install]: https://github.com/iwaltgen/github-dl#installation
