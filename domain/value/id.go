package value

import "fmt"

type ID struct {
	value int
}

func NewID(v int) (*ID, error) {
	if v < 0 {
		return nil, fmt.Errorf("[%w]IDが負数", ErrMinRange)
	}
	return &ID{v}, nil
}

func (v ID) Value() int {
	return v.value
}
