package antipatterns

import (
	"testing"

	"github.com/deliseev/letting-go/examples/a-fundament/values/course"
)

func TestValueReceiverDoesNotMutate(t *testing.T) {
	c := Course{title: "Go для питонистов"}
	c.SetTitle("Letting Go")
	if c.title != "Go для питонистов" {
		t.Fatalf("title = %q, want unchanged %q", c.title, "Go для питонистов")
	}
}

func TestEnrollLostDoesNotGrow(t *testing.T) {
	students := make([]course.Student, 0, 10)
	enroll(students, course.Student{})
	if len(students) != 0 {
		t.Fatalf("len(students) = %d, want 0", len(students))
	}
}
