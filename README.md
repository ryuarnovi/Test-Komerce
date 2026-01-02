# Test-Komerce

Sebuah aplikasi Go yang berisi dua program untuk menyelesaikan masalah algoritma:
1. **Character Separator** - Memisahkan karakter vokal dan konsonan dari input teks
2. **PSBB Family Bus Calculator** - Menghitung jumlah minimum bus yang diperlukan untuk mengangkut keluarga dengan aturan PSBB

## Fitur

### 1. Character Separator
Program ini memisahkan karakter dalam sebuah kalimat menjadi dua kategori:
- Huruf vokal (a, i, u, e, o)
- Huruf konsonan (semua huruf selain vokal)

Program akan memproses dua input dan menampilkan hasil pemisahan untuk masing-masing input.

### 2. PSBB Family Bus Calculator
Program ini menghitung jumlah minimum bus yang diperlukan untuk mengangkut keluarga-keluarga dengan batasan:
- Setiap bus maksimal dapat menampung 4 orang
- Keluarga dengan anggota sedikit dapat digabung dengan keluarga lain dalam satu bus
- Program memproses 3 test case

## Prerequisites

- Go 1.25.4 atau versi yang lebih baru
- Terminal/Command Prompt

## Instalasi

1. Clone repository ini:
```bash
git clone https://github.com/ryuarnovi/Test-Komerce.git
cd Test-Komerce
```

2. Download dependencies (jika diperlukan):
```bash
go mod download
```

## Cara Menjalankan

### Menjalankan Program

Jalankan program utama dengan perintah:
```bash
go run short_char.go
```

### Contoh Penggunaan

#### Character Separator
```
Input one line of words (S) : Sample Case
Vowel Characters : aaee
Consonant Characters : smplcs

Input one line of words (S) : Next Case
Vowel Characters : eae
Consonant Characters : nxtcs
```

#### PSBB Family Bus Calculator
```
Input #1
Input the number of families : 5
Input the number of members in the family (separated by a space): 1 2 4 3 3
Minimum bus required for input #1 is : 3

Input #2
Input the number of families : 5
Input the number of members in the family (separated by a space): 2 3 4 4 2
Minimum bus required for input #2 is : 4

Input #3
Input the number of families : 8
Input the number of members in the family (separated by a space): 2 2 2 2 2 2 2 2
Minimum bus required for input #3 is : 4
```

## Struktur Proyek

```
Test-Komerce/
├── go.mod                 # File konfigurasi Go module
├── short_char.go          # Program utama
├── psbbfamily/           
│   └── psbbfamily.go      # Package untuk kalkulasi bus PSBB
└── README.md              # Dokumentasi proyek
```

## Penjelasan Algoritma

### Character Separator
- Membaca input dari pengguna
- Mengkonversi semua huruf menjadi lowercase
- Memisahkan karakter berdasarkan apakah termasuk vokal atau konsonan
- Mengabaikan karakter non-alfabetis

### PSBB Family Bus Calculator
- Menggunakan algoritma greedy dengan two-pointer
- Mengurutkan jumlah anggota keluarga dari kecil ke besar
- Menggabungkan keluarga terkecil dengan keluarga terbesar jika memungkinkan
- Mengoptimalkan penggunaan bus dengan strategi pairing

## Build

Untuk membuat executable file:
```bash
go build -o test-komerce short_char.go
```

Kemudian jalankan:
```bash
./test-komerce
```

## Lisensi

Project ini dibuat untuk keperluan test dan pembelajaran.

## Author

[ryuarnovi](https://github.com/ryuarnovi)
