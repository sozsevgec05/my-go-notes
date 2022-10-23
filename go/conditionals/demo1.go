package conditionals

import (
	"fmt"
)

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("Bakiye Yetersiz!")
	}

	if cekilmekIstenen <= hesap {
		fmt.Println("Paranız Hazırlanıyor!")
		hesap = hesap - cekilmekIstenen
	}
	fmt.Println("Hesapta Kalan Bakiye:  " + fmt.Sprintf("%v", hesap))

	fmt.Println("İf Bitti!")

}
