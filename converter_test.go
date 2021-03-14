package converter

import "testing"

func TestConvert(t *testing.T) {
	t.Run("test simple struct without tags", func(t *testing.T) {
		type user struct {
			Name string
			Age  int
		}

		type other struct {
			Name string
			Age  int
		}

		u := user{Name: "ltoddy", Age: 23}
		o := new(other)

		Convert(&u, o)
		if o.Name != "ltoddy" || o.Age != 23 {
			t.Fatal()
		}
	})

	t.Run("test convert simple struct", func(t *testing.T) {
		type user struct {
			Name string
			Age  int
		}

		type other struct {
			Name string `convert:"from:Name,to:Name"`
			Age  int    `convert:"from:Age,to:Age"`
		}

		u := &user{"ltoddy", 23}
		o := new(other)

		Convert(u, o)
		if o.Name != "ltoddy" && o.Age != 23 {
			t.Fatal()
		}
	})
}
