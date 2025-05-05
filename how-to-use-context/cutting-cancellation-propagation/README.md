# Cases where you want to cut cancellation or timeout propagation

- Cases where you want to perform rollback or cleanup processing that should not be cancelled to maintain Atomicity in a function under the influence of a context
- Cases in which background processing triggered by an event is to be continued after the trigger process ends
