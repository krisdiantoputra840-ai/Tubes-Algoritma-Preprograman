package main

import "fmt"

// =====================
// SELECTION SORT
// BERDASARKAN TOTAL ABSENSI
// (DESCENDING)
// =====================

func selectionSortTotalAbsensi() {

	for i := 0; i < jumlah-1; i++ {

		max := i

		totalMax :=
			data[max].Hadir +
				data[max].Izin +
				data[max].Sakit +
				data[max].Alpha

		for j := i + 1; j < jumlah; j++ {

			totalJ :=
				data[j].Hadir +
					data[j].Izin +
					data[j].Sakit +
					data[j].Alpha

			if totalJ > totalMax {

				max = j
				totalMax = totalJ
			}
		}

		data[i], data[max] = data[max], data[i]
	}

	fmt.Println("Data berhasil diurutkan berdasarkan Total Absensi.")
}

// =====================
// INSERTION SORT
// BERDASARKAN NAMA
// (A - Z)
// =====================

func insertionSortNama() {

	for i := 1; i < jumlah; i++ {

		temp := data[i]

		j := i - 1

		for j >= 0 &&
			data[j].Nama > temp.Nama {

			data[j+1] = data[j]

			j--
		}

		data[j+1] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan Nama.")
}

// =====================
// SORT ALPHA TERBANYAK
// UNTUK STATISTIK
// =====================

func sortAlphaTerbanyak() {

	for i := 0; i < jumlah-1; i++ {

		max := i

		for j := i + 1; j < jumlah; j++ {

			if data[j].Alpha >
				data[max].Alpha {

				max = j
			}
		}

		data[i], data[max] = data[max], data[i]
	}
}