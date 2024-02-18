# getenv

A type-safe wrapper for `os.Getenv` that returns a default value if the environment variable is not set.


## Installation

```bash
go get github.com/josestg/getenv
```

## Usage

```go
package main

import (
    "fmt"
    "time"

    "github.com/josestg/getenv"
)

func main() {
    var (
        host             = getenv.String("HOST", "0.0.0.0")
        ports            = getenv.Ints("PORTS", []int{8080, 8081, 8082})
        debug            = getenv.Bool("DEBUG", false)
        threshold        = getenv.Float("THRESHOLD", 0.5)
        maxBodySize      = getenv.Int("MAX_BODY_SIZE", uint64(1<<20)) // 1MB
        shutdownTimeout  = getenv.Duration("SHUTDOWN_TIMEOUT", 5*time.Second)
        whitelistHeaders = getenv.Strings("WHITELIST_HEADERS", []string{"Content-Type", "Content-Length"})
    )

    fmt.Println(
        host,
        ports,
        debug,
        threshold,
        maxBodySize,
        shutdownTimeout,
        whitelistHeaders,
    )
}
```

## License

[MIT](LICENSE)