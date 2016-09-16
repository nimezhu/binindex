package binindex

type Coord struct {
	Chr   string
	Start int
	End   int
}
type BinIndexI interface {
	Add(o interface{}, name string) error
	Size() int
	Del(name string)
	Get(name string) (interface{}, bool)
	List() <-chan string
	AddCoord(name string, coord Coord)
	DelCoord(name string)
	Query(coord Coord) <-chan string
	//Writable() bool
}
