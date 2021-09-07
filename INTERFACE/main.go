package main

import "fmt"

type BangunDatar interface {
	hitungLuas() int
}

type Persegi struct{
	Sisi int
}

func (persegi Persegi) hitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar int
}

func (persegiPanjang PersegiPanjang) hitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func main(){
	persegi := Persegi{5}
	fmt.Println("Luas Persegi : ", persegi.hitungLuas())
	
	persegiPanjang := PersegiPanjang{4, 5}
	fmt.Println("Luas Persegi Panjang : ", persegiPanjang.hitungLuas())
}