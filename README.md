# filestr

Convert files to string or byte slice variables.

## Installing

```
go get github.com/yhat/filestr
```

## Example

Consider if you have two files a `main.go` and a `VERSION` file.

`main.go` might look like this.

```go
package main

import "fmt"

//go:generate filestr -trim VERSION version.go myVersion

func() {
    fmt.Println(myVersion)
}
```

And `VERSION` might look like this

```
1.2.3
```

Running `go generate` will create a file called `version.go` which looks like this.

```go
package main

var myVersion string = "1.2.3"
```
