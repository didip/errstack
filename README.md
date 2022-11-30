# errstack

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/didip/errstack/blob/main/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/didip/errstack.svg)](https://pkg.go.dev/github.com/didip/errstack)

A very small library to combine errors.

Furthermore, I want my errors to display the filename and line number.

```
pool.go:19="missing max value" users.go:25="missing password field"
```

## Example

```
package main

import (
    "log"
)

func main() {
    // 1. Something bad happened at the beginning of your stack.
    e := errstack.New("password field is missing")

    // 2. And then the second error occured.
    e.Append("company name is missing")

    // 3. And then the third error occured.
    e.Append("username is too short")

    log.Println(e.Error())
    // The log will look like this:
    // /path/to/project.go:68="username is too short" /path/to/project.go:65="company name is missing" /path/to/project.go:15="password field is missing"
}
```

## My other Go libraries

* [Tollbooth](https://github.com/didip/tollbooth): A generic middleware to rate-limit HTTP requests.

* [Stopwatch](https://github.com/didip/stopwatch): A small library to measure latency of things. Useful if you want to report latency data to Graphite.

* [LaborUnion](https://github.com/didip/laborunion): A dynamic worker pool library.

* [Gomet](https://github.com/didip/gomet): Simple HTTP client & server long poll library for Go. Useful for receiving live updates without needing Websocket.
