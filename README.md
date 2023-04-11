# Go Masking

Masking any character with **[Go](https://go.dev)** language.

## How to use

You must install this package into your module

```shell
go get github.com/pandudpn/go-masking
```

After installing, you can run code like this

```go
package main

import (
	masking "github.com/pandudpn/go-masking"
	"fmt"
)

func main() {
	// example want to mask password
	pwd := "abcdefghijklmnopqrstuvwxyz"
	result := masking.String(masking.Password, pwd) // will return *********
	fmt.Println(result)
}
```

Or you can check the [Example](example/main.go).
