# GOLANG REDIS STREAMS

## Tech Stack

- golang
- redis
- redis streams
- server send events

### Learning Redis Streams

#### Basic Commands

- **XADD**:
  - adds data to a redis stream.
  - cmd example: `XADD <stream_name> * <data>`
  - here `*` means the redis server will generate a unique id for us and return the id.

### Learning Server Send Events

- sse package: `go get github.com/alexandrevicenzi/go-sse`

### Things worth going to Rabbit Hole

- [https://en.wikipedia.org/wiki/Radix_tree?useskin=vector](https://en.wikipedia.org/wiki/Radix_tree?useskin=vector)
- learn more about this [sse package](go get github.com/alexandrevicenzi/go-sse)
-
