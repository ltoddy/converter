package converter

import "testing"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{name, age}
}

type Other struct {
	Name string `convert:"from:Name,to:Name"`
	Age  int    `convert:"from:Age,to:Age"`
}

func TestConvert(t *testing.T) {
	t.Run("test convert simple struct", func(t *testing.T) {
		user := NewUser("ltoddy", 23)
		other := new(Other)

		err := Convert(user, other)
		if err != nil {
			t.Fatal()
		}

		if other.Name != "ltoddy" && other.Age != 23 {
			t.Fatal()
		}
	})
}
