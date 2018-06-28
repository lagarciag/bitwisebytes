package bitwisebytes_test

import (
	"testing"
	"math/rand"
	"github.com/lagarciag/bitwisebytes"
	"encoding/binary"
)

const testLooops = 100
const debug = false

func TestUint16(t *testing.T) {
	for i:=0;i<testLooops;i++ {
		bytesSlice := make([]byte,2)
		randUint := uint16(rand.Intn(0xFFFFFFFF))
		bitwisebytes.LittleEndian.OrPutUint16(bytesSlice,randUint)
		resultUint := bitwisebytes.LittleEndian.Uint16(bytesSlice)
		if resultUint != randUint {
			t.Errorf("mistmatch: %d != %d", randUint, resultUint)
		}
	}
}

func TestUint24(t *testing.T) {
	for i:=0;i<testLooops;i++ {
		bytesSlice := make([]byte,3)
		randUint := uint32(rand.Intn(0xFFFFFF))
		bitwisebytes.LittleEndian.OrPutUint24(bytesSlice,randUint)
		resultUint := bitwisebytes.LittleEndian.Uint24(bytesSlice)
		if resultUint != randUint {
			t.Log("randUint: ",randUint)
			t.Error(bytesSlice)
			t.Errorf("mistmatch: %d != %d", randUint, resultUint)
		}



	}
}

func TestUint32(t *testing.T) {
	for i:=0;i<testLooops;i++ {
		bytesSlice := make([]byte,4)
		randUint := uint32(rand.Intn(0xFFFFFFFF))
		bitwisebytes.LittleEndian.OrPutUint32(bytesSlice,randUint)
		resultUint := bitwisebytes.LittleEndian.Uint32(bytesSlice)
		if resultUint != randUint {
			t.Log("randUint: ",randUint)
			t.Error(bytesSlice)
			t.Errorf("mistmatch: %d != %d", randUint, resultUint)
		}
	}
}

func TestUint40(t *testing.T) {
	const size = 5
	const bitsSize = size * 8
	const max = 0xFFFFFFFFFF
	for i:=0;i<testLooops;i++ {
		bytesSlice := make([]byte,size)
		randUint := uint64(rand.Intn(max))
		bitwisebytes.LittleEndian.OrPutUint40(bytesSlice,randUint)
		resultUint := bitwisebytes.LittleEndian.Uint40(bytesSlice)
		if resultUint != randUint {
			t.Log("randUint: ",randUint)
			t.Error(bytesSlice)
			t.Errorf("mistmatch: %d != %d", randUint, resultUint)
		}
	}
}

func TestUint48(t *testing.T) {
	const size = 6
	const bitsSize = size * 8
	const max = 0xFFFFFFFFFFFF
	for i:=0;i<testLooops;i++ {
		bytesSlice := make([]byte,size)
		randUint := uint64(rand.Intn(max))
		bitwisebytes.LittleEndian.OrPutUint48(bytesSlice,randUint)
		resultUint := bitwisebytes.LittleEndian.Uint48(bytesSlice)
		if resultUint != randUint {
			t.Log("randUint: ",randUint)
			t.Error(bytesSlice)
			t.Errorf("mistmatch: %d != %d", randUint, resultUint)
		}
	}
}

func TestUint56(t *testing.T) {
	const size = 7
	const bitsSize = size * 8
	const max = 0xFFFFFFFFFFFFFF
	for i:=0;i<testLooops;i++ {
		bytesSlice := make([]byte,size)
		randUint := uint64(rand.Intn(max))
		bitwisebytes.LittleEndian.OrPutUint56(bytesSlice,randUint)
		resultUint := bitwisebytes.LittleEndian.Uint56(bytesSlice)
		if resultUint != randUint {
			t.Log("randUint: ",randUint)
			t.Error(bytesSlice)
			t.Errorf("mistmatch: %d != %d", randUint, resultUint)
		}
	}
}

func TestUint64(t *testing.T) {
	const size = 8
	const bitsSize = size * 8
	const max = 0xFFFFFFFFFFFFFFF
	for i:=0;i<testLooops;i++ {
		bytesSlice := make([]byte,size)
		randUint := uint64(rand.Intn(max))
		bitwisebytes.LittleEndian.OrPutUint64(bytesSlice,randUint)
		resultUint := bitwisebytes.LittleEndian.Uint64(bytesSlice)
		if resultUint != randUint {
			t.Log("randUint: ",randUint)
			t.Error(bytesSlice)
			t.Errorf("mistmatch: %d != %d", randUint, resultUint)
		}
	}
}

func TestUint8Shifted(t *testing.T) {
	const intBytesSize = 1
    const bitsSize = intBytesSize * 8
    const max = 0xFF

	for i:=0;i<testLooops;i++ {
		destSlice := make([]byte, intBytesSize + 1)

		randUint := uint8(rand.Intn(max))
		randOffset := rand.Intn(bitsSize)
		shiftedRandUint := randUint << uint(randOffset)

		// test PutUintShifted
		bitwisebytes.LittleEndian.PutUint8ShiftedBytes(randOffset,destSlice, randUint)
		resultUint := bitwisebytes.LittleEndian.Uint8ShiftedBytes(max,randOffset,destSlice)
		if resultUint != randUint {
			t.Log("randUint: ",randUint)
			t.Log("randOffset: ",randOffset)
			t.Log(destSlice)
			t.Errorf("mistmatch: %d != %d", randUint, resultUint)
		}

		// test ShiftRight of bytes
		unShiftedSlice, err:= bitwisebytes.ShiftRight(destSlice,uint(randOffset))
		if err != nil {
			t.Error(err.Error())
		}
		newInt := bitwisebytes.LittleEndian.Uint8(unShiftedSlice[0:1])

		if newInt != randUint {
			t.Error(newInt,randUint)
		}
		reShiftedSlice, err:= bitwisebytes.ShiftLeft(unShiftedSlice,uint(randOffset))
		if err != nil {
			t.Error(err.Error())
		}
		shiftedInt := bitwisebytes.LittleEndian.Uint8(reShiftedSlice)
		//t.Log(shiftedInt, shiftedRandUint)
		if shiftedInt != shiftedRandUint {
			t.Error(shiftedInt, shiftedRandUint)
		}
		newInt = bitwisebytes.LittleEndian.Uint8ShiftedBytes(max,randOffset,reShiftedSlice)
		if newInt != randUint {
			t.Error(newInt,randUint)
		}


	}
}

func TestUint16Shifted(t *testing.T) {
	const intBytesSize = 2
	const bitsSize = intBytesSize * 8
	const max = 0xFFFF

	for i:=0;i<testLooops;i++ {
		destSlice := make([]byte, intBytesSize + 1)

		randUint := uint16(rand.Intn(max))
		randOffset := rand.Intn(8)
		shiftedRandUint := randUint << uint(randOffset)

		// test PutUintShifted
		bitwisebytes.LittleEndian.PutUint16ShiftedBytes(randOffset,destSlice,randUint)
		resultUint := bitwisebytes.LittleEndian.Uint16ShiftedBytes(max,randOffset,destSlice)
		if resultUint != randUint {
			t.Log("randUint: ",randUint)
			t.Log("randOffset: ",randOffset)
			t.Log(destSlice)
			t.Errorf("mistmatch: %d != %d", randUint, resultUint)
		}

		// test ShiftRight of bytes
		unShiftedSlice, err:= bitwisebytes.ShiftRight(destSlice,uint(randOffset))
		if err != nil {
			t.Error(err.Error())
		}

		newInt := bitwisebytes.LittleEndian.Uint16(unShiftedSlice[0:intBytesSize])

		if newInt != randUint {
			t.Error(newInt,randUint)
		}
		reShiftedSlice, err:= bitwisebytes.ShiftLeft(unShiftedSlice,uint(randOffset))
		if err != nil {
			t.Error(err.Error())
		}
		shiftedInt := bitwisebytes.LittleEndian.Uint16(reShiftedSlice)
		if shiftedInt != shiftedRandUint {
			t.Error(shiftedInt, shiftedRandUint)
		}
		newInt = bitwisebytes.LittleEndian.Uint16ShiftedBytes(max,randOffset,reShiftedSlice)
		if newInt != randUint {
			t.Error(newInt,randUint)
		}
	}
}

func TestUint24Shifted(t *testing.T) {
	const intBytesSize = 3
	const bitsSize = intBytesSize * 8
	const max = 0xFFFFFF

	for i:=0;i<testLooops;i++ {
		destSlice := make([]byte, intBytesSize + 1)

		randUint := uint32(rand.Intn(max))
		randOffset := rand.Intn(8)
		shiftedRandUint := randUint << uint(randOffset)

		// test PutUintShifted
		bitwisebytes.LittleEndian.PutUint24ShiftedBytes(randOffset,destSlice,randUint)
		resultUint := bitwisebytes.LittleEndian.Uint24ShiftedBytes(max,randOffset,destSlice)

		if resultUint != randUint {
			t.Logf("randUint: 0x%x",randUint)
			t.Logf("randOffset: 0x%x ",randOffset)
			t.Log(destSlice)
			t.Errorf("mistmatch: %d != %d", randUint, resultUint)
		}else if debug {
			t.Logf("resultUint : 0x%x", resultUint)
			t.Logf("destSlice : 0x%x", destSlice)
		}

		// test ShiftRight of bytes
		unShiftedSlice, err:= bitwisebytes.ShiftRight(destSlice,uint(randOffset))
		if err != nil {
			t.Error(err.Error())
		}

		if debug {
			t.Log("unsfhited slice: ",unShiftedSlice )
		}

		newInt := bitwisebytes.LittleEndian.Uint24(unShiftedSlice[0:intBytesSize])

		if newInt != randUint {
			t.Error(newInt,randUint)
		}
		reShiftedSlice, err:= bitwisebytes.ShiftLeft(unShiftedSlice,uint(randOffset))
		if err != nil {
			t.Error(err.Error())
		}

		for i , byte := range  reShiftedSlice {
			if byte != destSlice[i] {
				t.Errorf("byte 0x%x is different",i)
				t.Log(byte,destSlice[i] )
			}
		}

		if debug {
			t.Log("reshifted slice: ", reShiftedSlice)
		}

		shiftedInt := bitwisebytes.LittleEndian.Uint32(reShiftedSlice)
		if shiftedInt != shiftedRandUint {
			t.Error(shiftedInt, shiftedRandUint)
		}
		newInt = bitwisebytes.LittleEndian.Uint24ShiftedBytes(max,randOffset,reShiftedSlice)
		if newInt != randUint {
			t.Error(newInt,randUint)
		}
	}
}

func TestUint24Shifted2(t *testing.T) {
	const intBytesSize= 3
	const bitsSize= intBytesSize * 8
	const max= 0xFFFFFF

	for i := 0; i < testLooops; i++ {
		destSlice := make([]byte, intBytesSize+1)

		randUint := uint32(12089671) //uint32(rand.Intn(max))
		randOffset := rand.Intn(8)
		shiftedRandUint := randUint << uint(randOffset)

		// test PutUintShifted
		bitwisebytes.LittleEndian.PutUint24ShiftedBytes(randOffset, destSlice, randUint)
		resultUint := bitwisebytes.LittleEndian.Uint24ShiftedBytes(max, randOffset, destSlice)
		shiftedResultUint := bitwisebytes.LittleEndian.Uint32(destSlice)


		if resultUint != randUint {
			t.Log("randUint: ", randUint)
			t.Log("randOffset: ", randOffset)
			t.Log(destSlice)
			t.Errorf("mistmatch: %d != %d", randUint, resultUint)
		} else if debug {
			t.Log("resultUint : ", resultUint)
			t.Log("destSlice : ", destSlice)
		}


		if shiftedRandUint != shiftedResultUint {
			t.Logf("shiftedRandUint: 0x%x ", shiftedRandUint)
			t.Logf("shiftedResultUint: 0x%x ", shiftedResultUint)
			t.Log(destSlice)
			t.Errorf("mistmatch: 0x%x != 0x%x", shiftedRandUint, shiftedResultUint)
		} else if debug {
			t.Log("resultUint : ", resultUint)
			t.Log("destSlice : ", destSlice)
		}



	}

}

func TestUint32Shifted(t *testing.T) {
	const intBytesSize = 4
	const bitsSize = intBytesSize * 8
	const max = 0xFFFFFFFF

	for i:=0;i<testLooops;i++ {
		destSlice := make([]byte, intBytesSize + 1)

		randUint := uint32(rand.Intn(max))
		randOffset := rand.Intn(8)
		shiftedRandUint := uint64(randUint) << uint(randOffset)

		// test PutUintShifted
		bitwisebytes.LittleEndian.PutUint32ShiftedBytes(randOffset,destSlice,randUint)
		resultUint := bitwisebytes.LittleEndian.Uint32ShiftedBytes(max,randOffset,destSlice)

		if resultUint != randUint {
			t.Logf("randUint: 0x%x",randUint)
			t.Logf("randOffset: 0x%x ",randOffset)
			t.Log(destSlice)
			t.Errorf("mistmatch: %d != %d", randUint, resultUint)
		}else if debug {
			t.Logf("resultUint : 0x%x", resultUint)
			t.Logf("destSlice : 0x%x", destSlice)
		}

		// test ShiftRight of bytes
		unShiftedSlice, err:= bitwisebytes.ShiftRight(destSlice,uint(randOffset))
		if err != nil {
			t.Error(err.Error())
		}

		if debug {
			t.Log("unsfhited slice: ",unShiftedSlice )
		}

		newInt := bitwisebytes.LittleEndian.Uint32(unShiftedSlice[0:intBytesSize])

		if newInt != randUint {
			t.Error(newInt,randUint)
		}
		reShiftedSlice, err:= bitwisebytes.ShiftLeft(unShiftedSlice,uint(randOffset))
		if err != nil {
			t.Error(err.Error())
		}

		for i , byte := range  reShiftedSlice {
			if byte != destSlice[i] {
				t.Errorf("byte 0x%x is different",i)
				t.Log(byte,destSlice[i] )
			}
		}

		if debug {
			t.Log("reshifted slice: ", reShiftedSlice)
		}

		shiftedInt := bitwisebytes.LittleEndian.Uint40(reShiftedSlice)
		if shiftedInt != shiftedRandUint {
			t.Errorf("0x%x --> 0x%x", shiftedInt, shiftedRandUint)
		}
		newInt = bitwisebytes.LittleEndian.Uint32ShiftedBytes(max,randOffset,reShiftedSlice)
		if newInt != randUint {
			t.Error(newInt,randUint)
		}
	}
}

func TestUint40Shifted(t *testing.T) {
	const intBytesSize = 5
	const bitsSize = intBytesSize * 8
	const max = 0xFFFFFFFFFF

	for i:=0;i<testLooops;i++ {
		destSlice := make([]byte, intBytesSize + 1)

		randUint := uint64(rand.Intn(max))
		randOffset := rand.Intn(8)
		shiftedRandUint := uint64(randUint) << uint(randOffset)

		// test PutUintShifted
		bitwisebytes.LittleEndian.PutUint40ShiftedBytes(randOffset,destSlice,randUint)
		resultUint := bitwisebytes.LittleEndian.Uint40ShiftedBytes(max,randOffset,destSlice)

		if resultUint != randUint {
			t.Logf("randUint: 0x%x",randUint)
			t.Logf("randOffset: 0x%x ",randOffset)
			t.Log(destSlice)
			t.Errorf("mistmatch: %d != %d", randUint, resultUint)
		}else if debug {
			t.Logf("resultUint : 0x%x", resultUint)
			t.Logf("destSlice : 0x%x", destSlice)
		}

		// test ShiftRight of bytes
		unShiftedSlice, err:= bitwisebytes.ShiftRight(destSlice,uint(randOffset))
		if err != nil {
			t.Error(err.Error())
		}

		if debug {
			t.Log("unsfhited slice: ",unShiftedSlice )
		}

		newInt := bitwisebytes.LittleEndian.Uint40(unShiftedSlice[0:intBytesSize])

		if newInt != randUint {
			t.Error(newInt,randUint)
		}
		reShiftedSlice, err:= bitwisebytes.ShiftLeft(unShiftedSlice,uint(randOffset))
		if err != nil {
			t.Error(err.Error())
		}

		for i , byte := range  reShiftedSlice {
			if byte != destSlice[i] {
				t.Errorf("byte 0x%x is different",i)
				t.Log(byte,destSlice[i] )
			}
		}

		if debug {
			t.Log("reshifted slice: ", reShiftedSlice)
		}

		shiftedInt := bitwisebytes.LittleEndian.Uint48(reShiftedSlice)
		if shiftedInt != shiftedRandUint {
			t.Errorf("0x%x --> 0x%x", shiftedInt, shiftedRandUint)
		}
		newInt = bitwisebytes.LittleEndian.Uint40ShiftedBytes(max,randOffset,reShiftedSlice)
		if newInt != randUint {
			t.Error(newInt,randUint)
		}
	}
}

func TestUint48Shifted(t *testing.T) {
	const intBytesSize = 6
	const bitsSize = intBytesSize * 8
	const max = 0xFFFFFFFFFFFF

	for i:=0;i<testLooops;i++ {
		destSlice := make([]byte, intBytesSize + 1)

		randUint := uint64(rand.Intn(max))
		randOffset := rand.Intn(8)
		shiftedRandUint := uint64(randUint) << uint(randOffset)

		// test PutUintShifted
		bitwisebytes.LittleEndian.PutUint48ShiftedBytes(randOffset,destSlice,randUint)
		resultUint := bitwisebytes.LittleEndian.Uint48ShiftedBytes(max,randOffset,destSlice)

		if resultUint != randUint {
			t.Logf("randUint: 0x%x",randUint)
			t.Logf("randOffset: 0x%x ",randOffset)
			t.Log(destSlice)
			t.Errorf("mistmatch: %d != %d", randUint, resultUint)
		}else if debug {
			t.Logf("resultUint : 0x%x", resultUint)
			t.Logf("destSlice : 0x%x", destSlice)
		}

		// test ShiftRight of bytes
		unShiftedSlice, err:= bitwisebytes.ShiftRight(destSlice,uint(randOffset))
		if err != nil {
			t.Error(err.Error())
		}

		if debug {
			t.Log("unsfhited slice: ",unShiftedSlice )
		}

		newInt := bitwisebytes.LittleEndian.Uint48(unShiftedSlice[0:intBytesSize])

		if newInt != randUint {
			t.Error(newInt,randUint)
		}
		reShiftedSlice, err:= bitwisebytes.ShiftLeft(unShiftedSlice,uint(randOffset))
		if err != nil {
			t.Error(err.Error())
		}

		for i , byte := range  reShiftedSlice {
			if byte != destSlice[i] {
				t.Errorf("byte 0x%x is different",i)
				t.Log(byte,destSlice[i] )
			}
		}

		if debug {
			t.Log("reshifted slice: ", reShiftedSlice)
		}

		shiftedInt := bitwisebytes.LittleEndian.Uint56(reShiftedSlice)
		if shiftedInt != shiftedRandUint {
			t.Errorf("0x%x --> 0x%x", shiftedInt, shiftedRandUint)
		}
		newInt = bitwisebytes.LittleEndian.Uint48ShiftedBytes(max,randOffset,reShiftedSlice)
		if newInt != randUint {
			t.Error(newInt,randUint)
		}
	}
}

func TestUint56Shifted(t *testing.T) {
	const intBytesSize = 7
	const bitsSize = intBytesSize * 8
	const max = 0xFFFFFFFFFFFFFF

	for i:=0;i<testLooops;i++ {
		destSlice := make([]byte, intBytesSize + 1)

		randUint := uint64(rand.Intn(max))
		randOffset := rand.Intn(8)
		shiftedRandUint := uint64(randUint) << uint(randOffset)

		// test PutUintShifted
		bitwisebytes.LittleEndian.PutUint56ShiftedBytes(randOffset,destSlice,randUint)
		resultUint := bitwisebytes.LittleEndian.Uint56ShiftedBytes(max,randOffset,destSlice)

		if resultUint != randUint {
			t.Logf("randUint: 0x%x",randUint)
			t.Logf("randOffset: 0x%x ",randOffset)
			t.Log(destSlice)
			t.Errorf("mistmatch: %d != %d", randUint, resultUint)
		}else if debug {
			t.Logf("resultUint : 0x%x", resultUint)
			t.Logf("destSlice : 0x%x", destSlice)
		}

		// test ShiftRight of bytes
		unShiftedSlice, err:= bitwisebytes.ShiftRight(destSlice,uint(randOffset))
		if err != nil {
			t.Error(err.Error())
		}

		if debug {
			t.Log("unsfhited slice: ",unShiftedSlice )
		}

		newInt := bitwisebytes.LittleEndian.Uint56(unShiftedSlice[0:intBytesSize])

		if newInt != randUint {
			t.Error(newInt,randUint)
		}
		reShiftedSlice, err:= bitwisebytes.ShiftLeft(unShiftedSlice,uint(randOffset))
		if err != nil {
			t.Error(err.Error())
		}

		for i , byte := range  reShiftedSlice {
			if byte != destSlice[i] {
				t.Errorf("byte 0x%x is different",i)
				t.Log(byte,destSlice[i] )
			}
		}

		if debug {
			t.Log("reshifted slice: ", reShiftedSlice)
		}

		shiftedInt := bitwisebytes.LittleEndian.Uint64(reShiftedSlice)
		if shiftedInt != shiftedRandUint {
			t.Errorf("0x%x --> 0x%x", shiftedInt, shiftedRandUint)
		}
		newInt = bitwisebytes.LittleEndian.Uint56ShiftedBytes(max,randOffset,reShiftedSlice)
		if newInt != randUint {
			t.Error(newInt,randUint)
		}
	}
}

func TestUint64Shifted(t *testing.T) {
	const intBytesSize = 8
	const bitsSize = intBytesSize * 8
	const max = 0xFFFFFFFFFFFFFFF

	for i:=0;i<testLooops;i++ {
		destSlice := make([]byte, intBytesSize + 1)

		randUint := uint64(rand.Intn(max))
		randOffset := rand.Intn(8)
		//shiftedRandUint := uint64(randUint) << uint(randOffset)

		// test PutUintShifted
		bitwisebytes.LittleEndian.PutUint64ShiftedBytes(randOffset,destSlice,randUint)
		resultUint := bitwisebytes.LittleEndian.Uint64ShiftedBytes(max,randOffset,destSlice)
		//resultTmp := bitwisebytes.LittleEndian.Uint64ShiftedBytes(max,0,destSlice)

		if resultUint != randUint {
			t.Logf("randUint: 0x%x",randUint)
			t.Logf("randOffset: 0x%x ",randOffset)
			t.Log(destSlice)
			t.Errorf("mistmatch: 0x%X != 0x%X", randUint, resultUint)
		}else if debug {
			t.Logf("resultUint : 0x%x", resultUint)
			t.Logf("destSlice : 0x%x", destSlice)
			t.Logf("[0x%X 0x%X 0x%X 0x%X 0x%X 0x%X 0x%X 0x%X 0x%X]",destSlice[0],destSlice[1], destSlice[2],destSlice[3],destSlice[4], destSlice[5],destSlice[6],destSlice[7], destSlice[8])
		}

		// test ShiftRight of bytes
		unShiftedSlice, err:= bitwisebytes.ShiftRight(destSlice,uint(randOffset))
		if err != nil {
			t.Error(err.Error())
		}

		if debug {
			t.Logf("offset:%d",randOffset)
			t.Log("unshifted slice:")
			t.Logf("[0x%X 0x%X 0x%X 0x%X 0x%X 0x%X 0x%X 0x%X 0x%X]",unShiftedSlice[0],unShiftedSlice[1], unShiftedSlice[2],unShiftedSlice[3],unShiftedSlice[4], unShiftedSlice[5],unShiftedSlice[6],unShiftedSlice[7], unShiftedSlice[8])
		}

		//newInt := bitwisebytes.LittleEndian.Uint64(unShiftedSlice[0:intBytesSize])
		newInt := binary.LittleEndian.Uint64(unShiftedSlice[0:intBytesSize])

		if newInt != randUint {
			t.Errorf("0x%X -> 0x%X",newInt,randUint)
		}
		reShiftedSlice, err:= bitwisebytes.ShiftLeft(unShiftedSlice,uint(randOffset))
		if err != nil {
			t.Error(err.Error())
		}

		for i , byte := range  reShiftedSlice {
			if byte != destSlice[i] {
				t.Errorf("byte 0x%x is different",i)
				t.Log(byte,destSlice[i] )
			}
		}

		if debug {
			t.Log("reshifted slice: ", reShiftedSlice)
		}

		/*
		shiftedInt := bitwisebytes.LittleEndian.Uint64(reShiftedSlice)
		if shiftedInt != shiftedRandUint {
			t.Errorf("0x%x --> 0x%x", shiftedInt, shiftedRandUint)
		}
		newInt = bitwisebytes.LittleEndian.Uint56ShiftedBytes(max,randOffset,reShiftedSlice)
		if newInt != randUint {
			t.Error(newInt,randUint)
		}
		*/
	}
}

func TestBytesShifted(t *testing.T) {
	intBytesSize := rand.Intn(100) + 8
	//bitsSize := intBytesSize * 8

	randSlice := make([]byte,intBytesSize)
	maskSlice := make([]byte,intBytesSize)

	for i:=0;i<testLooops;i++ {
		destSlice := make([]byte, intBytesSize + 1)


		for i , _ := range randSlice {
			randSlice[i] = byte(rand.Intn(255))
			maskSlice[i] = byte(0xFF)
		}

		randOffset := rand.Intn(8)


		bitwisebytes.LittleEndian.PutBytesSliceShiftedBytes(randOffset,destSlice,randSlice)
		resultUint := bitwisebytes.LittleEndian.BytesSliceShiftedBytes(maskSlice,randOffset,destSlice)
		//resultTmp := bitwisebytes.LittleEndian.Uint64ShiftedBytes(max,0,destSlice)

		for i , aByte := range resultUint {

			if (aByte != randSlice[i]) && i != len(randSlice) - 1 {
				t.Logf("randUint: 0x%x",aByte)
				t.Logf("randOffset: 0x%x ",randOffset)
				t.Log(i, resultUint)
				t.Log(i, randSlice)
				t.Errorf("mistmatch: 0x%X != 0x%X", randSlice[i], aByte)
			}else if debug {
				t.Logf("resultUint : 0x%x", aByte)
				t.Logf("destSlice : 0x%x", destSlice)
				t.Logf("[0x%X 0x%X 0x%X 0x%X 0x%X 0x%X 0x%X 0x%X 0x%X]",destSlice[0],destSlice[1], destSlice[2],destSlice[3],destSlice[4], destSlice[5],destSlice[6],destSlice[7], destSlice[8])
			}


		}

	    /*
		// test ShiftRight of bytes
		unShiftedSlice, err:= bitwisebytes.ShiftRight(destSlice,uint(randOffset))
		if err != nil {
			t.Error(err.Error())
		}

		if debug {
			t.Logf("offset:%d",randOffset)
			t.Log("unshifted slice:")
			t.Logf("[0x%X 0x%X 0x%X 0x%X 0x%X 0x%X 0x%X 0x%X 0x%X]",unShiftedSlice[0],unShiftedSlice[1], unShiftedSlice[2],unShiftedSlice[3],unShiftedSlice[4], unShiftedSlice[5],unShiftedSlice[6],unShiftedSlice[7], unShiftedSlice[8])
		}

		//newInt := bitwisebytes.LittleEndian.Uint64(unShiftedSlice[0:intBytesSize])
		newInt := binary.LittleEndian.Uint64(unShiftedSlice[0:intBytesSize])

		if newInt != randUint {
			t.Errorf("0x%X -> 0x%X",newInt,randUint)
		}
		reShiftedSlice, err:= bitwisebytes.ShiftLeft(unShiftedSlice,uint(randOffset))
		if err != nil {
			t.Error(err.Error())
		}

		for i , byte := range  reShiftedSlice {
			if byte != destSlice[i] {
				t.Errorf("byte 0x%x is different",i)
				t.Log(byte,destSlice[i] )
			}
		}

		if debug {
			t.Log("reshifted slice: ", reShiftedSlice)
		}

		/*
		shiftedInt := bitwisebytes.LittleEndian.Uint64(reShiftedSlice)
		if shiftedInt != shiftedRandUint {
			t.Errorf("0x%x --> 0x%x", shiftedInt, shiftedRandUint)
		}
		newInt = bitwisebytes.LittleEndian.Uint56ShiftedBytes(max,randOffset,reShiftedSlice)
		if newInt != randUint {
			t.Error(newInt,randUint)
		}
		*/
	}
}

