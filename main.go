package main

import (
	"fmt"
	"os"
)

func main() {

	var pilihan int

	for {
		fmt.Println("1. tipedata")
		fmt.Println("2. loop")
		fmt.Println("3. do while")
		fmt.Println("4. tuser")
		fmt.Println("0. exit")
		fmt.Println("masukan pilihan")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("tipe data")
			tipedata()
		case 2:
			fmt.Println("loop")
			loop()
		case 3:
			fmt.Println("do while")
			dowhile()
		case 4:
			fmt.Println("tuser")
			Tuser()
		case 0:
			fmt.Println("terimakasih")
			os.Exit(0)
		default:
			fmt.Println("tidak valid")
		}
	}
}