package error_handling

import (
	"fmt"
	"os"
)

// type assertion
func Demo1() {
	f, err := os.Open("demo1.txt")

	// Dosya bulunursa err  => nil

	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya Bulunamad─▒: ", pErr.Path)
			return
		} else {
			fmt.Println("Dosya Bulunamad─▒.")
			return
		}
	} else {
		fmt.Println(f.Name())
	}
}
