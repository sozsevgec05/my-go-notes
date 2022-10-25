package defer_statement

import "fmt"

func Demo2(sayi int) string {
	defer DeferFunc()
	if sayi%2 == 0 {
		return "Çift Sayıdır"
	}

	if sayi%2 != 0 {
		return "Tek Sayıdır"
	}

	return "..."
}

func Test() {
	sonuc := Demo2(10)
	fmt.Println(sonuc)
}

func DeferFunc() {
	fmt.Println("Bitti!!!")
}
