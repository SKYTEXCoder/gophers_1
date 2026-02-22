package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processPayment(itemPrice float64, money float64) (bool, float64) {
	if money >= itemPrice {
		change := money - itemPrice
		return true, change
	} else {
		shortfall := itemPrice - money
		return false, shortfall
	}
}

func readPositiveFloat(scanner *bufio.Scanner, prompt string) float64 {
	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			fmt.Println("[ERROR]: Input tidak boleh kosong. Silakan coba lagi.")
			continue
		}

		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Printf("[ERROR]: \"%s\" bukan angka yang valid. Silakan coba lagi.\n", input)
			continue
		}

		if value <= 0 {
			fmt.Println("[ERROR]: Nilai harus lebih dari 0. Silakan coba lagi.")
			continue
		}

		return value
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("==========================================")
		fmt.Println("  PROGRAM PENCATAT TRANSAKSI PEMBAYARAN  ")
		fmt.Println("==========================================")

		itemPrice := readPositiveFloat(scanner, "Masukkan harga barang yang ingin dibeli oleh pembeli ke dalam program pencatat transaksi ini: ")
		money := readPositiveFloat(scanner, "Masukkan uang pembeli yang membeli barang tersebut ke dalam program pencatat transaksi ini: ")

		fmt.Println()
		fmt.Printf("Harga Barang:  Rp %.2f\n", itemPrice)
		fmt.Printf("Uang Pembeli:  Rp %.2f\n", money)
		fmt.Println()

		successStatus, amount := processPayment(itemPrice, money)
		if successStatus {
			if amount == 0 {
				fmt.Println("[SISTEM]: Transaksi Berhasil. Tidak ada kembalian yang harus dikembalikan ke pembeli.")
			} else {
				fmt.Printf("[SISTEM]: Transaksi Berhasil. Kembalian anda adalah sebesar: Rp %.2f\n", amount)
			}
		} else {
			fmt.Printf("[SISTEM]: Transaksi anda ditolak. Uang anda kurang sebesar: Rp %.2f\n", amount)
		}

		fmt.Println()
		fmt.Print("Apakah anda ingin melakukan transaksi lagi? (y/n): ")
		scanner.Scan()
		answer := strings.TrimSpace(strings.ToLower(scanner.Text()))

		if answer != "y" {
			fmt.Println("\nTerima kasih telah menggunakan program pencatat transaksi ini. Sampai jumpa! 👋")
			return
		}

		fmt.Println()
	}
}
