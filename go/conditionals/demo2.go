package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 1000

	if cekilmekIstenen > hesap {
		fmt.Println("Bakiye Yetersiz!")
	} else if cekilmekIstenen == hesap {
		fmt.Println("Paranız Hazırlanıyor!!!")
		fmt.Println("Dikkat! Hesabınızda Paranız Kalmadı!")

	} else {
		fmt.Println("Paranız Hazırlanıyor!!!")
	}

}
