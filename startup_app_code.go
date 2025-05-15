package main

import (
	"fmt"
	"strconv"
)

type AnggotaTim struct {
	Nama  string
	Peran string
}

type Startup struct {
	ID             int
	Nama           string
	BidangUsaha    string
	TahunBerdiri   int
	TotalPendanaan float64
	Anggota        []AnggotaTim
}

var startups []Startup
var idCounter = 1

func main() {
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Startup")
		fmt.Println("2. Tambah Anggota Tim")
		fmt.Println("3. Tampilkan Semua Startup")
		fmt.Println("4. Keluar")
		fmt.Println("5. Hapus Startup")
		fmt.Print("Pilih nomor menu: ")

		var pilihan string
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			tambahStartup()
		case "2":
			tambahAnggotaTim()
		case "3":
			tampilkanSemuaStartup()
		case "4":
			fmt.Println("Keluar dari program.")
			return
		case "5":
			hapusStartup()
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

func tambahStartup() {
	var nama, bidang, tahunStr, pendanaanStr string

	fmt.Print("Masukkan nama startup: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan bidang usaha: ")
	fmt.Scanln(&bidang)
	fmt.Print("Masukkan tahun berdiri: ")
	fmt.Scanln(&tahunStr)
	fmt.Print("Masukkan total pendanaan: ")
	fmt.Scanln(&pendanaanStr)

	tahun, err := strconv.Atoi(tahunStr)
	if err != nil {
		fmt.Println("Tahun berdiri harus angka!")
		return
	}
	pendanaan, err := strconv.ParseFloat(pendanaanStr, 64)
	if err != nil {
		fmt.Println("Total pendanaan harus angka!")
		return
	}

	startup := Startup{
		ID:             idCounter,
		Nama:           nama,
		BidangUsaha:    bidang,
		TahunBerdiri:   tahun,
		TotalPendanaan: pendanaan,
		Anggota:        []AnggotaTim{},
	}

	startups = append(startups, startup)
	idCounter++

	fmt.Println("Startup berhasil ditambahkan!")
}

func tambahAnggotaTim() {
	if len(startups) == 0 {
		fmt.Println("Belum ada startup, silakan tambah dulu.")
		return
	}

	var idStr, namaAnggota, peranAnggota string

	fmt.Print("Masukkan ID startup untuk tambah anggota: ")
	fmt.Scanln(&idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID harus angka!")
		return
	}

	index := -1
	for i, s := range startups {
		if s.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Startup tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan nama anggota: ")
	fmt.Scanln(&namaAnggota)
	fmt.Print("Masukkan peran anggota: ")
	fmt.Scanln(&peranAnggota)

	startups[index].Anggota = append(startups[index].Anggota, AnggotaTim{
		Nama:  namaAnggota,
		Peran: peranAnggota,
	})

	fmt.Println("Anggota berhasil ditambahkan ke startup", startups[index].Nama)
}

func tampilkanSemuaStartup() {
	if len(startups) == 0 {
		fmt.Println("Belum ada data startup.")
		return
	}

	for _, s := range startups {
		fmt.Println("-----------------------------")
		fmt.Println("ID:", s.ID)
		fmt.Println("Nama:", s.Nama)
		fmt.Println("Bidang Usaha:", s.BidangUsaha)
		fmt.Println("Tahun Berdiri:", s.TahunBerdiri)
		fmt.Printf("Total Pendanaan: Rp %.0f\n", s.TotalPendanaan)

		if len(s.Anggota) == 0 {
			fmt.Println("Anggota Tim: (Belum ada anggota)")
		} else {
			fmt.Println("Anggota Tim:")
			for _, a := range s.Anggota {
				fmt.Println(" -", a.Nama, "(", a.Peran, ")")
			}
		}
	}
	fmt.Println("-----------------------------")
}

func hapusStartup() {
	if len(startups) == 0 {
		fmt.Println("Belum ada startup untuk dihapus.")
		return
	}

	var idStr string
	fmt.Print("Masukkan ID startup yang ingin dihapus: ")
	fmt.Scanln(&idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID harus angka!")
		return
	}

	index := -1
	for i, s := range startups {
		if s.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Startup tidak ditemukan.")
		return
	}

	startups = append(startups[:index], startups[index+1:]...)
	fmt.Println("Startup berhasil dihapus.")
}
