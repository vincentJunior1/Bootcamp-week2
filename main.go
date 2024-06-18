package main

import "fmt"

type Siswa struct {
	Nama      string
	Kelas     int
	NilaiRata int
	Jurusan   string
}

func main() {

	// Tolong Penamaan variable di awalin sama huruf kecil aja jangan pake huruf besar
	Kelas := []Siswa{}
	Kelas = append(Kelas, Siswa{
		Nama:      "Botol",
		Kelas:     10,
		NilaiRata: 76,
	})

	Kelas = append(Kelas, Siswa{
		Nama:      "Sosro",
		Kelas:     10,
		NilaiRata: 10,
	})

	Kelas = append(Kelas, Siswa{
		Nama:      "Martin",
		Kelas:     10,
		NilaiRata: 100,
	})

	Kelas = append(Kelas, Siswa{
		Nama:      "Simon",
		Kelas:     10,
		NilaiRata: 45,
	})

	Kelas = append(Kelas, Siswa{
		Nama:      "Moe",
		Kelas:     10,
		NilaiRata: 74,
	})
	Kelas = append(Kelas, Siswa{
		Nama:      "Vincent",
		Kelas:     10,
		NilaiRata: 72,
	})

	for _, data := range Kelas {
		fmt.Println(data.Nama)
	}
	// for index, data := range Kelas {
	// 	if data.NilaiRata > 75 {
	// 		data.Jurusan = "IPA"
	// 	} else if data.NilaiRata > 50 {
	// 		data.Jurusan = "IPS"
	// 	} else {
	// 		data.Jurusan = "Bahasa"
	// 	}
	// 	Kelas[index] = data

	// 	fmt.Println(index)
	// 	fmt.Println("Siswa ini bernama", data.Nama)
	// 	fmt.Println("Nilai siswa ini adalah", data.NilaiRata)
	// 	fmt.Println("Siswa ini masuk kedalam", data.Jurusan)
	// 	fmt.Println("_________________________________")
	// }
}
