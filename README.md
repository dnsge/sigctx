# sigctx

A simple helper for finishing a `context.Context` when the process receives a `SIGINT` or `SIGTERM`.

Example:

```go
func main() {
    ctx, cancel := sigctx.NewShutdownContext()
    defer cancel()

    DoStuff(ctx) // Ctrl+C will cancel this context
}
```
