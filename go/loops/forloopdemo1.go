package loops


import "fmt"

var metin string = "Merhaba, Go!"

func Demo1() {
	i := 1
	for i <= 5 {
		fmt.Println(metin)
		i = i + 1

	}
	fmt.Println("Bitti!")

}