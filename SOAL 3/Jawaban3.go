package main

import "fmt"

type Klub struct {
	nama string
	skor int
}

func main() {
	var klubA, klubB Klub
	var nomorPertandingan int
	var hasil []string

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA.nama)

	fmt.Print("Klub B : ")
	fmt.Scan(&klubB.nama)

	for {
		nomorPertandingan++
		fmt.Printf("Pertandingan %d : ", nomorPertandingan)
		fmt.Scan(&klubA.skor, &klubB.skor)

		if klubA.skor < 0 || klubB.skor < 0 {
			break
		}

		if klubA.skor > klubB.skor {
			hasil = append(hasil, fmt.Sprintf("Hasil %d : %s", nomorPertandingan, klubA.nama))
		} else if klubB.skor > klubA.skor {
			hasil = append(hasil, fmt.Sprintf("Hasil %d : %s", nomorPertandingan, klubB.nama))
		} else {
			hasil = append(hasil, fmt.Sprintf("Hasil %d : Draw", nomorPertandingan))
		}
	}

	fmt.Println("\nHasil pertandingan:")
	for i := 0; i < len(hasil); i++ {
		fmt.Println(hasil[i])
	}

	fmt.Println("\nPertandingan selesai.")
}
