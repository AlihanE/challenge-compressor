package bit

import "testing"

func TestBit(t *testing.T) {
	testData := [][]int{{1, 1}, {1, 0}, {0, 0}, {0, 0, 1, 1, 0, 1}}
	b := NewBitWriter(nil)
	for _, td := range testData {
		b.AddBits(td)
	}

	b.flushCache()
	for b.Read() {
		t.Log(GetBitArray(b.CurrentByte()))
	}
	b.Close()
}
