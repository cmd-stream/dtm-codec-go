package dcodec

import (
	"github.com/cmd-stream/transport-go"
)

// Marshaller defines a Marshal method that should serialize DTM and data.
type Marshaller interface {
	Marshal(w transport.Writer) error
}
