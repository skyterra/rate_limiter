# Rate Limiter

A lightweight in-memory rate limiter manager for Go using LRU cache.  
Designed for controlling operation frequency per key to prevent overuse, abuse, or overload.

## Features

- **LRU Cache**: Stores up to 10,240 rate limiters, evicting least recently used ones.
- **Flexible Limiter Creation**: Configure requests per second and burst size.
- **Safe Access**: Get existing limiter or return `nil` if not present.
- **Use Cases**:
	- API request throttling
	- Anti-bot operation control
	- Internal service rate limiting

## Installation

```bash
go get github.com/skyterra/rate-limit
```


## Usage
```go
package main

import (
    "fmt"
    "time"
	"github.com/skyterra/rate_limiter"
)

func main() {
    // Create a new rate limiter for "user123"
    limiter := rate_limiter.NewRateLimiter("user123", 5, 10) // 5 req/sec, burst 10

    // Check if an operation is allowed
    if limiter.Allow() {
        fmt.Println("Request allowed")
    } else {
        fmt.Println("Request denied")
    }

    // Retrieve existing limiter
    sameLimiter := rate_limiter.GetRateLimiter("user123")
    fmt.Println("Limiter exists:", sameLimiter != nil)

    // Example: wait for next allowed request
    limiter.Wait(nil)
}
```
