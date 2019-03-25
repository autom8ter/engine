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

#### func  FromContext

```go
func FromContext(ctx context.Context, obj interface{}) string
```

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

#### func  ReadAsCSV

```go
func ReadAsCSV(val string) ([]string, error)
```

#### func  Render

```go
func Render(s string, data interface{}) string
```

#### func  RenderHTML

```go
func RenderHTML(s string, data interface{}) string
```

#### func  ScanAndReplace

```go
func ScanAndReplace(r io.Reader, replacements ...string) string
```

#### func  ToPrettyJson

```go
func ToPrettyJson(obj interface{}) []byte
```
ToPrettyJson encodes an item into a pretty (indented) JSON

#### func  ToPrettyJsonString

```go
func ToPrettyJsonString(obj interface{}) string
```
ToPrettyJsonString encodes an item into a pretty (indented) JSON string
