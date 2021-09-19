# Variables

Go dilinde oldukça farklı şekillerde değişken tanımı yapılabilir.

```go
var a int // sadece tanımlama
var a2 int = 1 // tanımlama ve değer atama

var b = "go turkiye" // veri tipini belirtmeden, tanımlama ve değer atama

var c, c2 int = 1, 2 // veri tipini belirterek, birden çok tanımlama ve değer atama

q, w, e := "go", "turkiye", 1 // var keywordunu kullanmadan ve veritipini belirtmeden tanımlama ve değer atama

var ( // var keywordu ile gruplayarak tanımlama
	dilAdi     = "go"
	konum      = "turkiye"
	kisiSayisi int
)

const ( // const keywordu ile gruplayarak tanımlama
	pi    = 3.14
	sabit = "sabit deger"
)
```

## Klasik tanımlamalar

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

## Multiple

```go
package main

import "fmt"

var ( // var keywordu ile gruplayarak tanımlama yapma
	dilAdi     = "go"
	konum      = "turkiye"
	kisiSayisi int
)

func main() {
	a, b, c, d := 1, 2, "go", true
	fmt.Println(a, b, c, d)
	a, b, c, d = 2, 1, "go turkiye", false
	fmt.Println(a, b, c, d)

	var i, j int = 1, 2
	fmt.Println(i, j)
	i, j = 2, 1
	fmt.Println(i, j)

	fmt.Println(dilAdi, konum, kisiSayisi)
}
```

Aynı satırda birden fazla değişken tanımlayabilirsiniz, bunlar birbirlerinden farklı tiplere veya sadece tek bir tipe sahip olabilir, tanımlama yaptığınız satırda değer atamasını da yapabilirsiniz. Birden çok değişkenin değerini tanımlarken olduğu gibi tek bir satırda değiştirebilirsiniz.

Bunun dışında değişken tanımlarınızı, fonksiyonların içinde veya dışında var keywordu altında gruplayarak tanımlayabilirsiniz.

## Constants

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