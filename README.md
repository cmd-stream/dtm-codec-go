# dtm-codec-go

[![Go Reference](https://pkg.go.dev/badge/github.com/cmd-stream/dtm-codec-go.svg)](https://pkg.go.dev/github.com/cmd-stream/dtm-codec-go)
[![GoReportCard](https://goreportcard.com/badge/cmd-stream/dtm-codec-go)](https://goreportcard.com/report/github.com/cmd-stream/dtm-codec-go)
[![codecov](https://codecov.io/gh/cmd-stream/dtm-codec-go/graph/badge.svg?token=6JVVHR8QHF)](https://codecov.io/gh/cmd-stream/dtm-codec-go)

dtm-codec-go provides client and server codecs for cmd-stream-go.

# How To
1. Define DTMs:
```go
import (
	com "github.com/mus-format/common-go"
)

const(
  Cmd1DTM com.DTM = iota
  Cmd2DTM
  ...
)

const(
  Result1DTM com.DTM = iota
  Result2DTM
  ...
)
```

2. Create DTM support variables for Commands and Results using [dts-stream-go](https://github.com/mus-format/dts-stream-go):
```go
import (
  dts "github.com/mus-format/dts-stream-go"
)

var (
  Cmd1DTS = dts.New[Cmd1DTS](Cmd1DTM, ...)
  Cmd2DTS = dts.New[Cmd2DTS](Cmd1DTM, ...)
)

var (
  Result1DTS = dts.New[Result1DTS](Result1DTM, ...)
  Result2DTS = dts.New[Result2DTS](Result1DTM, ...)
)
```

3. Create codecs:
```go
import (
  dcodec "github.com/cmd-stream/dtm-codec-go"
  "github.com/cmd-stream/base-go"
)

...
clientCodec, err := dcodec.NewClientCodec[Receiver](
  []dcodec.Unmarshaller[base.Result]{ // Elements must be arranged in ascending 
  // order based on their DTM values.
    dcodec.NewResultDTSAdapter(Result1DTS), // DTM == 0
    dcodec.NewResultDTSAdapter(Result2DTS), // DTM == 1
    ...
  },
)
...

serverCodec, err := dcodec.NewServerCodec(
  []dcodec.Unmarshaller[base.Cmd[Receiver]]{ // Elements must be arranged in ascending 
  // order based on their DTM values.
    dcodec.NewCmdDTSAdapter(Cmd1DTS), // DTM == 0
    dcodec.NewCmdDTSAdapter(Cmd2DTS), // DTM == 1
    ...
  },
)
...
```