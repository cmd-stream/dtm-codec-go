package dcodec

import (
	"github.com/cmd-stream/transport-go"
	dts "github.com/mus-format/dts-stream-go"
)

// New creates a new Codec.
func New[T, V any](us []Unmarshaller[V]) (codec Codec[T, V], err error) {
	err = check(us)
	if err != nil {
		err = NewIncorrectUnmarshallersError(err)
		return
	}
	codec = Codec[T, V]{us}
	return
}

// Codec implements the cmd-stream general Codec interface.
type Codec[T, V any] struct {
	us []Unmarshaller[V]
}

func (c Codec[T, V]) Encode(t T, w transport.Writer) (err error) {
	m, ok := any(t).(Marshaller)
	if !ok {
		return NewNotMarshallerError(t)
	}
	return m.Marshal(w)
}

func (c Codec[T, V]) Decode(r transport.Reader) (v V, err error) {
	dtm, _, err := dts.DTMSer.Unmarshal(r)
	if err != nil {
		return
	}
	i := int(dtm)
	if i < 0 || i > len(c.us)-1 {
		err = NewUnexpectedDTMError(dtm)
		return
	}
	v, _, err = c.us[i].Unmarshal(r)
	return
}

func check[T any](us []Unmarshaller[T]) (err error) {
	if len(us) == 0 {
		return EmptySliceErr
	}
	for i := 0; i < len(us); i++ {
		if us[i] == nil {
			return NewNilItemError(i)
		}
		dtm := us[i].DTM()
		if int(dtm) != i {
			err = NewDTMNotEqualIndexError(dtm, i)
			return
		}
	}
	return
}
