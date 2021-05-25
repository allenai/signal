# signal

The signal package implements helpers for signal handling in Go.

```golang
package main

func main() {
	ctx := signal.WithSignal(ctx)
	<-ctx.Done() // Blocks until Ctrl+C
}
```
