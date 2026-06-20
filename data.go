package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// =====================
// STRUCT MAHASISWA
// =====================

type Mahasiswa struct {
	NIM   string
	Nama  string
	Kelas string

	Hadir int
	Izin  int
	Sakit int
	Alpha int
}

// =====================
// DATABASE DUMMY
// =====================

var data [100]Mahasiswa
var jumlah int

// =====================
// LOAD DATA DARI FILE
// =====================

func loadData() {

	file, err := os.Open("mahasiswa.txt")

	if err != nil {

		// Jika file belum ada
		file, _ = os.Create("mahasiswa.txt")
		file.Close()

		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	jumlah = 0

	for scanner.Scan() {

		line := scanner.Text()

		part := strings.Split(line, "|")

		if len(part) != 7 {
			continue
		}

		data[jumlah].NIM = part[0]
		data[jumlah].Nama = part[1]
		data[jumlah].Kelas = part[2]

		data[jumlah].Hadir, _ = strconv.Atoi(part[3])
		data[jumlah].Izin, _ = strconv.Atoi(part[4])
		data[jumlah].Sakit, _ = strconv.Atoi(part[5])
		data[jumlah].Alpha, _ = strconv.Atoi(part[6])

		jumlah++
	}
}

// =====================
// SIMPAN DATA KE FILE
// =====================

func saveData() {

	file, err := os.Create("mahasiswa.txt")

	if err != nil {
		fmt.Println("Gagal menyimpan data.")
		return
	}

	defer file.Close()

	for i := 0; i < jumlah; i++ {

		fmt.Fprintf(
			file,
			"%s|%s|%s|%d|%d|%d|%d\n",

			data[i].NIM,
			data[i].Nama,
			data[i].Kelas,

			data[i].Hadir,
			data[i].Izin,
			data[i].Sakit,
			data[i].Alpha,
		)
	}
}