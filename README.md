# errstack

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/didip/errstack/blob/main/LICENSE)


A very small library to combine errors.

I want my errors to look like this:

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
    e := errstack.NewString("password field is missing")

    // 2. And then the second error occured.
    e.Append("company name is missing")

    // 3. And then the third error occured.
    e.Append("username is too short")

	log.Println(e.Error())
    // The log will look like this:
    // /path/to/project.go:68="username is too short" /path/to/project.go:65="company name is missing" /Users/didip/go/src/github.com/didip/errstack/errstack.go:15="password field is missing"
}
```