package main

import "fmt"

func tipedata() {

	var nama string = "bhaskoro prayoga"
	var umur int = 22
	var phi float32 = 3.14
	var i int;

	fmt.Println("hallo, nama saya", nama);
	fmt.Println("umur saya", umur);
	fmt.Println("bilangan desimal", phi);

	for i=0; i<=10; i++ {
		fmt.Printf("ini pengulangan ke %d\n", i);
	}
}