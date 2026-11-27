package main

import "fmt"

func main() {
	// comment nya dihapus aja kali ya??
	var teks, hasil, ulang, buang string
	var pilihan, kunci, i, totalHuruf, shift, k, n, mulai int
	var rasioEnkripsi float64
	var kodeOperasi, kar, hasilKar byte
	var pertahankanSimbol, valid, negatif bool

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
			buang = ""
			for {
				n, _ = fmt.Scanf("%c", &kar)
				if n == 0 {
					fmt.Println("\nInput berakhir. Program dihentikan.")
					return
				}
				if kar == '\n' {
					break
				}
				if kar != '\r' {
					buang = buang + string(kar)
				}
			}
			valid = len(buang) == 1 && buang[0] >= '1' && buang[0] <= '4'
			if !valid {
				fmt.Println("[ini ERROR]\nInput harus berupa angka 1-4!")
			} else {
				pilihan = int(buang[0] - '0')
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

			for {
				fmt.Print("Masukkan teks: ")
				teks = ""
				for {
					n, _ = fmt.Scanf("%c", &kar)
					if n == 0 {
						fmt.Println("\nInput berakhir. Program dihentikan.")
						return
					}
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
				buang = ""
				for {
					n, _ = fmt.Scanf("%c", &kar)
					if n == 0 {
						fmt.Println("\nInput berakhir. Program dihentikan.")
						return
					}
					if kar == '\n' {
						break
					}
					if kar != '\r' {
						buang = buang + string(kar)
					}
				}
				negatif = false
				mulai = 0
				if len(buang) > 0 && (buang[0] == '-' || buang[0] == '+') {
					negatif = buang[0] == '-'
					mulai = 1
				}
				valid = len(buang)-mulai >= 1 && len(buang)-mulai <= 9
				kunci = 0
				for i = mulai; i < len(buang) && valid; i++ {
					if buang[i] >= '0' && buang[i] <= '9' {
						kunci = kunci*10 + int(buang[i]-'0')
					} else {
						valid = false
					}
				}
				if valid {
					if negatif {
						kunci = -kunci
					}
					break
				}
				fmt.Println("[ini ERROR]\nKunci harus berupa angka bulat (maksimal 9 digit)!")
			}

			for {
				fmt.Print("Pertahankan simbol & angka? (true/false): ")
				buang = ""
				for {
					n, _ = fmt.Scanf("%c", &kar)
					if n == 0 {
						fmt.Println("\nInput berakhir. Program dihentikan.")
						return
					}
					if kar == '\n' {
						break
					}
					if kar != '\r' {
						buang = buang + string(kar)
					}
				}
				if buang == "true" {
					pertahankanSimbol = true
					break
				} else if buang == "false" {
					pertahankanSimbol = false
					break
				}
				fmt.Println("[ini ERROR]\nInput harus berupa true atau false!")
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

			for {
				fmt.Print("Masukkan ciphertext: ")
				teks = ""
				for {
					n, _ = fmt.Scanf("%c", &kar)
					if n == 0 {
						fmt.Println("\nInput berakhir. Program dihentikan.")
						return
					}
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

		for {
			fmt.Print("\nKembali ke menu utama? (y/n): ")
			ulang = ""
			for {
				n, _ = fmt.Scanf("%c", &kar)
				if n == 0 {
					fmt.Println("\nInput berakhir. Program dihentikan.")
					return
				}
				if kar == '\n' {
					break
				}
				if kar != '\r' {
					ulang = ulang + string(kar)
				}
			}
			if ulang == "y" || ulang == "Y" || ulang == "n" || ulang == "N" {
				break
			}
			fmt.Println("[ini ERROR]\nInput tidak valid! Masukkan 'y' atau 'n'.")
		}
	}

