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
var jumlahBuku int
var jumlahPinjam int

func tampilAwal() {
	var pilih int

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

	for {

		fmt.Println()
		fmt.Println("1. Masuk ke Menu")
		fmt.Println("2. Exit")

		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)

		if pilih == 1 {

			tampilMenu()

		} else if pilih == 2 {

			fmt.Println("Terimakasih Sudah Menggunakan Aplikasi")
			return

		} else {

			fmt.Println("Silahkan masukkan pilihan yang tersedia!")
		}
	}
}

func tampilMenu() {
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
		fmt.Println("7.  Pinjam Buku")
		fmt.Println("8.  Kembalikan Buku")
		fmt.Println("9.  Buku Favorit")
		fmt.Println("10. Exit")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahBuku()
		} else if pilihan == 2 {
			tampilBuku()
		} else if pilihan == 3 {
			cariBuku()
		} else if pilihan == 4 {
			editBuku()
		} else if pilihan == 5 {
			menuSorting()
		} else if pilihan == 6 {
			hapusBuku()
		} else if pilihan == 10 {
			fmt.Println()
			fmt.Println("Kembali ke halaman awal 🔙")
			tampilAwal()
		}
	}
}

func tambahBuku() {

	var lagi string

	for lagi != "n" {
		fmt.Println()
		fmt.Println("===== TAMBAH BUKU =====")

		fmt.Print("Masukkan ID Buku : ")
		fmt.Scan(&dataBuku[jumlahBuku].ID)

		input.ReadString('\n')

		fmt.Print("Masukkan Judul Buku : ")
		judul, _ := input.ReadString('\n')
		dataBuku[jumlahBuku].Judul = strings.TrimSpace(judul)

		fmt.Print("Masukkan Penulis Buku : ")
		penulis, _ := input.ReadString('\n')
		dataBuku[jumlahBuku].Penulis = strings.TrimSpace(penulis)

		fmt.Print("Masukkan Tahun : ")
		fmt.Scan(&dataBuku[jumlahBuku].Tahun)

		fmt.Print("Stok: ")
		fmt.Scan(&dataBuku[jumlahBuku].Stok)

		dataBuku[jumlahBuku].Dipinjam = 0

		jumlahBuku++

		fmt.Println()
		fmt.Println("Buku berhasil ditambahkan!")

		fmt.Print("Tambah buku lagi? (y/n): ")
		fmt.Scan(&lagi)
		fmt.Println()
	}
}

func tampilBuku() {

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
			fmt.Println("======================")
		}
	}

	fmt.Println("Tekan enter untuk kembali...")
	fmt.Scanln()
	fmt.Scanln()
}

func sequentialSearchJudul(cariJudul string) int {
	for i := 0; i < jumlahBuku; i++ {

		if dataBuku[i].Judul == cariJudul {

			return i
		}
	}

	return -1
}

func cariBuku() {

	var cariJudul string
	var ulang string

	for ulang != "n" {

		fmt.Println()
		fmt.Println("===== Cari Buku =====")

		input.ReadString('\n')

		fmt.Print("Masukkan Judul Buku yang dicari : ")
		cariJudul, _ = input.ReadString('\n')
		cariJudul = strings.TrimSpace(cariJudul)

		idx := sequentialSearchJudul(cariJudul)

		if idx != -1 {

			fmt.Println()
			fmt.Println("===== Buku ditemukan! =====")
			fmt.Println("ID       :", dataBuku[idx].ID)
			fmt.Println("Judul    :", dataBuku[idx].Judul)
			fmt.Println("Penulis  :", dataBuku[idx].Penulis)
			fmt.Println("Tahun    :", dataBuku[idx].Tahun)
			fmt.Println("Stok     :", dataBuku[idx].Stok)
			fmt.Println("==========================")

		} else {

			fmt.Println("Buku tidak ditemukan ❌")
		}

		fmt.Print("Apakah ingin mencari buku lagi? (y/n): ")
		fmt.Scan(&ulang)
	}
}

func editBuku() {
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

				fmt.Print("Masukkan Judul Baru :")
				judul, _ := input.ReadString('\n')
				dataBuku[i].Judul = strings.TrimSpace(judul)

				fmt.Print("Edit Penulis Baru :")
				penulis, _ := input.ReadString('\n')
				dataBuku[i].Penulis = strings.TrimSpace(penulis)

				fmt.Print("Edit Tahun :")
				fmt.Scan(&dataBuku[i].Tahun)

				fmt.Print("Edit Stok Buku :")
				fmt.Scan(&dataBuku[i].Stok)

				fmt.Println("Data berhasil diubah")

				ditemukan = true
			}
		}

		if ditemukan == false {
			fmt.Println("Tidak ada data yang ditemukan")
		}
		fmt.Print("Apakah ingin mengubah data lagi ? (y/n) :")
		fmt.Scan(&ulang)
	}
}

func hapusBuku() {
	var cariID int
	var idx int
	var ditemukan bool
	var ulang string

	for ulang != "n" {

		ditemukan = false

		fmt.Println()
		fmt.Println("===== Hapus Buku =====")
		fmt.Println("Masukkan ID Buku : ")
		fmt.Scan(&cariID)

		for i := 0; i < jumlahBuku; i++ {
			if dataBuku[i].ID == cariID {
				idx = i
				ditemukan = true
			}
		}

		if ditemukan == true {
			for i := idx; i < jumlahBuku-1; i++ {

				dataBuku[i] = dataBuku[i+1]
			}

			jumlahBuku--

			fmt.Println("Data berhasil dihapus!")
		} else {
			fmt.Println("ID buku tidak ditemukan")
		}
		fmt.Print("Apakah ingin menghapus lagi? (y/n): ")
		fmt.Scan(&ulang)
	}
}

func menuSorting() {
	var pilih int

	for pilih != 5 {
		fmt.Println("===== Menu Sorting =====")
		fmt.Println("1. Judul Ascending")
		fmt.Println("2. Judul Descending")
		fmt.Println("3. Stok Ascending")
		fmt.Println("4. Stok Descending")
		fmt.Println("5. Stok Ascending")

		fmt.Println("Masukkan Pilihan : ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			selectionSortJudulAsc()
		} else if pilih == 2 {
			selectionSortJudulDesc()
		} else if pilih == 3 {
			insertionSortStokAsc()
		} 
	}
}

func selectionSortJudulAsc() {
	var temp Buku
	var idx int

	for i := 0; i < jumlahBuku; i++ {
		idx = i

		for j := i + 1; j < jumlahBuku; i++ {
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

func selectionSortJudulDesc() {
	var temp Buku
	var idx int

	for i := 0; i < jumlahBuku; i++ {
		idx = i

		for j := i + 1; j < jumlahBuku; i++ {
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

func insertionSortStokAsc(){
	
}

func main() {
	tampilAwal()

}
