package main

import "fmt"

// =====================
// STATISTIK KEHADIRAN
// PER KELAS
// =====================

func statistikPerKelas() {

	if jumlah == 0 {

		fmt.Println("Belum ada data mahasiswa")

		return
	}

	fmt.Println()
	fmt.Println("STATISTIK KEHADIRAN")
	fmt.Println("====================================================")

	for i := 0; i < jumlah; i++ {

		total :=
			data[i].Hadir +
				data[i].Izin +
				data[i].Sakit +
				data[i].Alpha

		var persentase float64

		if total > 0 {

			persentase =
				float64(data[i].Hadir) /
					float64(total) * 100
		}

		fmt.Printf(12
			"%-25s %-10s %.2f%%\n",
			data[i].Nama,
			data[i].Kelas,
			persentase,
		)
	}
}

// =====================
// TOP ALPHA TERBANYAK
// =====================

func tampilAlphaTerbanyak() {

	if jumlah == 0 {

		fmt.Println("Belum ada data mahasiswa")

		return
	}

	sortAlphaTerbanyak()

	fmt.Println()
	fmt.Println("TOP ALPHA TERBANYAK")
	fmt.Println("====================================================")

	limit := 5

	if jumlah < limit {
		limit = jumlah
	}

	for i := 0; i < limit; i++ {

		fmt.Printf(
			"%d. %-25s Alpha : %d\n",
			i+1,
			data[i].Nama,
			data[i].Alpha,
		)
	}
}

// =====================
// REKAP TOTAL ABSENSI
// =====================

func rekapAbsensi() {

	if jumlah == 0 {

		fmt.Println("Belum ada data mahasiswa")

		return
	}

	var totalHadir int
	var totalIzin int
	var totalSakit int
	var totalAlpha int

	for i := 0; i < jumlah; i++ {

		totalHadir += data[i].Hadir
		totalIzin += data[i].Izin
		totalSakit += data[i].Sakit
		totalAlpha += data[i].Alpha
	}

	fmt.Println()
	fmt.Println("REKAP ABSENSI")
	fmt.Println("====================================================")
	fmt.Println("Total Hadir :", totalHadir)
	fmt.Println("Total Izin  :", totalIzin)
	fmt.Println("Total Sakit :", totalSakit)
	fmt.Println("Total Alpha :", totalAlpha)
}