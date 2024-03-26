package main

import "fmt"

func main() {
	// Data untuk Mark
	massaMark := 95.0  // kg
	tinggiMark := 1.88 // meter

	// Data untuk John
	massaJohn := 85.0  // kg
	tinggiJohn := 1.76 // meter

	// Menghitung BMI untuk Mark
	bmiMark := massaMark / (tinggiMark * tinggiMark)

	// Menghitung BMI untuk John
	bmiJohn := massaJohn / (tinggiJohn * tinggiJohn)

	// Menampilkan BMI Mark dan John
	fmt.Printf("BMI Mark: %.2f\n", bmiMark)
	fmt.Printf("BMI John: %.2f\n", bmiJohn)

	// Membuat variabel boolean untuk menentukan apakah Mark memiliki BMI lebih tinggi dari John
	markHigherBMI := bmiMark > bmiJohn

	// Menampilkan informasi apakah Mark memiliki BMI lebih tinggi dari John
	if markHigherBMI == true {
		fmt.Println("Mark memiliki nilai BMI lebih besar dari pada John")
	} else {
		fmt.Println("Mark tidak memiliki nilai BMI lebih besar dari pada John")
	}
}
