// Package course holds the idiomatic LMS domain code for the
// "Семантика значений и ссылок" module.
package course

type Student struct {
	id   int
	name string
}

type Course struct {
	id       int
	title    string
	students []Student
}

func (c *Course) Title() string {
	return c.title
}

// region: pointer-receiver
func (c *Course) SetTitle(t string) {
	c.title = t
}

// endregion: pointer-receiver

// region: titled-interface
type Titled interface {
	Title() string
	SetTitle(string)
}

func rename(t Titled, s string) {
	t.SetTitle(s)
}

// endregion: titled-interface

// region: rename-by-index
func renameByIndex(students []Student, i int, name string) {
	students[i].name = name
}

// endregion: rename-by-index

// region: enroll-returns
func enroll(students []Student, s Student) []Student {
	return append(students, s)
}

// endregion: enroll-returns

// region: new-course
func NewCourse(title string) *Course {
	c := Course{title: title}
	return &c
}

// endregion: new-course
