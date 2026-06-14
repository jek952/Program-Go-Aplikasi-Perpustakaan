# Flowchart Aplikasi Perpustakaan (Mermaid)
 
## 1. main, tampilAwal, tampilMenu
 
```mermaid
flowchart TD
    A([main]) --> B[tampilAwal]
    B --> C{Pilihan: 1=Masuk, 2=Exit}
    C -->|1| D[tampilMenu]
    C -->|2| E[Cetak ucapan terima kasih]
    E --> Z([Selesai])
 
    D --> F[Tampilkan 12 pilihan menu]
    F --> G{switch pilihan}
    G -->|1| H1[tambahBuku]
    G -->|2| H2[tampilBuku]
    G -->|3| H3[cariBuku]
    G -->|4| H4[editBuku]
    G -->|5| H5[hapusBuku]
    G -->|6| H6[menuSorting]
    G -->|7| H7[pinjamBuku]
    G -->|8| H8[kembalikanBuku]
    G -->|9| H9[tampilPeminjaman]
    G -->|10| H10[editPeminjaman]
    G -->|11| H11[bukuFavorit]
    G -->|12| H12[Kembali ke tampilAwal]
 
    H1 --> G
    H2 --> G
    H3 --> G
    H4 --> G
    H5 --> G
    H6 --> G
    H7 --> G
    H8 --> G
    H9 --> G
    H10 --> G
    H11 --> G
    H12 --> C
```
 
## 2. tambahBuku
 
```mermaid
flowchart TD
    A([tambahBuku]) --> B[Input ID, Judul, Penulis, Genre, Tahun, Stok, Tarif]
    B --> C[Set Dipinjam = 0, jumlahBuku++]
    C --> D[Cetak: Buku berhasil ditambahkan]
    D --> E{Tambah lagi? yes/no}
    E -->|yes| B
    E -->|no| F([Selesai])
```
 
## 3. tampilBuku
 
```mermaid
flowchart TD
    A([tampilBuku]) --> B{jumlahBuku == 0?}
    B -->|ya| C[Cetak: Belum ada data buku]
    B -->|tidak| D[Cetak header tabel]
    D --> E[Loop i = 0 .. jumlahBuku-1]
    E --> F[Cetak baris data buku ke-i]
    F --> E
    C --> G[Tunggu Enter]
    E --> G
    G --> H([Kembali ke menu])
```
 
## 4. cariBuku, sequentialSearchJudul, binarySearchID, urutkanBukuByID
 
```mermaid
flowchart TD
    A([cariBuku]) --> B{Pilih jenis pencarian}
    B -->|1: Judul| C[sequentialSearchJudul]
    B -->|2: ID| D[urutkanBukuByID lalu binarySearchID]
 
    C --> E{idx != -1?}
    D --> E
 
    E -->|ya| F[Tampilkan detail buku]
    E -->|tidak| G[Cetak: Buku tidak ditemukan]
 
    F --> H{Cari lagi? yes/no}
    G --> H
    H -->|yes| B
    H -->|no| I([Selesai])
 
    subgraph SEQ[sequentialSearchJudul]
        S1([Mulai]) --> S2[idx = -1]
        S2 --> S3[Loop i = 0..jumlahBuku-1]
        S3 --> S4{Judul i == cariJudul?}
        S4 -->|ya| S5[idx = i, stop loop]
        S4 -->|tidak| S3
        S5 --> S6([return idx])
        S3 --> S6
    end
 
    subgraph SORT[urutkanBukuByID - Insertion Sort]
        U1([Mulai]) --> U2[Loop i = 1..jumlahBuku-1]
        U2 --> U3[temp = data i, j = i-1]
        U3 --> U4{j>=0 dan ID j > temp.ID?}
        U4 -->|ya| U5[Geser data j ke j+1, j--]
        U5 --> U4
        U4 -->|tidak| U6[Simpan temp di j+1]
        U6 --> U2
        U2 --> U7([Selesai])
    end
 
    subgraph BIN[binarySearchID]
        B1([Mulai]) --> B2[low=0, high=jumlahBuku-1, idx=-1]
        B2 --> B3{low <= high dan idx==-1?}
        B3 -->|ya| B4[mid = low+high / 2]
        B4 --> B5{data mid .ID == cariID?}
        B5 -->|ya| B6[idx = mid]
        B5 -->|tidak, lebih kecil| B7[low = mid+1]
        B5 -->|tidak, lebih besar| B8[high = mid-1]
        B6 --> B3
        B7 --> B3
        B8 --> B3
        B3 -->|tidak| B9([return idx])
    end
```
 
## 5. editBuku
 
```mermaid
flowchart TD
    A([editBuku]) --> B[Input ID buku yang diedit]
    B --> C[Loop cari ID di dataBuku]
    C --> D{ID ditemukan?}
    D -->|tidak| E[Cetak: Tidak ada data ditemukan]
    D -->|ya| F[Tampilkan data buku saat ini]
    F --> G{Pilih field 1-9}
    G -->|1| H1[Edit ID]
    G -->|2| H2[Edit Judul]
    G -->|3| H3[Edit Penulis]
    G -->|4| H4[Edit Genre]
    G -->|5| H5[Edit Tahun]
    G -->|6| H6[Edit Stok]
    G -->|7| H7[Edit Tarif]
    G -->|8| H8[Edit semua field sekaligus]
    G -->|9| H9[Batal edit]
    G -->|lainnya| H10[Pilihan tidak valid]
 
    H1 --> I
    H2 --> I
    H3 --> I
    H4 --> I
    H5 --> I
    H6 --> I
    H7 --> I
    H8 --> I
    H9 --> I
    H10 --> I
    E --> I
 
    I{Ubah data lagi? y/n}
    I -->|y| B
    I -->|n| J([Selesai])
```
 
## 6. hapusBuku
 
```mermaid
flowchart TD
    A([hapusBuku]) --> B[Input ID buku]
    B --> C[Loop cari ID di dataBuku, simpan idx]
    C --> D{Ditemukan?}
    D -->|tidak| E[Cetak: ID buku tidak ditemukan]
    D -->|ya| F[Shift semua elemen setelah idx ke kiri]
    F --> G[jumlahBuku--]
    G --> H[Cetak: Data berhasil dihapus]
    E --> I
    H --> I
    I{Hapus lagi? y/n}
    I -->|y| B
    I -->|n| J([Selesai])
```
 
## 7. pinjamBuku
 
```mermaid
flowchart TD
    A([pinjamBuku]) --> B{Pilih: 1=Judul, 2=ID, 3=Kembali}
    B -->|1| C[sequentialSearchJudul -> idxBuku]
    B -->|2| D[urutkanBukuByID + binarySearchID -> idxBuku]
    B -->|3| Z([Kembali ke menu])
 
    C --> E{idxBuku != -1?}
    D --> E
 
    E -->|tidak| H
    E -->|ya| F{Stok > 0?}
    F -->|tidak| G[Cetak: stok kosong]
    F -->|ya| I[Input ID Pinjam, Nama, Lama Pinjam]
    I --> J[Set IDBuku, Denda=0]
    J --> K[Hitung total biaya = LamaPinjam x Tarif]
    K --> L[Stok--, Dipinjam++, jumlahPinjam++]
    L --> M[Cetak: Buku berhasil dipinjam]
 
    G --> H
    M --> H
    H{Input lagi? yes/no}
    H -->|yes| B
    H -->|no| Z2([Selesai])
```
 
## 8. kembalikanBuku
 
```mermaid
flowchart TD
    A([kembalikanBuku]) --> B[Input ID peminjaman]
    B --> C[Loop cari idxPinjam di dataPeminjam]
    C --> D{Ditemukan?}
    D -->|tidak| E[Cetak: Data peminjaman tidak ditemukan]
    D -->|ya| F[urutkanBukuByID + binarySearchID -> idxBuku]
    F --> G[Tampilkan info peminjam, judul, tenggat]
    G --> H[Input hari aktual pengembalian]
    H --> I{aktual > LamaPinjam?}
    I -->|ya| J[Hitung telat dan denda = telat x Tarif]
    I -->|tidak| K[Cetak: tepat waktu, tidak ada denda]
    J --> L
    K --> L
    L[Stok idxBuku ++]
    L --> M[Shift dataPeminjam, jumlahPinjam--]
    M --> N[Cetak: Buku berhasil dikembalikan]
 
    E --> O
    N --> O
    O{Kembalikan lagi? yes/no}
    O -->|yes| B
    O -->|no| P([Selesai])
```
 
## 9. tampilPeminjaman & editPeminjaman
 
```mermaid
flowchart TD
    A([tampilPeminjaman]) --> B{jumlahPinjam == 0?}
    B -->|ya| C[Cetak: Belum ada peminjaman aktif]
    B -->|tidak| D[Cetak header tabel]
    D --> E[Loop i = 0..jumlahPinjam-1]
    E --> F[Cari judul buku via IDBuku]
    F --> G[Cetak baris: ID Pinjam, Nama, Judul, Lama]
    G --> E
    C --> H[Tunggu Enter]
    E --> H
    H --> I([Kembali ke menu])
```
 
```mermaid
flowchart TD
    A([editPeminjaman]) --> B[Input ID peminjaman]
    B --> C[Loop cari ID di dataPeminjam]
    C --> D{Ditemukan?}
    D -->|tidak| E[Cetak: Data tidak ditemukan]
    D -->|ya| F[Tampilkan Nama dan Lama Pinjam saat ini]
    F --> G{Pilih 1-3}
    G -->|1| H1[Edit Nama Peminjam]
    G -->|2| H2[Edit Lama Pinjam]
    G -->|3| H3[Batal edit]
    G -->|lainnya| H4[Pilihan tidak valid]
 
    H1 --> I
    H2 --> I
    H3 --> I
    H4 --> I
    E --> I
 
    I{Ubah data lagi? yes/no}
    I -->|yes| B
    I -->|no| J([Selesai])
```
 
## 10. menuSorting & 4 fungsi sort
 
```mermaid
flowchart TD
    A([menuSorting]) --> B{Pilih 1-5}
    B -->|1| C[selectionSortJudulAsc]
    B -->|2| D[selectionSortJudulDesc]
    B -->|3| E[insertionSortStokAsc]
    B -->|4| F[insertionSortStokDesc]
    B -->|5| Z([Kembali ke menu])
 
    C --> G[tampilBuku hasil sort]
    D --> G
    E --> G
    F --> G
    G --> A
 
    subgraph SEL[selectionSortJudulAsc / Desc]
        S1([Mulai]) --> S2[Loop i = 0..n-1]
        S2 --> S3[idx = i]
        S3 --> S4[Loop j = i+1..n-1]
        S4 --> S5{Judul j lebih kecil -atau- lebih besar dari Judul idx?}
        S5 -->|ya| S6[idx = j]
        S6 --> S4
        S5 -->|tidak| S4
        S4 --> S7[Tukar data i dan data idx]
        S7 --> S2
        S2 --> S8([Selesai])
    end
 
    subgraph INS[insertionSortStokAsc / Desc]
        I1([Mulai]) --> I2[Loop i = 1..n-1]
        I2 --> I3[temp = data i, j = i-1]
        I3 --> I4{j>=0 dan Stok j lebih besar -atau- lebih kecil dari temp.Stok?}
        I4 -->|ya| I5[Geser data j ke j+1, j--]
        I5 --> I4
        I4 -->|tidak| I6[Simpan temp di j+1]
        I6 --> I2
        I2 --> I7([Selesai])
    end
```
 
## 11. bukuFavorit
 
```mermaid
flowchart TD
    A([bukuFavorit]) --> B{jumlahBuku == 0?}
    B -->|ya| C[Cetak: Data buku masih kosong]
    B -->|tidak| D[Salin dataBuku ke copyBuku]
    D --> E[Selection sort copyBuku berdasarkan Dipinjam, descending]
    E --> F[limit = min 5, jumlahBuku]
    F --> G[Cetak header tabel]
    G --> H[Loop idxPrint < limit dan Dipinjam > 0]
    H --> I[Cetak No, Judul, Penulis, Total Pinjam]
    I --> H
    H --> J{Ada peminjaman?}
    J -->|tidak| K[Cetak: Belum ada buku yang pernah dipinjam]
    J -->|ya| L[Tutup tabel]
    K --> L
    C --> M
    L --> M[Tunggu Enter]
    M --> N([Kembali ke menu])
```