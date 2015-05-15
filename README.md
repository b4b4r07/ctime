ctime 
==========

Usage
------------
File Create Times for Golang

Go has a hidden ctime function for most platforms, this repo makes it accessible.

```go
package main

import (
  "log"
  "github.com/b4b4r07/ctime"
)

func main() {
  ct, err := ctime.Stat("myfile")
  if err != nil {
    log.Fatal(err)
  }
  log.Println(ct)
}
```

Installation
------------
```sh
go get github.com/b4b4r07/ctime
```

License
------------

MIT