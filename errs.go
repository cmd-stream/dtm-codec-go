package dcodec

import (
	"errors"
	"fmt"
	"reflect"

	com "github.com/mus-format/common-go"
)

var (
	UnexpectedErr = errors.New("unexpected DTM")
	EmptySliceErr = errors.New("empty slice")
)

func NewIncorrectUnmarshallersError(cause error) error {
	return fmt.Errorf("incorrect unmarshallers, cause: %w", cause)
}

func NewDTMNotEqualIndexError(dtm com.DTM, index int) error {
	return fmt.Errorf("item's DTM=%v not equal to index=%v", dtm, index)
}

func NewNotMarshallerError[T any](t T) error {
	return fmt.Errorf("%v doesn't implement the Marshaller interface",
		reflect.TypeOf(t))
}

func NewNilItemError(index int) error {
	return fmt.Errorf("nil item at %v index", index)
}
