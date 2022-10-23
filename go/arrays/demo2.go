package arrays

import "fmt"

func Demo2() {
	var sehirler [5]string
	sehirler[0] = "Ankara"
	sehirler[1] = "İstanbul"
	sehirler[2] = "Bursa"
	sehirler[3] = "Aksaray"
	sehirler[4] = "İzmir"
	fmt.Println(sehirler)

	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])
	}

}
