package binindex

import (
	"github.com/nimezhu/ice"
)

type BinIndexMap struct {
	Data  *ice.Set
	Names map[string]int
	Index map[string]Bin
}

// Store Names and Index TODO

type BinIndex struct {
	Bins []Bin
}
type Bin []int

var (
	binOffsets    = []uint{512 + 64 + 8 + 1, 64 + 8 + 1, 8 + 1, 1, 0}
	binFirstShift = uint(17)
	binNextShift  = uint(3)
	binLength     = []uint{4096 + 512 + 64 + 8 + 1}
)

func NewBinIndexMap(set *ice.Set) *BinIndexMap {
	m := make(map[string]int)
	index := make(map[string]Bin)
	a := BinIndexMap{Data: set, Names: m, Index: index}
	return &a
}

func range2bin(start uint, end uint) uint {
	startBin := start
	endBin := end - 1
	startBin >>= binFirstShift
	endBin >>= binFirstShift
	for _, v := range binOffsets {
		if startBin == endBin {
			return v + startBin
		}
		startBin >>= binNextShift
		endBin >>= binNextShift

	}
	return 0
}
func iterRangeOverlapBins(start uint, end uint) <-chan uint {
	ch := make(chan uint)
	go func() {
		startBin := start
		endBin := end - 1
		startBin >>= binFirstShift
		endBin >>= binFirstShift
		for _, v := range binOffsets {
			for j := startBin; j < endBin+1; j++ {
				ch <- j + v
			}
			startBin >>= binNextShift
			endBin >>= binNextShift
		}
		close(ch)
	}()
	return ch
}

func bin2range(bin uint) (uint, uint) {
	binShift := binFirstShift
	for _, v := range binOffsets {
		if bin-v >= 0 {
			bin = bin - v
			break
		}
		binShift += binNextShift
	}
	return bin << binShift, (bin + 1) << binShift
}

func bin2length(bin uint) uint {
	start, end := bin2range(bin)
	return end - start
}

func bin2level(bin uint) int {
	for i, v := range binOffsets {
		if bin-v >= 0 {
			return 4 - i
		}
	}
	return 0
}

/*
func (m *BinIndexMap).Add(o interface{}, name string) error {

}
*/
