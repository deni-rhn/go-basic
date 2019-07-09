package main

import "fmt"

func main() {
	// seleksi
	var nilai = 9

	if nilai == 10 {
		fmt.Println("Lulus dengan sempurna")
	} else if nilai == 9 {
		fmt.Println("Lulus memuaskan")
	} else if nilai > 7 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("anda tidak lulus, nilai anda %d\n", nilai)
	} 

	// variable temporary : variabel yang hanya bisa digunakan dalam skop suatu kondisi
	var angka = 8840.0

	if percent := angka / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// switch case
	var val1 = 6

	switch val1 {
		case 9:
			fmt.Println("kondisi 1")
		case 5:
			fmt.Println("kondisi 2")
		default:
			fmt.Println("default kondisi")
	}

	// switch case banyak kondisi
	var val2 = 7

	switch val2 {
		case 9:
			fmt.Println("kondisi 1")
		case 6,7,8:
			fmt.Println("Kondisi 2")
		default:
			fmt.Println("default")	
	}

	var val3 = 8

	switch val3 {
		case 9:
			fmt.Println("kondisi 1")
		case 6,7:
			fmt.Println("Kondisi 2")
		default:
			{
				fmt.Println("default 1")
				fmt.Println("default 2")
			}
	}

	var val4 = 4

	switch {
		case val4 == 9:
			fmt.Println("kondisi 1")
		case (val4 > 3) && (val4 <8):
			fmt.Println("kondisi 2 di case gaya if - else")
		default:
			{
				fmt.Println("default 1")
				fmt.Println("default 2")
			}	
	}

	/*
		fallthrough dalam case digunakan untuk mengecek kondisi setelahnya, 
		dimana tanpa fallthrough di Golang pengecekan kondisi akan berhenti setelah terpenuhi
	*/
	// case ini akan ada 2 nilai yang terpenuhi jika pakai fallthrough
	var val5 = 12

	switch {
		case val5 == 10:
			fmt.Println("default 1")
		case (val5 > 10) && (val5 < 13):
			fmt.Println("awesome")
			fallthrough
		case (val5 < 20) && (val5 >= 12):
			fmt.Println("great")
		default:
			fmt.Println("default")
	}

	var val6 = 20

	if val6 > 17 {
		switch val6 {
			case 18:
				fmt.Println("nice")
			case 20:
				fmt.Println("perfectly")
			default:
				fmt.Println("default")	
		}
	} else {
		if val6 == 7 {
			fmt.Println("angka 7")
		} else if (val6 > 10) && (val6 < 15) {
			fmt.Println("angka di range 11 - 14")
		} else {
			if val6 == 16 {
				fmt.Println("angka 16")
			} else if val6 < 5 {
				fmt.Println("angka < 5")
			}
		}
	}
}