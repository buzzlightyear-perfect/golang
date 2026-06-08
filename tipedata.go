package main

import "fmt"

func tipedata() {

	var nama string = "bhaskoro prayoga"
	var umur int = 22
	var phi float32 = 3.14

	fmt.Println("hallo, nama saya", nama)
	fmt.Println("umur saya", umur)
	fmt.Println("bilangan desimal", phi)

}

func loop(){

	var i int

	for i=1; i<=5; i++{
		fmt.Printf("ini pengulangan ke %d\n", i)
	}
}

func dowhile(){

	i:= 1

	for {
		fmt.Println("do while")
		i++
		if i > 5 {
			break
		}
	}
}

func Tuser() {

	var pilih string

	for {
		fmt.Println("program berjalan")

		fmt.Println("ulangi? (y/t)")
		fmt.Scan(&pilih)

		if pilih != "y"{
			break
		}
	}
}