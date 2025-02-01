package dcodec

import (
	"testing"

	"github.com/cmd-stream/dtm-codec-go/testdata/mock"

	"github.com/cmd-stream/base-go"
	com "github.com/mus-format/common-go"
	assert_error "github.com/ymz-ncnk/assert/error"
)

func TestDTMCodec(t *testing.T) {

	t.Run("If client codec creation fails with an error, Create should return it",
		func(t *testing.T) {
			var (
				wantErr = NewIncorrectUnmarshallersError(NewDTMNotEqualIndexError(1, 0))
				us      = []Unmarshaller[base.Result]{
					mock.NewUnmarshaller[base.Result]().RegisterDTM(func() (dtm com.DTM) {
						return 1
					}),
				}
			)
			_, err := NewClientCodec[any](us)
			assert_error.EqualError(err, wantErr, t)
		})

	t.Run("If server codec creation fails with an error, Create should return it",
		func(t *testing.T) {
			var (
				wantErr = NewIncorrectUnmarshallersError(NewDTMNotEqualIndexError(1, 0))
				us      = []Unmarshaller[base.Cmd[any]]{
					mock.NewUnmarshaller[base.Cmd[any]]().RegisterDTM(func() (dtm com.DTM) {
						return 1
					}),
				}
			)
			_, err := NewServerCodec[any](us)
			assert_error.EqualError(err, wantErr, t)
		})

	t.Run("NewClientCodec should return codec", func(t *testing.T) {
		var (
			us = []Unmarshaller[base.Result]{
				mock.NewUnmarshaller[base.Result]().RegisterN_DTM(2, func() (dtm com.DTM) {
					return 0
				}),
			}
			wantCodec, _ = New[base.Cmd[int]](us)
			wantErr      error
		)
		codec, err := NewClientCodec[any](us)
		assert_error.EqualError(err, wantErr, t)

		for i := 0; i < len(codec.us); i++ {
			if codec.us[i] != wantCodec.us[i] {
				t.Fatal("not equal")
			}
		}
	})

	t.Run("NewServerCodec should return codec", func(t *testing.T) {
		var (
			us = []Unmarshaller[base.Cmd[any]]{
				mock.NewUnmarshaller[base.Cmd[any]]().RegisterN_DTM(2, func() (dtm com.DTM) {
					return 0
				}),
			}
			wantCodec, _ = New[base.Result](us)
			wantErr      error
		)
		codec, err := NewServerCodec[any](us)
		assert_error.EqualError(err, wantErr, t)

		for i := 0; i < len(codec.us); i++ {
			if codec.us[i] != wantCodec.us[i] {
				t.Fatal("not equal")
			}
		}
	})

}
