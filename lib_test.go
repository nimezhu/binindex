package binindex

import (
	"testing"
)

func TestIO(t *testing.T) {
	t.Log(range2bin(1000, 2000))
	t.Log(bin2range(585))
}
