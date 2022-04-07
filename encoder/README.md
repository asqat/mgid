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
	t1 := encoder.Template{
		Type: 1,
		Result: []string{
			"res1", "res2", "res3",
		},
	}
	enc, err := encoder.Encode(t1)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(enc)) // out: {"type":1,"result":["res1","res2","res3"]}

	t2 := encoder.Template{
		Type: 2,
		Result: map[string]string{
			"0": "res1",
			"1": "res2",
			"2": "res3",
		},
	}
	enc, err = encoder.Encode(t2)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(enc)) // out: {"type":2,"result":{"0":"res1","1":"res2","2":"res3"}}
}
```