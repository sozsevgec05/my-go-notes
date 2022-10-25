package string_functions

import (
	"fmt"
	"strings"
)

func Demo2() {
	isim := "Selehattin ÖZSEVGEÇ"

	fmt.Println(strings.HasPrefix(isim, "Sel"))
	fmt.Println(strings.HasSuffix(isim, "Sel"))

	fmt.Println(strings.Index(isim, "SEV"))

	harfler := []string{"s", "e", "l", "o"}

	sonuc := strings.Join(harfler, "-")
	fmt.Println(sonuc)

	fmt.Println(strings.Replace(sonuc, "-", "/", -1))
	fmt.Println(strings.Replace(sonuc, "-", "/", 1))
	fmt.Println(strings.Replace(sonuc, "-", "/", 2))
	fmt.Println(strings.Replace(sonuc, "-", "/", 3))

	fmt.Println(strings.Split(sonuc, "-"))
	fmt.Println(strings.Split(sonuc, "*"))

	fmt.Println(strings.Repeat(sonuc, 5))

}
