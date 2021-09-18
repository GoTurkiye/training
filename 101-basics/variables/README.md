# Variables

Go dilinde üç farklı şekilde değişken tanımı yapılabilir.

```go
package main

import "fmt"

var degisken_1 string
var degisken_2 = "deger 2"

func main() {
	degisken_1 = "deger 1"
	degisken_3 := "deger 3"

	fmt.Println(degisken_1,degisken_2,degisken_3)
}
```
Yukarıdaki örnekte görüldüğü üzere `var` keywordu ile belirli bir tipte değişken tanımı yapabiliriz.
`var` keywordü ile fonksiyon çağrısı içerisinde olmaksızın uygulamada ihtiyaç duyguduğumuz yerde değişken tanımı yapabilirken,
`:=` operatörü ile sadece fonksiyon çağrıları içerisinde değişken tanımı yapılabiliriz.
Eğer değişkenin değeri belli değil ise `degisken_1`de olduğu üzere tipini belirtmemiz gerekmekte.
Fakat değişken değerimiz belirli ise `degisken_2`de olduğu üzere tip belirtmedende değişken tanımı yapabiliriz.

# Constants

```go
package main

import "fmt"

const sabit_1 = "deger 1"

const (
	sabit_2 = "deger 2"
	sabit_3 = "deger 3"
	sabit_4 = "deger 4"
)

const (
	sabit_5 = iota
	sabit_6
	sabit_7
)

func main() {
	//sabit_1 = "deger 1.1"

	fmt.Println(sabit_1, sabit_2, sabit_3, sabit_4, sabit_5, sabit_6, sabit_7)
}

```