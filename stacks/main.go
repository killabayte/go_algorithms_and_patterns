package main

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {

}

func (s *Stack) Pop() int {
	return 0
}

func main() {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
}
