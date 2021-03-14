# converter

Go语言结构体互相转换的库

### 例子:

```go
package main

import (
	"fmt"

	"github.com/ltoddy/converter"
)

type User struct {
	Nickname string `convert:"field:Name"`
	Age  int
}

type Person struct {
	Name string
	Old  int `convert:"from:Age"`
}

func main() {
	u := &User{"ltoddy", 23}
	p := new(Person)
	converter.Convert(u, p)

	fmt.Printf("Person(Name = %v, Old = %v)\n", p.Name, p.Old)
	// Person(Name = ltoddy, Old = 23)
}
```

TODO:

- 嵌套结构体
- array/slice
