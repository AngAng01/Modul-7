# MODUL TUJUH
 ## SOAL 1
   Program di atas adalah program recursive yang digunakan untuk menghitung deret Fibonacci. Fungsi fibonacci menggunakan rekursi untuk menghitung nilai Fibonacci dari suatu bilangan n, yaitu jumlah dari dua angka sebelumnya dalam deret Fibonacci.
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Fungsi fibonacci(n int) int, Menghitung nilai Fibonacci untuk n secara rekursif.
      - Paket fmt, Digunakan untuk mencetak hasil ke layar menggunakan fmt.Printf.
      
   ## Code Explanation
   ```go
   func fibonacci(n int) int {
		`if n <= 1 {
		 	return n
		 }
		 return fibonacci(n-1) + fibonacci(n-2)
	}
   ```
   Kode di atas adalah fungsi rekursif untuk menghitung nilai Fibonacci dari bilangan n. Fungsi fibonacci menerima satu parameter integer n dan mengembalikan nilai Fibonacci yang sesuai. 
   
   ```go
      for i := 0; i <= 10; i++ {
	        fmt.Printf("Fibonacci(%d) = %d\n", i, fibonacci(i))
	    }
   ```
   Kode di atas adalah Fungsi permutasi yang digunakan untuk menghitung permutasi dari n elemen yang diambil sebanyak r. Fungsi ini pertama-tama menghitung faktorial dari n (disimpan dalam variabel pembilang) dan faktorial dari n - r (disimpan dalam variabel penyebut) dengan memanfaatkan fungsi factorial. 
   

