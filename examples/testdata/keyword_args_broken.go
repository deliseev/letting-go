//go:build broken

package testdata

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) User {
	return User{Name: name, Age: age}
}

func Example() User {
	return NewUser(name: "Ada", age: 36)
}
