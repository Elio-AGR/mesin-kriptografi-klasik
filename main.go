package main

import "fmt"

func main() {
	// ini var (kamus)
	var teks, hasil, ulang string
	var pilihan, kunci, totalHuruf int
	var rasioEnkripsi float64
	var kodeOperasi byte
	var pertahankanSimbol, valid bool

	// ini algoritma
	ulang = "y"

	for ulang == "y" || ulang == "Y" {
		fmt.Println("===================================")
		fmt.Println("   MESIN KRIPTOGRAFI KLASIK   ")
		fmt.Println("===================================")
		fmt.Println("1. Enkripsi Caesar Cipher (ROT-N)")
		fmt.Println("2. Dekripsi Caesar Cipher (ROT-N)")
		fmt.Println("3. Brute-Force Shift (Mode CTF 1-25)")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu (1-4): ")
		fmt.Scan(&pilihan)

		// cek opsi yang dipilih
		valid = (pilihan >= 1 && pilihan <= 4)
		if !valid {
			fmt.Println("[Error] Pilihan tidak valid! Masukkan angka 1-4.")
		} else if pilihan == 4 {
			fmt.Println("Program selesai.")
			break
		} else {
			// gambaran kedepannya nih
			fmt.Println("[Info] Fitur menu", pilihan, "sedang dalam pengembangan...")
			_ = teks
			_ = hasil
			_ = kunci
			_ = totalHuruf
			_ = rasioEnkripsi
			_ = kodeOperasi
			_ = pertahankanSimbol
		}

		fmt.Print("\nKembali ke menu utama? (y/n): ")
		fmt.Scan(&ulang)
	}
}
