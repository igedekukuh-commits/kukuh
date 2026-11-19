package main

import "fmt"

const TARIF_PER_KG = 9500
const BIAYA_ANTAR = 3000

func main() {

	fmt.Println("Nama: I GEDE KUKUH DANAOARTHA | NIM: 42630021")
	fmt.Println("KUKUH LAUNDRY")
	fmt.Println("KHUSUS PELAJAR UNDIKNAS")
	var berat float64
	var uang int

	fmt.Print("Berat cucian (kg): ")
	fmt.Scan(&berat)
	fmt.Print("Uang yang diterima (Rp): ")
	fmt.Scan(&uang)
	biayaCuci := int(berat*TARIF_PER_KG + 0.5)
	total := biayaCuci + BIAYA_ANTAR
	totalBayar := (total + 99) / 100 * 100
	kembalian := uang - totalBayar

	fmt.Printf("Berat cucian  : %.2f kg\n", berat)
	fmt.Printf("Tarif per kg  : Rp%d\n", TARIF_PER_KG)
	fmt.Printf("Biaya cuci    : Rp%d\n", biayaCuci)
	fmt.Printf("Biaya antar   : Rp%d\n", BIAYA_ANTAR)
	fmt.Printf("Total         : Rp%d\n", total)
	fmt.Printf("Total bayar   : Rp%d\n", totalBayar)
	fmt.Printf("Uang diterima : Rp%d\n", uang)
	fmt.Printf("Kembalian     : Rp%d\n", kembalian)
}
