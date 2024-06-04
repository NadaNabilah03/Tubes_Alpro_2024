package main

import "fmt"

const NMAX = 1000 // n
const MK = 1000   // m
const q = 15      // qz

type mahasiswa struct {
	nim, nama, kelas string
	matakuliah       [MK]matkul
	m                int
	transkrip        hasil
	tot_nilai        float64
	tot_sks          int
}

type matkul struct {
	kode, nama string
	sks        int
	nilai      point
}

type point struct {
	uas, uts hasil
	quiz     [q]hasil
	qz       int
}

type hasil struct {
	grade string
	angka float64
}

type tabArr [NMAX]mahasiswa

func main() {
	var n, nData int
	var data tabArr
	var pw string
	fmt.Println("     Selamat Datang! Silahkan masukkan password ")
	fmt.Print("     Password:")
	fmt.Scan(&pw)
	for pw != "TelkomUniversityDatabase" {
		fmt.Println("     Password yang anda masukkan salah.")
		fmt.Print("     Silahkan masukkan password yang benar:")
		fmt.Scan(&pw)
	}
	for n != 7 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("    Selamat Datang di Pusat Informasi Mahasiswa   ")
		fmt.Println("       Silahkan pilih sesuai kebutuhan Anda       ")
		fmt.Println("                                                  ")
		fmt.Println("    1. Data Mahasiswa                             ")
		fmt.Println("    2. Pengurutan data                            ")
		fmt.Println("    3. Data seluruh mahasiswa                     ")
		fmt.Println("    4. Data mata kuliah yang diambil mahasiswa    ")
		fmt.Println("    5. Pencarian data mahasiswa tertentu          ")
		fmt.Println("    6. Transkrip seluruh mahasiswa                ")
		fmt.Println("    7. Logout                                     ")
		fmt.Println("       Masukkan angka 1/2/3/4/5/6                 ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("       Input: ")
		fmt.Scan(&n)
		if n == 1 {
			dataMahasiswa(&data, &nData)
		} else if n == 2 {
			pencarianData(&data, nData)
		} else if n == 3 {
			tampilmhs(&data, nData)
		} else if n == 4 {
			tampilmkdiambilmhs(&data, nData)
		} else if n == 5 {
			cariDataMahasiswa(data, nData)
		} else if n == 6 {
			transkrip(&data, nData)
		}
	}
	fmt.Println("     Anda sudah logout")
	fmt.Print("     Terima kasih sudah menggunakan aplikasi ini!")
}

func dataMahasiswa(A *tabArr, n *int) {
	var pilihan int
	for pilihan != 5 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("          Apa yang sedang anda butuhkan?          ")
		fmt.Println("                                                  ")
		fmt.Println("       1. Input Data Mahasiswa                    ")
		fmt.Println("       2. Hapus Data Mahasiswa                    ")
		fmt.Println("       3. Edit Data Mahasiswa                     ")
		fmt.Println("       4. Data Matakuliah Mahasiswa               ")
		fmt.Println("       5. Kembali ke menu utama                   ")
		fmt.Println("       Masukkan angka 1/2/3/4/5                   ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("       Input: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			inputDataMahasiswa(A, n)
		} else if pilihan == 2 {
			hapusDataMahasiswa(A, n)
		} else if pilihan == 3 {
			editDataMahasiswa(A, n)
		} else if pilihan == 4 {
			DataMatakuliahMahasiswa(A, *n)
		} else if pilihan == 5 {
			fmt.Println("       Baik, anda akan segera menuju ke menu utama")
			fmt.Println("       Mohon tunggu sebentar")
			fmt.Println()
		}
	}
}

func inputDataMahasiswa(A *tabArr, n *int) {
	var jawaban string
	var nilai int
	fmt.Println()
	fmt.Print("Berapa yang ingin anda masukkan: ")
	fmt.Scan(&nilai)
	fmt.Println("Silahkan masukkan NIM, nama dengan format namaDepan_namaBelakang, dan kelas!")
	for i := 0; i < nilai; i++ {
		fmt.Println("Mahasiswa ", i+1)
		fmt.Print("Masukkan NIM: ")
		fmt.Scan(&A[*n].nim)
		fmt.Print("Masukkan Nama: ")
		fmt.Scan(&A[*n].nama)
		fmt.Print("Masukkan Kelas: ")
		fmt.Scan(&A[*n].kelas)
		*n += 1
	}
	fmt.Println("Data telah diinputkan!")
	fmt.Println("Apakah anda ingin input data lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		inputDataMahasiswa(A, n)
	}
}

func hapusDataMahasiswa(A *tabArr, n *int) {
	var nim, jawaban string
	fmt.Println()
	fmt.Println("Masukkan data NIM mahasiswa yang ingin dihapus:")
	fmt.Scan(&nim)
	a := findDataNim(*A, *n, nim)
	// fmt.Println(a)
	if a != -1 {
		fmt.Println("Data NIM Mahasiswa dengan NIM ", A[a].nim, " ditemukan")
		fmt.Println("Apakah anda ingin menghapus data mahasiswa tersebut?")
		fmt.Print("Ya/Tidak:")
		fmt.Scan(&jawaban)
		if jawaban == "ya" || jawaban == "Ya" || jawaban == "YA" {
			for a < *n {
				A[a] = A[a+1]
				a++
			}
			*n--
		}
		fmt.Println("Data telah dihapus!")
	} else {
		fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
	}
	fmt.Println("Apakah anda ingin hapus data lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		hapusDataMahasiswa(A, n)
	}
}

func editDataMahasiswa(A *tabArr, n *int) {
	var nim, jawaban string
	var a int
	fmt.Println()
	fmt.Println("Data Mahasiswa yang mana yang ingin anda edit?")
	fmt.Print("Masukkan NIM Mahasiswa yang ingin di edit: ")
	fmt.Scan(&nim)
	a = findDataNim(*A, *n, nim)
	if a != -1 {
		fmt.Println("Data Mahasiswa dengan NIM ", A[a].nim, " ditemukan")
		fmt.Println("Silahkan masukkan NIM, nama, dan kelas mahasiswa!")
		fmt.Print("Masukkan NIM: ")
		fmt.Scan(&A[a].nim)
		fmt.Print("Masukkan Nama: ")
		fmt.Scan(&A[a].nama)
		fmt.Print("Masukkan Kelas: ")
		fmt.Scan(&A[a].kelas)
		fmt.Println("Data telah diedit!")
	} else {
		fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
	}
	fmt.Println("Apakah anda ingin edit data mata kuliah lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		editDataMahasiswa(A, n)
	}
}

func DataMatakuliahMahasiswa(A *tabArr, n int) {
	var pilihan int
	for pilihan != 4 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("         Data apa yang ingin anda edit?           ")
		fmt.Println("                                                  ")
		fmt.Println("         1. Input / Edit Mata Kuliah              ")
		fmt.Println("         2. Nilai UTS dan UAS                     ")
		fmt.Println("         3. Nilai Quiz                            ")
		fmt.Println("         4. Kembali ke menu sebelumnya            ")
		fmt.Println("         Masukkan angka 1/2/3/4                   ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("       Input: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			mataKuliah(A, n)
		} else if pilihan == 2 {
			UtsUas(A, n)
		} else if pilihan == 3 {
			Quiz(A, n)
		} else if pilihan == 4 {
			fmt.Println("       Baik, anda akan segera menuju ke menu sebelumnya")
			fmt.Println("       Mohon tunggu sebentar")
			fmt.Println()
		}
	}
}

func mataKuliah(A *tabArr, n int) {
	var jawaban int
	for jawaban != 4 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("           Menu apa yang anda butuhkan?          ")
		fmt.Println("           1. Input Data Mata Kuliah              ")
		fmt.Println("           2. Edit Data Mata Kuliah               ")
		fmt.Println("           3. Hapus Data Mata Kuliah              ")
		fmt.Println("           4. Kembali ke menu sebelumnya          ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("    Input: ")
		fmt.Scan(&jawaban)
		if jawaban == 1 {
			inputMataKuliah(A, n)
		} else if jawaban == 2 {
			editMataKuliah(A, n)
		} else if jawaban == 3 {
			hapusDataMatkul(A, n)
		} else if jawaban == 4 {
			fmt.Println("       Baik, anda akan segera menuju ke menu sebelumnya")
			fmt.Println("       Mohon tunggu sebentar")
			fmt.Println()
		}
	}
}

func inputMataKuliah(A *tabArr, n int) {
	var nim, jawaban string
	var a, nilai int
	fmt.Println()
	fmt.Println("Data Mahasiswa yang mana yang ingin anda input?")
	fmt.Print("Masukkan NIM Mahasiswa: ")
	fmt.Scan(&nim)
	a = findDataNim(*A, n, nim)
	if a != -1 {
		fmt.Println("Data Mahasiswa dengan NIM ", A[a].nim, " ditemukan")
		fmt.Println("Berapa data mata kuliah yang ingin anda masukkan?")
		fmt.Print("Masukkan: ")
		fmt.Scan(&nilai)
		fmt.Println("Silahkan masukkan kode, nama, dan sks mata kuliah!")
		for i := 0; i < nilai; i++ {
			fmt.Println("Mata Kuliah ", i+1)
			fmt.Print("Masukkan Kode MK: ")
			fmt.Scan(&A[a].matakuliah[i].kode)
			fmt.Print("Masukkan Nama MK: ")
			fmt.Scan(&A[a].matakuliah[i].nama)
			fmt.Print("Masukkan SKS MK: ")
			fmt.Scan(&A[a].matakuliah[i].sks)
			A[a].m++
		}
		fmt.Println("Data telah diinputkan!")
	} else {
		fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
	}
	fmt.Println("Apakah anda ingin input data mata kuliah lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		inputMataKuliah(A, n)
	}
}

func editMataKuliah(A *tabArr, n int) {
	var nim, namkul, kodkul, jawaban, kode string
	var a, b, skskul int
	fmt.Println()
	fmt.Println("Data Mahasiswa yang mana yang ingin anda edit?")
	fmt.Print("Masukkan NIM Mahasiswa dan kode mata kuliah yang ingin di edit: ")
	fmt.Scan(&nim, &kode)
	a = findDataNim(*A, n, nim)
	if a != -1 {
		b = findDataMatkul(A[a], kode)
		if b != -1 {
			fmt.Println("Data Mahasiswa ditemukan")
			fmt.Println("Silahkan masukkan kode, nama, dan sks mata kuliah: ")
			fmt.Scan(&kodkul, &namkul, &skskul)
			A[a].matakuliah[b].kode = kodkul
			A[a].matakuliah[b].nama = namkul
			A[a].matakuliah[b].sks = skskul
			fmt.Println("Data telah diedit!")
		} else {
			fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
		}
	} else {
		fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
	}
	fmt.Println("Apakah anda ingin edit data mata kuliah lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		editMataKuliah(A, n)
	}
}

func hapusDataMatkul(A *tabArr, n int) {
	var nim, jawaban, kode string
	var a, b int
	fmt.Println()
	fmt.Println("Masukkan data NIM dan KODE mata kuliah mahasiswa yang ingin dihapus:")
	fmt.Scan(&nim, &kode)
	a = findDataNim(*A, n, nim)
	if a != -1 {
		b = findDataMatkul(A[a], kode)
		if b != -1 {
			fmt.Println("Data NIM dan mata kuliah Mahasiswa ditemukan")
			fmt.Println("Apakah anda ingin menghapus data mata kuliah mahasiswa tersebut?")
			fmt.Print("Ya/Tidak:")
			fmt.Scan(&jawaban)
			if jawaban == "ya" || jawaban == "Ya" || jawaban == "YA" {
				for b < A[a].m-1 {
					A[a].matakuliah[b] = A[a].matakuliah[b+1]
					b++
				}
				fmt.Println("Data telah dihapus!")
			}
		} else {
			fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
		}
	} else {
		fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
	}
	fmt.Println("Apakah anda ingin hapus data mata kuliah lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		hapusDataMatkul(A, n)
	}
}

func UtsUas(A *tabArr, n int) {
	var jawaban int
	for jawaban != 3 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("          Menu apa yang anda butuhkan?           ")
		fmt.Println("          1. Input / Edit Data UTS dan UAS        ")
		fmt.Println("          2. Hapus Data UTS dan UAS               ")
		fmt.Println("          3. Kembali ke menu sebelumnya           ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("    Input: ")
		fmt.Scan(&jawaban)
		if jawaban == 1 {
			inputUtsUas(A, n)
		} else if jawaban == 2 {
			hapusUtsUas(A, n)
		} else if jawaban == 3 {
			fmt.Println("       Baik, anda akan segera menuju ke menu sebelumnya")
			fmt.Println("       Mohon tunggu sebentar")
			fmt.Println()
		}
	}
}

func inputUtsUas(A *tabArr, n int) {
	var nim, kode, jawaban string
	var a, b int
	fmt.Println()
	fmt.Println("Data Mahasiswa yang mana yang ingin anda input?")
	fmt.Print("Masukkan NIM Mahasiswa dan kode Mata kuliah: ")
	fmt.Scan(&nim, &kode)
	a = findDataNim(*A, n, nim)
	if a != -1 {
		b = findDataMatkul(A[a], kode)
		if b != -1 {
			fmt.Println("Data Mahasiswa dan Mata Kuliah ditemukan")
			fmt.Println("Silahkan masukkan nilai UTS dan UAS beserta grade masing-masing terurut: ")
			fmt.Print("Masukkan nilai angka dan grade UTS: ")
			fmt.Scan(&A[a].matakuliah[b].nilai.uts.angka, &A[a].matakuliah[b].nilai.uts.grade)
			fmt.Print("Masukkan nilai angka dan grade UAS: ")
			fmt.Scan(&A[a].matakuliah[b].nilai.uas.angka, &A[a].matakuliah[b].nilai.uas.grade)
			fmt.Println("Data telah diinputkan!")
		} else {
			fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
		}
	}
	fmt.Println("Apakah anda ingin input/edit data mata kuliah lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		inputUtsUas(A, n)
	}
}

func hapusUtsUas(A *tabArr, n int) {
	var nim, kode, jawaban string
	var a, b int
	fmt.Println()
	fmt.Println("Data Mahasiswa yang mana yang ingin anda hapus?")
	fmt.Print("Masukkan NIM Mahasiswa dan kode Mata kuliah: ")
	fmt.Scan(&nim, &kode)
	a = findDataNim(*A, n, nim)
	if a != -1 {
		b = findDataMatkul(A[a], kode)
		if b != -1 {
			fmt.Println("Data NIM dan mata kuliah Mahasiswa ditemukan")
			fmt.Println("Apakah anda ingin menghapus data nilai UTS dan UAS mahasiswa tersebut?")
			fmt.Print("Ya/Tidak:")
			fmt.Scan(&jawaban)
			if jawaban == "ya" || jawaban == "Ya" || jawaban == "YA" {
				A[a].matakuliah[b].nilai.uts.angka = 0
				A[a].matakuliah[b].nilai.uts.grade = "-"
				A[a].matakuliah[b].nilai.uas.angka = 0
				A[a].matakuliah[b].nilai.uas.grade = "-"
			}
			fmt.Println("Data telah dihapus!")
		} else {
			fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
		}
	}
	fmt.Println("Apakah anda ingin hapus data mata kuliah lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		hapusUtsUas(A, n)
	}
}

func Quiz(A *tabArr, n int) {
	var jawaban int
	for jawaban != 4 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("         Menu apa yang anda butuhkan?            ")
		fmt.Println("         1. Input Data Quiz                       ")
		fmt.Println("         2. Edit Data Quiz                        ")
		fmt.Println("         3. Hapus Data Quiz                       ")
		fmt.Println("         4. Kembali ke menu sebelumnya            ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("    Input: ")
		fmt.Scan(&jawaban)
		if jawaban == 1 {
			inputQuiz(A, n)
		} else if jawaban == 2 {
			editQuiz(A, n)
		} else if jawaban == 3 {
			hapusQuiz(A, n)
		} else if jawaban == 4 {
			fmt.Println("Baik, anda akan segera menuju ke menu sebelumnya")
			fmt.Println("Mohon tunggu sebentar")
			fmt.Println()
		}
	}
}

func inputQuiz(A *tabArr, n int) {
	var a, b, c int
	var nim, kode, jawaban string
	fmt.Println()
	fmt.Println("Data Mahasiswa mana yang ingin anda input nilai quiznya?")
	fmt.Print("Masukkan NIM Mahasiswa dan kode Mata kuliah: ")
	fmt.Scan(&nim, &kode)
	a = findDataNim(*A, n, nim)
	if a != -1 {
		b = findDataMatkul(A[a], kode)
		if b != -1 {
			fmt.Println("Data Mahasiswa ditemukan")
			fmt.Print("Ada berapa quiz yang ingin di input nilainya?:")
			fmt.Scan(&c)
			fmt.Println("Silahkan masukkan nilai quiz beserta grade: ")
			if c > q {
				c = q
			}
			for i := 0; i < c; i++ {
				fmt.Println("Quiz ", i+1)
				fmt.Print("Masukkan nilai angka: ")
				fmt.Scan(&A[a].matakuliah[b].nilai.quiz[i].angka)
				fmt.Print("Masukkan nilai grade: ")
				fmt.Scan(&A[a].matakuliah[b].nilai.quiz[i].grade)
				A[a].matakuliah[b].nilai.qz++
			}
			fmt.Println("Data telah diinputkan!")
		} else {
			fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
		}
	}
	fmt.Println("Apakah anda ingin input data quiz lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		inputQuiz(A, n)
	}
}

func editQuiz(A *tabArr, n int) {
	var nim, kode, jawaban string
	var nilai, a, b int
	fmt.Println()
	fmt.Println("Data Mahasiswa yang mana yang ingin anda edit?")
	fmt.Print("Masukkan NIM Mahasiswa dan kode Mata kuliah: ")
	fmt.Scan(&nim, &kode)
	a = findDataNim(*A, n, nim)
	if a != -1 {
		b = findDataMatkul(A[a], kode)
		if b != -1 {
			fmt.Println("Data Mahasiswa dan MK ditemukan")
			fmt.Print("Silahkan masukkan quiz ke berapa yang ingin anda edit: ")
			fmt.Scan(&nilai)
			fmt.Println("Silahkan masukkan nilai quiz beserta grade: ")
			fmt.Print("Masukkan nilai angka: ")
			fmt.Scan(&A[a].matakuliah[b].nilai.quiz[nilai-1].angka)
			fmt.Print("Masukkan nilai grade: ")
			fmt.Scan(&A[a].matakuliah[b].nilai.quiz[nilai-1].grade)
			fmt.Println("Data telah diedit!")
		} else {
			fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
		}
	}
	fmt.Println("Apakah anda ingin edit data quiz lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		editQuiz(A, n)
	}
}

func hapusQuiz(A *tabArr, n int) {
	var nim, kode, jawaban string
	var a, b, quiz int
	fmt.Println()
	fmt.Println("Masukkan NIM Mahasiswa dan kode Mata kuliah yang ingin anda hapus nilai quiznya: ")
	fmt.Scan(&nim, &kode)
	a = findDataNim(*A, n, nim)
	if a != -1 {
		b = findDataMatkul(A[a], kode)
		if b != -1 {
			fmt.Println("Data Mahasiswa dan MK ditemukan")
			fmt.Print("Masukkan quiz ke berapa yang ingin anda hapus?: ")
			fmt.Scan(&quiz)
			if quiz-1 < A[a].matakuliah[b].nilai.qz {
				fmt.Println("Apakah anda ingin menghapus data mata kuliah mahasiswa tersebut?")
				fmt.Print("Ya/Tidak:")
				fmt.Scan(&jawaban)
				if jawaban == "ya" || jawaban == "Ya" || jawaban == "YA" {
					A[a].matakuliah[b].nilai.quiz[quiz-1].angka = 0
					A[a].matakuliah[b].nilai.quiz[quiz-1].grade = "-"
					fmt.Println("Data telah dihapus!")
				}
			} else {
				fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
			}
		} else {
			fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
		}
	} else {
		fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
	}
	fmt.Println("Apakah anda ingin hapus data mata kuliah lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		hapusQuiz(A, n)
	}
}

func findDataNim(A tabArr, n int, x string) int {
	var idx int
	idx = -1
	i := 0
	for i < n && idx == -1 {
		if x == A[i].nim {
			idx = i
		}
		i++
	}
	return idx
}

func findDataMatkul(mhs mahasiswa, x string) int {
	var idx int
	idx = -1
	i := 0
	for i < mhs.m && idx == -1 {
		if x == mhs.matakuliah[i].kode {
			idx = i
		}
		i++
	}
	return idx
}

func rataQuiz(kuis point) float64 {
	var rerataQuiz, total float64
	for i := 0; i < kuis.qz; i++ {
		total += kuis.quiz[i].angka
	}
	rerataQuiz = total / float64(kuis.qz)
	return rerataQuiz
}

func totalNilai(mhs mahasiswa) float64 {
	var rerataQuiz, total float64
	for i := 0; i < mhs.m; i++ {
		rerataQuiz = rataQuiz(mhs.matakuliah[i].nilai)
		total += (mhs.matakuliah[i].nilai.uts.angka + mhs.matakuliah[i].nilai.uas.angka + rerataQuiz) / float64(3)
	}
	total = total / float64(mhs.m)
	return total
}

func transkrip(A *tabArr, n int) {
	fmt.Printf("%15s %40s %10s %6s\n", "NIM", "NAMA", "NILAI", "GRADE")
	for i := 0; i < n; i++ {
		a := totalSks(A[i])
		A[i].tot_nilai = totalNilai(A[i])
		A[i].transkrip.angka = (A[i].tot_nilai * float64(a)) / float64(a)
		if A[i].transkrip.angka >= 80 {
			A[i].transkrip.grade = "A"
		} else if A[i].transkrip.angka >= 70 && A[i].transkrip.angka < 80 {
			A[i].transkrip.grade = "AB"
		} else if A[i].transkrip.angka >= 60 && A[i].transkrip.angka < 70 {
			A[i].transkrip.grade = "B"
		} else if A[i].transkrip.angka >= 55 && A[i].transkrip.angka < 60 {
			A[i].transkrip.grade = "BC"
		} else if A[i].transkrip.angka >= 45 && A[i].transkrip.angka < 55 {
			A[i].transkrip.grade = "C"
		} else if A[i].transkrip.angka >= 40 && A[i].transkrip.angka < 45 {
			A[i].transkrip.grade = "D"
		} else if A[i].transkrip.angka >= 0 && A[i].transkrip.angka < 40 {
			A[i].transkrip.grade = "E"
		}
		fmt.Printf("%15s %40s %10.2f %6s\n", A[i].nim, A[i].nama, A[i].transkrip.angka, A[i].transkrip.grade)
	}
	fmt.Println()
}

func totalSks(mhs mahasiswa) int {
	var total int
	for i := 0; i < mhs.m; i++ {
		total += mhs.matakuliah[i].sks
	}
	return total
}

func pencarianData(A *tabArr, n int) {
	var pilihan int
	for pilihan != 4 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("           Menu apa yang anda butuhkan?           ")
		fmt.Println("                                                  ")
		fmt.Println("           1. Data menurut Nilai                  ")
		fmt.Println("           2. Data menurut SKS                    ")
		fmt.Println("           3. Data menurut NIM                    ")
		fmt.Println("           4. Kembali ke menu sebelumnya          ")
		fmt.Println("           Masukkan angka 1/2/3/4                 ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("      Input:")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			dataMenurutNilai(A, n)
		} else if pilihan == 2 {
			dataMenurutSks(A, n)
		} else if pilihan == 3 {
			dataMenurutNim(A, n)
		} else if pilihan == 4 {
			fmt.Println("       Baik, anda akan segera menuju ke menu sebelumnya")
			fmt.Println("       Mohon tunggu sebentar")
			fmt.Println()
		}
	}
}

func dataMenurutNilai(A *tabArr, n int) {
	var pilihan int
	for pilihan != 3 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("    Jenis data terurut apa yang anda inginkan?    ")
		fmt.Println("                                                  ")
		fmt.Println("           1. Terurut membesar                    ")
		fmt.Println("           2. Terurut mengecil                    ")
		fmt.Println("           3. Kembali ke menu sebelumnya          ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("      Input:")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			nilaiUrutMembesar(A, n)
		} else if pilihan == 2 {
			nilaiUrutMengecil(A, n)
		} else if pilihan == 3 {
			fmt.Println("Baik, anda akan segera menuju ke menu sebelumnya")
			fmt.Println("Mohon tunggu sebentar")
			fmt.Println()
		}
	}
}

func nilaiUrutMembesar(A *tabArr, n int) {
	var i, j, idx int
	var t mahasiswa
	i = 1
	for i <= n-1 {
		idx = i - 1
		j = 1 + 1
		for j < n {
			if (A[idx].tot_nilai > A[j].tot_nilai) || (A[idx].tot_nilai == A[j].tot_nilai && A[idx].transkrip.angka > A[j].transkrip.angka) {
				idx = j
			}
			j++
		}
		t = A[idx]
		A[idx] = A[i-1]
		A[i-1] = t
		i++
	}
	tampilnilai(A, n)
}

func nilaiUrutMengecil(A *tabArr, n int) {
	var i, j, idx int
	var t mahasiswa
	i = 1
	for i <= n-1 {
		idx = i - 1
		j = 1 + 1
		for j < n {
			if (A[idx].tot_nilai < A[j].tot_nilai) || (A[idx].tot_nilai == A[j].tot_nilai && A[idx].transkrip.angka < A[j].transkrip.angka) {
				idx = j
			}
			j++
		}
		t = A[idx]
		A[idx] = A[i-1]
		A[i-1] = t
		i++
	}
	tampilnilai(A, n)
}

func dataMenurutSks(A *tabArr, n int) {
	var pilihan int
	for pilihan != 3 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("    Jenis data terurut apa yang anda inginkan?    ")
		fmt.Println("                                                  ")
		fmt.Println("           1. Terurut membesar                    ")
		fmt.Println("           2. Terurut mengecil                    ")
		fmt.Println("           3. Kembali ke menu sebelumnya          ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("      Input:")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			sksUrutMembesar(A, n)
		} else if pilihan == 2 {
			sksUrutMengecil(A, n)
		} else if pilihan == 3 {
			fmt.Println("Baik, anda akan segera menuju ke menu sebelumnya")
			fmt.Println("Mohon tunggu sebentar")
			fmt.Println()
		}
	}
}

func sksUrutMembesar(A *tabArr, n int) {
	var i, j, idx int
	var t mahasiswa
	i = 1
	for i <= n-1 {
		idx = i - 1
		j = 1 + 1
		A[idx].tot_sks = totalSks(A[idx])
		A[j].tot_sks = totalSks(A[j])
		for j < n {
			if (A[idx].tot_sks > A[j].tot_sks) || (A[idx].tot_sks == A[j].tot_sks && A[idx].transkrip.angka > A[j].transkrip.angka) {
				idx = j
			}
			j++
		}
		t = A[idx]
		A[idx] = A[i-1]
		A[i-1] = t
		i++
	}
	tampilsks(A, n)
}

func sksUrutMengecil(A *tabArr, n int) {
	var i, j, idx int
	var t mahasiswa
	i = 1
	for i <= n-1 {
		idx = i - 1
		j = 1 + 1
		A[idx].tot_sks = totalSks(A[idx])
		A[j].tot_sks = totalSks(A[j])
		for j < n {
			if (A[idx].tot_sks < A[j].tot_sks) || (A[idx].tot_sks == A[j].tot_sks && A[idx].transkrip.angka < A[j].transkrip.angka) {
				idx = j
			}
			j++
		}
		t = A[idx]
		A[idx] = A[i-1]
		A[i-1] = t
		i++
	}
	tampilsks(A, n)
}

func dataMenurutNim(A *tabArr, n int) {
	var pilihan int
	for pilihan != 3 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("    Jenis data terurut apa yang anda inginkan?    ")
		fmt.Println("                                                  ")
		fmt.Println("           1. Terurut membesar                    ")
		fmt.Println("           2. Terurut mengecil                    ")
		fmt.Println("           3. Kembali ke menu sebelumnya          ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("       Masukkan: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			nimUrutMembesar(A, n)
		} else if pilihan == 2 {
			nimUrutMengecil(A, n)
		} else if pilihan == 3 {
			fmt.Println("Baik, anda akan segera menuju ke menu sebelumnya")
			fmt.Println("Mohon tunggu sebentar")
			fmt.Println()
		}
	}
}

func nimUrutMembesar(A *tabArr, n int) {
	var i, j int
	var temp mahasiswa
	i = 1
	for i <= n-1 {
		j = i
		temp = A[j]
		for j > 0 && temp.nim < A[j-1].nim {
			A[j] = A[j-1]
			j = j - 1
		}
		A[j] = temp
		i++
	}
	tampilmhs(A, n)
}

func nimUrutMengecil(A *tabArr, n int) {
	var i, j int
	var temp mahasiswa
	i = 1
	for i <= n-1 {
		j = i
		temp = A[j]
		for j > 0 && temp.nim > A[j-1].nim {
			A[j] = A[j-1]
			j = j - 1
		}
		A[j] = temp
		i++
	}
	tampilmhs(A, n)
}

func pokoknyaUrutbdamatlahcapek(A *tabArr, n int) {
	var i, j int
	var temp mahasiswa
	i = 1
	for i <= n-1 {
		j = i
		temp = A[j]
		for j > 0 && temp.nim > A[j-1].nim {
			A[j] = A[j-1]
			j = j - 1
		}
		A[j] = temp
		i++
	}
}

func cariDataMahasiswa(A tabArr, n int) {
	var nim string
	fmt.Println()
	fmt.Println("Silahkan masukkan NIM mahasiswa yang ingin anda cari")
	fmt.Print("Masukkan :")
	fmt.Scan(&nim)
	fmt.Println()
	pokoknyaUrutbdamatlahcapek(&A, n)
	hasil := binarySearch(A, n, nim)
	if hasil != -1 {
		fmt.Println("Data mahasiswa ditemukan!")
		fmt.Println("Jika tidak ada mata kuliah atau nilai yang diinputkan,")
		fmt.Println("tampilan akan kosong")
		fmt.Println("Berikut adalah data mahasiswa tersebut:")
		fmt.Printf("%15s %40s %10s %15s %20s %3s %10s %10s %10s\n", "NIM ", " NAMA", "KELAS", "KODE MATKUL", "NAMA MATKUL", "SKS", "UTS", "UAS", "QUIZ")
		for j := 0; j < A[hasil].m; j++ {
			for k := 0; k < A[hasil].matakuliah[j].nilai.qz; k++ {
				fmt.Printf("%15s %40s %10s %15s %20s %3d %8.2f %2s %8.2f %2s %8.2f %2s\n", A[hasil].nim, A[hasil].nama, A[hasil].kelas, A[hasil].matakuliah[j].kode, A[hasil].matakuliah[j].nama, A[hasil].matakuliah[j].sks, A[hasil].matakuliah[j].nilai.uts.angka, A[hasil].matakuliah[j].nilai.uts.grade, A[hasil].matakuliah[j].nilai.uas.angka, A[hasil].matakuliah[j].nilai.uas.grade, A[hasil].matakuliah[j].nilai.quiz[k].angka, A[hasil].matakuliah[j].nilai.quiz[k].grade)
			}
		}
	} else {
		fmt.Println("Mohon maaf, data tersebut tidak ditemukan")
	}
}

func binarySearch(A tabArr, n int, nim string) int {
	var found int
	var kiri, kanan, tengah int
	found = -1
	kiri = 0
	kanan = n - 1
	for kiri <= kanan && found == -1 {
		tengah = (kiri + kanan) / 2
		if nim > A[tengah].nim {
			kanan = tengah - 1
		} else if nim < A[tengah].nim {
			kiri = tengah + 1
		} else {
			found = tengah
		}
	}
	return found
}

func datanim(A tabArr, n int) {
	var nim string
	var hasil, i int
	hasil = -1
	i = 0
	for i < n && hasil == -1 {
		if nim == A[i].nim {
			hasil = i
		}
	}
}

func tampilmhs(A *tabArr, n int) {
	fmt.Printf("%15s %40s %10s\n", "NIM ", " NAMA", "KELAS")
	for i := 0; i < n; i++ {
		fmt.Printf("%15s %40s %10s\n", A[i].nim, A[i].nama, A[i].kelas)
	}
}

func tampilmkdiambilmhs(A *tabArr, n int) {
	var kode, jawaban string
	fmt.Println()
	fmt.Print("Masukkan kode mata kuliah yang ingin anda tampilkan datanya: ")
	fmt.Scan(&kode)
	fmt.Println("Jika mata kuliah tidak ada atau tidak ada mahasiswa yang")
	fmt.Println("mengambil mata kuliah ini, maka hasilnya akan kosong")
	fmt.Println("Berikut adalah data mahasiswa yang mengambil mata kuliah ini")
	fmt.Println()

	fmt.Println("Data Mahasiswa yang mengambil matakuliah ")
	fmt.Printf("%15s %40s %10s\n", "NIM", "NAMA", "KELAS")
	for i := 0; i < n; i++ {
		a := findDataMatkul(A[i], kode)
		if a != -1 {
			fmt.Printf("%15s %40s %10s\n", A[i].nim, A[i].nama, A[i].kelas)
		}
	}
	fmt.Println("Apakah anda ingin menampilkan data mata kuliah yang lain?")
	fmt.Print("Ya/Tidak: ")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		tampilmkdiambilmhs(A, n)
	}
}

func tampilsks(A *tabArr, n int) {
	fmt.Printf("%15s %40s %10s %3s\n", "NIM ", " NAMA", "KELAS", "SKS")
	for i := 0; i < n; i++ {
		fmt.Printf("%15s %40s %10s %3d\n", A[i].nim, A[i].nama, A[i].kelas, A[i].tot_sks)
	}
}

func tampilnilai(A *tabArr, n int) {
	fmt.Printf("%15s %40s %10s %12s\n", "NIM ", " NAMA", "KELAS", "TOTAL NILAI")
	for i := 0; i < n; i++ {
		fmt.Printf("%15s %40s %10s %7.2f\n", A[i].nim, A[i].nama, A[i].kelas, A[i].tot_nilai)
	}
}
