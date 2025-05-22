package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
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
		fmt.Println("\n=== APLIKASI MANAJEMEN STARTUP ===")
		fmt.Println("1. Tambah Startup")
		fmt.Println("2. Tambah Anggota Tim")
		fmt.Println("3. Tampilkan Semua Startup")
		fmt.Println("4. Cari Startup")
		fmt.Println("5. Urutkan Startup")
		fmt.Println("6. Laporan Startup per Bidang Usaha")
		fmt.Println("7. Hapus Startup")
		fmt.Println("8. Keluar")
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
			cariStartup()
		case "5":
			urutkanStartup()
		case "6":
			laporanPerBidang()
		case "7":
			hapusStartup()
		case "8":
			fmt.Println("Keluar dari program.")
			return
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

func cariStartup() {
	if len(startups) == 0 {
		fmt.Println("Belum ada data startup.")
		return
	}
	
	var kriteria, keyword string
	fmt.Println("Pilih kriteria pencarian:")
	fmt.Println("1. Nama")
	fmt.Println("2. Bidang Usaha")
	fmt.Print("Pilih (1/2): ")
	fmt.Scanln(&kriteria)
	
	fmt.Print("Masukkan kata kunci: ")
	fmt.Scanln(&keyword)
	
	found := false
	keyword = strings.ToLower(keyword)
	
	for _, s := range startups {
		var match bool
		if kriteria == "1" {
			match = strings.Contains(strings.ToLower(s.Nama), keyword)
		} else if kriteria == "2" {
			match = strings.Contains(strings.ToLower(s.BidangUsaha), keyword)
		}
		
		if match {
			if !found {
				fmt.Println("\n=== HASIL PENCARIAN ===")
				found = true
			}
			fmt.Println("-----------------------------")
			fmt.Println("ID:", s.ID)
			fmt.Println("Nama:", s.Nama)
			fmt.Println("Bidang Usaha:", s.BidangUsaha)
			fmt.Println("Tahun Berdiri:", s.TahunBerdiri)
			fmt.Printf("Total Pendanaan: Rp %.0f\n", s.TotalPendanaan)
		}
	}
	
	if !found {
		fmt.Println("Startup tidak ditemukan.")
	}
}

func urutkanStartup() {
	if len(startups) == 0 {
		fmt.Println("Belum ada data startup.")
		return
	}
	
	var kriteria, urutan string
	fmt.Println("Pilih kriteria pengurutan:")
	fmt.Println("1. Total Pendanaan")
	fmt.Println("2. Tahun Berdiri")
	fmt.Print("Pilih (1/2): ")
	fmt.Scanln(&kriteria)
	
	fmt.Println("Pilih urutan:")
	fmt.Println("1. Ascending (kecil ke besar)")
	fmt.Println("2. Descending (besar ke kecil)")
	fmt.Print("Pilih (1/2): ")
	fmt.Scanln(&urutan)
	
	// Buat salinan untuk pengurutan
	sortedStartups := make([]Startup, len(startups))
	copy(sortedStartups, startups)
	
	if kriteria == "1" { // Total Pendanaan
		if urutan == "1" { // Ascending
			sort.Slice(sortedStartups, func(i, j int) bool {
				return sortedStartups[i].TotalPendanaan < sortedStartups[j].TotalPendanaan
			})
		} else { // Descending
			sort.Slice(sortedStartups, func(i, j int) bool {
				return sortedStartups[i].TotalPendanaan > sortedStartups[j].TotalPendanaan
			})
		}
	} else if kriteria == "2" { // Tahun Berdiri
		if urutan == "1" { // Ascending
			sort.Slice(sortedStartups, func(i, j int) bool {
				return sortedStartups[i].TahunBerdiri < sortedStartups[j].TahunBerdiri
			})
		} else { // Descending
			sort.Slice(sortedStartups, func(i, j int) bool {
				return sortedStartups[i].TahunBerdiri > sortedStartups[j].TahunBerdiri
			})
		}
	}
	
	fmt.Println("\n=== HASIL PENGURUTAN ===")
	for _, s := range sortedStartups {
		fmt.Println("-----------------------------")
		fmt.Println("ID:", s.ID)
		fmt.Println("Nama:", s.Nama)
		fmt.Println("Bidang Usaha:", s.BidangUsaha)
		fmt.Println("Tahun Berdiri:", s.TahunBerdiri)
		fmt.Printf("Total Pendanaan: Rp %.0f\n", s.TotalPendanaan)
	}
}

func laporanPerBidang() {
	if len(startups) == 0 {
		fmt.Println("Belum ada data startup.")
		return
	}
	
	// Menghitung jumlah startup per bidang usaha
	bidangCount := make(map[string]int)
	
	for _, s := range startups {
		bidangCount[s.BidangUsaha]++
	}
	
	fmt.Println("\n=== LAPORAN STARTUP PER BIDANG USAHA ===")
	for bidang, jumlah := range bidangCount {
		fmt.Printf("%s: %d startup\n", bidang, jumlah)
	}
	fmt.Println("========================================")
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
