package bit

import (
	"io"
	"math"
)

type BitWriter struct {
	bytes           []byte
	byteCache       []int
	currentPosition int
	writer          io.WriteCloser
}

func NewBitWriter(w io.WriteCloser) *BitWriter {
	return &BitWriter{
		currentPosition: -1,
		writer:          w,
	}
}

func (bw *BitWriter) AddBits(arr []int) {
	freeMem := 8 - len(bw.byteCache)
	for _, b := range arr {
		if freeMem == 0 {
			bw.flushCache()
			freeMem = 8 - len(bw.byteCache)
		}
		bw.byteCache = append(bw.byteCache, b)
		freeMem--
	}
}

func (bw *BitWriter) flushCache() {
	bw.bytes = append(bw.bytes, makeByte(bw.byteCache))
	bw.byteCache = []int{}
	if len(bw.bytes) == 1000 && bw.writer != nil {
		_, err := bw.writer.Write(bw.bytes)
		if err != nil {
			panic(err)
		}
		bw.bytes = []byte{}
	}
}

func (bw *BitWriter) Read() bool {
	bw.currentPosition++
	return bw.currentPosition < len(bw.bytes)
}

func (bw *BitWriter) CurrentByte() byte {
	return bw.bytes[bw.currentPosition]
}

func (bw *BitWriter) Close() {
	bw.flushCache()
	if bw.writer != nil {
		_, err := bw.writer.Write(bw.bytes)
		if err != nil {
			panic(err)
		}
	}
}

func makeByte(arr []int) byte {
	res := byte(0)
	for i := len(arr) - 1; i >= 0; i-- {
		res = res<<1 | byte(arr[i])
	}
	return res
}

func GetBitArray(b byte) [8]byte {
	res := [8]byte{}
	for i := 0; i < 8; i++ {
		sec := byte(math.Pow(2, float64(i)))
		if (b & sec) > 0 {
			res[i] = 1
		} else {
			res[i] = 0
		}

	}
	return res
}
