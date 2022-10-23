package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"Bilgisayar", 5000, "Asus"})
	fmt.Println(product{"Bilgisayar", 5000, "x"})
	fmt.Println(product{name: "Bilgisayar", unitPrice: 5000})
	fmt.Println(product{name: "Bilgisayar"})

}

type product struct {
	name      string
	unitPrice float32
	brand     string
}
