package main

import "fmt"

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
	IDPinjam       int
	NamaPeminjam   string
	IDBuku         int
	LamaPinjam     int
	Denda          int
	TanggalPinjam  int
	TanggalKembali int
}

var dataBuku [MAXBUKU]Buku
var dataPeminjam [MAXPINJAM]Peminjaman

func tampilAwal(dataBuku *[MAXBUKU]Buku, jumlahBuku *int, dataPeminjam *[MAXPINJAM]Peminjaman, jumlahPinjam *int) {
	var pilih int
	var jalan bool = true

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

	for jalan {
		fmt.Println()
		fmt.Println("1. Masuk ke Menu")
		fmt.Println("2. Exit")

		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tampilMenu(dataBuku, jumlahBuku, dataPeminjam, jumlahPinjam)
		} else if pilih == 2 {
			fmt.Println("Terimakasih Sudah Menggunakan Aplikasi")
			jalan = false
		} else {
			fmt.Println("\nSilahkan masukkan pilihan yang tersedia!")
		}
	}
}

func tampilMenu(dataBuku *[MAXBUKU]Buku, jumlahBuku *int, dataPeminjam *[MAXPINJAM]Peminjaman, jumlahPinjam *int) {
	var pilihan int

	for pilihan != 12 {
		fmt.Println("\n==================================================")
		fmt.Println("                   MENU UTAMA                     ")
		fmt.Println("==================================================")
		fmt.Println()
		fmt.Println("1.  Tambah Buku")
		fmt.Println("2.  Tampilkan Buku")
		fmt.Println("3.  Cari Buku")
		fmt.Println("4.  Edit Buku")
		fmt.Println("5.  Hapus Buku")
		fmt.Println("6.  Sorting Buku")
		fmt.Println("7.  Pinjam Buku")
		fmt.Println("8.  Kembalikan Buku")
		fmt.Println("9.  Daftar Peminjaman Aktif")
		fmt.Println("10. Edit Data Peminjaman")
		fmt.Println("11. 5 Buku Terfavorit")
		fmt.Println("12. Exit")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahBuku(dataBuku, jumlahBuku)
		case 2:
			tampilBuku(dataBuku, *jumlahBuku)
		case 3:
			cariBuku(dataBuku, *jumlahBuku)
		case 4:
			editBuku(dataBuku, *jumlahBuku)
		case 5:
			hapusBuku(dataBuku, jumlahBuku)
		case 6:
			menuSorting(dataBuku, *jumlahBuku)
		case 7:
			pinjamBuku(dataBuku, *jumlahBuku, dataPeminjam, jumlahPinjam)
		case 8:
			kembalikanBuku(dataBuku, *jumlahBuku, dataPeminjam, jumlahPinjam)
		case 9:
			tampilPeminjaman(dataPeminjam, *jumlahPinjam, dataBuku, *jumlahBuku)
		case 10:
			editPeminjaman(dataPeminjam, *jumlahPinjam)
		case 11:
			bukuFavorit(dataBuku, *jumlahBuku)
		case 12:
			fmt.Println("\nKembali ke halaman awal")
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func tambahBuku(dataBuku *[MAXBUKU]Buku, jumlahBuku *int) {
	var lagi string

	for lagi != "no" && lagi != "n" {
		fmt.Println()
		fmt.Println("===== MENU TAMBAH BUKU =====")

		fmt.Print("Masukkan ID Buku : ")
		fmt.Scan(&dataBuku[*jumlahBuku].ID)

		fmt.Print("Masukkan Judul Buku : ")
		fmt.Scan(&dataBuku[*jumlahBuku].Judul)

		fmt.Print("Masukkan Penulis Buku : ")
		fmt.Scan(&dataBuku[*jumlahBuku].Penulis)

		fmt.Print("Masukkan Genre : ")
		fmt.Scan(&dataBuku[*jumlahBuku].Genre)

		fmt.Print("Masukkan Tahun : ")
		fmt.Scan(&dataBuku[*jumlahBuku].Tahun)

		fmt.Print("Stok : ")
		fmt.Scan(&dataBuku[*jumlahBuku].Stok)

		fmt.Print("Tarif (Per Hari) : ")
		fmt.Scan(&dataBuku[*jumlahBuku].Tarif)

		dataBuku[*jumlahBuku].Dipinjam = 0

		*jumlahBuku++

		fmt.Println()
		fmt.Println("\nBuku berhasil ditambahkan!")

		fmt.Print("Tambah buku lagi? (yes/no) : ")
		fmt.Scan(&lagi)
	}
}

func tampilBuku(dataBuku *[MAXBUKU]Buku, jumlahBuku int) {
	fmt.Println("\n=== [ DAFTAR DATA BUKU ] ===")

	if jumlahBuku == 0 {
		fmt.Println("Belum ada data buku di perpustakaan.")
	} else {
		fmt.Println("===========================================================================================================")
		fmt.Printf("| %-4s | %-25s | %-15s | %-15s | %-5s | %-4s | %-6s | %-9s |\n",
			"ID", "Judul Buku", "Genre", "Penulis", "Tahun", "Stok", "Pinjam", "Tarif/Hr")
		fmt.Println("===========================================================================================================")

		for i := 0; i < jumlahBuku; i++ {
			fmt.Printf("| %-4d | %-25s | %-15s | %-15s | %-5d | %-4d | %-6d | Rp %-6d |\n",
				dataBuku[i].ID,
				dataBuku[i].Judul,
				dataBuku[i].Genre,
				dataBuku[i].Penulis,
				dataBuku[i].Tahun,
				dataBuku[i].Stok,
				dataBuku[i].Dipinjam,
				dataBuku[i].Tarif)
		}
		fmt.Println("===========================================================================================================")
	}
	fmt.Println("Tekan enter untuk kembali....")
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

func binarySearchID(dataBuku *[MAXBUKU]Buku, jumlahBuku int, cariID int) int {
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

func cariBuku(dataBuku *[MAXBUKU]Buku, jumlahBuku int) {
	var ulang string
	var pilihan int
	var idx int = -1
	var cariID int
	var cariJudul string

	for ulang != "n" && ulang != "no" {
		idx = -1
		fmt.Println("\n===== [ CARI BUKU ] =====")
		fmt.Println("1. Cari Berdasarkan Judul (Sequential Search)")
		fmt.Println("2. Cari Berdasarkan ID (Binary Search)")
		fmt.Print("Pilih jenis pencarian : ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Print("Masukkan Judul yang dicari : ")
			fmt.Scan(&cariJudul)
			idx = sequentialSearchJudul(dataBuku, jumlahBuku, cariJudul)
		} else if pilihan == 2 {
			urutkanBukuByID(dataBuku, jumlahBuku)
			fmt.Print("Masukkan ID yang dicari : ")
			fmt.Scan(&cariID)
			idx = binarySearchID(dataBuku, jumlahBuku, cariID)
		}

		if idx != -1 {
			fmt.Println("\n Buku ditemukan!")
			fmt.Println("--------------------------------------------------")
			fmt.Printf("ID       : %d\n", dataBuku[idx].ID)
			fmt.Printf("Judul    : %s\n", dataBuku[idx].Judul)
			fmt.Printf("Genre    : %s\n", dataBuku[idx].Genre)
			fmt.Printf("Penulis  : %s\n", dataBuku[idx].Penulis)
			fmt.Printf("Tahun    : %d\n", dataBuku[idx].Tahun)
			fmt.Printf("Stok     : %d\n", dataBuku[idx].Stok)
			fmt.Println("--------------------------------------------------")
		} else {
			fmt.Println("Buku tidak ditemukan")
		}
		fmt.Print("\nApakah ingin mencari buku lagi? (yes/no) : ")
		fmt.Scan(&ulang)
	}
}

func editBuku(dataBuku *[MAXBUKU]Buku, jumlahBuku int) {
	var cariID int
	var ditemukan bool
	var ulang string
	var pilihanEdit int

	for ulang != "n" && ulang != "no" {
		ditemukan = false

		fmt.Println()
		fmt.Println("===== [ EDIT BUKU ] =====")
		fmt.Print("Masukkan ID Buku yang ingin diedit : ")
		fmt.Scan(&cariID)

		for i := 0; i < jumlahBuku; i++ {
			if dataBuku[i].ID == cariID {
				ditemukan = true

				fmt.Println("\nData Buku Saat Ini:")
				fmt.Printf("1. ID       : %d\n", dataBuku[i].ID)
				fmt.Printf("2. Judul    : %s\n", dataBuku[i].Judul)
				fmt.Printf("3. Penulis  : %s\n", dataBuku[i].Penulis)
				fmt.Printf("4. Genre    : %s\n", dataBuku[i].Genre)
				fmt.Printf("5. Tahun    : %d\n", dataBuku[i].Tahun)
				fmt.Printf("6. Stok     : %d\n", dataBuku[i].Stok)
				fmt.Printf("7. Tarif/Hr : Rp %d\n", dataBuku[i].Tarif)
				fmt.Println("8. Edit Keseluruhan")
				fmt.Println("9. Batal Edit")
				fmt.Print("\nPilih bagian yang ingin diedit (1-9) : ")
				fmt.Scan(&pilihanEdit)

				if pilihanEdit == 1 {
					fmt.Print("Masukkan ID Baru : ")
					fmt.Scan(&dataBuku[i].ID)
					fmt.Println("\n ID berhasil diubah!")
				} else if pilihanEdit == 2 {
					fmt.Print("Masukkan Judul Baru : ")
					fmt.Scan(&dataBuku[i].Judul)
					fmt.Println("\n Judul berhasil diubah!")
				} else if pilihanEdit == 3 {
					fmt.Print("Masukkan Penulis Baru : ")
					fmt.Scan(&dataBuku[i].Penulis)
					fmt.Println("\n Penulis berhasil diubah!")
				} else if pilihanEdit == 4 {
					fmt.Print("Masukkan Genre Baru : ")
					fmt.Scan(&dataBuku[i].Genre)
					fmt.Println("\n Genre berhasil diubah!")
				} else if pilihanEdit == 5 {
					fmt.Print("Masukkan Tahun Baru : ")
					fmt.Scan(&dataBuku[i].Tahun)
					fmt.Println("\n Tahun berhasil diubah!")
				} else if pilihanEdit == 6 {
					fmt.Print("Masukkan Stok Baru : ")
					fmt.Scan(&dataBuku[i].Stok)
					fmt.Println("\n Stok berhasil diubah!")
				} else if pilihanEdit == 7 {
					fmt.Print("Masukkan Tarif Baru : ")
					fmt.Scan(&dataBuku[i].Tarif)
					fmt.Println("\n Tarif berhasil diubah!")
				} else if pilihanEdit == 8 {
					fmt.Print("Masukkan ID Baru : ")
					fmt.Scan(&dataBuku[i].ID)

					fmt.Print("Masukkan Judul Baru : ")
					fmt.Scan(&dataBuku[i].Judul)

					fmt.Print("Masukkan Genre Baru : ")
					fmt.Scan(&dataBuku[i].Genre)

					fmt.Print("Edit Penulis Baru : ")
					fmt.Scan(&dataBuku[i].Genre)

					fmt.Print("Edit Tahun : ")
					fmt.Scan(&dataBuku[i].Tahun)

					fmt.Print("Edit Stok Buku : ")
					fmt.Scan(&dataBuku[i].Stok)

					fmt.Print("Edit Tarif Denda : ")
					fmt.Scan(&dataBuku[i].Tarif)

					fmt.Println("\n Seluruh data berhasil diubah!")
				} else if pilihanEdit == 9 {
					fmt.Println("\n Proses edit dibatalkan.")
				} else {
					fmt.Println("\n Pilihan tidak valid.")
				}
			}
		}
		if !ditemukan {
			fmt.Println("Tidak ada data yang ditemukan")
		}
		fmt.Print("Apakah ingin mengubah data lagi ? (y/n) : ")
		fmt.Scan(&ulang)
	}
}

func hapusBuku(dataBuku *[MAXBUKU]Buku, jumlahBuku *int) {
	var cariID int
	var idx int
	var ditemukan bool
	var ulang string

	for ulang != "n" && ulang != "no" {
		ditemukan = false

		fmt.Println()
		fmt.Println("===== [ HAPUS BUKU ] =====")
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
		fmt.Print("Apakah ingin menghapus lagi? (y/n) : ")
		fmt.Scan(&ulang)
	}
}

func pinjamBuku(dataBuku *[MAXBUKU]Buku, jumlahBuku int, dataPeminjam *[MAXPINJAM]Peminjaman, jumlahPinjam *int) {
	var ulang string
	var pilihan int
	var cariIDBuku int
	var cariJudul string

	for ulang != "n" && ulang != "no" {
		fmt.Println("\n===== [ PINJAM BUKU ] =====")
		fmt.Println("1. Cari Berdasarkan Judul")
		fmt.Println("2. Cari Berdasarkan ID")
		fmt.Println("3. Kembali ke menu utama")
		fmt.Print("Pilih jenis pencarian : ")
		fmt.Scan(&pilihan)

		var idxBuku int = -1

		if pilihan == 1 {
			fmt.Print("Masukkan Judul Buku yang ingin dipinjam : ")
			fmt.Scan(&cariJudul)
			idxBuku = sequentialSearchJudul(dataBuku, jumlahBuku, cariJudul)
		} else if pilihan == 2 {
			fmt.Print("Masukkan ID Buku yang akan dipinjam : ")
			fmt.Scan(&cariIDBuku)
			urutkanBukuByID(dataBuku, jumlahBuku)
			idxBuku = binarySearchID(dataBuku, jumlahBuku, cariIDBuku)
		} else if pilihan == 3 {
			ulang = "n"
		} else {
			fmt.Println("\nPilihan pencarian tidak valid")
		}

		if idxBuku != -1 {
			if dataBuku[idxBuku].Stok > 0 {
				fmt.Print("ID Peminjaman : ")
				fmt.Scan(&dataPeminjam[*jumlahPinjam].IDPinjam)

				fmt.Print("Nama Peminjam : ")
				fmt.Scan(&dataPeminjam[*jumlahPinjam].NamaPeminjam)

				fmt.Print("Tanggal Pinjam (cth: 12/10/2026) >> ")
				fmt.Scan(&dataPeminjam[*jumlahPinjam].TanggalPinjam)

				fmt.Print("Tanggal Kembali (cth: 15/10/2026) >> ")
				fmt.Scan(&dataPeminjam[*jumlahPinjam].TanggalKembali)

				fmt.Print("Lama Pinjam (hari) : ")
				fmt.Scan(&dataPeminjam[*jumlahPinjam].LamaPinjam)

				dataPeminjam[*jumlahPinjam].IDBuku = dataBuku[idxBuku].ID
				dataPeminjam[*jumlahPinjam].Denda = 0

				totalBiaya := dataPeminjam[*jumlahPinjam].LamaPinjam * dataBuku[idxBuku].Tarif
				fmt.Printf("Total Biaya Peminjaman: Rp %d\n", totalBiaya)

				dataBuku[idxBuku].Stok--
				dataBuku[idxBuku].Dipinjam++

				*jumlahPinjam++
				fmt.Println(" \nBuku berhasil dipinjam")
			} else {
				fmt.Println("\nMohon maaf, stok buku ini sedang kosong!")
			}
		}
		fmt.Print("Masukkan data peminjam lagi ? (yes/no) : ")
		fmt.Scan(&ulang)
	}
}

func kembalikanBuku(dataBuku *[MAXBUKU]Buku, jumlahBuku int, dataPeminjam *[MAXPINJAM]Peminjaman, jumlahPinjam *int) {
	var ulang string
	var cariIDPinjam int
	var aktual int

	for ulang != "n" && ulang != "no" {
		fmt.Println("\n===== [ KEMBALIKAN BUKU ] =====")
		fmt.Print("Masukkan ID peminjaman : ")
		fmt.Scan(&cariIDPinjam)

		idxPinjam := -1
		i := 0
		for i < *jumlahPinjam && idxPinjam == -1 {
			if dataPeminjam[i].IDPinjam == cariIDPinjam {
				idxPinjam = i
			}
			i++
		}

		if idxPinjam != -1 {
			urutkanBukuByID(dataBuku, jumlahBuku)
			idxBuku := binarySearchID(dataBuku, jumlahBuku, dataPeminjam[idxPinjam].IDBuku)

			fmt.Println("---------------------------------------------------")
			fmt.Println("Peminjam           : ", dataPeminjam[idxPinjam].NamaPeminjam)
			fmt.Println("Buku yang dipinjam : ", dataBuku[idxBuku].Judul)
			fmt.Println("Tanggal Pinjam     : ", dataPeminjam[idxPinjam].TanggalPinjam)
			fmt.Println("Tanggal Kembali    : ", dataPeminjam[idxPinjam].TanggalKembali)
			fmt.Println("Tenggat (Hari)     : ", dataPeminjam[idxPinjam].LamaPinjam)
			fmt.Println("---------------------------------------------------")

			fmt.Print("Berapa hari buku ini dikembalikan (aktual)? : ")
			fmt.Scan(&aktual)
			if aktual > dataPeminjam[idxPinjam].LamaPinjam {
				telat := aktual - dataPeminjam[idxPinjam].LamaPinjam
				denda := telat * dataBuku[idxBuku].Tarif
				fmt.Printf("Terlambat %d hari! Denda: Rp %d\n", telat, denda)
			} else {
				fmt.Println("Dikembalikan tepat waktu. Tidak ada denda.")
			}

			if idxBuku != -1 {
				dataBuku[idxBuku].Stok++
			}

			for i := idxPinjam; i < *jumlahPinjam-1; i++ {
				dataPeminjam[i] = dataPeminjam[i+1]
			}
			*jumlahPinjam--
			fmt.Println("Buku berhasil dikembalikan.")
		} else {
			fmt.Println("Data Peminjaman tidak ditemukan.")
		}
		fmt.Print("\nKembalikan buku yang lain? (yes/no) : ")
		fmt.Scan(&ulang)
	}
}

func tampilPeminjaman(dataPeminjam *[MAXPINJAM]Peminjaman, jumlahPinjam int, dataBuku *[MAXBUKU]Buku, jumlahBuku int) {
	fmt.Println("\n=== [ DAFTAR PEMINJAMAN AKTIF ] ===")
	if jumlahPinjam == 0 {
		fmt.Println("Belum ada data peminjaman yang sedang berlangsung.")
	} else {
		fmt.Println("======================================================================================================")
		fmt.Printf("| %-10s | %-15s | %-17s | %-12s | %-12s | %-11s |\n", "ID Pinjam", "Nama Peminjam", "Judul Buku", "Tgl Pinjam", "Tgl Kembali", "Lama Pinjam")
		fmt.Println("======================================================================================================")

		for i := 0; i < jumlahPinjam; i++ {
			idxBuku := -1
			j := 0
			for j < jumlahBuku && idxBuku == -1 {
				if dataBuku[j].ID == dataPeminjam[i].IDBuku {
					idxBuku = j
				}
				j++
			}

			judul := "Tidak Ditemukan"
			if idxBuku != -1 {
				judul = dataBuku[idxBuku].Judul
			}

			fmt.Printf("| %-10d | %-15.15s | %-17.17s | %-12.12s | %-12.12s | %-6d Hari |\n",
				dataPeminjam[i].IDPinjam,
				dataPeminjam[i].NamaPeminjam,
				judul,
				dataPeminjam[i].TanggalPinjam,
				dataPeminjam[i].TanggalKembali,
				dataPeminjam[i].LamaPinjam)
		}
		fmt.Println("======================================================================================")
	}
	fmt.Println("Tekan enter untuk kembali...")
	fmt.Scanln()
	fmt.Scanln()
}

func editPeminjaman(dataPeminjam *[MAXPINJAM]Peminjaman, jumlahPinjam int) {
	var cariID int
	var ditemukan bool
	var ulang string
	var pilihanEdit int

	for ulang != "n" && ulang != "no" {
		ditemukan = false
		fmt.Println("\n===== [ EDIT DATA PEMINJAMAN ] =====")
		fmt.Print("Masukkan ID Peminjaman yang ingin diedit : ")
		fmt.Scan(&cariID)

		for i := 0; i < jumlahPinjam; i++ {
			if dataPeminjam[i].IDPinjam == cariID {
				ditemukan = true
				fmt.Println("\nData Peminjaman Saat Ini:")
				fmt.Printf("1. Nama Peminjam : %s\n", dataPeminjam[i].NamaPeminjam)
				fmt.Printf("2. Tanggal Pinjam: %s\n", dataPeminjam[i].TanggalPinjam)
				fmt.Printf("3. Tanggal Kmbli : %s\n", dataPeminjam[i].TanggalKembali)
				fmt.Printf("4. Lama Pinjam   : %d Hari\n", dataPeminjam[i].LamaPinjam)
				fmt.Println("5. Batal Edit")
				fmt.Print("\nPilih bagian yang ingin diedit (1-5) >> ")
				fmt.Scan(&pilihanEdit)

				if pilihanEdit == 1 {
					fmt.Print("Masukkan Nama Baru : ")
					fmt.Scan(&dataPeminjam[i].NamaPeminjam)
					fmt.Println("\n Nama peminjam berhasil diubah!")
				} else if pilihanEdit == 2 {
					fmt.Print("Masukkan Tanggal Pinjam Baru (cth: 12/10/2026) >> ")
					fmt.Scan(&dataPeminjam[i].TanggalPinjam)
					fmt.Println("\n✅ Tanggal pinjam berhasil diubah!")
				} else if pilihanEdit == 3 {
					fmt.Print("Masukkan Tanggal Kembali Baru (cth: 15/10/2026) >> ")
					fmt.Scan(&dataPeminjam[i].TanggalKembali)
					fmt.Println("\n✅ Tanggal kembali berhasil diubah!")
				} else if pilihanEdit == 4 {
					fmt.Print("Masukkan Lama Pinjam Baru (Hari) >> ")
					fmt.Scan(&dataPeminjam[i].LamaPinjam)
					fmt.Println("\n✅ Lama pinjam berhasil diubah!")
				} else if pilihanEdit ==  5 {
					fmt.Println("edit dibatalkan")
				} else {
					fmt.Println("Pilihan tidak valid")
				}
			}
		}
		if !ditemukan {
			fmt.Println("Data Peminjaman tidak ditemukan.")
		}
		fmt.Print("Apakah ingin mengubah data lagi? (yes/no) : ")
		fmt.Scan(&ulang)
	}
}

func bukuFavorit(dataBuku *[MAXBUKU]Buku, jumlahBuku int) {
	var copyBuku [MAXBUKU]Buku

	fmt.Println("\n==== 5 BUKU FAVORIT ====")
	if jumlahBuku == 0 {
		fmt.Println("Data Buku masih kosong!")
	} else {
		for i := 0; i < jumlahBuku; i++ {
			copyBuku[i] = dataBuku[i]
		}

		for i := 0; i < jumlahBuku-1; i++ {
			idxMax := i
			for j := i + 1; j < jumlahBuku; j++ {
				if copyBuku[j].Dipinjam > copyBuku[idxMax].Dipinjam {
					idxMax = j
				}
			}
			temp := copyBuku[i]
			copyBuku[i] = copyBuku[idxMax]
			copyBuku[idxMax] = temp
		}

		fmt.Println("===========================================================================")
		fmt.Printf("| %-3s | %-25s | %-20s | %-12s |\n", "No", "Judul Buku", "Penulis", "Total Pinjam")
		fmt.Println("===========================================================================")

		limit := 5
		if jumlahBuku < 5 {
			limit = jumlahBuku
		}

		idxPrint := 0
		adaPeminjaman := false
		for idxPrint < limit && copyBuku[idxPrint].Dipinjam > 0 {
			adaPeminjaman = true
			judulTampil := copyBuku[idxPrint].Judul
			if len(judulTampil) > 24 {
				judulTampil = judulTampil[:22] + "..."
			}

			penulisTampil := copyBuku[idxPrint].Penulis
			if len(penulisTampil) > 19 {
				penulisTampil = penulisTampil[:17] + "..."
			}

			fmt.Printf("| %-3d | %-25s | %-20s | %-7d Kali |\n", idxPrint+1, judulTampil, penulisTampil, copyBuku[idxPrint].Dipinjam)
			idxPrint++
		}

		if !adaPeminjaman {
			fmt.Println("| Belum ada buku yang pernah dipinjam.                                    |")
		}
		fmt.Println("===========================================================================")
	}
	fmt.Println("\nTekan Enter untuk kembali ke menu...")
	fmt.Scanln()
	fmt.Scanln()
}

func menuSorting(dataBuku *[MAXBUKU]Buku, jumlahBuku int) {
	var pilih int

	for pilih != 5 {
		fmt.Println("\n===== Menu Sorting =====")
		fmt.Println("1. Judul Ascending (A-Z)")
		fmt.Println("2. Judul Descending (Z-A)")
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
