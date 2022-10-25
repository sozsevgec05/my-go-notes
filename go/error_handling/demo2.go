package error_handling

import (
	"errors"
	"fmt"
)

func tahminEt(tahmin int) (string, error) {
	aklimdakiSayi := 50

	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir sayı giriniz!")
	}

	if tahmin > aklimdakiSayi {
		return "Daha Küçük Bir Sayı Giriniz", nil
	}

	if tahmin < aklimdakiSayi {
		return "Daha Büyük Bir Sayı Giriniz", nil
	}

	return "Tebrikler. Tahmininiz Doğru.", nil
}

func Demo2() {
	mesaj, hata := tahminEt(101)
	fmt.Println(mesaj)
	fmt.Println(hata)

	// fmt.Println(tahminEt(101))
	// fmt.Println(tahminEt(-10))

}
