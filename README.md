# signal

The signal package implements helpers for signal handling in Go.

```go
package main

import (
	"context"

	"github.com/allenai/signal"
)

func main() {
	ctx, cancel := signal.WithSignal(context.Background())
	defer cancel()

	<-ctx.Done() // Blocks until Ctrl+C
}
```
