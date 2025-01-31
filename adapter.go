package dcodec

import (
	"github.com/cmd-stream/base-go"
	com "github.com/mus-format/common-go"
	dts "github.com/mus-format/mus-stream-dts-go"
	muss "github.com/mus-format/mus-stream-go"
)

// Creates a new CmdDTSAdapter.
func NewCmdDTSAdapter[T base.Cmd[V], V any](d dts.DTS[T]) CmdDTSAdapter[T, V] {
	return CmdDTSAdapter[T, V]{d}
}

// CmdDTSAdapter implements the Unmarshaller interface and serves as an adapter
// for Command DTS.
type CmdDTSAdapter[T base.Cmd[V], V any] struct {
	d dts.DTS[T]
}

func (u CmdDTSAdapter[T, V]) DTM() com.DTM {
	return u.d.DTM()
}

func (u CmdDTSAdapter[T, V]) Unmarshal(r muss.Reader) (v base.Cmd[V], n int,
	err error) {
	return u.d.UnmarshalData(r)
}

// NewResultDTSAdapter creates a new ResultDTSAdapter.
func NewResultDTSAdapter[T base.Result](d dts.DTS[T]) ResultDTSAdapter[T] {
	return ResultDTSAdapter[T]{d}
}

// ResultDTSAdapter implements the Unmarshaller interface and serves as an
// adapter for Result DTS.
type ResultDTSAdapter[T base.Result] struct {
	d dts.DTS[T]
}

func (u ResultDTSAdapter[T]) DTM() com.DTM {
	return u.d.DTM()
}

func (u ResultDTSAdapter[T]) Unmarshal(r muss.Reader) (v base.Result, n int,
	err error) {
	return u.d.UnmarshalData(r)
}
