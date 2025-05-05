# How to use context

## Cancel processing using context

- A context for propagating the cancellation process can be created with the `context.WithCancel` function.
- Cancellation can be indicated with the `cancel` function obtained from the `context.WithCancel` function.
- When canceled by the `cancel` function, the channel obtained from the `Done` method of the context will be closed to determine whether it is canceled or not.

## Values that should not be passed to context

- Arguments that can change function behavior
- Values that should not be of type unsafe
- Variable values
- Values that are not goroutine safe

## Values to be passed to context

- Request-scoped value

### What is request scope?

Request scope is a property that is "shared while a single request is being processed".

Example:

- User ID extracted from the header
- Authentication token
- Process ID attached on the server side for tracing purpose
- etc...

## Reference

- https://zenn.dev/hsaki/books/golang-context/viewer/done
