package functions

import "fmt"

func SelamVer(kullanıcıAdi string) {
	fmt.Println("Selam! ", kullanıcıAdi)
}

func Topla(sayi1 int, sayi2 int) int {
	var toplam = sayi1 + sayi2
	return toplam
}
