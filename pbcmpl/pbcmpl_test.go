package pbcmpl

import (
	"testing"

	"github.com/openacid/errors"

	"github.com/openacid/low/iohelper"
	"github.com/stretchr/testify/require"
)

func TestMarshalUnMarshal(t *testing.T) {

	ta := require.New(t)

	// use a header as a proto.Message
	msg := newHeader("123", 0xff)

	ta.Equal(32, HeaderSize(msg))
	ta.Equal(64, Size(msg))

	rw := newBuf(512)

	n, err := Marshal(iohelper.AtToWriter(rw, 0), msg)
	ta.Nil(err)
	ta.Equal(int64(64), n)

	// unmarshal

	m2 := &header{}

	n, ver, err := Unmarshal(iohelper.AtToReader(rw, 0), m2)
	ta.Nil(err)
	ta.Equal(int64(64), n)
	ta.Equal("1.0.0", ver, "default ver")
	ta.Equal(uint64(0xff), m2.BodySize)
}

type vh struct {
	header
}

func (h *vh) GetVersion() string {
	return "1.2.3"
}

func TestMarshalUnMarshal_withVersion(t *testing.T) {

	ta := require.New(t)

	msg := &vh{
		header: header{
			Version:    [16]byte{1, 2, 3, 4},
			HeaderSize: 456,
			BodySize:   789,
		},
	}

	rw := newBuf(512)

	n, err := Marshal(iohelper.AtToWriter(rw, 0), msg)
	ta.Nil(err)
	ta.Equal(int64(64), n)

	// unmarshal

	m2 := &vh{}

	n, ver, err := Unmarshal(iohelper.AtToReader(rw, 0), m2)
	ta.Nil(err)
	ta.Equal(int64(64), n)
	ta.Equal("1.2.3", ver)

	ta.Equal([16]byte{1, 2, 3, 4}, m2.Version)
	ta.Equal(uint64(456), m2.HeaderSize)
	ta.Equal(uint64(789), m2.BodySize)
}

type IncomleteReaderWriter struct {
	Buf []byte
}

func (rw *IncomleteReaderWriter) Read(p []byte) (n int, err error) {
	// Read 1 byte per Read()
	b := rw.Buf[0]
	rw.Buf = rw.Buf[1:]
	p[0] = b
	return 1, nil
}

func (rw *IncomleteReaderWriter) Write(p []byte) (n int, err error) {
	rw.Buf = append(rw.Buf, p...)
	return len(p), nil
}

type badMarshal struct {
	header
}

func (b *badMarshal) Marshal() ([]byte, error) {
	return nil, errors.New("badMarshal")
}

type badUnmarshal struct {
	header
}

func (b *badUnmarshal) Unmarshal(buf []byte) error {
	return errors.New("badUnmarshal")
}

func TestMarshal_Error(t *testing.T) {

	ta := require.New(t)

	msg := &badMarshal{
		header: header{
			Version:    [16]byte{1, 2, 3, 4},
			HeaderSize: 456,
			BodySize:   789,
		},
	}

	rw := newBuf(512)
	n, err := Marshal(iohelper.AtToWriter(rw, 0), msg)
	ta.Equal(int64(0), n)
	ta.Equal("badMarshal", err.Error())
}

func TestMarshal_WriteError(t *testing.T) {

	ta := require.New(t)

	msg := &vh{
		header: header{
			Version:    [16]byte{1, 2, 3, 4},
			HeaderSize: 456,
			BodySize:   789,
		},
	}

	// fail at first write

	rw := newBuf(31)
	n, err := Marshal(iohelper.AtToWriter(rw, 0), msg)
	ta.Equal(int64(31), n)
	ta.Equal("OutOfBoundary", err.Error())

	// fail at second write

	rw = newBuf(33)
	n, err = Marshal(iohelper.AtToWriter(rw, 0), msg)
	ta.Equal(int64(33), n)
	ta.Equal("OutOfBoundary", err.Error())
}

func TestUnmarshal_Error(t *testing.T) {

	ta := require.New(t)

	msg := &badUnmarshal{}

	rw := newBuf(64)
	_, err := Marshal(iohelper.AtToWriter(rw, 0), msg)
	ta.Nil(err)

	m2 := &badUnmarshal{}
	n, ver, err := Unmarshal(iohelper.AtToReader(rw, 0), m2)
	ta.Equal(int64(64), n)
	ta.Equal("1.0.0", ver)
	ta.Equal("badUnmarshal", err.Error())
}

func TestUnmarshal_ReadError(t *testing.T) {

	ta := require.New(t)

	msg := &vh{
		header: header{
			Version:    [16]byte{1, 2, 3, 4},
			HeaderSize: 456,
			BodySize:   789,
		},
	}

	rw := newBuf(64)
	_, err := Marshal(iohelper.AtToWriter(rw, 0), msg)
	ta.Nil(err)

	{
		// shrink buf to make a Read error at body
		tmp := rw.b
		rw.b = make([]byte, 33)
		copy(rw.b, tmp)

		m2 := &vh{}
		n, ver, err := Unmarshal(iohelper.AtToReader(rw, 0), m2)
		ta.Equal(int64(33), n)
		ta.Equal("1.2.3", ver)
		ta.Equal("EOF", err.Error())
	}

	{
		// shrink buf to make a Read error at body
		tmp := rw.b
		rw.b = make([]byte, 31)
		copy(rw.b, tmp)

		m2 := &vh{}
		n, ver, err := Unmarshal(iohelper.AtToReader(rw, 0), m2)
		ta.Equal(int64(31), n)
		ta.Equal("", ver)
		ta.Equal("EOF", err.Error())
	}

}

func TestUnmarshal_invalidHeaderSize(t *testing.T) {

	ta := require.New(t)

	msg := &vh{
		header: header{
			Version:    [16]byte{1, 2, 3, 4},
			HeaderSize: 456,
			BodySize:   789,
		},
	}

	rw := newBuf(64)
	_, err := Marshal(iohelper.AtToWriter(rw, 0), msg)
	ta.Nil(err)

	// headerSize starts from 16-th byte
	rw.b[16] = 0x21

	m2 := &vh{}
	n, ver, err := Unmarshal(iohelper.AtToReader(rw, 0), m2)
	ta.Equal(int64(32), n, "unmarshaled header but yet not body")
	ta.Equal("1.2.3", ver, "version is still correct")
	ta.Equal(ErrInvalidHeaderSize, errors.Cause(err))
}

func TestUnMarshal_IncompleteReader(t *testing.T) {

	ta := require.New(t)

	// Marshal() must work correctly
	// with an io.Reader:Read(p []byte)
	// returns n < len(p).

	msg := &vh{
		header: header{
			Version:    [16]byte{1, 2, 3, 4},
			HeaderSize: 456,
			BodySize:   789,
		},
	}

	rw := &IncomleteReaderWriter{}

	n, err := Marshal(rw, msg)
	ta.Nil(err)
	ta.Equal(int64(64), n)

	// unmarshal

	m2 := &vh{}

	n, ver, err := Unmarshal(rw, m2)
	ta.Nil(err)
	ta.Equal(int64(64), n)
	ta.Equal("1.2.3", ver)

	ta.Equal([16]byte{1, 2, 3, 4}, m2.Version)
	ta.Equal(uint64(456), m2.HeaderSize)
	ta.Equal(uint64(789), m2.BodySize)
}

func newBuf(n int) *atBuffer {
	return &atBuffer{
		b: make([]byte, n),
	}
}

type atBuffer struct {
	b []byte
}

func (t *atBuffer) WriteAt(b []byte, off int64) (n int, err error) {
	length := len(b)
	for i := 0; i < length; i++ {
		if i+int(off) >= len(t.b) {
			return i, errors.New("OutOfBoundary")
		}
		t.b[int64(i)+off] = b[i]
	}
	return length, nil
}

func (t *atBuffer) ReadAt(b []byte, off int64) (n int, err error) {
	length := len(b)

	min := off + int64(length)
	if min > int64(len(t.b)) {
		min = int64(len(t.b))
	}

	n = copy(b, t.b[off:min])
	if n < length {
		return n, errors.New("EOF")
	}
	return length, nil
}
