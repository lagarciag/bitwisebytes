package bitwisebytes_test


import (
	"testing"
	"os"
	"github.com/lagarciag/bitwisebytes"
	"time"
	"fmt"
	"math/rand"
)

func TestMain(t *testing.M) {
	seed := int64(time.Now().Nanosecond())
	rand.Seed(seed)
	fmt.Println("SEED:", seed)
	v := t.Run()
	os.Exit(v)

}

func TestShiftRight(t *testing.T) {
	inputSlice := []byte{1,1}
	outputSlice, _ := bitwisebytes.ShiftRight(inputSlice,1)

	t.Log("outputslice shiftRight: ",outputSlice)

	outputSlice, _ = bitwisebytes.ShiftLeft(outputSlice,1)

	t.Log("outputslice shiftLeft: ",outputSlice)

	if outputSlice[1] != 1 {
		t.Error("last byte shoulf be 1")
	}
}

func TestShiftLeft(t *testing.T) {
	inputSlice := []byte{1,0,0,0,0,0,0,0 ,0,0,0,0,0,0,0,0x80,0,0}
	t.Logf("input bitfield:         %-3v", inputSlice)
	outputSlice, _ := bitwisebytes.ShiftLeft(inputSlice,1)

	t.Logf("outputslice shiftLeft:  %-3v",outputSlice)

	outputSlice, _ = bitwisebytes.ShiftRight(outputSlice,1)

	t.Logf("outputslice shiftRight: %-3v",outputSlice)

}

func TestShiftLeft2(t *testing.T) {
	inputSlice := []byte{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}
	t.Logf("input bitfield:         %-3v", inputSlice)
	outputSlice, _ := bitwisebytes.ShiftLeft(inputSlice,32)

	t.Logf("outputslice shiftLeft:  %-3v",outputSlice)

	outputSlice, _ = bitwisebytes.ShiftRight(outputSlice,32)

	t.Logf("outputslice shiftRight: %-3v",outputSlice)

	lostBytes := 4

	for i := 0;i < (len(inputSlice) - lostBytes);i++ {
		if inputSlice[i] != outputSlice[i] {
			t.Error("mistmatch in byte comparison")
		}
	}
}

func TestBitfieldShiftLeftRight(t *testing.T) {
	inputBitfield := []byte{1,0,0,0,0,0,0,0 ,0,0,0,0,0,0,0,0x80,0,0,0,0,0,0,0,0,0}
	t.Log(" input bitfield:", inputBitfield)

	outPutBitfield , err := bitwisebytes.ShiftLeft(inputBitfield, 1)

	t.Log("output bitfield:", outPutBitfield)

	if err != nil {
		t.Error(err.Error())
	}

	if len(inputBitfield) != len(outPutBitfield) {
		t.Error("Sizes do not match: ",len(inputBitfield), len(outPutBitfield))
	}

	if outPutBitfield[0] != 2 {
		t.Error("Byte 0 should be 2")
	}

	if outPutBitfield[1] != 0 {
		t.Error("Byte 1 should be 0")
	}

	t.Log("output: ",outPutBitfield)


	// Shift Right

	outPutBitfield, err = bitwisebytes.ShiftRight(outPutBitfield, 1)
	t.Log("shiftRight:",outPutBitfield)



	if len(inputBitfield) != len(outPutBitfield) {
		t.Error("Sizes do not match: ",len(inputBitfield), len(outPutBitfield))
	}


	if outPutBitfield[0] != inputBitfield[0] {
		t.Error("Must be equal")
	}


	inputBitfield = []byte{1,1}
	t.Log(" input bitfield:", inputBitfield)

	outPutBitfield , err = bitwisebytes.ShiftLeft(inputBitfield, 1)

	t.Log("output bitfield:", outPutBitfield)

	if err != nil {
		t.Error(err.Error())
	}

	if len(inputBitfield) != len(outPutBitfield) {
		t.Error("Sizes do not match: ",len(inputBitfield), len(outPutBitfield))
	}


	if outPutBitfield[1] != 2 {
		t.Error("Byte 0 should be 2")
	}


	t.Log("output: ",outPutBitfield)



}

func TestByteSliceToWordSlice(t *testing.T) {
	inputByteSlice := make([]byte,bitwisebytes.BitsWordSize)

	if bitwisebytes.BytesWordSize == 8 {
		inputByteSlice = []byte{1,0,0,0,0,0,0,0,1}
	}else if bitwisebytes.BytesWordSize == 4 {
		inputByteSlice = []byte{1,0,0,0,1}
	}else{
		t.Error("platform no supported")
	}
	outputSlice := bitwisebytes.ByteSliceToWordSlice(inputByteSlice)

	if outputSlice[0] != 1 || outputSlice[1] != 1  {
		t.Error("Word value not expected", outputSlice)
	}
}

func TestShiftWordsSliceRightLeft(t *testing.T) {
	inputSlice := []uint{1,2,3,4,5}

	outputSlice := bitwisebytes.ShiftWordsSliceLeft(inputSlice, 1)
	t.Log("outpusSlice shiftLeft: ",outputSlice)


	outputSlice = bitwisebytes.ShiftWordsSliceRight(outputSlice, 1)
	t.Log("outpusSlice shifRight: ",outputSlice)


	if len(inputSlice) != len(outputSlice) {
		t.Error("lengths differ")
	}

	if (inputSlice[0] != outputSlice[0]) {
		t.Error("first words should be equal")
	}

}

func TestShiftWordsSliceRightLeftLong(t *testing.T) {
	inputSlice := []uint{2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0}

	outputSlice := bitwisebytes.ShiftWordsSliceLeft(inputSlice, 1)
	t.Log("outpusSlice shiftLeft: ",outputSlice)


	outputSlice = bitwisebytes.ShiftWordsSliceRight(outputSlice, 1)
	t.Log("outpusSlice shifRight: ",outputSlice)


	if len(inputSlice) != len(outputSlice) {
		t.Error("lengths differ")
	}

	if (inputSlice[0] != outputSlice[0]) {
		t.Error("first words should be equal")
	}

}
