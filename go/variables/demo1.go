package variables

import "fmt"

//camelCase
//PascalCase
func Demo1() {
	var metin string = "Merhaba Go!"
	fmt.Println(metin)

	var kdv int = 10
	fmt.Println(kdv)
	fmt.Println(100 + (100 * 10 / 100))

	var kdv2 float32 = 0.1
	fmt.Println(kdv2)
	fmt.Println(100 + (100 * kdv2))

	kdv3 := 25 //int
	fmt.Println(kdv3)
	fmt.Printf("Veri Tipi: %T", kdv3)

	kdv4 := 2.5 //float
	fmt.Println(kdv3)
	fmt.Printf("Veri Tipi: %T\n", kdv4)

	//Mantıksal Veri Tipleri

	// ==
	var durum bool = true // false

	var metin1 string = "Selahattin"
	var metin2 string = "ÖZSEVGEÇ"

	durum = metin1 == metin2
	fmt.Println(durum)

	// !=
	var durum2 bool = true // false

	var metin3 string = "Selahattin"
	var metin4 string = "ÖZSEVGEÇ"

	durum2 = metin3 == metin4
	fmt.Println(durum2)
}
