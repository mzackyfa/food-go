package main

import (
	"fmt"
	"strings"
)

func subMenu1(makanan string) {
	fmt.Printf("Daftar Makanan: %s\n", makanan)
}

func subMenu2(minuman string) {
	fmt.Printf("Daftar Minuman: %s\n", minuman)
}

func total(harga string, jumlah int) int {
	switch strings.ToLower(harga) {
	case "airputih":
		return airPutih * jumlah
	case "esteh":
		return esTeh * jumlah
	case "jusmangga":
		return jusMangga * jumlah
	case "jusalpukat":
		return jusAlpukat * jumlah
	case "jussirsak":
		return jusSirsak * jumlah
	case "nasigoreng":
		return nasiGoreng * jumlah
	case "miegoreng":
		return mieGoreng * jumlah
	case "keraktelor":
		return kerakTelor * jumlah
	case "ayamgoreng":
		return ayamGoreng * jumlah
	default:
		fmt.Println("Menu tidak valid")
		return 0
	}
}

var (
	nasiGoreng = 34000
	mieGoreng  = 23000
	kerakTelor = 12000
	mieRebus   = 23000
	ayamGoreng = 35000

	airPutih  = 8000
	esTeh     = 10000
	jusMangga = 24000
	jusAlpukat = 20000
	jusSirsak  = 21000
)

func main() {
	fmt.Println("====Pilih Menu====")
	fmt.Println("1.Makanan")
	fmt.Println("2.Minuman")
	fmt.Println("===================")

	var daftarMenu int
	fmt.Print("Pilih Menu Diatas 1/2:")
	fmt.Scan(&daftarMenu)

	switch daftarMenu {
	case 1:
		subMenu1("Rp. 34,000 nasiGoreng")
		subMenu1("Rp. 23,000 mieGoreng")
		subMenu1("Rp. 12,000 kerakTelor")
		subMenu1("Rp. 23,000 mieRebus")
		subMenu1("Rp. 35,000 ayamGoreng")
	case 2:
		subMenu2("Rp. 8,000 airPutih")
		subMenu2("Rp. 10,000 esTeh")
		subMenu2("Rp. 24,000 jusMangga")
		subMenu2("Rp. 20,000 jusAlpukat")
		subMenu2("Rp. 21,000 jusSirsak")
	default:
		fmt.Println("Pilihan menu tidak valid")
		return
	}

	var harga string
	fmt.Print("Makanan/Minuman yang anda pilih: ")
	fmt.Scan(&harga)

	var jumlah int
	fmt.Print("Masukkan Jumlah: ")
	fmt.Scan(&jumlah)

	Total := total(harga, jumlah)

	if Total > 500000 {
		Total = Total - int(0.25*float64(Total))
		fmt.Printf("Nominal Pembayaran Diskon 25%%: Rp. %d\n", Total)
	} else if Total > 250000 {
		Total = Total - int(0.15*float64(Total))
		fmt.Printf("Nominal Pembayaran Diskon 15%%: Rp. %d\n", Total)
	} else if Total > 100000 {
		Total = Total - int(0.10*float64(Total))
		fmt.Printf("Nominal Pembayaran Diskon 10%%: Rp. %d\n", Total)
	} else {
		fmt.Printf("Nominal Pembayaran: Rp. %d\n", Total)
	}

	var Nim int
	fmt.Print("Masukkan NIM:")
	fmt.Scan(&Nim)

	var Nama string
	fmt.Print("Nama Anda:")
	fmt.Scan(&Nama)

	fmt.Println("====================================")

	var Bayar int
	fmt.Print("Jumlah Nominal Pembayaran:")
	fmt.Scan(&Bayar)

	Kembalian := Bayar - Total

	fmt.Println("====================================")
	fmt.Printf("Nama: %s | Nim: %d\n", Nama, Nim)
	fmt.Printf("Uang Kembalian: Rp. %d\n", Kembalian)
}
