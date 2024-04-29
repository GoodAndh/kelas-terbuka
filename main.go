package main

import "fmt"

type mahasiswa struct {
	nama    string
	jurusan string
	usia    int
}

func main() {
	fmt.Println("TASK 1 ")
	// 1.
	jumlah := sum(1, 2)
	fmt.Println("jumlah dari angka :", jumlah)

	fmt.Println()
	fmt.Println("TASK 2 ")
	// 2.
	for i := 0; i <= +10; i++ {
		fmt.Println("angka ke :", i)
	}

	fmt.Println()
	fmt.Println("TASK 3")
	// 3.
	mahasiswa := mahasiswa{
		nama:    "andy",
		jurusan: "elektro",
		usia:    19,
	}
	fmt.Println("nama :", mahasiswa.nama)
	fmt.Println("jurusan :", mahasiswa.jurusan)
	fmt.Println("usia :", mahasiswa.usia)

	fmt.Println()
	fmt.Println("TASK 4")
	// 4.
	mahasiswaInSlice := []string{mahasiswa.nama, mahasiswa.jurusan}
	fmt.Println("mahasiswa in slice :", mahasiswaInSlice)
	fmt.Println("panjang mahasiswa in slice :", len(mahasiswaInSlice))

	fmt.Println()
	fmt.Println("TASK 5")
	// 5.

	/*
		di golang ada 2 jenis yaitu pass by reference dan pass by value ,di bawah ini mencoba pass by value atau disebut copy an dari defaultname jadi nilai asli dari variabel defaultName tdk pernah berubah
	*/
	var defaultName string = "andy"
	var newName string = defaultName
	fmt.Println("newName :", newName)

	//mengubah nilai defaultName
	defaultName = "new Andy"
	fmt.Println("copyDefaultName :", defaultName)

	/*
		pass by references
	*/
	newDefaultname := &defaultName
	fmt.Println("newDefaultName :", *newDefaultname)
	*newDefaultname=newName
	fmt.Println("newDefaultName :", *newDefaultname)

}

func sum(a, b int) int {
	return a + b
}
