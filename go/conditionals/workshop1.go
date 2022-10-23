package conditionals

import "fmt"

func Demo3() {
	// üç adet int değişken tanımlayız.
	// ekrana en büyük olanı yazdırınız

	var x, y, z int = 132554, 532542, 34645

	var enBuyuk int = x

	if y > enBuyuk {
		enBuyuk = y
	}

	if z > enBuyuk {
		enBuyuk = z
	}

	fmt.Printf("En Büyük Sayı: %v", enBuyuk)
}
