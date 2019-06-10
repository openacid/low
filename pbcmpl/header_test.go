package pbcmpl

import (
	"testing"

	proto "github.com/golang/protobuf/proto"

	"github.com/stretchr/testify/require"
)

func TestVerStr(t *testing.T) {

	ta := require.New(t)

	var str string = verStr(nil)
	ta.Equal("", str)

	ta.Equal("", verStr([]byte{}))
	ta.Equal("abc", verStr([]byte{'a', 'b', 'c'}))
	ta.Equal("1.0.0", verStr([]byte{'1', '.', '0', '.', '0', 0}))

	bBuf := []byte{'1', '.', '0', '.', '0', 0}
	str = verStr(bBuf)
	bBuf[0] = '2'
	ta.Equal("1.0.0", str)
}

func TestNewHeader(t *testing.T) {

	ta := require.New(t)

	ver := "0.0.1"
	dataSize := uint64(1000)
	headerSize := uint64(32)
	h := newHeader(ver, dataSize)

	ta.Equal(dataSize, h.BodySize)
	ta.Equal(headerSize, h.HeaderSize)
	ta.Equal(ver, verStr(h.Version[:]))

	ta.Equal("0.0.1 32 1000", h.String())

	// 16 byte ver, no panic
	_ = newHeader("111111111111.2.6", 10)
	ta.Panics(func() { newHeader("111111111111.2.62", 10) })
}

func TestHeader_MarshalUnmarshal(t *testing.T) {

	ta := require.New(t)

	ver := "0.0.1"
	dataSize := uint64(0xffff01)
	h := newHeader(ver, dataSize)

	b, err := proto.Marshal(h)
	ta.Nil(err)

	want := append([]byte("0.0.1"), make([]byte, 11)...)
	want = append(want,
		32, 0, 0, 0,
		0, 0, 0, 0)
	want = append(want,
		1, 0xff, 0xff, 0,
		0, 0, 0, 0)
	ta.Equal(want, b)

	h2 := &header{}
	err = proto.Unmarshal(b, h2)
	ta.Nil(err)
	ta.Equal(h, h2)

	h.ProtoMessage()
}

func TestReadHeader(t *testing.T) {

	ta := require.New(t)

	msg := &vh{}
	rw := &IncomleteReaderWriter{}
	n, err := Marshal(rw, msg)
	ta.Nil(err)
	ta.Equal(int64(64), n)

	n, h, err := ReadHeader(rw)

	ta.Nil(err)
	ta.Equal(int64(32), n)
	ta.Equal("1.2.3", h.GetVersion())
	ta.Equal(int64(32), h.GetHeaderSize())
	ta.Equal(int64(32), h.GetBodySize())
}
