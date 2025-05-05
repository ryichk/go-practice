# Cleanup

## Usage of the AfterFunc function

- Release all processes registered in `sync.Cond` when the context is canceled.
- Stop reading from a connection when the context is canceled.
- When context 1 is canceled, make context 2 also canceled.
