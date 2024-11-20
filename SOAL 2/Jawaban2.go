package main

import "fmt"

func Array(data []int, n, x int) (bool, int) {
	var ketemu bool
	var k int
	var idx int = -1

	ketemu = false

	for k < n && idx == -1 {
		if data[k] == x && !ketemu {
			ketemu = true
			idx = k 
		}
		k = k + 1
	}
	return ketemu, idx
}

func main() {
	var n int

	fmt.Print("Masukkan jumlah data array : ")
	fmt.Scan(&n)

	data := make([]int, n)

	fmt.Println("Masukkan elemen : ")
	for i := 0; i < n; i++ { 
		fmt.Printf("Indeks ke-%d : ", i)
		fmt.Scan(&data[i])
	}

	fmt.Println("\na. Menampilkan keseluruhan isi array : ")
	fmt.Println(data)

	fmt.Println("\nb. Menampilkan elemen elemen array dengan indeks ganjil saja : ")
	for i := 0; i < n; i++ {
		if i%2 != 0 { 
			fmt.Printf("Indeks %d : %d\n", i, data[i])
		}
	}

	fmt.Println("\nc. Menampilkan elemen elemen array dengan indeks genap saja : ")
	for i := 0; i < n; i++ {
		if i%2 == 0 { 
			fmt.Printf("Indeks %d : %d\n", i, data[i])
		}
	}

	fmt.Println("\nd. Menampilkan elemen-elemen array dengan indeks kelipatan x : ")
	var x int
	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&x)
	if x == 0 {
		fmt.Println("Nilai x jangan 0 yaa sayangg.")
	} else {
		for i := 0; i < n; i++ {
			if i%x == 0 {
				fmt.Printf("Indeks %d : %d\n", i, data[i])
			}
		}
	}

	fmt.Println("\ne. Menghapus elemen array pada indeks tertentu dan tampilkan data setelah penghapusan : ")
	var hapus int
	fmt.Print("Masukan indeks yang ingin dihapus : ")
	fmt.Scan(&hapus)
	data = append(data[:hapus], data[hapus+1:]...)
	n = len(data)
	fmt.Println(data)

	fmt.Println("\nf. menampilkan rata rata dari bilangan yang ada didalam array : ")
	var total int
	for j := 0; j < n; j++ {
		total += data[j] 
	}

	ratarata := float64(total) / float64(n) 
	fmt.Println(data)
	fmt.Println("Rata Rata : ", ratarata) 

	fmt.Println("\ng. menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut : ")
	var jumlahKuadrat float64
	for i := 0; i < n; i++ {
		selisih := float64(data[i]) - ratarata
		jumlahKuadrat += selisih * selisih
	}

	var stdDev float64
	stdDev = jumlahKuadrat / float64(n)
	for i := 0; i < 10; i++ { 
		stdDev = (stdDev + jumlahKuadrat/stdDev) / 2
	}
	fmt.Printf("Standar Deviasi : %.2f\n", stdDev)

	fmt.Println("\nh. menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut : ")
	var y int
	fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya : ")
	fmt.Scan(&y) 

	frekuensi := 0
	for i := 0; i < n; i++ {
		if data[i] == y {
			frekuensi++
		}
	}

	if frekuensi > 0 {
		fmt.Printf("Bilangan %d muncul sebanyak %d kali.\n", y, frekuensi)
	} else {
		fmt.Printf("Bilangan %d tidak ditemukan dalam array.\n", y)
	}


}
