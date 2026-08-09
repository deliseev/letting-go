package main

type Course struct {
	id    int
	title string
}

func (c *Course) Title() string {
	return c.title
}

func (c *Course) SetTitle(t string) {
	c.title = t
}

type Titled interface {
	Title() string
	SetTitle(string)
}

func rename(t Titled, s string) {
	t.SetTitle(s)
}

func main() {
	c := Course{id: 1}
	rename(c, "Letting Go")
}
