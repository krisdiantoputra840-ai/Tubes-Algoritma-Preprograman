package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// =====================
// INPUT HELPER
// =====================

func inputString(prompt string) string {

	fmt.Print(prompt)

	text, _ := reader.ReadString('\n')

	return strings.TrimSpace(text)
}

func inputInt(prompt string) int {

	fmt.Print(prompt)

	text, _ := reader.ReadString('\n')

	text = strings.TrimSpace(text)

	num, err := strconv.Atoi(text)

	if err != nil {
		return 0
	}

	return num
}

// =====================
// TAMBAH MAHASISWA
// =====================

func tambahMahasiswa() {

	if jumlah >= 100 {

		fmt.Println("Data mahasiswa penuh")

		return
	}

	nim := inputString("NIM   : ")

	if sequentialSearch(nim) != -1 {

		fmt.Println("NIM sudah terdaftar")

		return
	}

	nama := inputString("Nama  : ")
	kelas := inputString("Kelas : ")

	data[jumlah] = Mahasiswa{
		NIM:   nim,
		Nama:  nama,
		Kelas: kelas,
		Hadir: 0,
		Izin:  0,
		Sakit: 0,
		Alpha: 0,
	}

	jumlah++

	saveData()

	fmt.Println("Mahasiswa berhasil ditambahkan")
}

// =====================
// UBAH MAHASISWA
// =====================

func ubahMahasiswa() {

	nim := inputString("Masukkan NIM : ")

	idx := sequentialSearch(nim)

	if idx == -1 {

		fmt.Println("Data tidak ditemukan")

		return
	}

	fmt.Println("\nData Ditemukan")
	fmt.Println("Nama Lama  :", data[idx].Nama)
	fmt.Println("Kelas Lama :", data[idx].Kelas)

	data[idx].Nama = inputString("Nama Baru  : ")
	data[idx].Kelas = inputString("Kelas Baru : ")

	saveData()

	fmt.Println("Data berhasil diubah")
}

// =====================
// HAPUS MAHASISWA
// =====================

func hapusMahasiswa() {

	nim := inputString("Masukkan NIM : ")

	idx := sequentialSearch(nim)

	if idx == -1 {

		fmt.Println("Data tidak ditemukan")

		return
	}

	for i := idx; i < jumlah-1; i++ {

		data[i] = data[i+1]
	}

	jumlah--

	saveData()

	fmt.Println("Data berhasil dihapus")
}

// =====================
// INPUT ABSENSI
// =====================

func inputAbsensi() {

	nim := inputString("Masukkan NIM : ")

	idx := sequentialSearch(nim)

	if idx == -1 {

		fmt.Println("Mahasiswa tidak ditemukan")

		return
	}

	fmt.Println()
	fmt.Println("1. Hadir")
	fmt.Println("2. Izin")
	fmt.Println("3. Sakit")
	fmt.Println("4. Alpha")

	pilih := inputInt("Pilih Status : ")

	switch pilih {

	case 1:
		data[idx].Hadir++

	case 2:
		data[idx].Izin++

	case 3:
		data[idx].Sakit++

	case 4:
		data[idx].Alpha++

	default:
		fmt.Println("Pilihan tidak valid")
		return
	}

	saveData()

	fmt.Println("Absensi berhasil dicatat")
}

// =====================
// TAMPIL DATA
// =====================

func tampilkanData() {

	if jumlah == 0 {

		fmt.Println("Belum ada data")

		return
	}

	fmt.Println()
	fmt.Println("================================================================================")
	fmt.Printf("%-12s %-25s %-10s %-5s %-5s %-5s %-5s\n",
		"NIM", "Nama", "Kelas", "H", "I", "S", "A")
	fmt.Println("================================================================================")

	for i := 0; i < jumlah; i++ {

		fmt.Printf("%-12s %-25s %-10s %-5d %-5d %-5d %-5d\n",
			data[i].NIM,
			data[i].Nama,
			data[i].Kelas,
			data[i].Hadir,
			data[i].Izin,
			data[i].Sakit,
			data[i].Alpha)
	}
}