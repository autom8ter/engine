# util
--
    import "github.com/autom8ter/engine/util"


## Usage

#### func  ChannelzClient

```go
func ChannelzClient(addr string) channelz.ChannelzClient
```
ChannelzClient creates a new grpc channelz client for connecting to a registered
channelz server for debugging.

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

#### func  LoadPlugins

```go
func LoadPlugins() []driver.Plugin
```
