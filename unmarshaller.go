package dcodec

import (
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
)

// Unmarshaller defines an Unmarshal method that should deserialize only data,
// without DTM.
type Unmarshaller[T any] interface {
	DTM() com.DTM
	Unmarshal(r muss.Reader) (t T, n int, err error)
}
