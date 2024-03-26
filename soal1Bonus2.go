package main

import "fmt"

func main() {
	//Buat variable dengan menggunakan array
	hasilLumba := [3]int{97, 112, 101}
	hasilKoala := [3]int{109, 95, 106}

	//buat variable yang berisi func rata" yang sudah dibuat
	hasilAkhirLumba := hitungRataBonus2(hasilLumba)
	hasilAkhirKoala := hitungRataBonus2(hasilKoala)

	//print ke console untuk liat nilai rata"
	fmt.Printf("Hasil Rata - Rata tim Lumba - Lumba Adalah %d\n", hasilAkhirLumba)
	fmt.Printf("Hasil Rata - Rata tim Koala Adalah %d\n", hasilAkhirKoala)

	//menggunakan if untuk melihat kondisi
	if hasilAkhirLumba > hasilAkhirKoala && hasilAkhirLumba >= 100 {
		fmt.Println("Tim Lumba Lumba Berhasil Menang")
	} else if hasilAkhirKoala > hasilAkhirLumba && hasilAkhirKoala >= 100 {
		fmt.Println("Tim Koala Berhasil Menang")
	} else if hasilAkhirLumba == hasilAkhirKoala && hasilAkhirLumba >= 100 && hasilAkhirKoala >= 100 {
		fmt.Println("Hasil seri")
	} else {
		fmt.Println("Tidak ada pemenang, skor di bawah 100.")
	}
}

//func untuk menghitung rata"
func hitungRataBonus2(scores [3]int) int {
	total := 0
	for _, score := range scores {
		total += score
	}
	return total / len(scores)
}
