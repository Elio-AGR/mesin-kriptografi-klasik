package main

import "fmt"

func main() {
	// comment nya dihapus aja kali ya??
	var teks, hasil, ulang, buang string
	var pilihan, kunci, i, totalHuruf, shift, k int
	var rasioEnkripsi float64
	var kodeOperasi, kar, hasilKar byte
	var pertahankanSimbol, valid bool

	ulang = "y"

	for ulang == "y" || ulang == "Y" {
		fmt.Println("\n=============================")
		fmt.Println("   MESIN KRIPTOGRAFI KLASIK  ")
		fmt.Println("=============================")
		fmt.Println("1. Enkripsi Caesar Cipher (ROT-N)")
		fmt.Println("2. Dekripsi Caesar Cipher (ROT-N)")
		fmt.Println("3. Brute-Force Shift 1-25")
		fmt.Println("4. Keluar")

		// validasi opsi
		for {
			fmt.Print("Pilih menu (1-4): ")
			_, err := fmt.Scan(&pilihan)
			if err != nil {
				fmt.Println("[ini ERROR]\nInput harus berupa angka 1-4!")
				fmt.Scan(&buang)
				continue
			}
			valid = (pilihan >= 1 && pilihan <= 4)
			if !valid {
				fmt.Println("[ini ERROR]\nPilihan di luar jangkauan! Masukkan 1-4.")
			} else {
				break
			}
		}

		if pilihan == 4 {
			fmt.Println("Program selesai. Terima kasih!")
			break
		}

		if pilihan == 1 || pilihan == 2 {
			if pilihan == 1 {
				kodeOperasi = 'E'
			} else {
				kodeOperasi = 'D'
			}

			// blok buat handle sisa newline
			for {
				fmt.Scanf("%c", &kar)
				if kar == '\n' {
					break
				}
			}

			for {
				fmt.Print("Masukkan teks: ")
				teks = ""
				for {
					fmt.Scanf("%c", &kar)
					if kar == '\n' {
						break
					}
					if kar != '\r' {
						teks = teks + string(kar)
					}
				}
				if len(teks) == 0 {
					fmt.Println("[ini ERROR]\nTeks tidak boleh kosong!")
				} else {
					break
				}
			}

			// input kunci
			for {
				fmt.Print("Masukkan kunci pergeseran (bilangan bulat, boleh negatif): ")
				_, err := fmt.Scan(&kunci)
				if err != nil {
					fmt.Println("[ini ERROR]\nKunci harus berupa angka bulat!")
					fmt.Scan(&buang)
					continue
				}
				break
			}

			for {
				fmt.Print("Pertahankan simbol & angka? (true/false): ")
				_, err := fmt.Scan(&pertahankanSimbol)
				if err != nil {
					fmt.Println("[ini ERROR]\nInput harus berupa true atau false!")
					fmt.Scan(&buang)
					continue
				}
				break
			}

			// normalisasi shift
			shift = ((kunci % 26) + 26) % 26
			hasil = ""
			totalHuruf = 0

			for i = 0; i < len(teks); i++ {
				kar = teks[i]

				if kar >= 'A' && kar <= 'Z' {
					totalHuruf++
					if kodeOperasi == 'E' {
						hasilKar = 'A' + (kar-'A'+byte(shift))%26
					} else {
						hasilKar = 'A' + (kar-'A'+26-byte(shift))%26
					}
					hasil = hasil + string(hasilKar)
				} else if kar >= 'a' && kar <= 'z' {
					totalHuruf++
					if kodeOperasi == 'E' {
						hasilKar = 'a' + (kar-'a'+byte(shift))%26
					} else {
						hasilKar = 'a' + (kar-'a'+26-byte(shift))%26
					}
					hasil = hasil + string(hasilKar)
				} else if pertahankanSimbol {
					hasil = hasil + string(kar)
				}
			}

			rasioEnkripsi = (float64(totalHuruf) / float64(len(teks))) * 100.0

			fmt.Println("------------------------------------------")
			fmt.Printf("Mode Operasi     : %c (Shift efektif: %d)\n", kodeOperasi, shift)
			fmt.Println("Hasil Teks       :", hasil)
			fmt.Printf("Rasio Alfabet    : %.2f%% (%d dari %d karakter)\n", rasioEnkripsi, totalHuruf, len(teks))
			fmt.Println("------------------------------------------")

		} else if pilihan == 3 {
			// brute force 1-25
			kodeOperasi = 'B'

			// blok buat handle sisa \n
			for {
				fmt.Scanf("%c", &kar)
				if kar == '\n' {
					break
				}
			}

			for {
				fmt.Print("Masukkan ciphertext: ")
				teks = ""
				for {
					fmt.Scanf("%c", &kar)
					if kar == '\n' {
						break
					}
					if kar != '\r' {
						teks = teks + string(kar)
					}
				}
				if len(teks) == 0 {
					fmt.Println("[ini ERROR]\nCiphertext tidak boleh kosong!")
				} else {
					break
				}
			}

			fmt.Println("------------------------------------------")
			fmt.Printf("Mode %c - Hasil Brute-Force (ROT 1 - 25):\n", kodeOperasi)
			for k = 1; k <= 25; k++ {
				hasil = ""
				for i = 0; i < len(teks); i++ {
					kar = teks[i]
					if kar >= 'A' && kar <= 'Z' {
						hasilKar = 'A' + (kar-'A'+26-byte(k))%26
						hasil = hasil + string(hasilKar)
					} else if kar >= 'a' && kar <= 'z' {
						hasilKar = 'a' + (kar-'a'+26-byte(k))%26
						hasil = hasil + string(hasilKar)
					} else {
						hasil = hasil + string(kar)
					}
				}
				fmt.Printf("Shift -%2d : %s\n", k, hasil)
			}
			fmt.Println("------------------------------------------")
		}

		fmt.Print("\nKembali ke menu utama? (y/n): ")
		fmt.Scan(&ulang)
	}
}
