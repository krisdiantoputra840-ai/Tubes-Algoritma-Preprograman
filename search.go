package main

import (
	"fmt"
	"strings"
)

// =====================
// SEQUENTIAL SEARCH
// =====================

func sequentialSearch(nim string) int {

	for i := 0; i < jumlah; i++ {

		if data[i].NIM == nim {
			return i
		}
	}

	return -1
}

// =====================
// SORT NIM
// DIGUNAKAN SEBELUM
// BINARY SEARCH
// =====================

func sortNIM() {

	for i := 1; i < jumlah; i++ {

		temp := data[i]

		j := i - 1

		for j >= 0 && data[j].NIM > temp.NIM {

			data[j+1] = data[j]

			j--
		}

		data[j+1] = temp
	}
}

// =====================
// BINARY SEARCH
// =====================

func binarySearch(nim string) int {

	left := 0
	right := jumlah - 1

	for left <= right {

		mid := (left + right) / 2

		if data[mid].NIM == nim {
			return mid
		}

		if data[mid].NIM < nim {

			left = mid + 1

		} else {

			right = mid - 1
		}
	}

	return -1
}

// =====================
// CARI BERDASARKAN STATUS
// =====================

func cariBerdasarkanStatus() {

	if jumlah == 0 {

		fmt.Println("Belum ada data mahasiswa")

		return
	}

	fmt.Println()
	fmt.Println("H = Hadir")
	fmt.Println("I = Izin")
	fmt.Println("S = Sakit")
	fmt.Println("A = Alpha")

	status := strings.ToUpper(
		inputString("Masukkan Status : "),
	)

	fmt.Println()
	fmt.Println("HASIL PENCARIAN")
	fmt.Println("====================================================")

	ditemukan := false

	for i := 0; i < jumlah; i++ {

		switch status {

		case "H":

			if data[i].Hadir > 0 {

				fmt.Printf(
					"%s - %s - Hadir: %d\n",
					data[i].NIM,
					data[i].Nama,
					data[i].Hadir,
				)

				ditemukan = true
			}

		case "I":

			if data[i].Izin > 0 {

				fmt.Printf(
					"%s - %s - Izin: %d\n",
					data[i].NIM,
					data[i].Nama,
					data[i].Izin,
				)

				ditemukan = true
			}

		case "S":

			if data[i].Sakit > 0 {

				fmt.Printf(
					"%s - %s - Sakit: %d\n",
					data[i].NIM,
					data[i].Nama,
					data[i].Sakit,
				)

				ditemukan = true
			}

		case "A":

			if data[i].Alpha > 0 {

				fmt.Printf(
					"%s - %s - Alpha: %d\n",
					data[i].NIM,
					data[i].Nama,
					data[i].Alpha,
				)

				ditemukan = true
			}

		default:

			fmt.Println("Status tidak valid")
			return
		}
	}

	if !ditemukan {

		fmt.Println("Tidak ada data yang sesuai")
	}
}