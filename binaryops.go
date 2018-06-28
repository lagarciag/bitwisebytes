package bitwisebytes


import (
	"fmt"
	"encoding/binary"
)


//ByteOrder specifies how to convert byte sequences into
// 16-, 32-, or 64-bit unsigned integers.
type ByteOrder interface {
	Uint8([]byte) uint8
	Uint16([]byte) uint16
	Uint24([]byte) uint32
	Uint32([]byte) uint32
	Uint40([]byte) uint64
	Uint48([]byte) uint64
	Uint56([]byte) uint64
	Uint64([]byte) uint64
	OrPutUint8([]byte, uint8)
	OrPutUint16([]byte, uint16)
	OrPutUint24([]byte, uint32)
	OrPutUint32([]byte, uint32)
	OrPutUint40([]byte, uint64)
	OrPutUint48([]byte, uint64)
	OrPutUint56([]byte, uint64)
	OrPutUint64([]byte, uint64)
}

// LittleEndian is the little-endian implementation of ByteOrder.
var LittleEndian littleEndian

// BigEndian is the big-endian implementation of ByteOrder.
var BigEndian bigEndian

type littleEndian struct{}

// -----------------------
// Little Endian
// -----------------------


func (littleEndian) Uint8(b []byte) uint8 {
	_ = b[0] // bounds check hint to compiler; see golang.org/issue/14808
	return uint8(b[0])
}

func (littleEndian) OrPutUint8(b []byte, v uint8) {
	_ = b[0] // early bounds check to guarantee safety of writes below
	b[0] |= byte(v)
}

func (littleEndian) Uint16(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

func (littleEndian) OrPutUint16(b []byte, v uint16) {
	_ = b[1] // early bounds check to guarantee safety of writes below
	b[0] |= byte(v)
	b[1] |= byte(v >> 8)
}

func (littleEndian) Uint24(b []byte) uint32 {
	_ = b[2] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16
}

func (littleEndian) OrPutUint24(b []byte, v uint32) {
	_ = b[2] // early bounds check to guarantee safety of writes below
	b[0] |= byte(v)
	b[1] |= byte(v >> 8)
	b[2] |= byte(v >> 16)
}

func (littleEndian) Uint32(b []byte) uint32 {
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func (littleEndian) OrPutUint32(b []byte, v uint32) {
	_ = b[3] // early bounds check to guarantee safety of writes below
	b[0] |= byte(v)
	b[1] |= byte(v >> 8)
	b[2] |= byte(v >> 16)
	b[3] |= byte(v >> 24)
}

func (littleEndian) Uint40(b []byte) uint64 {
	_ = b[4] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32
}

func (littleEndian) OrPutUint40(b []byte, v uint64) {
	_ = b[4] // early bounds check to guarantee safety of writes below
	b[0] |= byte(v)
	b[1] |= byte(v >> 8)
	b[2] |= byte(v >> 16)
	b[3] |= byte(v >> 24)
	b[4] |= byte(v >> 32)
}

func (littleEndian) Uint48(b []byte) uint64 {
	_ = b[5] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40
}

func (littleEndian) OrPutUint48(b []byte, v uint64) {
	_ = b[5] // early bounds check to guarantee safety of writes below
	b[0] |= byte(v)
	b[1] |= byte(v >> 8)
	b[2] |= byte(v >> 16)
	b[3] |= byte(v >> 24)
	b[4] |= byte(v >> 32)
	b[5] |= byte(v >> 40)
}

func (littleEndian) Uint56(b []byte) uint64 {
	_ = b[6] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48
}

func (littleEndian) OrPutUint56(b []byte, v uint64) {
	_ = b[6] // early bounds check to guarantee safety of writes below
	b[0] |= byte(v)
	b[1] |= byte(v >> 8)
	b[2] |= byte(v >> 16)
	b[3] |= byte(v >> 24)
	b[4] |= byte(v >> 32)
	b[5] |= byte(v >> 40)
	b[6] |= byte(v >> 48)
}

func (littleEndian) Uint64(b []byte) uint64 {
	_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

func (littleEndian) OrPutUint64(b []byte, v uint64) {
	_ = b[7] // early bounds check to guarantee safety of writes below
	b[0] |= byte(v)
	b[1] |= byte(v >> 8)
	b[2] |= byte(v >> 16)
	b[3] |= byte(v >> 24)
	b[4] |= byte(v >> 32)
	b[5] |= byte(v >> 40)
	b[6] |= byte(v >> 48)
	b[7] |= byte(v >> 56)
}

// ------------------------------------------------------------------
//                   Operations on/to shifted bytes
// ------------------------------------------------------------------

func (l littleEndian) Uint8ShiftedBytes(mask, offset int, b []byte) uint8 {
	if offset > 7 {
		panic("offset > 7")
	}
	return uint8(l.Uint16(b)>>uint(offset)) & uint8(mask)
}

func (l littleEndian) PutUint8ShiftedBytes(offset int, b []byte, v uint8) {
	if offset > 7 {
		panic("offset > 7")
	}
	switch len(b) {
	case 2:
		l.OrPutUint16(b, uint16(v)<<uint(offset))
	case 1:
		l.OrPutUint8(b, uint8(v)<<uint(offset))
	default:
		panic(fmt.Sprintf("incorrect size:%s", len(b)))
	}

}

func (l littleEndian) Uint16ShiftedBytes(mask, offset int, b []byte) uint16 {
	if offset > 7 {
		panic("offset > 7")
	}
	_ = b[2]
	return uint16(l.Uint24(b) >> uint(offset)) & uint16(mask)
}

func (l littleEndian) PutUint16ShiftedBytes(offset int, b []byte, v uint16) {
	if offset > 7 {
		panic("offset > 7")
	}
	switch len(b) {
	case 3:
		_ = b[2]
		l.OrPutUint24(b, uint32(v)<<uint(offset))
	case 2:
		_ = b[1]
		l.OrPutUint16(b, uint16(v)<<uint(offset))
	default:
		panic(fmt.Sprintf("incorrect size:%s", len(b)))
	}

}

func (l littleEndian) Uint24ShiftedBytes(mask, offset int, b []byte) uint32 {
	if offset > 7 {
		panic("offset > 7")
	}
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(l.Uint32(b) >> uint(offset)) & uint32(mask)
}

func (l littleEndian) PutUint24ShiftedBytes( offset int, b []byte, v uint32) {
	if offset > 7 {
		panic("offset > 7")
	}
	switch len(b) {
	case 4:
		_ = b[3]
		l.OrPutUint32(b, uint32(v)<<uint(offset))
	case 2:
		_ = b[2]
		l.OrPutUint24(b, uint32(v)<<uint(offset))
	default:
		panic(fmt.Sprintf("incorrect size:%s", len(b)))
	}


}

func (l littleEndian) Uint32ShiftedBytes(mask, offset int, b []byte) uint32 {
	if offset > 7 {
		panic("offset > 7")
	}
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(l.Uint40(b) >> uint(offset)) & uint32(mask)
}

func (l littleEndian) PutUint32ShiftedBytes(offset int, b []byte, v uint32) {
	if offset > 7 {
		panic("offset > 7")
	}
	switch len(b) {
	case 5:
		_ = b[4]
		l.OrPutUint40(b, uint64(v)<<uint(offset))
	case 4:
		_ = b[3]
		l.OrPutUint32(b, uint32(v)<<uint(offset))
	default:
		panic(fmt.Sprintf("incorrect size:%s", len(b)))
	}



}

func (l littleEndian) Uint40ShiftedBytes(mask, offset int, b []byte) uint64 {
	if offset > 7 {
		panic("offset > 7")
	}
	_ = b[4] // bounds check hint to compiler; see golang.org/issue/14808

	return uint64(l.Uint48(b) >> uint(offset)) & uint64(mask)
}

func (l littleEndian) PutUint40ShiftedBytes(offset int, b []byte, v uint64) {
	if offset > 7 {
		panic("offset > 7")
	}
	switch len(b) {
	case 6:
		_ = b[5]
		l.OrPutUint48(b, uint64(v)<<uint(offset))
	case 5:
		_ = b[4]
		l.OrPutUint40(b, uint64(v)<<uint(offset))
	default:
		panic(fmt.Sprintf("incorrect size:%s", len(b)))
	}

}

func (l littleEndian) Uint48ShiftedBytes(mask, offset int, b []byte) uint64 {
	if offset > 7 {
		panic("offset > 7")
	}
	_ = b[5] // bounds check hint to compiler; see golang.org/issue/14808

	return uint64(l.Uint56(b) >> uint(offset)) & uint64(mask)
}

func (l littleEndian) PutUint48ShiftedBytes(offset int, b []byte, v uint64) {
	if offset > 7 {
		panic("offset > 7")
	}
	switch len(b) {
	case 7:
		_ = b[6]
		l.OrPutUint56(b, uint64(v)<<uint(offset))
	case 6:
		_ = b[5]
		l.OrPutUint48(b, uint64(v)<<uint(offset))
	default:
		panic(fmt.Sprintf("incorrect size:%s", len(b)))
	}
}

func (l littleEndian) Uint56ShiftedBytes(mask, offset int, b []byte) uint64 {
	if offset > 7 {
		panic("offset > 7")
	}
	_ = b[6] // bounds check hint to compiler; see golang.org/issue/14808

	return uint64(l.Uint64(b) >> uint(offset)) & uint64(mask)
}

func (l littleEndian) PutUint56ShiftedBytes(offset int, b []byte, v uint64) {
	if offset > 7 {
		panic("offset > 7")
	}
	switch len(b) {
	case 8:
		_ = b[7]
		l.OrPutUint64(b, uint64(v)<<uint(offset))
	case 7:
		_ = b[6]
		l.OrPutUint56(b, uint64(v)<<uint(offset))
	default:
		panic(fmt.Sprintf("incorrect size:%s", len(b)))
	}

}

func (l littleEndian) Uint64ShiftedBytes(mask, offset int, b []byte) uint64 {
	if offset > 7 {
		panic("offset > 7")
	}
	_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808

	tmpB := make([]byte,len(b))
	copy(tmpB,b)
	tmpC , err := ShiftRight(tmpB,uint(offset))

	if err != nil {
		panic(err.Error())
	}

	return uint64(l.Uint64(tmpC) & uint64(mask))
}

func (l littleEndian) PutUint64ShiftedBytes(offset int, b []byte, v uint64) {
	if offset > 7 {
		panic("offset > 7")
	}
	switch len(b) {
	case 9: {
		_ = b[8]

		tmpV := make([]byte, 9)
		l.OrPutUint64(tmpV, uint64(v))
		tmpV2 , err  := ShiftLeft(tmpV, uint(offset))

		if err != nil {
			panic(err.Error())
		}

		for i, aByte := range tmpV2 {
			b[i] |= aByte
		}
	}

	case 8:
		_ = b[7]
		l.OrPutUint64(b, uint64(v)<<uint(offset))
	default:
		panic(fmt.Sprintf("incorrect size:%s", len(b)))
	}


}


func (l littleEndian) PutBytesSliceShiftedBytes(offset int, out,in []byte) {
	if offset > 7 {
		panic("offset > 7")
	}
	tmpOut , err := ShiftLeft(in,uint(offset))
	if err != nil {
		panic(err.Error())
	}

	for i , _ := range tmpOut {
		out[i] |= tmpOut[i]
	}
}


func (l littleEndian) BytesSliceShiftedBytes(mask []byte, offset int, b []byte) []byte {
	if offset > 7 {
		panic("offset > 7")
	}
	returnBytes , err := ShiftRight(b,uint(offset))
	if err != nil {
		panic(err.Error())
	}

	for i , aByteMask := range mask {
		returnBytes[i] =  returnBytes[i] & aByteMask
	}
	return returnBytes[0:len(mask)]
}



// --------------------------------
// BigEndian TODO: incomplete code
// --------------------------------
type bigEndian struct{}

func (bigEndian) Uint16(b []byte) uint16 {
	_ = b[1] // bounds check hint to compiler; see golang.org/issue/14808
	return uint16(b[1]) | uint16(b[0])<<8
}

func (bigEndian) PutUint16(b []byte, v uint16) {
	_ = b[1] // early bounds check to guarantee safety of writes below
	b[0] = byte(v >> 8)
	b[1] = byte(v)
}

func (bigEndian) Uint32(b []byte) uint32 {
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

func (bigEndian) PutUint32(b []byte, v uint32) {
	_ = b[3] // early bounds check to guarantee safety of writes below
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
}

func (bigEndian) Uint64(b []byte) uint64 {
	_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
}

func (bigEndian) PutUint64(b []byte, v uint64) {
	_ = b[7] // early bounds check to guarantee safety of writes below
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
}
