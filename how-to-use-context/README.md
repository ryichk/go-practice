# How to use context

## Cancel processing using context

- A context for propagating the cancellation process can be created with the `context.WithCancel` function.
- Cancellation can be indicated with the `cancel` function obtained from the `context.WithCancel` function.
- When canceled by the `cancel` function, the channel obtained from the `Done` method of the context will be closed to determine whether it is canceled or not.

## Reference

- https://zenn.dev/hsaki/books/golang-context/viewer/done
