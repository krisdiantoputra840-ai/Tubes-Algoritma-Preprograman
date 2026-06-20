package main

import "fmt"

func main() {

	loadData()

	for {

		fmt.Println()
		fmt.Println("======================================")
		fmt.Println("           SIPRESENSI CLI")
		fmt.Println("======================================")
		fmt.Println("Total Mahasiswa :", jumlah)
		fmt.Println()
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Ubah Mahasiswa")
		fmt.Println("3. Hapus Mahasiswa")
		fmt.Println("4. Input Absensi")
		fmt.Println("5. Cari NIM (Sequential Search)")
		fmt.Println("6. Cari NIM (Binary Search)")
		fmt.Println("7. Cari Berdasarkan Status")
		fmt.Println("8. Urutkan Total Absensi")
		fmt.Println("9. Urutkan Nama")
		fmt.Println("10. Statistik Kehadiran")
		fmt.Println("11. Alpha Terbanyak")
		fmt.Println("12. Tampilkan Data")
		fmt.Println("0. Keluar")

		pilih := inputInt("Pilih Menu : ")

		switch pilih {

		case 1:
			tambahMahasiswa()

		case 2:
			ubahMahasiswa()

		case 3:
			hapusMahasiswa()

		case 4:
			inputAbsensi()

		case 5:

			nim := inputString("Masukkan NIM : ")

			idx := sequentialSearch(nim)

			if idx == -1 {

				fmt.Println("Data tidak ditemukan")

			} else {

				fmt.Println("\nData Ditemukan")
				fmt.Println("NIM   :", data[idx].NIM)
				fmt.Println("Nama  :", data[idx].Nama)
				fmt.Println("Kelas :", data[idx].Kelas)
			}

		case 6:

			sortNIM()

			nim := inputString("Masukkan NIM : ")

			idx := binarySearch(nim)

			if idx == -1 {

				fmt.Println("Data tidak ditemukan")

			} else {

				fmt.Println("\nData Ditemukan")
				fmt.Println("NIM   :", data[idx].NIM)
				fmt.Println("Nama  :", data[idx].Nama)
				fmt.Println("Kelas :", data[idx].Kelas)
			}

		case 7:
			cariBerdasarkanStatus()

		case 8:
			selectionSortTotalAbsensi()
			tampilkanData()

		case 9:
			insertionSortNama()
			tampilkanData()

		case 10:
			statistikPerKelas()

		case 11:
			tampilAlphaTerbanyak()

		case 12:
			tampilkanData()

		case 0:

			saveData()

			fmt.Println("Program selesai.")
			return

		default:
			fmt.Println("Menu tidak tersedia.")
		}
	}
}