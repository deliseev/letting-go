package course

import (
	"fmt"
	"testing"
)

func TestSetTitle(t *testing.T) {
	c := NewCourse("Go для питонистов")
	c.SetTitle("Letting Go")
	if c.Title() != "Letting Go" {
		t.Fatalf("Title() = %q, want %q", c.Title(), "Letting Go")
	}
}

func TestTitledInterface(t *testing.T) {
	c := NewCourse("Go")
	rename(c, "Letting Go")
	if c.Title() != "Letting Go" {
		t.Fatalf("Title() = %q, want %q", c.Title(), "Letting Go")
	}
}

func TestRenameByIndex(t *testing.T) {
	students := []Student{{name: "Аня"}}
	renameByIndex(students, 0, "Другое имя")
	if students[0].name != "Другое имя" {
		t.Fatalf("name = %q, want %q", students[0].name, "Другое имя")
	}
}

func TestEnroll(t *testing.T) {
	students := enroll(nil, Student{name: "Аня"})
	if len(students) != 1 {
		t.Fatalf("len(students) = %d, want 1", len(students))
	}
}

func Example_subsliceAliasing() {
	// region: subslice-aliasing
	c := Course{students: []Student{
		{name: "Аня"}, {name: "Боря"}, {name: "Вика"},
	}}
	first := c.students[:2]
	first[0].name = "Другое имя"
	fmt.Println(c.students[0].name)
	// endregion: subslice-aliasing

	// Output: Другое имя
}

func Example_nilSliceAppend() {
	// region: nil-slice-append
	var students []Student
	students = append(students, Student{name: "Аня"})
	fmt.Println(len(students))
	// endregion: nil-slice-append

	// Output: 1
}

func Example_nilMapRead() {
	// region: nil-map-read
	var enrolled map[int]Student
	s, ok := enrolled[42]
	fmt.Printf("%+v %v\n", s, ok)
	// endregion: nil-map-read

	// Output: {id:0 name:} false
}

func TestNilMapWritePanics(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("assignment to nil map did not panic")
		}
		got := fmt.Sprint(r)
		want := "assignment to entry in nil map"
		if got != want {
			t.Fatalf("panic = %q, want %q", got, want)
		}
	}()

	var enrolled map[int]Student
	enrolled[42] = Student{name: "Аня"}
}

func Example_studentEquality() {
	// region: struct-equality
	a := Student{id: 1, name: "Аня"}
	b := Student{id: 1, name: "Аня"}
	fmt.Println(a == b)
	// endregion: struct-equality

	// Output: true
}
