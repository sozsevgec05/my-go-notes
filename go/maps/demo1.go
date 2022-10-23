package maps

import "fmt"

func Demo1() {
	// key / value
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["computer"] = "Bilgisayar"
	sozluk["pencil"] = "Kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("Eleman Sayısı: ", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)
	delete(sozluk, "book")
	fmt.Println("Eleman Sayısı: ", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)

	deger := sozluk["table"]
	fmt.Println(deger)

	deger, varMi := sozluk["book"]
	fmt.Println(deger)
	fmt.Println("Listede Olma Durumu: ", varMi)

}
