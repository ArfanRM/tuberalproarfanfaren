//Aplikasi ini membantu pengguna membuat resume dan surat lamaran kerja secara otomatis berdasarkan:
// Input pengguna: keahlian, pengalaman, pendidikan, pekerjaan yang dilamar.
// Algoritma AI: untuk memberikan saran penyusunan resume dan surat lamaran yang relevan.
// Fitur pencarian: untuk menemukan posisi kerja berdasarkan kata kunci.
// Fitur pengurutan: pekerjaan bisa diurut berdasarkan gaji atau relevansi.
package main

import "fmt"

type Pengalaman struct {
	Posisi     string
	Perusahaan string
	Tahun      int
}

type Pendidikan struct {
	Gelar      string
	Institusi  string
	TahunLulus int
}

type Pengguna struct {
	Nama               string
	Pengalaman         [10]Pengalaman
	JumlahPengalaman   int
	Keterampilan       [10]string
	JumlahSkill        int
	Pendidikan         [10]Pendidikan
	JumlahPendidikan   int
}

func (p *Pengguna) TambahPengalaman() {
	if p.JumlahPengalaman >= 10 {
		fmt.Println("Maksimal pengalaman tercapai.")
		return
	}
	var posisi, perusahaan string
	var tahun int
	fmt.Print("Masukkan posisi: ")
	fmt.Scanln(&posisi)
	fmt.Print("Masukkan perusahaan: ")
	fmt.Scanln(&perusahaan)
	fmt.Print("Masukkan tahun: ")
	fmt.Scanln(&tahun)

	p.Pengalaman[p.JumlahPengalaman] = Pengalaman{posisi, perusahaan, tahun}
	p.JumlahPengalaman++
}

func (p *Pengguna) TambahKeterampilan() {
	if p.JumlahSkill >= 10 {
		fmt.Println("Maksimal keterampilan tercapai.")
		return
	}
	var skill string
	fmt.Print("Masukkan keterampilan: ")
	fmt.Scanln(&skill)
	p.Keterampilan[p.JumlahSkill] = skill
	p.JumlahSkill++
}

func (p *Pengguna) TambahPendidikan() {
	if p.JumlahPendidikan >= 10 {
		fmt.Println("Maksimal pendidikan tercapai.")
		return
	}
	var gelar, institusi string
	var tahun int
	fmt.Print("Masukkan gelar: ")
	fmt.Scanln(&gelar)
	fmt.Print("Masukkan institusi: ")
	fmt.Scanln(&institusi)
	fmt.Print("Masukkan tahun lulus: ")
	fmt.Scanln(&tahun)
	p.Pendidikan[p.JumlahPendidikan] = Pendidikan{gelar, institusi, tahun}
	p.JumlahPendidikan++
}

func main() {
	var user Pengguna

	fmt.Print("Masukkan nama pengguna: ")
	fmt.Scanln(&user.Nama)

	fmt.Println("\n--- Tambah Pengalaman Kerja ---")
	user.TambahPengalaman()

	fmt.Println("\n--- Tambah Keterampilan ---")
	user.TambahKeterampilan()

	fmt.Println("\n--- Tambah Pendidikan ---")
	user.TambahPendidikan()

	fmt.Println("\n--- Data yang Diinput ---")
	fmt.Println("Nama:", user.Nama)

	fmt.Println("Pengalaman:")
	for i := 0; i < user.JumlahPengalaman; i++ {
		p := user.Pengalaman[i]
		fmt.Println("-", p.Posisi, "di", p.Perusahaan, "tahun", p.Tahun)
	}

	fmt.Println("Keterampilan:")
	for i := 0; i < user.JumlahSkill; i++ {
		fmt.Println("-", user.Keterampilan[i])
	}

	fmt.Println("Pendidikan:")
	for i := 0; i < user.JumlahPendidikan; i++ {
		edu := user.Pendidikan[i]
		fmt.Println("-", edu.Gelar, "dari", edu.Institusi, "lulus tahun", edu.TahunLulus)
	}
}
func panjang(s string) int {
	i := 0
	for ; i < 1000; i++ {
		if s[i:i+1] == "" {
			return i
		}
	}
	return i
}

// Ubah huruf ke huruf kecil
func toLower(s string) string {
	var hasil string
	i := 0
	for ; i < 1000; i++ {
		ch := s[i : i+1]
		if ch == "" {
			return hasil
		}
		kode := ch[0]
		if kode >= 'A' && kode <= 'Z' {
			kode += 32
		}
		hasil += string(kode)
	}
	return hasil
}
func contains(teks string, substr string) bool {
	n := panjang(teks)
	m := panjang(substr)
	i := 0
	for ; i <= n-m; i++ {
		j := 0
		sama := 1
		for ; j < m; j++ {
			if teks[i+j:i+j+1] != substr[j:j+1] {
				sama = 0
			}
		}
		if sama == 1 {
			return true
		}
	}
	return false
}
func panjang(s string) int {
	i := 0
	for ; i < 1000; i++ {
		if s[i:i+1] == "" {
			return i
		}
	}
	return i
}

// Ubah ke huruf kecil
func toLower(s string) string {
	var hasil string
	i := 0
	for ; i < 1000; i++ {
		ch := s[i : i+1]
		if ch == "" {
			return hasil
		}
		kode := ch[0]
		if kode >= 'A' && kode <= 'Z' {
			kode += 32
		}
		hasil += string(kode)
	}
	return hasil
}

// Cek substring tanpa strings.Contains
func contains(teks string, substr string) bool {
	n := panjang(teks)
	m := panjang(substr)
	i := 0
	for ; i <= n-m; i++ {
		j := 0
		match := 1
		for ; j < m; j++ {
			if teks[i+j:i+j+1] != substr[j:j+1] {
				match = 0
			}
		}
		if match == 1 {
			return true
		}
	}
	return false
}

// Sequential Search
func sequentialSearch(template [10]string, resume string) {
	fmt.Println("\nHasil Sequential Search:")
	i := 0
	for ; i < 10; i++ {
		if template[i] == "" {
			i = 10
		} else {
			if contains(resume, template[i]) {
				fmt.Println("- Ditemukan:", template[i])
			}
		}
	}
}

// Binary Search (harus urut)
func binarySearch(template [10]string, resume string) {
	fmt.Println("\nHasil Binary Search (manual urut):")
	kiri := 0
	kanan := 4 // hanya 5 data contoh

	for ; kiri <= kanan; {
		tengah := (kiri + kanan) / 2
		if contains(resume, template[tengah]) {
			fmt.Println("- Ditemukan:", template[tengah])
			kiri = kanan + 1 // keluar
		} else {
			// urutan manual alfabet
			if template[tengah] < "data analyst" {
				kiri = tengah + 1
			} else {
				kanan = tengah - 1
			}
		}
	}
}
