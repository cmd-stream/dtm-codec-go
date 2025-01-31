package mock

import (
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

func NewUnmarshaller[T any]() Unmarshaller[T] {
	return Unmarshaller[T]{
		Mock: mok.New("Unmarshaller"),
	}
}

type Unmarshaller[T any] struct {
	*mok.Mock
}

func (u Unmarshaller[T]) RegisterN_DTM(n int, fn func() (dtm com.DTM)) Unmarshaller[T] {
	u.RegisterN("DTM", n, fn)
	return u
}

func (u Unmarshaller[T]) RegisterDTM(fn func() (dtm com.DTM)) Unmarshaller[T] {
	u.Register("DTM", fn)
	return u
}

func (u Unmarshaller[T]) RegisterUnmarshal(
	fn func(r muss.Reader) (t T, n int, err error)) Unmarshaller[T] {
	u.Register("Unmarshal", fn)
	return u
}

func (u Unmarshaller[T]) DTM() (dtm com.DTM) {
	vals, err := u.Call("DTM")
	if err != nil {
		panic(err)
	}
	return vals[0].(com.DTM)
}

func (u Unmarshaller[T]) Unmarshal(r muss.Reader) (t T, n int, err error) {
	vals, err := u.Call("Unmarshal", r)
	if err != nil {
		panic(err)
	}
	t = vals[0].(T)
	n = vals[1].(int)
	err, _ = vals[2].(error)
	return
}
