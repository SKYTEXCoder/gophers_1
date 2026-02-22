# 🧾 Program Pencatat Transaksi Pembayaran

Program CLI interaktif untuk mencatat transaksi pembayaran yang dibangun menggunakan Go. Program ini mensimulasikan sistem kasir sederhana di mana pengguna dapat memasukkan harga barang dan uang pembeli, lalu menerima hasil transaksi secara instan — termasuk kembalian atau kekurangan dalam mata uang IDR (Rupiah Indonesia).

Dibuat sebagai **Task 1 (Dasar-Dasar)** untuk mempelajari fundamental Go: I/O, kontrol alur, fungsi dengan multiple return values, validasi input, dan pengujian otomatis.

## ✨ Fitur

- **CLI Interaktif** — Meminta input harga barang dan jumlah pembayaran melalui `bufio.Scanner`
- **Pemrosesan pembayaran** — Menghitung kembalian atau kekurangan dengan fungsi yang mengembalikan dua nilai
- **Validasi input** — Menolak input kosong, bukan angka, nol, dan negatif dengan pesan error yang jelas serta otomatis meminta ulang
- **Perulangan transaksi** — Setelah setiap transaksi, program menanyakan apakah pengguna ingin melakukan transaksi lagi tanpa harus menjalankan ulang program
- **Format mata uang IDR** — Semua nilai uang ditampilkan dengan awalan `Rp` dan 2 angka desimal
- **Pengujian otomatis** — 9 unit test berbasis tabel yang mencakup semua cabang logika bisnis dan kasus tepi

## 📁 Struktur Proyek

```
task_1_basics/
├── cmd/                          # Dicadangkan untuk entrypoint CLI di masa depan
├── internal/
│   └── greeting/
│       ├── greeting.go           # Paket utilitas greeting
│       └── greeting_test.go      # Unit test greeting
├── main.go                       # Entry point aplikasi & logika utama
├── main_test.go                  # Pengujian otomatis untuk processPayment & readPositiveFloat
├── go.mod                        # Definisi modul Go (Go 1.25.6)
└── README.md
```

## 🚀 Cara Memulai

### Prasyarat

- [Go](https://go.dev/dl/) versi 1.25+ sudah terinstal

### Menjalankan aplikasi

```bash
go run .
```

### Membuat binary mandiri

```bash
go build -o pencatat_transaksi.exe .
.\pencatat_transaksi.exe
```

## 💡 Contoh Penggunaan

```
==========================================
  PROGRAM PENCATAT TRANSAKSI PEMBAYARAN
==========================================
Masukkan harga barang yang ingin dibeli oleh pembeli ke dalam program pencatat transaksi ini: 35000
Masukkan uang pembeli yang membeli barang tersebut ke dalam program pencatat transaksi ini: 45000

Harga Barang:  Rp 35000.00
Uang Pembeli:  Rp 45000.00

[SISTEM]: Transaksi Berhasil. Kembalian anda adalah sebesar: Rp 10000.00

Apakah anda ingin melakukan transaksi lagi? (y/n): n

Terima kasih telah menggunakan program pencatat transaksi ini. Sampai jumpa! 👋
```

### Validasi Input

```
Masukkan harga barang ...: abc
[ERROR]: "abc" bukan angka yang valid. Silakan coba lagi.
Masukkan harga barang ...: -5000
[ERROR]: Nilai harus lebih dari 0. Silakan coba lagi.
Masukkan harga barang ...:
[ERROR]: Input tidak boleh kosong. Silakan coba lagi.
Masukkan harga barang ...: 25000    ← diterima ✅
```

## 🏗️ Fungsi Utama

| Fungsi              | Signature                                      | Deskripsi                                                                                                                    |
| ------------------- | ---------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `processPayment`    | `(itemPrice, money float64) → (bool, float64)` | Mengembalikan `true` + kembalian jika pembayaran cukup, `false` + kekurangan jika tidak                                      |
| `readPositiveFloat` | `(scanner, prompt) → float64`                  | Terus meminta input sampai pengguna memasukkan angka positif yang valid. Menolak input kosong, bukan angka, nol, dan negatif |

## 🧪 Pengujian

### Menjalankan semua test

```bash
go test -v ./...
```

### Cakupan pengujian

| Test Suite              | Jumlah | Yang Diuji                                                                     |
| ----------------------- | ------ | ------------------------------------------------------------------------------ |
| `TestProcessPayment`    | 5      | Pembayaran pas, lebih bayar, kurang bayar, lebih bayar besar, kekurangan kecil |
| `TestReadPositiveFloat` | 4      | Input valid, tidak valid→valid, kosong→negatif→valid, nol→valid                |
| `TestHello`             | 1      | Paket utilitas greeting                                                        |

### Analisis statis

```bash
go vet ./...
```

## 🧠 Konsep Go yang Didemonstrasikan

- **Multiple return values** — `processPayment` mengembalikan `(bool, float64)`
- **`bufio.Scanner`** — Membaca input stdin secara interaktif baris per baris
- **`strconv.ParseFloat`** — Konversi string ke angka dengan penanganan error
- **Loop validasi input** — Perulangan `for` dengan `continue` untuk logika coba ulang
- **Table-driven tests** — Pola pengujian Go idiomatik dengan subtest `t.Run`
- **Simulasi stdin dalam test** — Menggunakan `strings.NewReader` untuk mengirim input palsu
- **Internal packages** — `internal/greeting` untuk kode utilitas yang terenkapsulasi
