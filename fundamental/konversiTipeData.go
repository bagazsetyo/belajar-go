package main

import "fmt"

func main() {
	var nilai32 int32 = 10000
	fmt.Println(nilai32)

	var nilai64 int64 = int64(nilai32)
	fmt.Println(nilai64)

	var nilai8 int8 = int8(nilai32)
	fmt.Println(nilai8)

	fmt.Println("=====================================")


	var nama = "Bagas"
	var e byte = nama[0]
	var eString string = string(e)
	fmt.Println(nama)
	fmt.Println(eString)
}