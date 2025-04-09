# gomanuf

[![Discord Server Badge](https://img.shields.io/discord/1358456011316396295?logo=discord)](https://discord.gg/xj6y5ZaTMr) [![Go Reference](https://pkg.go.dev/badge/github.com/kkrypt0nn/gomanuf.svg)](https://pkg.go.dev/github.com/kkrypt0nn/gomanuf) ![Repository License](https://img.shields.io/github/license/kkrypt0nn/gomanuf?style=flat-square) ![Code Size](https://img.shields.io/github/languages/code-size/kkrypt0nn/gomanuf?style=flat-square) ![Last Commit](https://img.shields.io/github/last-commit/kkrypt0nn/gomanuf?style=flat-square)

gomanuf is a very simple library to get the manufacturer of a specific MAC address.

## Installation
If you want to use this library for one of your projects, you can install it like any other Go library

```shell
go get github.com/kkrypt0nn/gomanuf
```

## Example
```go
package main

import (
	"fmt"
	"github.com/kkrypt0nn/gomanuf"
)

func main() {
	manuf, _ := gomanuf.Search("C4:A8:1D:73:D7:8C")
	fmt.Println(manuf)
}

```

## License
This library was made with 💜 by Krypton and is under the [MIT](LICENSE.md) license.
