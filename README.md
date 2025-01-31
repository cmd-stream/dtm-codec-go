# dtm-codec-go
dtm-codec-go contains predefined client and server codecs for cmd-stream-go.

# How To Use
```go
clientCodec, err := codec.CreateClientCodec(
  []codec.Unmarshaller[base.Result]{
    codec.NewResultDTSAdapter(Result1DTS),
    codec.NewResultDTSAdapter(Result2DTS),
    ...
  }
)
...
serverCodec, err := codec.CreateServerCodec(
  []codec.Unmarshaller[base.Result]{
    codec.NewCmdDTSAdapter(Cmd1DTS),
    codec.NewCmdDTSAdapter(Cmd2DTS),
    ...
  }
)
...
```