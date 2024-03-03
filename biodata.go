package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {

	absen, _ := strconv.ParseInt(os.Args[1], 10, 0)

	person := biodata(absen - 1)

	fmt.Printf("Berikut biodata dari absen %d :\n\n", absen)
	fmt.Println("Nama: ", person.nama)
	fmt.Println("Alamat: ", person.alamat)
	fmt.Println("Pekerjaan: ", person.pekerjaan)
	fmt.Println("Alasan: ", person.alasan)
}

func biodata(absen int64) Person {

	person := []Person{

		{
			nama:      "Muhammad Seman",
			alamat:    "Jalan Merpati Timur Nomor 3",
			pekerjaan: "Mahasiswa",
			alasan:    "Ingin mencoba hal baru",
		},

		{
			nama:      "Bayu Aji",
			alamat:    "Jalan Kertajaya No. 5 A",
			pekerjaan: "Selebgram",
			alasan:    "Ingin menjadi backend developer",
		},
	}

	return person[absen]
}
