# Variables

Go dilinde farklı şekillerde değişken tanımı yapılabilir.

```go
// Veri tipi belirterek sadece tanımlama.
var a int

// Değer atarken veri tipini belirterek tanımlama.
var a2 int = 1

// Veri tipi belirtmeden, tanımlama ve değer atama. b değişkeni artık string veri tipine sahiptir.
var b = "go turkiye"

/* Veri tipini belirtmeden, tek satırda birden çok değişken tanımlama ve değer atama. 
c ve c2 değişkenleri int, c3 değişkeni ise string veri tipine sahip olacak. */
var c, c2, c3 = 1, 2, "go turkiye"

/* Veri tipini belirterek, tek satırda birden çok değişken tanımlama ve değer atama.
Bu tanımlama şekli ile tek satırda farklı tiplere sahip olacak değişken tanımlaması yapamazsınız. */
var c3, c4 int = 1, 2

/* Var keywordu kullanmadan, tek satırda farklı tiplerde bir veya birden çok değişkeni tanımlayabilirsiniz. 
q ve w string, e ise int veri tipini alacak */
q, w, e := "go", "turkiye", 1

/* Gruplayarak tanımlama, veri tipini belirtmek veya belirtmemek opsiyoneldir.
Veri tipi belirtilmezse, değişkene bir değer ataması yapılmalıdır. */
var (
	dilAdi     = "go"
	konum      = "turkiye"
	kisiSayisi int
)

// Const keywordu ile constant değerleri gruplayarak tanımlama.
const (
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

var (
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