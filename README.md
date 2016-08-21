# Concurrent Map

> A thread-safe concurrent map for the Go programming language.

[![License][License-Image]][License-Url] [![ReportCard][ReportCard-Image]][ReportCard-Url] [![GoDoc][GoDoc-Image]][GoDoc-Url]

## Install

``` bash
$ go get -u github.com/LyricTian/cmap
```

## Usage

``` go
package main

import (
	"fmt"

	"github.com/LyricTian/cmap"
)

func main() {
	m := cmap.NewShardMap()
	// or
	// m := cmap.NewMap()
	m.Set("foo", "bar")
	if v, ok := m.Get("foo"); ok {
		fmt.Println("foo=", v.(string))
	}
	m.Remove("foo")
}
```

## MIT License

```
Copyright (c) 2016 Lyric
```

[License-Url]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/npm/l/express.svg
[ReportCard-Url]: https://goreportcard.com/report/github.com/LyricTian/cmap
[ReportCard-Image]: https://goreportcard.com/badge/github.com/LyricTian/cmap
[GoDoc-Url]: https://godoc.org/github.com/LyricTian/cmap
[GoDoc-Image]: https://godoc.org/github.com/LyricTian/cmap?status.svg