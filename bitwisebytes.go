package bitwisebytes

import (
	"fmt"
	"encoding/binary"
	"math"
	"unsafe"
	log "github.com/sirupsen/logrus"
)

//BytesWordSize holds the size in bytes of a uint word
const BytesWordSize = uint(unsafe.Sizeof(uint(0)))
//BitsWordSize holds the number of bits in a uint word
const BitsWordSize = BytesWordSize *8

//ShiftLeft shifts a slice of bytes shiftCount bits to the left
func ShiftLeft(inputBuffer []byte, shiftCount uint) (outputBuffer []byte, err error) {
	bitsShift := shiftCount % BitsWordSize
	wordsShift := shiftCount / BitsWordSize

	widthInBytes := uint(len(inputBuffer))
	outputBuffer = make([]byte, len(inputBuffer))

	widthInWords := uint(widthInBytes / BytesWordSize)
	modulus := widthInBytes % BytesWordSize

	if modulus != 0 {
		widthInWords++
	}

	carry := uint(0)

	// --------------------------------------------------
	// Convert the slice of bytes into a slice of words
	// ---------------------------------------------------
	sliceOfWords := ByteSliceToWordSlice(inputBuffer)

	tmpWordsBuffer := make([]uint, widthInWords)  // here is where shifted words will be tmp saved

	// ---------------------------------------------------------
	// Itereate through the slice of words and shift each word
	// Store the shifted word in the tmp slice
	// --------------------------------------------------------
	for i , word := range sliceOfWords {

		shiftedWord := (uint(word) << bitsShift) | carry
		carry = uint(word)
		carryMask := uint(math.Pow(2,float64(bitsShift))) - 1
		maskShift := uint(BitsWordSize - bitsShift)
		carryMask = carryMask << maskShift
		carry = (carry & carryMask) >> uint(BitsWordSize -bitsShift)
		tmpWordsBuffer[i] = shiftedWord //| (carry & carryMask)
	}
	// ---------------------------
	// Now do the Words shifting
	// --------------------------
	outputBufferWords := ShiftWordsSliceLeft(tmpWordsBuffer, int(wordsShift))

	// ---------------------------------------------------------------------
	// Convert the final slice of words into its final representation as a
	// slice of bytes
	// --------------------------------------------------------------------
	outputBuffer = WordSliceToByteSlice(outputBufferWords)

	return outputBuffer[0:widthInBytes] , err
}

//ShiftRight shifts a slice of bytes shiftCount bits to the right
func ShiftRight(inputBuffer []byte, shiftCount uint) (outputBuffer []byte, err error) {
	bitsShift := shiftCount % BitsWordSize
	wordsShift := shiftCount / BitsWordSize

	widthInBytes := uint(len(inputBuffer))
	outputBuffer = make([]byte, len(inputBuffer))

	widthInWords := uint(widthInBytes / BytesWordSize)
	modulus := widthInBytes % BytesWordSize

	if modulus != 0 {
		widthInWords++
	}

	carry := uint(0)

	// --------------------------------------------------
	// Convert the slice of bytes into a slice of words
	// ---------------------------------------------------
	sliceOfWords := ByteSliceToWordSlice(inputBuffer)

	tmpWordsBuffer := make([]uint, widthInWords)  // here is where shifted words will be tmp saved

	// ---------------------------------------------------------
	// Itereate through the slice of words and shift each word
	// Store the shifted word in the tmp slice
	// --------------------------------------------------------
	//for i , word := range sliceOfWords {
	for i := len(sliceOfWords) -1 ; i>=0 ;i-- {
		word := sliceOfWords[i]
		shiftedWord := (uint(word) >> bitsShift) | carry
		tmpWordsBuffer[i] = shiftedWord
		carry = uint(word)
		carryMask := uint(math.Pow(2,float64(bitsShift))) - 1
		carryShift := uint(BitsWordSize - bitsShift)
		carry = (carry & carryMask) << carryShift
	}
	// ---------------------------
	// Now do the Words shifting
	// --------------------------

	outputBufferWords := ShiftWordsSliceRight(tmpWordsBuffer, int(wordsShift))

	// ---------------------------------------------------------------------
	// Convert the final slice of words into its final representation as a
	// slice of bytes
	// --------------------------------------------------------------------

	outputBuffer = WordSliceToByteSlice(outputBufferWords)

	return outputBuffer[0:widthInBytes] , err
}

//ByteSliceToWordSlice converts a slice of bytes into a slice of words
func ByteSliceToWordSlice(inputBytes []byte) (outputWords []uint) {

	widthInBytes := uint(len(inputBytes))
	widthInWords := widthInBytes / BytesWordSize
	modulus := widthInBytes % BytesWordSize

	if modulus != 0 {
		widthInWords++
	}

	// ------------------------------------
	// Instatiate the output slice of words
	// ------------------------------------
	outputWords = make([]uint,widthInWords)


	// --------------------------------------------------------
	// Iterate through the input slice of bytes in word sizes
	// extract the bytes
	// --------------------------------------------------------
	for i := uint(0); i< widthInWords; i++ {
		start := i * BytesWordSize
		end := start + BytesWordSize

		if end >= widthInBytes {
			end = widthInBytes
		}

		// ------------------------------------------
		// Create temporal byte slice that will store
		// the bytes to be converted to 1 word
		// ------------------------------------------
		tmpSlice := make([]byte, BytesWordSize)

		// Copy the corresponding bytes to the tmp slice
		for i, aByte := range inputBytes[start:end] {
			tmpSlice[i] = aByte
		}

		// -----------------------------------------
		// Do the slice of bytes to word conversion
		// -----------------------------------------
		var word uint
		switch BytesWordSize {
		case 8:
			word = uint(binary.LittleEndian.Uint64(tmpSlice))
		case 4:
			word = uint(binary.LittleEndian.Uint32(tmpSlice))
		default:
			log.Panic("Word size of %d no supported!!", BytesWordSize)
		}
		outputWords[i] = word
	}
		return outputWords
}

//BytesSliceToWordSlice converts a slice of words into a slice of bytes
func WordSliceToByteSlice(inputWords []uint) (outputBytes []byte) {
	outputBytes = make([]byte,0,uint(len(inputWords))*BytesWordSize)

	for _ , word := range inputWords {
		// Create subslice here
		tmpBuf := make([]byte, BytesWordSize)
		binary.LittleEndian.PutUint64(tmpBuf,uint64(word))
		outputBytes = append(outputBytes,tmpBuf...)

	}
	return outputBytes
}

//ShiftWordsSliceLeft shifts a slice of words x words left
func ShiftWordsSliceLeft(buff []uint, shiftWords int) (returnBuff []uint) {
	returnBuff = make([]uint, len(buff))
	for i , word := range buff {
		if i+shiftWords < len(buff) {
			returnBuff[i+shiftWords] = uint(word)
		}else{
			break
		}
	}
	return returnBuff
}

//ShiftWordsSliceRight shifts a slice of words x words left
func ShiftWordsSliceRight(buff []uint, shiftWords int) (returnBuff []uint) {
	returnBuff = make([]uint, len(buff))
	for i , _ := range buff {
		if i+shiftWords < len(buff) {
			returnBuff[i] = buff[i+shiftWords]
		}else{
			break
		}
	}
	return returnBuff
}


func And(inputOutput []byte, operand []byte) (err error){
	if len(inputOutput) != len(operand) {
		return fmt.Errorf("input and operand must of of the same lenth")
	}
	for i , op := range  operand {
		inputOutput[i] = inputOutput[i] & op
	}
	return err
}


func Or(inputOutput []byte, operand []byte) (err error){
	if len(inputOutput) != len(operand) {
		return fmt.Errorf("input and operand must of of the same lenth")
	}
	for i , op := range  operand {
		inputOutput[i] = inputOutput[i] | op
	}
	return err
}

func MakeMask(size uint, width uint, offset uint) (outputMask []byte) {

	maskWords := size / BytesWordSize
	modulus := size % BytesWordSize
	if modulus >0 {
		maskWords ++
	}
	wordsSlice := make([]uint,maskWords)
	widthModulus := width % BitsWordSize

	for i, _ := range wordsSlice {
		if uint(i) == (maskWords - 1) {
			if widthModulus > 0 {
				wordsSlice[i] = uint(math.Pow(2,float64(widthModulus))) - 1
			}else{
				wordsSlice[i] = uint(math.Pow(2,float64(BitsWordSize))) - 1
			}
		}else {
			wordsSlice[i] = uint(math.Pow(2,float64(BitsWordSize))) - 1
		}
	}
	outputMask = WordSliceToByteSlice(wordsSlice)

	if offset > 0 {
		var err error
		outputMask , err = ShiftLeft(outputMask,offset)
		if err != nil {
			log.Panic(err.Error())
		}
	}

	return outputMask[0:size]
}