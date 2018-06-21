# M3lsh

A library for exception handling in Golang, it simply uses `panic` with the ability to `Try`/`TryCatch` exceptions and `Throw` custom exceptions with details of stack trace and thread info.
You can use multiple catch "statements", only the first matching statement will be executed.
Due to golang's limitations, exceptions will be sent as interface{} to your method, but they will always be castable to the type you specify.
Catching a `&m3lsh.BaseException{}` will catch any exception.

## Installation

`go get github.com/mohamed-essam/m3lsh`

Or using dep, add to your Gopkg.toml

```
[[constraint]]
  name = "github.com/mohamed-essam/m3lsh"
  revision = "<latest master revision>"
```

## Usage

```
import (
  "m3lsh"
  "fmt"
)

type MyException struct {
  m3lsh.BaseException
}

func main() {
	m3lsh.TryCatch(func() {
    m3lsh.Throw(&MyException{}, "Cats")
	}, m3lsh.Catcher(&MyException{}, func(e interface{}){
    fmt.Printf("%T", ex) // *MyException
    ex := e.(*MyException)
    fmt.Printf(ex.Message) // Cats
  }), m3lsh.Catcher(&m3lsh.BaseException{}, func(e interface{}){
    // This will catch all exceptions other than MyException
  })
}
```
