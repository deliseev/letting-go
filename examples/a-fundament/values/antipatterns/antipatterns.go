// Package antipatterns holds code that compiles but does the wrong
// thing — the antipatterns discussed in "Семантика значений и
// ссылок" before their idiomatic fix in package course.
package antipatterns

import (
	"github.com/deliseev/letting-go/examples/a-fundament/values/course"
)

// region: value-receiver
type Course struct {
	id    int
	title string
}

func (c Course) SetTitle(t string) {
	c.title = t
}

// endregion: value-receiver

// region: enroll-lost
func enroll(students []course.Student, s course.Student) {
	students = append(students, s)
}

// endregion: enroll-lost
