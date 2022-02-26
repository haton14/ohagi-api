package anycast

func ToIntP(i interface{}) *int {
	num, ok := i.(int)
	if !ok {
		return nil
	}
	return &num
}
