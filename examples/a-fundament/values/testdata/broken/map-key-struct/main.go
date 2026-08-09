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
	grades := map[Course]int{}
	_ = grades
}
