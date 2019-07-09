package main

import "fmt"

func main(){
	var firstName string = "rafky"
	var lastName string
	lastName = "aditiya"
	tanggal := 2
	bulan, tahun := "Desember", 2010
	var ex bool = true
	var decimal = 2.12
	var paragraph = `Ini adalah variable untuk "Paragraph"
	di "Golang".`
	var value = ((( 2+6 ) % 3) * 4 - 2) / 3
	var isEqual = (value == 2)

	// dengan Printf maka %s akan di ganti oleh variable
	fmt.Println("Hello World");
	fmt.Printf("Hello %s %s !\n", firstName, lastName)
	fmt.Printf("Lahir %d %s %d !\n", tanggal, bulan, tahun)
	fmt.Printf("Boolean %t \n", ex)
	fmt.Printf("Decimal %f \n", decimal)
	fmt.Printf("Decimal %.2f \n", decimal)
	fmt.Println(paragraph)
	fmt.Printf("hasil %d (%t) \n", value, isEqual)
}
