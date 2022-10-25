package string_functions

import (
	"fmt"
	// "strings"

	// alias
	s "strings"
)

// case sensitive (Büyük küçük harf duyarlı)

func Demo1() {
	isim := "Selehattin"
	fmt.Println(s.Count(isim, "e"))
	fmt.Println(s.Count(isim, "E"))

	fmt.Println(s.Contains(isim, "a"))
	fmt.Println(s.Contains(isim, "A"))

	fmt.Println(s.Index(isim, "t"))
	fmt.Println(s.Index(isim, "c"))

	sonuc := s.Index(isim, "a")
	if sonuc != -1 {
		fmt.Println("a harfi var")
	} else {
		fmt.Println("a harfi yok")
	}

}
