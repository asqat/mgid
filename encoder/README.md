# Получение encoder

```bash
go get -v "github.com/asqat/mgid/encoder"
```

# Запуск тестов
```bash
go test -v
```

# Использование
```go
package main

import (
	"fmt"
	"github.com/asqat/mgid/encoder"
)

func main() {
	encoded, err := encoder.Encode(1, []string{"res1", "res2", "res3"})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(encoded)) // out: {"type":1,"result":["res1","res2","res3"]}

	encoded, err = encoder.Encode(2, []string{"res1", "res2", "res3"})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(encoded)) // // out: {"type":2,"result":{"0":"res1","1":"res2","2":"res3"}}
}
```