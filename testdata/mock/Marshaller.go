package mock

import (
	"github.com/cmd-stream/transport-go"
	"github.com/ymz-ncnk/mok"
)

func NewMarshaller() MarshallerCmd {
	return MarshallerCmd{
		Mock: mok.New("CmdDTS"),
	}
}

type MarshallerCmd struct {
	*mok.Mock
}

func (c MarshallerCmd) RegisterMarshal(
	fn func(w transport.Writer) error) MarshallerCmd {
	c.Register("Marshal", fn)
	return c
}

func (c MarshallerCmd) Marshal(w transport.Writer) (err error) {
	vals, err := c.Call("Marshal", mok.SafeVal[transport.Writer](w))
	if err != nil {
		panic(err)
	}
	err, _ = vals[0].(error)
	return
}
