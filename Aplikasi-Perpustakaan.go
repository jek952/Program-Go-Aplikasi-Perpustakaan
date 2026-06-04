package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var input = bufio.NewReader(os.Stdin)

const MAXBUKU int = 100
const MAXPINJAM int = 100

type Buku struct {
	ID       int
	Judul    string
	Genre    string
	Penulis  string
	Tahun    int
	Stok     int
	Dipinjam int
	Tarif    int
}

type Peminjaman struct {
	IDPinjam     int
	NamaPeminjam string
	IDBuku       int
	LamaPinjam   int
	Denda        int
}

var dataBuku [MAXBUKU]Buku
var dataPeminjam [MAXPINJAM]Peminjaman

func tampilAwal(dataBuku *[MAXBUKU]Buku, jumlahBuku *int, dataPeminjam *[MAXPINJAM]Peminjaman, jumlahPinjam *int) {
	var pilih int
	var jalan bool = true

	for jalan {
		fmt.Println("╔══════════════════════════════════════════════╗")
		fmt.Println("║                                              ║")
		fmt.Println("║             APLIKASI PERPUSTAKAAN            ║")
		fmt.Println("║                                              ║")
		fmt.Println("╠══════════════════════════════════════════════╣")
		fmt.Println("║                  Created By                  ║")
		fmt.Println("║                                              ║")
		fmt.Println("║    • Dzakky Al Firdaus   (103032500023)      ║")
		fmt.Println("║    • Reyhan Putra D.A.H  (103032500023)      ║")
		fmt.Println("║                                              ║")
		fmt.Println("╚══════════════════════════════════════════════╝")
		fmt.Println()
		fmt.Println("1. Masuk ke Menu")
		fmt.Println("2. Exit")

		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tampilMenu(dataBuku, jumlahBuku, dataPeminjam, jumlahPinjam)
		} else if pilih == 2 {
			fmt.Println("Terimakasih Sudah Menggunakan Aplikasi")
			jalan = false // Menggantikan 'break'
		} else {
			fmt.Println("Silahkan masukkan pilihan yang tersedia!")
		}
	}
}

func tampilMenu(dataBuku *[MAXBUKU]Buku, jumlahBuku *int, dataPeminjam *[MAXPINJAM]Peminjaman, jumlahPinjam *int) {
	var pilihan int

	for pilihan != 10 {
		fmt.Println("╔══════════════════════════════════════════════╗")
		fmt.Println("║                                              ║")
		fmt.Println("║          📚APLIKASI PERPUSTAKAAN📚           ║")
		fmt.Println("║                                              ║")
		fmt.Println("╠══════════════════════════════════════════════╣")
		fmt.Println("║                  Created By                  ║")
		fmt.Println("║                                              ║")
		fmt.Println("║  • Dzakky Al Firdaus   (103032500023)        ║")
		fmt.Println("║  • Reyhan Putra D.A.H  (103032500023)        ║")
		fmt.Println("║                                              ║")
		fmt.Println("╚══════════════════════════════════════════════╝")
		fmt.Println("1.  Tambah Buku")
		fmt.Println("2.  Tampilkan Buku")
		fmt.Println("3.  Cari Buku")
		fmt.Println("4.  Edit Buku")
		fmt.Println("5.  Sorting Buku")
		fmt.Println("6.  Hapus Buku")
		fmt.Println("7.  Pinjam Buku (Belum Tersedia)")
		fmt.Println("8.  Kembalikan Buku (Belum Tersedia)")
		fmt.Println("9.  Buku Favorit (Belum Tersedia)")
		fmt.Println("10. Exit")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahBuku(dataBuku, jumlahBuku)
		} else if pilihan == 2 {
			tampilBuku(dataBuku, *jumlahBuku)
		} else if pilihan == 3 {
			cariBuku(dataBuku, *jumlahBuku)
		} else if pilihan == 4 {
			editBuku(dataBuku, *jumlahBuku)
		} else if pilihan == 5 {
			menuSorting(dataBuku, *jumlahBuku)
		} else if pilihan == 6 {
			hapusBuku(dataBuku, jumlahBuku)
		} else if pilihan == 7 {
			pinjamBuku(dataBuku, *jumlahBuku, dataPeminjam, jumlahPinjam)
		} else if pilihan == 8 {
			kembalikanBuku(dataBuku, *jumlahBuku, dataPeminjam, jumlahPinjam)
		} else if pilihan == 9 {
			// bukuFavorit(dataBuku, *jumlahBuku)
		} else if pilihan == 10 {
			fmt.Println("\nKembali ke halaman awal 🔙")
		}
	}
}

func tambahBuku(dataBuku *[MAXBUKU]Buku, jumlahBuku *int) {
	var lagi string

	for lagi != "no" && lagi != "n" {
		fmt.Println()
		fmt.Println("===== TAMBAH BUKU =====")

		fmt.Print("Masukkan ID Buku : ")
		fmt.Scan(&dataBuku[*jumlahBuku].ID)

		input.ReadString('\n')

		fmt.Print("Masukkan Judul Buku : ")
		judul, _ := input.ReadString('\n')
		dataBuku[*jumlahBuku].Judul = strings.TrimSpace(judul)

		fmt.Print("Masukkan Penulis Buku : ")
		penulis, _ := input.ReadString('\n')
		dataBuku[*jumlahBuku].Penulis = strings.TrimSpace(penulis)

		fmt.Print("Masukkan Tahun : ")
		fmt.Scan(&dataBuku[*jumlahBuku].Tahun)

		fmt.Print("Stok: ")
		fmt.Scan(&dataBuku[*jumlahBuku].Stok)

		dataBuku[*jumlahBuku].Dipinjam = 0

		*jumlahBuku++

		fmt.Println()
		fmt.Println("Buku berhasil ditambahkan!")

		fmt.Print("Tambah buku lagi? (yes/no): ")
		fmt.Scan(&lagi)
		fmt.Println()
	}
}

func tampilBuku(dataBuku *[MAXBUKU]Buku, jumlahBuku int) {
	var i int

	fmt.Println()
	fmt.Println("===== DATA BUKU =====")

	if jumlahBuku == 0 {
		fmt.Println("Belum ada data buku")
	} else {
		for i = 0; i < jumlahBuku; i++ {
			fmt.Println("ID       :", dataBuku[i].ID)
			fmt.Println("Judul    :", dataBuku[i].Judul)
			fmt.Println("Penulis  :", dataBuku[i].Penulis)
			fmt.Println("Tahun    :", dataBuku[i].Tahun)
			fmt.Println("Stok     :", dataBuku[i].Stok)
			fmt.Println("Dipinjam :", dataBuku[i].Dipinjam, "kali")
			fmt.Println("Tarif/Hr :", dataBuku[i].Tarif)
			fmt.Println("======================")
		}
	}

	fmt.Println("Tekan enter untuk kembali...")
	fmt.Scanln()
	fmt.Scanln()
}

func sequentialSearchJudul(dataBuku *[MAXBUKU]Buku, jumlahBuku int, cariJudul string) int {
	var idx int = -1
	for i := 0; i < jumlahBuku && idx == -1; i++ {
		if dataBuku[i].Judul == cariJudul {
			idx = i
		}
	}
	return idx
}

func binarySearchID(databuku *[MAXBUKU]Buku, jumlahBuku int, cariID int) int {
	var low, mid int
	var high int = jumlahBuku - 1
	var idx int = -1

	for low <= high && idx == -1 {
		mid = (low + high) / 2
		if dataBuku[mid].ID == cariID {
			idx = mid
		} else if dataBuku[mid].ID < cariID {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return idx
}

func urutkanBukuByID(dataBuku *[MAXBUKU]Buku, jumlahBuku int) {
	var temp Buku
	var j int

	for i := 1; i < jumlahBuku; i++ {
		temp = dataBuku[i]
		j = i - 1
		for j >= 0 && dataBuku[j].ID > temp.ID {
			dataBuku[j+1] = dataBuku[j]
			j--
		}
		dataBuku[j+1] = temp
	}
}

func cariBuku(dataBuku *[MAXBUKU]Buku, jumlahBuku int) {
	var ulang string
	var pilihan int

	for ulang != "n" && ulang != "no" {
		fmt.Println("\n===== Cari Buku =====")
		fmt.Println("1. Cari Berdasarkan Judul (Sequential Search)")
		fmt.Println("2. Cari Berdasarkan ID (Binary Search)")
		fmt.Print("Pilih jenis pencarian: ")
		fmt.Scan(&pilihan)

		input.ReadString('\n')
		var idx int = -1

		if pilihan == 1 {
			fmt.Print("Masukkan Judul yang dicari : ")
			cariJudul, _ := input.ReadString('\n')
			idx = sequentialSearchJudul(dataBuku, jumlahBuku, strings.TrimSpace(cariJudul))
		} else if pilihan == 2 {
			urutkanBukuByID(dataBuku, jumlahBuku)
			var cariID int
			fmt.Print("Masukkan ID yang dicari : ")
			fmt.Scan(&cariID)
			idx = binarySearchID(dataBuku, jumlahBuku, cariID)
		}

		if idx != -1 {
			fmt.Println("\n===== Buku ditemukan! =====")
			fmt.Println("ID       :", dataBuku[idx].ID)
			fmt.Println("Judul    :", dataBuku[idx].Judul)
			fmt.Println("Genre    :", dataBuku[idx].Genre)
			fmt.Println("Penulis  :", dataBuku[idx].Penulis)
			fmt.Println("Tahun    :", dataBuku[idx].Tahun)
			fmt.Println("Stok     :", dataBuku[idx].Stok)
			fmt.Println("==========================")
		} else {
			fmt.Println("Buku tidak ditemukan ❌")
		}
		fmt.Print("\nApakah ingin mencari buku lagi? (yes/no): ")
		fmt.Scan(&ulang)
	}
}

func editBuku(dataBuku *[MAXBUKU]Buku, jumlahBuku int) {
	var cariID int
	var ditemukan bool
	var ulang string

	for ulang != "n" {
		ditemukan = false

		fmt.Println()
		fmt.Println("===== Edit Buku =====")
		fmt.Print("Masukkan ID Buku :")
		fmt.Scan(&cariID)

		for i := 0; i < jumlahBuku; i++ {
			if dataBuku[i].ID == cariID {
				input.ReadString('\n')

				fmt.Print("Masukkan Judul Baru : ")
				judul, _ := input.ReadString('\n')
				dataBuku[i].Judul = strings.TrimSpace(judul)

				fmt.Print("Masukkan Genre Baru : ")
				fmt.Scan(&dataBuku[i].Genre)

				fmt.Print("Edit Penulis Baru : ")
				penulis, _ := input.ReadString('\n')
				dataBuku[i].Penulis = strings.TrimSpace(penulis)

				fmt.Print("Edit Tahun : ")
				fmt.Scan(&dataBuku[i].Tahun)

				fmt.Print("Edit Stok Buku : ")
				fmt.Scan(&dataBuku[i].Stok)

				fmt.Println("Data berhasil diubah")
				ditemukan = true
			}
		}

		if !ditemukan {
			fmt.Println("Tidak ada data yang ditemukan")
		}
		fmt.Print("Apakah ingin mengubah data lagi ? (y/n) :")
		fmt.Scan(&ulang)
	}
}

func hapusBuku(dataBuku *[MAXBUKU]Buku, jumlahBuku *int) {
	var cariID int
	var idx int
	var ditemukan bool
	var ulang string

	for ulang != "n" {
		ditemukan = false

		fmt.Println()
		fmt.Println("===== Hapus Buku =====")
		fmt.Print("Masukkan ID Buku : ")
		fmt.Scan(&cariID)

		for i := 0; i < *jumlahBuku; i++ {
			if dataBuku[i].ID == cariID {
				idx = i
				ditemukan = true
			}
		}

		if ditemukan {
			for i := idx; i < *jumlahBuku-1; i++ {
				dataBuku[i] = dataBuku[i+1]
			}
			*jumlahBuku--
			fmt.Println("Data berhasil dihapus!")
		} else {
			fmt.Println("ID buku tidak ditemukan")
		}
		fmt.Print("Apakah ingin menghapus lagi? (y/n): ")
		fmt.Scan(&ulang)
	}
}

func pinjamBuku(dataBuku *[MAXBUKU]Buku, jumlahBuku int, dataPeminjam *[MAXPINJAM]Peminjaman, jumlahPinjam *int) {
	var ulang string
	var pilihan int

	for ulang != "n" && ulang != "no" {
		fmt.Println("\n===== PINJAM BUKU =====")
		fmt.Println("1. Cari Berdasarkan Judul")
		fmt.Println("2. Cari Berdasarkan ID")
		fmt.Println("Pilih jenis pencarian : ")
		fmt.Scan(&pilihan)

		input.ReadString('\n')

		var idxBuku int = -1

		if pilihan == 1 {
			fmt.Print("Masukkan Judul Buku yang ingin dipinjam : ")
			cariJudul, _ := input.ReadString('\n')
			idxBuku = sequentialSearchJudul(dataBuku, jumlahBuku, strings.TrimSpace(cariJudul))
		} else if pilihan == 2 {
			var cariIDBuku int
			fmt.Print("Masukkan ID Buku yang akan dipinjam: ")
			fmt.Scan(&cariIDBuku)
			urutkanBukuByID(dataBuku, jumlahBuku)
			idxBuku = binarySearchID(dataBuku, jumlahBuku, cariIDBuku)
		} else {
			fmt.Println("Pilihan pencarian tidak valid")
		}

		if idxBuku != -1 {
			fmt.Print("ID Peminjaman : ")
			fmt.Scan(&dataPeminjam[*jumlahPinjam].IDPinjam)

			input.ReadString('\n')

			fmt.Print("Nama Peminjam : ")
			namaPeminjam, _ := input.ReadString('\n')
			dataPeminjam[*jumlahPinjam].NamaPeminjam = strings.TrimSpace(namaPeminjam)

			dataPeminjam[*jumlahPinjam].IDBuku = dataBuku[idxBuku].ID
			dataPeminjam[*jumlahPinjam].Denda = 0

			dataBuku[idxBuku].Stok--
			dataBuku[idxBuku].Dipinjam++

			*jumlahPinjam++
			fmt.Println("Buku berhasil dipinjam!")
		} else {
			fmt.Println("Mohon maaf, stok buku sedang habis")
		}
	}
}

func kembalikanBuku(dataBuku *[MAXBUKU]Buku, jumlahBuku int, dataPeminjam *[MAXPINJAM]Peminjaman, jumlahPinjam *int) {
	var ulang string
	var cariIDPinjam int

	for ulang != "n" && ulang != "no" {
		fmt.Println("\n===== KEMBALIKAN BUKU =====")
		fmt.Print("Masukkan ID peminjaman : ")
		fmt.Scan(&cariIDPinjam)

		idxPinjam := -1

		for i := 0; i < *jumlahPinjam && idxPinjam == -1; i++ {
			if dataPeminjam[i].IDPinjam == cariIDPinjam {
				idxPinjam = i
			}
		}

		if idxPinjam != -1 {

		}

	}

}

func menuSorting(dataBuku *[MAXBUKU]Buku, jumlahBuku int) {
	var pilih int

	for pilih != 5 {
		fmt.Println("===== Menu Sorting =====")
		fmt.Println("1. Judul Ascending")
		fmt.Println("2. Judul Descending")
		fmt.Println("3. Stok Ascending")
		fmt.Println("4. Stok Descending")
		fmt.Println("5. Kembali ke menu")

		fmt.Print("Masukkan Pilihan : ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			selectionSortJudulAsc(dataBuku, jumlahBuku)
			tampilBuku(dataBuku, jumlahBuku)
		} else if pilih == 2 {
			selectionSortJudulDesc(dataBuku, jumlahBuku)
			tampilBuku(dataBuku, jumlahBuku)
		} else if pilih == 3 {
			insertionSortStokAsc(dataBuku, jumlahBuku)
			tampilBuku(dataBuku, jumlahBuku)
		} else if pilih == 4 {
			insertionSortStokDesc(dataBuku, jumlahBuku)
			tampilBuku(dataBuku, jumlahBuku)
		}
	}
}

func selectionSortJudulAsc(dataBuku *[MAXBUKU]Buku, jumlahBuku int) {
	var temp Buku
	var idx int

	for i := 0; i < jumlahBuku; i++ {
		idx = i
		for j := i + 1; j < jumlahBuku; j++ {
			if dataBuku[j].Judul < dataBuku[idx].Judul {
				idx = j
			}
		}
		temp = dataBuku[i]
		dataBuku[i] = dataBuku[idx]
		dataBuku[idx] = temp
	}
	fmt.Println("Data judul berhasil diurutkan secara ascending")
}

func selectionSortJudulDesc(dataBuku *[MAXBUKU]Buku, jumlahBuku int) {
	var temp Buku
	var idx int

	for i := 0; i < jumlahBuku; i++ {
		idx = i
		for j := i + 1; j < jumlahBuku; j++ {
			if dataBuku[j].Judul > dataBuku[idx].Judul {
				idx = j
			}
		}
		temp = dataBuku[i]
		dataBuku[i] = dataBuku[idx]
		dataBuku[idx] = temp
	}
	fmt.Println("Data judul berhasil diurutkan secara descending")
}

func insertionSortStokAsc(dataBuku *[MAXBUKU]Buku, jumlahBuku int) {
	var temp Buku
	var j int

	for i := 1; i < jumlahBuku; i++ {
		temp = dataBuku[i]
		j = i - 1
		for j >= 0 && dataBuku[j].Stok > temp.Stok {
			dataBuku[j+1] = dataBuku[j]
			j--
		}
		dataBuku[j+1] = temp
	}
	fmt.Println("Data stok berhasil diurutkan secara ascending")
}

func insertionSortStokDesc(dataBuku *[MAXBUKU]Buku, jumlahBuku int) {
	var temp Buku
	var j int

	for i := 1; i < jumlahBuku; i++ {
		temp = dataBuku[i]
		j = i - 1
		for j >= 0 && dataBuku[j].Stok < temp.Stok {
			dataBuku[j+1] = dataBuku[j]
			j--
		}
		dataBuku[j+1] = temp
	}
	fmt.Println("Data stok berhasil diurutkan secara descending")
}

func main() {
	var jumlahBuku int
	var jumlahPinjam int

	tampilAwal(&dataBuku, &jumlahBuku, &dataPeminjam, &jumlahPinjam)

}
