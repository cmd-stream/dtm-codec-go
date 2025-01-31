package dcodec

import "github.com/cmd-stream/base-go"

// CreateClientCodec initializes a cmd-stream client codec.
//
// It accepts a slice of Unmarshallers, where each Unmarshaller's DTM value
// must match its position in the slice. Example of a correct slice:
//
//		[]Unmarshaller[base.Result]{
//		  u0, // u0.DTM() = 0, index = 0
//		  u1, // u1.DTM() = 1, index = 1
//		  u2, // u2.DTM() = 2, index = 2
//	   ...
//		}
//
// Incorrect slices:
//
//	[]Unmarshaller{} // Empty slice.
//	[]Unmarshaller{  // Contains an item where DTM != index.
//	  u0, // u0.DTM() = 1, index = 0
//	}
//	[]Unmarshaller{  // Contains a nil item.
//	  nil,
//	}
//
// Returns an error if the provided slice is invalid.
func CreateClientCodec[T any](us []Unmarshaller[base.Result]) (
	Codec[base.Cmd[T], base.Result], error) {
	return NewCodec[base.Cmd[T], base.Result](us)
}

// CreateServerCodec initializes a cmd-stream server codec.
//
// It accepts a slice of Unmarshallers, where each Unmarshaller's DTM value
// must match its position in the slice. Example of a correct slice:
//
//		[]Unmarshaller[base.Result]{
//		  u0, // u0.DTM() = 0, index = 0
//		  u1, // u1.DTM() = 1, index = 1
//		  u2, // u2.DTM() = 2, index = 2
//	    ...
//		}
//
// Incorrect slices:
//
//	[]Unmarshaller{} // Empty slice.
//	[]Unmarshaller{  // Contains an item where DTM != index.
//	  u0, // u0.DTM() = 1, index = 0
//	}
//	[]Unmarshaller{  // Contains a nil item.
//	  nil,
//	}
//
// Returns an error if the provided slice is invalid.
func CreateServerCodec[T any](us []Unmarshaller[base.Cmd[T]]) (
	Codec[base.Result, base.Cmd[T]], error) {
	return NewCodec[base.Result, base.Cmd[T]](us)
}
