[![go.mod](https://img.shields.io/github/go-mod/go-version/koki-develop/hebr)](./go.mod)
[![release](https://img.shields.io/github/v/release/koki-develop/hebr)](https://github.com/koki-develop/hebr/releases/latest)
[![GitHub Actions](https://github.com/koki-develop/hebr/actions/workflows/main.yml/badge.svg)](https://github.com/koki-develop/hebr/actions/workflows/main.yml)
[![Maintainability](https://api.codeclimate.com/v1/badges/142e9cd2d5a120f2a6dc/maintainability)](https://codeclimate.com/github/koki-develop/hebr/maintainability)
[![LICENSE](https://img.shields.io/github/license/koki-develop/hebr)](./LICENSE)
[![Twitter Follow](https://img.shields.io/twitter/follow/koki_develop?style=social)](https://twitter.com/koki_develop)

# hebr

`hebr` is a command line tool or library for encoding/decoding data in Hebrew characters.

## Contents

- [CLI](#cli)
    - [Installation](#installation)
    - [Usage](#usage)
        - [Encode](#encode)
        - [Decode](#decode)
- [Library](#library)
    - [Installation](#installation-1)
    - [Import](#import)
    - [Usage](#usage-1)
- [LICENSE](#license)

## CLI

### Installation

```sh
$ go install github.com/koki-develop/hebr/cmd/hebr@latest
```

### Usage

#### Encode

Data to be encoded can be passed from stdin.

```sh
$ echo hello | hebr
םבעצר,דןבט
```

It is also possible to specify the target file to be encoded.

```sh
$ hebr ./hello.txt
םבעצר,דןבט
```

#### Decode

You can decode data encoded in Hebrew characters by specifying the `-d` or `-decode` flag.  
Usage is the same as Encode.

```sh
$ echo 'םבעצר,דןבט' | hebr -d
hello
```

```sh
$ hebr --decode ./hello.txt.hebr
hello
```

## Library

### Installation

```
$ go get github.com/koki-develop/hebr
```

### Import

```go
import "github.com/koki-develop/hebr"
```

### Usage

`hebr` package provides `Encode()` and `Decode()` functions.

```go
package main

import (
	"fmt"
	"log"

	"github.com/koki-develop/hebr"
)

func main() {
	data := []byte("data")

	// encode data
	encoded, err := hebr.Encode(data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(encoded))
	// => לסנקטרט

	// decode data
	decoded, err := hebr.Decode(encoded)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(decoded))
	// => data
}

```

## LICENSE

[MIT](./LICENSE)
