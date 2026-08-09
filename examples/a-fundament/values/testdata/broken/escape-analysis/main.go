package main

type Course struct {
	title string
}

func NewCourse(title string) *Course {
	c := Course{title: title}
	return &c
}

func main() {
	_ = NewCourse("Go")
}
