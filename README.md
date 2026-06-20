# SIPRESENSI CLI

SIPRESENSI CLI adalah aplikasi Sistem Informasi Presensi Mahasiswa berbasis Command Line Interface (CLI) yang dikembangkan menggunakan bahasa pemrograman Go (Golang). Aplikasi ini bertujuan untuk membantu pengelolaan data presensi mahasiswa secara lebih terstruktur, cepat, dan efisien.

## Fitur Utama

* Menambah data mahasiswa.
* Mengubah data mahasiswa.
* Menghapus data mahasiswa.
* Input absensi mahasiswa (Hadir, Izin, Sakit, Alpha).
* Menampilkan seluruh data mahasiswa.
* Pencarian data berdasarkan NIM menggunakan:

  * Sequential Search
  * Binary Search
* Pencarian berdasarkan status kehadiran.
* Pengurutan data berdasarkan:

  * Nama (Insertion Sort)
  * Total Absensi (Selection Sort)
* Menampilkan statistik kehadiran mahasiswa.
* Menampilkan mahasiswa dengan jumlah alpha terbanyak.
* Penyimpanan data otomatis ke file.

## Struktur Project

```text
SIPRESENSI/
│
├── presensi.go      # Program utama
├── crud.go          # Operasi CRUD dan input pengguna
├── search.go        # Algoritma pencarian
├── sort.go          # Algoritma pengurutan
├── statistik.go     # Statistik dan rekap absensi
├── data.go          # Struktur data dan penyimpanan file
├── mahasiswa.txt    # File penyimpanan data
└── README.md        # Dokumentasi project
```

## Teknologi yang Digunakan

* Bahasa Pemrograman: Go (Golang)
* Interface: Command Line Interface (CLI)
* Penyimpanan Data: File Text (.txt)

## Algoritma yang Digunakan

### Searching

* Sequential Search
* Binary Search

### Sorting

* Selection Sort
* Insertion Sort

## Cara Menjalankan Program

1. Clone repository ini.

```bash
git clone https://github.com/username/sipresensi-cli.git
```

2. Masuk ke folder project.

```bash
cd sipresensi-cli
```

3. Jalankan program.

```bash
go run *.go
```

## Tampilan Menu

```text
======================================
           SIPRESENSI CLI
======================================
1. Tambah Mahasiswa
2. Ubah Mahasiswa
3. Hapus Mahasiswa
4. Input Absensi
5. Cari NIM (Sequential Search)
6. Cari NIM (Binary Search)
7. Cari Berdasarkan Status
8. Urutkan Total Absensi
9. Urutkan Nama
10. Statistik Kehadiran
11. Alpha Terbanyak
12. Tampilkan Data
0. Keluar
```

## Tujuan Pengembangan

Project ini dibuat sebagai Tugas Besar mata kuliah Algoritma dan Pemrograman dengan tujuan:

* Menerapkan konsep algoritma dan struktur data.
* Mengimplementasikan algoritma searching dan sorting.
* Melatih kemampuan analisis, perancangan, dan implementasi program.

## Pengembang

**Krisdianto Meissa Aji Putra And Ringgas Noble Fristion - IF 05 04**

## Lisensi

Project ini dibuat untuk keperluan akademik dan pembelajaran.
