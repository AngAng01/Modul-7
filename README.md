# MODUL TUJUH
 ## SOAL 1
Program di atas adalah program untuk menentukan apakah sebuah titik berada di dalam salah satu atau kedua lingkaran, atau di luar keduanya.
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt' dan 'math').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Struct Titik menyimpan koordinat (x, y) suatu titik.
      - Struct Lingkaran menyimpan koordinat pusat lingkaran (x, y) dan jari-jari r.
      - Fungsi jarak(p, q) menghitung jarak Euclidean antara dua titik p dan q.
      - Fungsi diDalam(lingkaran, titik) mengevaluasi apakah jarak dari titik ke pusat lingkaran lebih kecil atau sama dengan jari-jari lingkaran, yang berarti titik berada di dalam lingkaran.
      
   ## Code Explanation
   ```go
   type Titik struct {
	x, y int
   }
   ```
   Kode di atas adalah definisi struct di dalam bahasa pemrograman Go (Golang). Struct ini digunakan untuk merepresentasikan data titik pada koordinat kartesius dengan dua atribut x dan y
  
   ```go
   type Lingkaran struct {
	x, y, r int
   }
   ```
   Kode di atas adalah definisi sebuah struct di dalam bahasa pemrograman Go (Golang). Struct ini digunakan untuk merepresentasikan sebuah lingkaran dalam ruang 2D dengan atribut x, y ,dan r

   ```go
   func jarak(p, q Titik) float64 {
	 return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))
   }
   ```
   Kode di atas adalah fungsi untuk menghitung jarak antara dua titik dalam ruang 2D menggunakan rumus jarak Euclidean.

   ```go
   func diDalam(lingkaran Lingkaran, titik Titik) bool {
	 jarakKePusat := jarak(Titik{lingkaran.x, lingkaran.y}, titik)
	 return jarakKePusat <= float64(lingkaran.r)
   }
   ```
   Kode di atas adalah fungsi untuk menentukan apakah sebuah titik berada di dalam atau di tepi lingkaran tertentu.

