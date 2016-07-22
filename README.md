# fullstory

[![wercker status](https://app.wercker.com/status/5c617e0ba84e532e22029444f79d835f/s/master "wercker status")](https://app.wercker.com/project/bykey/5c617e0ba84e532e22029444f79d835f)
[![GoDoc](https://godoc.org/github.com/nishanths/fullstory?status.svg)](https://godoc.org/github.com/nishanths/fullstory)

Package `fullstory` implements a client for the
[fullstory.com](https://fullstory.com) API.

# Docs

See [godoc](https://godoc.org/github.com/nishanths/fullstory).

# Test

```
go test -race 
```

# Example

```
package main

import (
	"fmt"
	"log"

	"github.com/nishanths/fullstory"
)

func main() {
	client := fullstory.NewClient("API token")

	s, err := client.Sessions(15, "foo", "hikingfan@gmail.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}
```

# License

[MIT](https://nishanths.mit-license.org).
