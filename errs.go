package dcodec

import (
	"errors"
	"fmt"
	"reflect"

	com "github.com/mus-format/common-go"
)

// EmptySliceErr occurs when the Codec is initialized with an empty slice.
var EmptySliceErr = errors.New("empty slice")

// NewUnexpectedDTMError creates an error that occurs when Codec.Decode
// encounters an unexpected DTM.
func NewUnexpectedDTMError(dtm com.DTM) error {
	return fmt.Errorf("unexpected DTM %v", dtm)
}

// NewIncorrectUnmarshallersError creates an error that occurs when the codec is
// initialized with an incorrect slice.
func NewIncorrectUnmarshallersError(cause error) error {
	return fmt.Errorf("incorrect unmarshallers, cause: %w", cause)
}

// NewDTMNotEqualIndexError creates an error is the cause of
// NewIncorrectUnmarshallersError.
func NewDTMNotEqualIndexError(dtm com.DTM, index int) error {
	return fmt.Errorf("item's DTM=%v not equal to index=%v", dtm, index)
}

// NewNilItemError creates an error is the cause of NewIncorrectUnmarshallersError.
func NewNilItemError(index int) error {
	return fmt.Errorf("nil item at %v index", index)
}

// NewNotMarshallerError creates an error that occurs when Codec.Encode
// parameter doesn't implement the Marshaller interface.
func NewNotMarshallerError[T any](t T) error {
	return fmt.Errorf("%v doesn't implement the Marshaller interface",
		reflect.TypeOf(t))
}
