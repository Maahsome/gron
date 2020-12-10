# gron

The [gron](https://github.com/tomnomnom/gron) go cli tool is very cool, and I've
discovered some uses for the conversion between gron and json in some other code,
so I wanted to make the functionality reusable as an importable module, rather
than simply pull the go files into the various projects.

## Usage

The interface for the this uses `.ToGron()` and `.ToJSON()` to represent the
gron and ungron methods of the original cli tool.

We create a gron instance by passing an io.Reader and an io.Writer.

`func NewGron(reader io.Reader, writer io.Writer) *Gron`

When we are wanting to manipulate the output data, we aren't likely to pass an
io.Writer that outputs to anything other than a byte buffer to capture the data
to a variable.  Though we can simply pass an io.Writer that writes to stdout
similar to the way the cli tool does with `colorable.NewColorableStdout()`.

### Pull down the module

`go get github.com/maahsome/gron`

### Example for .ToGron()

```go
package main

import (
    "bytes"
    "fmt"
    "strings"
    "github.com/maahsome/gron"
)

func main() {

    // json object to gron
    jsonData := `{
"book":[
    {
        "id":"444",
        "language":"C",
        "edition":"First",
        "author":"Dennis Ritchie"
    },
    {
        "id":"555",
        "language":"C++",
        "edition":"Second",
        "author":"Bjarne Stroustrup"
    }
]
}
`
    reader := strings.NewReader(jsonData)
    out := &bytes.Buffer{}
    ge := gron.NewGron(reader, out)
    ge.SetMonochrome(true)
    err := ge.ToGron()
    if err != nil {
        fmt.Println(err)
    }
    // Our out variable now contains the gron notation lines, we can work with
    // that data.
    // We will output the data to verify that the data was generated as expected.
    fmt.Println(out)
}
```

### Example for .ToJSON()

```go
package main

import (
    "bytes"
    "fmt"
    "strings"
    "github.com/maahsome/gron"
)

func main() {
    // gron object to ungron
    gronData := `json = {};
json.book = [];
json.book[0] = {};
json.book[0].author = "Dennis Ritchie";
json.book[0].edition = "First";
json.book[0].id = "444";
json.book[0].language = "C";
json.book[1] = {};
json.book[1].author = "Bjarne Stroustrup";
json.book[1].edition = "second";
json.book[1].id = "555";
json.book[1].language = "C++";
`

    reader := strings.NewReader(gronData)
    out := &bytes.Buffer{}
    ge := gron.NewGron(reader, out)
    ge.SetMonochrome(true)
    err := ge.ToJSON()
    if err != nil {
        fmt.Println(err)
    }

    // Our out variable contains json, which we can manipulate as we need.
    // We will just output to see that everything worked as expected.
    fmt.Println(out)
}
```
