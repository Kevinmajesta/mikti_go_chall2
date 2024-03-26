package main

import "fmt"

func main() {
	//Buat variable dengan menggunakan array
	hasilLumba := [3]int{96, 108, 89}
	hasilKoala := [3]int{88, 91, 110}
	//buat variable yang berisi func rata" yang sudah dibuat
	hasilAkhirLumba := hitungRataRata(hasilLumba)
	hasilAkhirKoala := hitungRataRata(hasilKoala)

	//print ke console untuk liat nilai rata"
	fmt.Printf("Hasil Rata-Rata Tim Lumba-Lumba Adalah: %d\n", hasilAkhirLumba)
	fmt.Printf("Hasil Rata-Rata Tim Koala Adalah: %d\n", hasilAkhirKoala)

	//menggunakan if untuk melihat kondisi
	if hasilAkhirLumba > hasilAkhirKoala {
		fmt.Println("Tim Lumba-Lumba Berhasil Menang")
	} else if hasilAkhirLumba < hasilAkhirKoala {
		fmt.Println("Tim Koala Berhasil Menang")
	} else {
		fmt.Println("Hasil Seri")
	}
}

//func untuk menghitung rata"
func hitungRataRata(scores [3]int) int {
	total := 0
	for _, score := range scores {
		total += score
	}
	return total / len(scores)
}
