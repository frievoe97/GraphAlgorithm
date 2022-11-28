package main

type Stack struct {
	Elements []*Vertex
}

func (s *Stack) Push(elem *Vertex) {
	s.Elements = append(s.Elements, elem)
}

func (s *Stack) Peek() *Vertex {
	return s.Elements[len(s.Elements)-1]
}

/*func (s *Stack) Pop() *Vertex {
	returnVertex := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return returnVertex
}*/

func (s *Stack) Pop() *Vertex {
	returnVertex := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return returnVertex
}

func (s Stack) Empty() bool {
	return len(s.Elements) == 0
}
