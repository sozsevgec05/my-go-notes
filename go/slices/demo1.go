package slices

import (
	"fmt"
)

func Demo1() {
	isimler := make([]string, 3)
	fmt.Println(isimler)
	isimler[0] = "Selahattin"
	isimler[1] = "Halil"
	isimler[2] = "İbrahim"
	isimler = append(isimler, "Büşra")

	fmt.Println(isimler)
	fmt.Println(len(isimler))

}
