
- Using the -race flag to detect race condition issues

`go run -race main.go`


## Unbuffered Channels

An unbuffered channel does not have any capacity to store messages. The sender and receiver must synchronize directly

When to use:

- When goroutines need to synchronize directly.
- When you need tight coupling between the producer and consumer.
- Example: Signaling between goroutines.

## Buffered Channels

A buffered channel has a capacity to store messages. This means:

When to use:

- When goroutines can operate asynchronously, and some delay in processing is acceptable.
- When you want to reduce blocking for the sender as long as the buffer isn’t full.
- Example: Log processing, task queues, or decoupling producer-consumer systems.