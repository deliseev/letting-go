package main

type Student struct {
	id   int
	name string
}

type Course struct {
	id       int
	students []Student
}

func main() {
	a := Course{id: 1}
	b := Course{id: 1}
	_ = a == b
}
