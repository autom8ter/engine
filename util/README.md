# util
--
    import "github.com/autom8ter/engine/util"


## Usage

#### func  Debugf

```go
func Debugf(format string, args ...interface{})
```
Debugf is grpclog.Infof(format, args...) but only executes if debug=true is set
in your config or environmental variables

#### func  Debugln

```go
func Debugln(args ...interface{})
```
Debugln is grpclog.Infoln(args...) but only executes if debug=true is set in
your config or environmental variables

#### func  NewMultiStreamServerInterceptor

```go
func NewMultiStreamServerInterceptor(sints ...grpc.StreamServerInterceptor) grpc.StreamServerInterceptor
```

#### func  NewMultiUnaryServerInterceptor

```go
func NewMultiUnaryServerInterceptor(uints ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor
```

#### func  NewServerStreamWithContext

```go
func NewServerStreamWithContext(stream grpc.ServerStream, ctx context.Context) grpc.ServerStream
```
