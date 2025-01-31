package dcodec

import (
	"dtm-codec-go/testdata/mock"
	"errors"
	"testing"

	"github.com/cmd-stream/transport-go"
	transport_mock "github.com/cmd-stream/transport-go/testdata/mock"
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	assert_error "github.com/ymz-ncnk/assert/error"
	"github.com/ymz-ncnk/mok"
)

func TestCodec(t *testing.T) {

	t.Run("NewCodec should fail if receives invalid slice", func(t *testing.T) {
		var (
			wantErr = NewIncorrectUnmarshallersError(NewDTMNotEqualIndexError(2, 1))
			u0      = mock.NewUnmarshaller[any]().RegisterDTM(func() (dtm com.DTM) {
				return 0
			})
			u1 = mock.NewUnmarshaller[any]().RegisterDTM(func() (dtm com.DTM) {
				return 2
			})
			mocks  = []*mok.Mock{u0.Mock, u1.Mock}
			_, err = NewCodec[any, any]([]Unmarshaller[any]{u0, u1})
		)
		assert_error.EqualError(err, wantErr, t)

		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}
	})

	t.Run("NewCodec should fail if receives empty slice", func(t *testing.T) {
		wantErr := NewIncorrectUnmarshallersError(NewNilItemError(0))
		_, err := NewCodec[any, any]([]Unmarshaller[any]{nil})
		assert_error.EqualError(err, wantErr, t)
	})

	t.Run("NewCodec should fail if receives a nil Unmarshaller", func(t *testing.T) {
		wantErr := NewIncorrectUnmarshallersError(EmptySliceErr)
		_, err := NewCodec[any, any](nil)
		assert_error.EqualError(err, wantErr, t)
	})

	t.Run("Codec.Encode should fail param doesn't implement Marshaller",
		func(t *testing.T) {
			var (
				param   = 1
				wantErr = NewNotMarshallerError(param)
				u0      = mock.NewUnmarshaller[any]().RegisterDTM(func() (dtm com.DTM) {
					return 0
				})
				mocks    = []*mok.Mock{u0.Mock}
				codec, _ = NewCodec[int, any]([]Unmarshaller[any]{u0})
			)
			err := codec.Encode(param, nil)
			assert_error.EqualError(err, wantErr, t)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

	t.Run("If Marshaller returns an error, Codec.Encode should return it",
		func(t *testing.T) {
			var (
				wantErr = errors.New("Marshal error")
				u0      = mock.NewUnmarshaller[any]().RegisterDTM(
					func() (dtm com.DTM) { return 0 })
				m = mock.NewMarshaller().RegisterMarshal(
					func(w transport.Writer) error { return wantErr })
				mocks    = []*mok.Mock{u0.Mock, m.Mock}
				codec, _ = NewCodec[any]([]Unmarshaller[any]{u0})
			)
			err := codec.Encode(m, nil)
			assert_error.EqualError(err, wantErr, t)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

	t.Run("If DTM unmarshalling fails with an error, Codec.Decode should return it",
		func(t *testing.T) {
			var (
				wantResult any
				wantErr    = errors.New("DTM unmarshal error")
				r          = transport_mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						err = wantErr
						return
					},
				)
				u0 = mock.NewUnmarshaller[any]().RegisterDTM(func() (dtm com.DTM) {
					return 0
				})
				mocks    = []*mok.Mock{r.Mock, u0.Mock}
				codec, _ = NewCodec[any]([]Unmarshaller[any]{u0})
			)
			result, err := codec.Decode(r)
			assert_error.Equal(result, wantResult, t)
			assert_error.EqualError(err, wantErr, t)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

	t.Run("Codec.Decode should return UnexpectedErr for too large DTM value",
		func(t *testing.T) {
			var (
				wantResult any
				wantErr    = UnexpectedErr
				r          = transport_mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						b = 10
						return
					},
				)
				u0 = mock.NewUnmarshaller[any]().RegisterDTM(func() (dtm com.DTM) {
					return 0
				})
				mocks    = []*mok.Mock{r.Mock, u0.Mock}
				codec, _ = NewCodec[any]([]Unmarshaller[any]{u0})
			)
			result, err := codec.Decode(r)
			assert_error.Equal(result, wantResult, t)
			assert_error.EqualError(err, wantErr, t)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

	t.Run("ClientCodec.Decode should return a result", func(t *testing.T) {
		var (
			wantResult = 2
			wantErr    error
			u0         = mock.NewUnmarshaller[int]().RegisterDTM(func() (dtm com.DTM) {
				return 0
			})
			u1 = mock.NewUnmarshaller[int]().RegisterDTM(func() (dtm com.DTM) {
				return 1
			}).RegisterUnmarshal(func(r muss.Reader) (result int, n int, err error) {
				result = wantResult
				return
			})
			u2 = mock.NewUnmarshaller[int]().RegisterDTM(func() (dtm com.DTM) {
				return 2
			})
			r = transport_mock.NewReader().RegisterReadByte(func() (b byte, err error) {
				b = 1
				return
			})
			mocks    = []*mok.Mock{u0.Mock, u1.Mock, u2.Mock}
			codec, _ = NewCodec[any]([]Unmarshaller[int]{u0, u1, u2})
		)
		result, err := codec.Decode(r)
		assert_error.Equal(result, wantResult, t)
		assert_error.EqualError(err, wantErr, t)

		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}
	})

}
