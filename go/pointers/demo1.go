package pointers

import "fmt"

// * = address

func Demo1(sayi *int) {
	*sayi = *sayi + 1
	fmt.Println("Demodaki sayı :", sayi)
}
