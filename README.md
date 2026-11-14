# Mesin Kriptografi Klasik

[![Go](https://img.shields.io/badge/Go-1.24-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
![Dependencies](https://img.shields.io/badge/dependencies-none-brightgreen)

Aplikasi *command line* berbahasa Go untuk **mengenkripsi, mendekripsi, dan memecahkan (brute-force) Caesar Cipher / ROT-N**. Dibuat tanpa library tambahan, hanya `fmt` dari pustaka standar, dengan penanganan input yang ketat sehingga tidak *crash* maupun terjebak *infinite loop*.

## Daftar Isi

- [Fitur](#fitur)
- [Memulai](#memulai)
- [Cara Penggunaan](#cara-penggunaan)
- [Contoh Sesi](#contoh-sesi)
- [Cara Kerja](#cara-kerja)
- [Penanganan Kesalahan Input](#penanganan-kesalahan-input)
- [Struktur Proyek](#struktur-proyek)
- [Batasan yang Diketahui](#batasan-yang-diketahui)
- [Konteks Proyek](#konteks-proyek)
- [Lisensi](#lisensi)

## Fitur

- **Enkripsi Caesar Cipher** dengan kunci bilangan bulat positif, nol, maupun negatif (dinormalisasi dengan modulo 26).
- **Dekripsi Caesar Cipher** dengan kunci yang sama saat enkripsi.
- **Brute-force** seluruh 25 kemungkinan pergeseran untuk pesan yang kuncinya tidak diketahui.
- **Opsi simbol:** pilih untuk mempertahankan atau membuang angka, spasi, dan tanda baca pada hasil.
- **Statistik rasio alfabet:** persentase karakter huruf terhadap total panjang teks.
- **Validasi input menyeluruh:** setiap kesalahan input menampilkan pesan dan meminta pengguna mengulang.
- **Konsisten di Linux, macOS, dan Windows:** input dibaca per baris sehingga akhir baris `\n` maupun `\r\n` diperlakukan sama.
- **Nol dependensi:** hanya `import "fmt"`.

## Memulai

### Prasyarat

- [Go](https://go.dev/dl/) (dikembangkan dan diuji dengan Go 1.24)

### Mengunduh dan menjalankan

```bash
git clone https://github.com/Elio-AGR/mesin-kriptografi-klasik.git
cd mesin-kriptografi-klasik
go run main.go
```

### Membangun *binary*

```bash
# Linux / macOS
go build -o mesin-kriptografi main.go
./mesin-kriptografi

# Windows
go build -o mesin-kriptografi.exe main.go
mesin-kriptografi.exe
```

## Cara Penggunaan

Program menampilkan menu utama. Ketik angka pilihan lalu tekan `Enter`.

| Menu | Fungsi | Input yang diminta |
| :-: | :--- | :--- |
| `1` | Enkripsi Caesar Cipher (ROT-N) | teks, kunci, pertahankan simbol (`true`/`false`) |
| `2` | Dekripsi Caesar Cipher (ROT-N) | teks sandi, kunci, pertahankan simbol (`true`/`false`) |
| `3` | Brute-force shift 1-25 | teks sandi |
| `4` | Keluar | - |

Setelah menu 1, 2, atau 3 selesai, program menanyakan apakah ingin kembali ke menu utama (`y`/`n`).

**Tentang kunci:** boleh bilangan bulat negatif, nol, atau positif dengan maksimal 9 digit, misalnya `3`, `-5`, `+7`, atau `29`. Kunci apa pun dinormalisasi menjadi pergeseran efektif `0-25`, sehingga kunci `29` setara dengan `3`, dan `-3` setara dengan `23`.

## Contoh Sesi

### Enkripsi

```text
=============================
   MESIN KRIPTOGRAFI KLASIK  
=============================
1. Enkripsi Caesar Cipher (ROT-N)
2. Dekripsi Caesar Cipher (ROT-N)
3. Brute-Force Shift 1-25
4. Keluar
Pilih menu (1-4): 1
Masukkan teks: Halo Dunia 2026!
Masukkan kunci pergeseran (bilangan bulat, boleh negatif): 3
Pertahankan simbol & angka? (true/false): true
------------------------------------------
Mode Operasi     : E (Shift efektif: 3)
Hasil Teks       : Kdor Gxqld 2026!
Rasio Alfabet    : 56.25% (9 dari 16 karakter)
------------------------------------------
```

### Brute-force

Pesan `Khoor Zruog` dengan kunci yang tidak diketahui:

```text
Pilih menu (1-4): 3
Masukkan ciphertext: Khoor Zruog
------------------------------------------
Mode B - Hasil Brute-Force (ROT 1 - 25):
Shift - 1 : Jgnnq Yqtnf
Shift - 2 : Ifmmp Xpsme
Shift - 3 : Hello World
Shift - 4 : Gdkkn Vnqkc
...
Shift -25 : Lipps Asvph
------------------------------------------
```

Hasil yang terbaca (`Hello World`) muncul pada `Shift - 3`, artinya kunci aslinya adalah `3`.

### Penanganan input yang salah

```text
Pilih menu (1-4): 9
[ini ERROR]
Input harus berupa angka 1-4!
Pilih menu (1-4): 1
Masukkan teks: 
[ini ERROR]
Teks tidak boleh kosong!
Masukkan teks: Rahasia
Masukkan kunci pergeseran (bilangan bulat, boleh negatif): 3abc
[ini ERROR]
Kunci harus berupa angka bulat (maksimal 9 digit)!
Masukkan kunci pergeseran (bilangan bulat, boleh negatif): -5
Pertahankan simbol & angka? (true/false): maybe
[ini ERROR]
Input harus berupa true atau false!
Pertahankan simbol & angka? (true/false): false
------------------------------------------
Mode Operasi     : E (Shift efektif: 21)
Hasil Teks       : Mvcvndv
Rasio Alfabet    : 100.00% (7 dari 7 karakter)
------------------------------------------
```

## Cara Kerja

Hanya huruf `A-Z` dan `a-z` yang digeser. Huruf kapital dan huruf kecil diproses terpisah sehingga kapitalisasi tetap terjaga.

| Langkah | Rumus |
| :--- | :--- |
| Normalisasi kunci | `shift = ((kunci mod 26) + 26) mod 26` |
| Enkripsi | `hasil = 'A' + (huruf - 'A' + shift) mod 26` |
| Dekripsi | `hasil = 'A' + (huruf - 'A' + 26 - shift) mod 26` |
| Brute-force | dekripsi dengan `k = 1..25` |
| Rasio alfabet | `(jumlah huruf / panjang teks) * 100` |

Rumus enkripsi dan dekripsi di atas ditulis untuk huruf kapital. Huruf kecil memakai pola yang sama dengan basis `'a'`.

Karakter non-huruf (angka, spasi, tanda baca) tidak pernah digeser. Pada menu 1 dan 2 karakter tersebut dipertahankan atau dibuang sesuai opsi `true`/`false`; pada brute-force selalu dipertahankan agar hasil mudah dibaca.

## Penanganan Kesalahan Input

Setiap input divalidasi, dan jika salah, program menampilkan pesan lalu meminta input ulang.

| Prompt | Input yang ditolak | Contoh |
| :--- | :--- | :--- |
| Menu | selain tepat satu angka `1-4` | `x`, `9`, `0`, `2abc`, `1.5`, baris kosong |
| Teks / ciphertext | kosong | langsung `Enter` |
| Kunci | bukan bilangan bulat, atau lebih dari 9 digit | `abc`, `3abc`, `1.5`, `99999999999999999999` |
| Pertahankan simbol | selain tepat `true` atau `false` | `maybe`, `ya`, `True` |
| Kembali ke menu | selain `y`, `Y`, `n`, `N` | `yes`, `z` |
| Semua prompt | input terputus (EOF, misalnya `Ctrl+D`) | program berhenti dengan pesan "Input berakhir" |

## Struktur Proyek

```text
mesin-kriptografi-klasik/
├── main.go      # seluruh program
├── LICENSE      # lisensi MIT
└── README.md
```

## Batasan yang Diketahui

- Hanya huruf ASCII (`A-Z`, `a-z`) yang diproses sebagai alfabet. Karakter non-ASCII seperti `é` atau `€` tidak didukung, dan hasilnya dapat tidak sesuai bila opsi pertahankan simbol bernilai `true`.
- Caesar Cipher adalah sandi klasik yang **tidak aman** untuk melindungi data nyata. Brute-force pada program ini membuktikannya: seluruh kunci dapat dicoba dalam sekejap. Proyek ini ditujukan untuk pembelajaran.

## Konteks Proyek

Proyek ini dibuat sebagai tugas proyek mata kuliah **Berpikir Komputasional dan Pengenalan Pemrograman (BKPP)**. Ketentuan teknis yang dipenuhi:

- Memakai tipe data dasar `int`, `float64`, `bool`, `byte`, dan `string`, serta percabangan dan perulangan.
- Tidak memakai library tambahan; hanya `import "fmt"`.
- Memiliki *error handling* untuk kesalahan input tanpa *crash*, dan pengguna dapat mengisi ulang nilai yang sesuai.

## Lisensi

Dirilis di bawah [Lisensi MIT](LICENSE).
