package loops

import (
	"fmt"
)

func Demo3() {
	var aklimdakiSayi int = 54
	var tahminEdilenSayi int = 0
	var tahminSayisi int = 0

	fmt.Println("Bir Sayı Tahmin Ediniz: ")
	fmt.Scanln(&tahminEdilenSayi)
	tahminSayisi = tahminSayisi + 1

	for aklimdakiSayi != tahminEdilenSayi {
		if tahminEdilenSayi > aklimdakiSayi {
			fmt.Println("Daha Küçük Bir Sayı Giriniz!")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1

		}
		if tahminEdilenSayi < aklimdakiSayi {
			fmt.Println("Daha Büyük Bir Sayı Giriniz!")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1
		}
	}
	basariDurumu := ""
	if (tahminSayisi > 0) && (tahminSayisi <= 3) {
		basariDurumu = "Süper :-)"
	} else if tahminSayisi <= 6 {
		basariDurumu = "Güzel :)"
	} else {
		basariDurumu = "Kötü :("
	}

	fmt.Printf("Bravo. Doğru Tahmin!.  %v tahmin: %v", tahminSayisi, basariDurumu)

}
