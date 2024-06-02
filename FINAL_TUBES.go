package main

import (
	"fmt"
	"strings"
)

type PaketMCU struct {
	idMCU, namaPaket string
	harga            int
}

type Pasien struct {
	idPasien, nama, tglMCU, idMCU string
}

var dataPAKET [100]PaketMCU
var dataPASIEN [100]Pasien

func inputPaket() PaketMCU {
	var p PaketMCU
	fmt.Print("Masukkan ID Paket MCU: ")
	fmt.Scan(&p.idMCU)
	fmt.Print("Masukkan Nama Paket: ")
	fmt.Scan(&p.namaPaket)
	fmt.Print("Masukkan Harga Paket: ")
	fmt.Scan(&p.harga)
	return p
}

func inputPasien() Pasien {
	var p Pasien
	fmt.Print("Masukkan ID Pasien: ")
	fmt.Scan(&p.idPasien)
	fmt.Print("Masukkan Nama Pasien: ")
	fmt.Scan(&p.nama)
	fmt.Print("Masukkan ID Paket MCU: ")
	fmt.Scan(&p.idMCU)
	fmt.Print("Masukkan Tanggal MCU (dd-mm-yyyy): ")
	fmt.Scan(&p.tglMCU)

	return p
}

func tambahPaket(data *[100]PaketMCU, p PaketMCU, jumlah int) int {
	data[jumlah] = p
	return jumlah + 1
}

func tambahPasien(data *[100]Pasien, p Pasien, jumlah int) int {
	data[jumlah] = p
	return jumlah + 1
}

func ubahPaket(data *[100]PaketMCU, jumlah int) {
	var id string
	fmt.Print("Masukkan ID Paket yang akan diubah: ")
	fmt.Scan(&id)
	for i := 0; i < jumlah; i++ {
		if data[i].idMCU == id {
			fmt.Println("Paket ditemukan. Masukkan data baru.")
			data[i] = inputPaket()
			return
		}
	}
	fmt.Println("Paket tidak ditemukan.")
}

func ubahPasien(data *[100]Pasien, jumlah int) {
	var id string
	fmt.Print("Masukkan ID Pasien yang akan diubah: ")
	fmt.Scan(&id)
	for i := 0; i < jumlah; i++ {
		if data[i].idPasien == id {
			fmt.Println("Pasien ditemukan. Masukkan data baru.")
			data[i] = inputPasien()
			return
		}
	}
	fmt.Println("Pasien tidak ditemukan.")
}

func hapusPaket(data *[100]PaketMCU, jumlah *int) {
	var id string
	fmt.Print("Masukkan ID Paket yang akan dihapus: ")
	fmt.Scan(&id)
	for i := 0; i < *jumlah; i++ {
		if data[i].idMCU == id {
			for j := i; j < *jumlah-1; j++ {
				data[j] = data[j+1]
			}
			*jumlah--
			fmt.Println("Paket dihapus.")
			return
		}
	}
	fmt.Println("Paket tidak ditemukan.")
}

func hapusPasien(data *[100]Pasien, jumlah *int) {
	var id string
	fmt.Print("Masukkan ID Pasien yang akan dihapus: ")
	fmt.Scan(&id)
	for i := 0; i < *jumlah; i++ {
		if data[i].idPasien == id {
			for j := i; j < *jumlah-1; j++ {
				data[j] = data[j+1]
			}
			*jumlah--
			fmt.Println("Pasien dihapus.")
			return
		}
	}
	fmt.Println("Pasien tidak ditemukan.")
}

func seqSearch_PAKET(data [100]PaketMCU, jumlah int, id string) {
	for i := 0; i < jumlah; i++ {
		if data[i].idMCU == id {
			fmt.Printf("Paket ditemukan: %v\n", data[i])
			return
		}
	}
	fmt.Println("Paket tidak ditemukan.")
}

func seqSearch_PASIEN(data [100]Pasien, jumlah int, id string) {
	for i := 0; i < jumlah; i++ {
		if data[i].idPasien == id {
			fmt.Printf("Pasien ditemukan: %v\n", data[i])
			return
		}
	}
	fmt.Println("Pasien tidak ditemukan.")
}

func selectSort_PAKET(data *[100]PaketMCU, jumlah int, ascending bool) {
	for i := 0; i < jumlah-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlah; j++ {
			if (ascending && data[j].idMCU < data[minIdx].idMCU) || (!ascending && data[j].idMCU > data[minIdx].idMCU) {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
}

func insertSort_PASIEN(data *[100]Pasien, jumlah int, ascending bool) {
	for i := 1; i < jumlah; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && ((ascending && data[j].idPasien > key.idPasien) || (!ascending && data[j].idPasien < key.idPasien)) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func tampilkanPaket(data [100]PaketMCU, jumlah int) {
	for i := 0; i < jumlah; i++ {
		fmt.Printf("ID: %s, Nama: %s, Harga: %d\n", data[i].idMCU, data[i].namaPaket, data[i].harga)
	}
}

func tampilkanPasien(data [100]Pasien, jumlah int) {
	for i := 0; i < jumlah; i++ {
		fmt.Printf("ID: %s, Nama: %s, Tanggal MCU: %s, ID Paket: %s\n",
			data[i].idPasien, data[i].nama, data[i].tglMCU, data[i].idMCU)
	}
}

func laporanPemasukanMCU(dataPAKET [100]PaketMCU, jumlahPAKET int, dataPASIEN [100]Pasien, jumlahPASIEN int) {
	var startDate, endDate string
	var total int

	fmt.Print("Masukkan tanggal mulai (dd-mm-yyyy): ")
	fmt.Scan(&startDate)
	fmt.Print("Masukkan tanggal akhir (dd-mm-yyyy): ")
	fmt.Scan(&endDate)

	for i := 0; i < jumlahPASIEN; i++ {
		if dataPASIEN[i].tglMCU >= startDate && dataPASIEN[i].tglMCU <= endDate {
			for j := 0; j < jumlahPAKET; j++ {
				if dataPASIEN[i].idMCU == dataPAKET[j].idMCU {
					total += dataPAKET[j].harga
				}
			}
		}
	}
	fmt.Printf("Total pemasukan dari tanggal %s hingga %s adalah: %d\n", startDate, endDate, total)
}

func cariPasienByPaket(data [100]Pasien, jumlah int, idPaket string) {
	for i := 0; i < jumlah; i++ {
		if data[i].idMCU == idPaket {
			fmt.Printf("ID: %s, Nama: %s, Tanggal MCU: %s, ID Paket: %s\n",
				data[i].idPasien, data[i].nama, data[i].tglMCU, data[i].idMCU)
		}
	}
}

func search_sub(s, sub string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(sub))
}

func cariPasienByNama(data [100]Pasien, jumlah int, nama string) {
	for i := 0; i < jumlah; i++ {
		if search_sub(data[i].nama, nama) {
			fmt.Printf("ID: %s, Nama: %s, Tanggal MCU: %s, ID Paket: %s\n",
				data[i].idPasien, data[i].nama, data[i].tglMCU, data[i].idMCU)
		}
	}
}

func main() {
	var jumlahPAKET, jumlahPASIEN int

	for {
		var pilihan int
		fmt.Println("------------------------------------------------------")
		fmt.Println("|                       MENU                         |")
		fmt.Println("------------------------------------------------------")
		fmt.Println("1. Tambah Paket MCU                                   ")
		fmt.Println("2. Tambah Pasien                                      ")
		fmt.Println("3. Ubah Paket MCU                                     ")
		fmt.Println("4. Ubah Pasien                                        ")
		fmt.Println("5. Hapus Paket MCU                                    ")
		fmt.Println("6. Hapus Pasien                                       ")
		fmt.Println("7. Cari Paket MCU                                     ")
		fmt.Println("8. Cari Pasien                                        ")
		fmt.Println("9. Tampilkan Semua Paket MCU                          ")
		fmt.Println("10. Tampilkan Semua Pasien                            ")
		fmt.Println("11. Urutkan Paket MCU                                 ")
		fmt.Println("12. Urutkan Pasien                                    ")
		fmt.Println("13. Laporan Pemasukan MCU                             ")
		fmt.Println("14. Cari Pasien Berdasarkan Paket                     ")
		fmt.Println("15. Cari Pasien Berdasarkan Nama                      ")
		fmt.Println("0. Keluar                                            ")
		fmt.Println("------------------------------------------------------")

		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			paket := inputPaket()
			jumlahPAKET = tambahPaket(&dataPAKET, paket, jumlahPAKET)
		case 2:
			pasien := inputPasien()
			jumlahPASIEN = tambahPasien(&dataPASIEN, pasien, jumlahPASIEN)
		case 3:
			ubahPaket(&dataPAKET, jumlahPAKET)
		case 4:
			ubahPasien(&dataPASIEN, jumlahPASIEN)
		case 5:
			hapusPaket(&dataPAKET, &jumlahPAKET)
		case 6:
			hapusPasien(&dataPASIEN, &jumlahPASIEN)
		case 7:
			var id string
			fmt.Print("Masukkan ID Paket yang dicari: ")
			fmt.Scan(&id)
			seqSearch_PAKET(dataPAKET, jumlahPAKET, id)
		case 8:
			var id string
			fmt.Print("Masukkan ID Pasien yang dicari: ")
			fmt.Scan(&id)
			seqSearch_PASIEN(dataPASIEN, jumlahPASIEN, id)
		case 9:
			tampilkanPaket(dataPAKET, jumlahPAKET)
		case 10:
			tampilkanPasien(dataPASIEN, jumlahPASIEN)
		case 11:
			var asc bool
			fmt.Print("Urutkan secara ascending (true/false)? ")
			fmt.Scan(&asc)
			selectSort_PAKET(&dataPAKET, jumlahPAKET, asc)
		case 12:
			var asc bool
			fmt.Print("Urutkan secara ascending (true/false)? ")
			fmt.Scan(&asc)
			insertSort_PASIEN(&dataPASIEN, jumlahPASIEN, asc)
		case 13:
			laporanPemasukanMCU(dataPAKET, jumlahPAKET, dataPASIEN, jumlahPASIEN)
		case 14:
			var idPaket string
			fmt.Print("Masukkan ID Paket: ")
			fmt.Scan(&idPaket)
			cariPasienByPaket(dataPASIEN, jumlahPASIEN, idPaket)
		case 15:
			var nama string
			fmt.Print("Masukkan Nama Pasien: ")
			fmt.Scan(&nama)
			cariPasienByNama(dataPASIEN, jumlahPASIEN, nama)
		case 0:
			fmt.Println("Sukses keluar dari program.")
			fmt.Println("Terimakasih telah menggunakan!")
			fmt.Println()
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
