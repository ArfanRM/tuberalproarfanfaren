//Aplikasi ini membantu pengguna membuat resume dan surat lamaran kerja secara otomatis berdasarkan:
// Input pengguna: keahlian, pengalaman, pendidikan, pekerjaan yang dilamar.
// Algoritma AI: untuk memberikan saran penyusunan resume dan surat lamaran yang relevan.
// Fitur pencarian: untuk menemukan posisi kerja berdasarkan kata kunci.
// Fitur pengurutan: pekerjaan bisa diurut berdasarkan gaji atau relevansi.
package main

import "fmt"

type RiwayatKerja struct {
	Jabatan    string
	Organisasi string
	TahunMasuk int
}
type RiwayatStudi struct {
	Tingkat    string
	Kampus     string
	TahunLulus int
}
type Kandidat struct {
	NamaLengkap      string
	PengalamanKerja  [10]RiwayatKerja
	TotalPengalaman  int
	Keahlian         [10]string
	TotalKeahlian    int
	PendidikanFormal [10]RiwayatStudi
	TotalPendidikan  int
	NaskahResume     string
}

func hitungPanjang(teks string) int {
	i := 0
	for i != 1000 {
		if teks[i:i+1] == "" {
			return i
		}
		i++
	}
	return i
}
func keHurufKecil(teks string) string {
	var hasil string
	i := 0
	for i != 1000 {
		if teks[i:i+1] == "" {
			return hasil
		}
		kar := teks[i]
		if kar >= 'A' && kar <= 'Z' {
			kar += 32
		}
		hasil += string(kar)
		i++
	}
	return hasil
}
func mengandung(teks string, subteks string) bool {
	n := hitungPanjang(teks)
	m := hitungPanjang(subteks)
	i := 0
	for i <= n-m {
		j := 0
		sama := 1
		for j < m {
			if teks[i+j:i+j+1] != subteks[j:j+1] {
				sama = 0
			}
			j++
		}
		if sama == 1 {
			return true
		}
		i++
	}
	return false
}
func deteksiIndustri(teks string, daftar [20]string) string {
	i := 0
	for i < 20 {
		if daftar[i] != "" && mengandung(teks, daftar[i]) {
			return daftar[i]
		}
		i++
	}
	kiri := 0
	kanan := 19
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if daftar[tengah] == "" {
			kanan = tengah - 1
		} else if mengandung(teks, daftar[tengah]) {
			return daftar[tengah]
		} else if daftar[tengah] < "data analyst" {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return ""
}
func main() {
	var pelamar Kandidat
	fmt.Print("Masukkan nama lengkap: ")
	fmt.Scanln(&pelamar.NamaLengkap)
	menu := -1
	for menu != 0 {
		fmt.Println("\n===== MENU =====")
		fmt.Println("1. Tambah Riwayat Kerja")
		fmt.Println("2. Tambah Keahlian")
		fmt.Println("3. Tambah Pendidikan")
		fmt.Println("4. Edit Riwayat Kerja")
		fmt.Println("5. Hapus Riwayat Kerja")
		fmt.Println("6. Edit Keahlian")
		fmt.Println("7. Hapus Keahlian")
		fmt.Println("8. Edit Pendidikan")
		fmt.Println("9. Hapus Pendidikan")
		fmt.Println("10. Masukkan Resume + Surat Lamaran")
		fmt.Println("0. Analisa dan Lanjut")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&menu)
		if menu == 0 {
			fmt.Println("\n[Analisa Resume dan Surat Lamaran]")
			resume := keHurufKecil(pelamar.NaskahResume)
			kategori := [5][5]string{
				{"sql", "python", "tableau", "cloud", "excel"},
				{"jira", "powerbi", "google data studio", "confluence", "slack"},
				{"problem solving", "communication", "teamwork", "adaptability", "attention to detail"},
				{"impact", "insight", "value", "optimize", "transform"},
				{"meningkatkan", "efisiensi", "mengurangi", "hemat", "skala"},
			}
			judul := [5]string{
				"- Keahlian teknis yang belum disebutkan:",
				"- Tools/platform tidak disebutkan:",
				"- Soft skill belum terlihat:",
				"- Kata kerja kuat yang bisa digunakan:",
				"- Pencapaian kuantitatif kurang terlihat:",
			}
			for k := 0; k < 5; k++ {
				fmt.Println(judul[k])
				i := 0
				terlewat := 0
				for i < 5 {
					if !mengandung(resume, kategori[k][i]) {
						fmt.Println("  •", kategori[k][i])
						terlewat = 1
					}
					i++
				}
				if terlewat == 0 {
					fmt.Println("  (Sudah lengkap)")
				}
			}
			var daftarBidang [20]string = [20]string{
				"ai", "banking", "cloud computing", "cybersecurity", "data analyst",
				"devops", "e-commerce", "finance", "fmcg", "hr", "legal", "logistics",
				"manufacturing", "marketing", "network engineer", "product manager",
				"project manager", "software engineer", "startup", "telecommunication",
			}
			hasil := deteksiIndustri(resume, daftarBidang)
			if hasil != "" {
				fmt.Println("Cocok dengan bidang:", hasil)
			} else {
				fmt.Println("Tidak ditemukan kecocokan.")
			}
		}
	}
}
