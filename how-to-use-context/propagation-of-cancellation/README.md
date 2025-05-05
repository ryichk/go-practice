# Propagation fo cancellation

- Same behavior when the same context is passed to multiple goroutines, whether in series or parallel.
- Since the context that controls groutine life/death is the same, the cancellation timing is also linked.
