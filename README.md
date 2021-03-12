# converter

```go
package main

import (
	"fmt"
	"log"

	"github.com/ltoddy/converter"
)

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

func main() {
	user := NewUser("ltoddy", 23)
	other := new(Other)

	err := converter.Convert(user, other)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Other(Name = %v, Age = %v)\n", other.Name, other.Age)
}
```
