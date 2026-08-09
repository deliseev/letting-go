// Package lms is the exercise starting point for "Семантика
// значений и ссылок" — intentionally does not compile; testdata is
// ignored by go tooling, so go build ./... skips it.
package lms

// region: exercise-start
type Student struct {
	id   int
	name string
}

type Course struct {
	id       int
	title    string
	students []Student
}

func (c Course) Enroll(s Student) {
	c.students = append(c.students, s)
}

func (c Course) Rename(title string) {
	c.title = title
}

func report(courses []Course) map[Course]int {
	stats := map[Course]int{}
	for _, c := range courses {
		stats[c] = len(c.students)
	}
	return stats
}

// endregion: exercise-start
