package main

import "fmt"

func main(){
	// perulangan
	for i := 0; i <= 5; i++ {
		fmt.Println("Angka", i)
	}

	var a = 0
	for a <= 5 {
		fmt.Println("Ini angka", a)
		a++
	}

	var b = 0
	for {
		fmt.Println("Value", b)
		b++

		if b == 5 {
			break
		}
	}

	for d := 0; d <= 10; d++ {
		// jika d adalah ganjil (d % 2) maka akan dilanjutkan perulangan
		if d % 2 == 1 {
			continue
		}

		// jika d lebih besar dari 8 maka akan dipaksa berhenti perulangannya
		if d > 8 {
			break
		}

		fmt.Println("Angkaa", d)

	}

	for s := 0; s < 5; s++ {
		for k := s; k < 5; k++ {
			fmt.Print(k, " ")
		}

		fmt.Println()
	}

	// label (outerLoop) dalam perulangan
	outerLoop:
	for i := 0; i < 5; i++ {
		for j :=0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matrik [", i ,"][", j ,"]", "\n")	
		}
	}
}