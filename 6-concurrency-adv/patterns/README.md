# Concurrency Patterns

### Examples

- [Ping Pong](https://github.com/golang-basics/concurrency/blob/master/patterns/ping-pong/main.go)
- [Quit](https://github.com/golang-basics/concurrency/blob/master/patterns/quit/main.go)
- [Or Channel](https://github.com/golang-basics/concurrency/blob/master/patterns/or-channel/main.go)
- [OrDone Channel - Simple Done](https://github.com/golang-basics/concurrency/blob/master/patterns/or-done-channel/simple-done/main.go)
- [OrDone Channel - Or Done](https://github.com/golang-basics/concurrency/blob/master/patterns/or-done-channel/or-done/main.go)
- [Bridge Channel](https://github.com/golang-basics/concurrency/blob/master/patterns/bridge-channel/main.go)
- [Tee Channel](https://github.com/golang-basics/concurrency/blob/master/patterns/tee-channel/main.go)
- [Timers/Tickers](https://github.com/golang-basics/concurrency/blob/master/patterns/timers-tickers/main.go)
- [Cancellation - Simple Functions](https://github.com/golang-basics/concurrency/blob/master/patterns/cancellation/simple-functions/main.go)
- [Cancellation - Method Chaining](https://github.com/golang-basics/concurrency/blob/master/patterns/cancellation/method-chaining/main.go)
- [Circuit Breaker - Simple Example](https://github.com/golang-basics/concurrency/blob/master/patterns/circuit-breaker/simple/main.go)
- [Circuit Breaker - HTTP Example](https://github.com/golang-basics/concurrency/blob/master/patterns/circuit-breaker/http/main.go)
- [Code Generation](https://github.com/golang-basics/concurrency/blob/master/patterns/codegen/main.go)
- [Context - Simple Example](https://github.com/golang-basics/concurrency/blob/master/patterns/context/simple/main.go)
- [Context - Custom Context](https://github.com/golang-basics/concurrency/blob/master/patterns/context/mycontext/mycontext.go)
- [Context - Keys Collision](https://github.com/golang-basics/concurrency/blob/master/patterns/context/context-keys/collision/main.go)
- [Context - Simple Map Keys](https://github.com/golang-basics/concurrency/blob/master/patterns/context/simple-map-keys/main.go)
- [Context - Private Keys](https://github.com/golang-basics/concurrency/blob/master/patterns/context/context-keys/private-keys/main.go)
- [Context - Nested](https://github.com/golang-basics/concurrency/blob/master/patterns/context/nested/main.go)
- [Context - Over HTTP](https://github.com/golang-basics/concurrency/blob/master/patterns/context/http/main.go)
- [Context vs Done - done channel](https://github.com/golang-basics/concurrency/blob/master/patterns/context/context-vs-done-channel/channel/main.go)
- [Context vs Done - context](https://github.com/golang-basics/concurrency/blob/master/patterns/context/context-vs-done-channel/ctx/ctx/main.go)
- [Daisy Chain - Simple Example](https://github.com/golang-basics/concurrency/blob/master/patterns/daisy-chain/simple/main.go)
- [Daisy Chain - Functions Chain](https://github.com/golang-basics/concurrency/blob/master/patterns/daisy-chain/functions-chain/main.go)
- [Error Basics - Basic Wrap/Unwrap](https://github.com/golang-basics/concurrency/blob/master/patterns/error-handling/error-basics/basic-wrap-unwrap/main.go)
- [Error Basics - Simple Custom Error](https://github.com/golang-basics/concurrency/blob/master/patterns/error-handling/error-basics/simple-custom-error/main.go)
- [Error Basics - Complex Custom Error](https://github.com/golang-basics/concurrency/blob/master/patterns/error-handling/error-basics/complex-custom-error/main.go)
- [Error Basics - Errors.Is/Errors.As](https://github.com/golang-basics/concurrency/blob/master/patterns/error-handling/error-basics/errors-is-as/main.go)
- [Error Basics - Errors Wrap/Unwrap](https://github.com/golang-basics/concurrency/blob/master/patterns/error-handling/error-basics/errors-wrap-unwrap/main.go)
- [Error Basics - Propagation](https://github.com/golang-basics/concurrency/blob/master/patterns/error-handling/error-basics/propagation/main.go)
- [Errors - Concurrent: Bad](https://github.com/golang-basics/concurrency/blob/master/patterns/error-handling/concurrent/bad/main.go)
- [Errors - Concurrent: Good](https://github.com/golang-basics/concurrency/blob/master/patterns/error-handling/concurrent/good/main.go)
- [Errors - Concurrent: Better](https://github.com/golang-basics/concurrency/blob/master/patterns/error-handling/concurrent/better/main.go)
- [Errors - HTTP REST API](https://github.com/golang-basics/concurrency/blob/master/patterns/error-handling/http/main.go)
- [Fan In](https://github.com/golang-basics/concurrency/blob/master/patterns/fanin/fanin.go)
- [Fan Out](https://github.com/golang-basics/concurrency/blob/master/patterns/fanout/fanout.go)
- [Fan In-Out - Simple](https://github.com/golang-basics/concurrency/blob/master/patterns/fan-in-out/main.go)
- [Fan In-Out - Efficient](https://github.com/golang-basics/concurrency/blob/master/patterns/fan-in-out/efficient/main.go)
- [Fan In-Out - Inefficient](https://github.com/golang-basics/concurrency/blob/master/patterns/fan-in-out/inefficient/main.go)
- [Generator](https://github.com/golang-basics/concurrency/blob/master/patterns/generator/generator.go)
- [Useful Generators - Generic](https://github.com/golang-basics/concurrency/blob/master/patterns/generators/generators.go)
- [Useful Generators - Typed](https://github.com/golang-basics/concurrency/blob/master/patterns/generators/typed_generators.go)
- [Useful Generators - Benchmarks](https://github.com/golang-basics/concurrency/blob/master/patterns/generators/benchmarks/typed_vs_generic_test.go)
- [Heartbeats - Simple Example](https://github.com/golang-basics/concurrency/blob/master/patterns/heartbeats/simple/main.go)
- [Heartbeats - Broken Example](https://github.com/golang-basics/concurrency/blob/master/patterns/heartbeats/broken/main.go)
- [Heartbeats - 1 to 1](https://github.com/golang-basics/concurrency/blob/master/patterns/heartbeats/1-to-1-heartbeats/main.go)
- [Heartbeats - Interval Based](https://github.com/golang-basics/concurrency/blob/master/patterns/heartbeats/interval-based-heartbeats/main.go)
- [Healing Unhealthy Go Routines](https://github.com/golang-basics/concurrency/blob/master/patterns/healing-unhealthy-goroutines/main.go)
- [Pipeline - Simple](https://github.com/golang-basics/concurrency/blob/master/patterns/pipeline/simple/main.go)
- [Pipeline - Channels](https://github.com/golang-basics/concurrency/blob/master/patterns/pipeline/channels/main.go)
- [Pipeline - Functions/Methods](https://github.com/golang-basics/concurrency/blob/master/patterns/pipeline/cmd/main.go)
- [Pipeline - Stream vs Batch Example](https://github.com/golang-basics/concurrency/blob/master/patterns/pipeline/stream-vs-batch/main.go)
- [Pipeline - Stream vs Batch Benchmarks](https://github.com/golang-basics/concurrency/blob/master/patterns/pipeline/stream-vs-batch/stream_vs_batch_test.go)
- [Pipeline - Digest Tree Example](https://github.com/golang-basics/concurrency/blob/master/patterns/pipeline/digest-tree/main.go)
- [Pipeline - Digest Tree Benchmarks](https://github.com/golang-basics/concurrency/blob/master/patterns/pipeline/digest-tree/benchmarks_test.go)
- [Rate Limiting - Simple Rate Limiting](https://github.com/golang-basics/concurrency/blob/master/patterns/rate-limiting/simple-rate-limiting/main.go)
- [Rate Limiting - Token Bucket](https://github.com/golang-basics/concurrency/blob/master/patterns/rate-limiting/token-bucket/main.go)
- [Rate Limiting - Single Rate Limiting](https://github.com/golang-basics/concurrency/blob/master/patterns/rate-limiting/single-rate-limiting/main.go)
- [Rate Limiting - Multi Rate Limiting](https://github.com/golang-basics/concurrency/blob/master/patterns/rate-limiting/multi-rate-limiting/main.go)
- [Rate Limiting - Finer Rate Limiting](https://github.com/golang-basics/concurrency/blob/master/patterns/rate-limiting/finer-rate-limiting/main.go)
- [Replicated Requests](https://github.com/golang-basics/concurrency/blob/master/patterns/replicated-requests/main.go)
- [Search](https://github.com/golang-basics/concurrency/blob/master/patterns/search/main.go)
- [Server](https://github.com/golang-basics/concurrency/blob/master/patterns/server/main.go)
- [Worker](https://github.com/golang-basics/concurrency/blob/master/patterns/worker/main.go)
- [Server & Worker](https://github.com/golang-basics/concurrency/blob/master/patterns/server-and-worker/main.go)

[Home](https://github.com/golang-basics/concurrency)
