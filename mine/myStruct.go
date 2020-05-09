package mine

type MyStruct struct {
}

func (m *MyStruct) Add(one, two int) int {
	return one + two
}
