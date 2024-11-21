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

   ```go
   var lingkaran1, lingkaran2 Lingkaran
   ```
   Kode di atas adalah kode deklarasi variabel dalam bahasa pemrograman Go yang digunakan untuk mendeklarasikan dua variabel lingkaran1 dan lingkaran2 dengan tipe data Lingkaran. 

   ```go
   var titik Titik
   ```
   Kode di atas adalah kode deklarasi variabel yang digunakan untuk mendeklarasikan variabel titik dengan tipe data Titik. 

   ```go
   var lingkaran1, lingkaran2 Lingkaran
   ```
   Kode di atas adalah kode deklarasi variabel yang digunakan untuk mendeklarasikan dua variabel lingkaran1 dan lingkaran2 dengan tipe data Lingkaran. 

   ```go
   fmt.Scan(&lingkaran1.x, &lingkaran1.y, &lingkaran1.r)
   fmt.Scan(&lingkaran2.x, &lingkaran2.y, &lingkaran2.r)
   fmt.Scan(&titik.x, &titik.y)
   ```
   Kode di atas adalah kode yang digunakan untuk membaca input dari pengguna melalui fungsi fmt.Scan.

   ```go
   diDalam1 := diDalam(lingkaran1, titik)
   diDalam2 := diDalam(lingkaran2, titik)
   ```
   Kode di atas adalah kode yang memanggil fungsi diDalam untuk memeriksa apakah sebuah titik berada di dalam lingkaran tertentu. Variabel diDalam1 menyimpan hasil pengecekan untuk lingkaran1 dan titik, sementara diDalam2 menyimpan hasil pengecekan untuk lingkaran2 dan titik. 

   ```go
   if diDalam1 && diDalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
   ```
   Kode di atas adalah kode yang memanggil fungsi diDalam untuk memeriksa apakah sebuah titik berada di dalam lingkaran tertentu. Variabel diDalam1 menyimpan hasil pengecekan untuk lingkaran1 dan titik, sementara diDalam2 menyimpan hasil pengecekan untuk lingkaran2 dan titik. 


## SOAL 2
Program di atas adalah program yang menyediakan berbagai operasi pada array. Program ini meminta pengguna untuk memasukkan elemen array, lalu menampilkan isi array, elemen dengan indeks ganjil, genap, dan kelipatan nilai tertentu. Program juga dapat menghapus elemen pada indeks tertentu, menghitung rata-rata elemen, menghitung standar deviasi, serta menentukan frekuensi kemunculan suatu bilangan di dalam array. Operasi-operasi ini diimplementasikan menggunakan loop dan manipulasi array, seperti slicing dan perhitungan statistik sederhana.
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt' dan 'math').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Input Data Array, Pengguna diminta memasukkan jumlah elemen array dan elemen-elemen tersebut satu per satu.
      - Operasi pada Array
        	1. Menampilkan semua elemen array.
		2. Menampilkan elemen dengan indeks ganjil dan genap.
	 	3. Menampilkan elemen pada indeks kelipatan tertentu berdasarkan input pengguna.
     - Penghapusan Elemen Array, Menghapus elemen pada indeks tertentu menggunakan slicing array dan menampilkan hasilnya.
     - Statistik Array
	       1. Menghitung rata-rata elemen array.
	       2. Menghitung standar deviasi elemen array menggunakan rumus simpangan baku.
     - Frekuensi Elemen, Menghitung dan menampilkan jumlah kemunculan suatu bilangan tertentu di dalam array berdasarkan input pengguna.
      
   ## Code Explanation
   ```go
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
   ```
   Kode di atas adalah kode yang mendefinisikan fungsi Array untuk mencari elemen tertentu dalam sebuah array. Fungsi ini menerima tiga parameter: array data, panjang array n, dan elemen yang dicari x. Fungsi menggunakan perulangan untuk memeriksa apakah elemen x ada dalam array. Jika ditemukan, fungsi mengembalikan nilai true untuk variabel ketemu dan indeks elemen tersebut (idx). Jika tidak ditemukan, fungsi mengembalikan false dan -1 untuk indeks. Fungsi ini berguna untuk pencarian sederhana dalam array dengan hasil berupa status keberadaan elemen dan indeksnya.
  
   ```go
   var n int
   ```
   Kode var n int adalah deklarasi variabel. Variabel n dideklarasikan dengan tipe data int (bilangan bulat) tanpa inisialisasi nilai awal. 

   ```go
   fmt.Print("Masukkan jumlah data array : ")
   fmt.Scan(&n)
   ```
   Kode di atas adalah kode dalam bahasa Go yang digunakan untuk meminta input dari pengguna. Fungsi fmt.Print menampilkan pesan "Masukkan jumlah data array :" ke layar, sementara fmt.Scan membaca input dari pengguna dan menyimpannya ke variabel n. Variabel n biasanya digunakan untuk menentukan jumlah elemen dalam sebuah array atau struktur data lainnya.

   ```go
   data := make([]int, n)
   ```
   Kode di atas adalah kode yang digunakan untuk membuat sebuah slice bernama data dengan tipe data int dan panjang sebanyak n. Fungsi make digunakan untuk mengalokasikan memori untuk slice, sehingga slice dapat langsung digunakan untuk menyimpan elemen-elemen integer. Panjang slice n biasanya sudah ditentukan sebelumnya berdasarkan input pengguna.

   ```go
   fmt.Println("Masukkan elemen : ")
	for i := 0; i < n; i++ { 
		fmt.Printf("Indeks ke-%d : ", i)
		fmt.Scan(&data[i])
	}
   ```
   Kode di atas meminta pengguna untuk memasukkan elemen-elemen array. Menggunakan perulangan for, setiap elemen diminta sesuai dengan indeksnya dan disimpan dalam slice data dengan fmt.Scan.

   ```go
  fmt.Println("\na. Menampilkan keseluruhan isi array : ")
  fmt.Println(data)
   ```
  Kode di atas menampilkan seluruh isi array atau slice data menggunakan fmt.Println. Pesan "a. Menampilkan keseluruhan isi array :" ditampilkan terlebih dahulu, diikuti dengan output array yang berisi elemen-elemen yang sudah dimasukkan sebelumnya.

   ```go
   fmt.Println("\nb. Menampilkan elemen elemen array dengan indeks ganjil saja : ")
	for i := 0; i < n; i++ {
		if i%2 != 0 { 
			fmt.Printf("Indeks %d : %d\n", i, data[i])
		}
	}
   ```
   Kode di atas menampilkan elemen-elemen array yang memiliki indeks ganjil. Pesan "b. Menampilkan elemen-elemen array dengan indeks ganjil saja :" ditampilkan terlebih dahulu. Kemudian, perulangan for memeriksa setiap indeks, dan jika indeks tersebut ganjil (diperiksa dengan i%2 != 0), elemen pada indeks tersebut ditampilkan menggunakan fmt.Printf.

   ```go
   fmt.Println("\nc. Menampilkan elemen elemen array dengan indeks genap saja : ")
	for i := 0; i < n; i++ {
		if i%2 == 0 { 
			fmt.Printf("Indeks %d : %d\n", i, data[i])
		}
	}
   ```
   Kode di atas menampilkan elemen-elemen array yang memiliki indeks genap. Pesan "c. Menampilkan elemen-elemen array dengan indeks genap saja :" ditampilkan terlebih dahulu. Kemudian, perulangan for memeriksa setiap indeks, dan jika indeks tersebut genap (diperiksa dengan i%2 == 0), elemen pada indeks tersebut ditampilkan menggunakan fmt.Printf.

   ```go
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
   ```
   Kode di atas menampilkan elemen-elemen array yang memiliki indeks kelipatan x. Pertama, pengguna diminta untuk memasukkan nilai x. Jika nilai x adalah 0, program menampilkan pesan peringatan. Jika tidak, perulangan for akan memeriksa setiap indeks dan menampilkan elemen yang memiliki indeks kelipatan x (diperiksa dengan i%x == 0). Elemen tersebut kemudian ditampilkan menggunakan fmt.Printf.

   ```go
   fmt.Println("\ne. Menghapus elemen array pada indeks tertentu dan tampilkan data setelah penghapusan : ")
	var hapus int
	fmt.Print("Masukan indeks yang ingin dihapus : ")
	fmt.Scan(&hapus)
	data = append(data[:hapus], data[hapus+1:]...)
	n = len(data)
	fmt.Println(data)
   ```
   Kode di atas menghapus elemen array pada indeks tertentu dan menampilkan data setelah penghapusan. Pengguna diminta untuk memasukkan indeks yang ingin dihapus. Fungsi append digunakan untuk menghapus elemen pada indeks yang diminta dengan cara menggabungkan bagian array sebelum dan setelah indeks tersebut. Setelah elemen dihapus, panjang array diperbarui dengan n = len(data), dan array yang sudah diperbarui ditampilkan menggunakan fmt.Println.

   ```go
   fmt.Println("\nf. menampilkan rata rata dari bilangan yang ada didalam array : ")
	var total int
	for j := 0; j < n; j++ {
		total += data[j] 
	}

	ratarata := float64(total) / float64(n) 
	fmt.Println(data)
	fmt.Println("Rata Rata : ", ratarata) 
   ```
   Kode di atas menghitung dan menampilkan rata-rata elemen-elemen yang ada dalam array. Pertama, variabel total digunakan untuk menjumlahkan seluruh elemen array menggunakan perulangan for. Setelah itu, rata-rata dihitung dengan membagi total dengan jumlah elemen n, dan hasilnya disimpan dalam variabel ratarata. Program kemudian menampilkan array yang sudah dimasukkan dan nilai rata-rata menggunakan fmt.Println.

   ```go
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
   ```
   Kode di atas menghitung dan menampilkan standar deviasi (simpangan baku) dari elemen-elemen dalam array. Pertama, variabel jumlahKuadrat digunakan untuk menghitung jumlah kuadrat selisih antara setiap elemen dan rata-rata. Kemudian, standar deviasi dihitung dengan membagi jumlahKuadrat dengan jumlah elemen n. Untuk meningkatkan akurasi, dilakukan iterasi 10 kali menggunakan metode iterasi numerik untuk memperbaiki nilai standar deviasi. Hasil akhirnya ditampilkan dengan format dua desimal menggunakan fmt.Printf.

   ```go
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
   ```
   Kode di atas menampilkan frekuensi kemunculan suatu bilangan tertentu dalam array. Pertama, pengguna diminta untuk memasukkan bilangan yang ingin dicari frekuensinya. Program kemudian menggunakan perulangan for untuk memeriksa setiap elemen array dan menghitung berapa kali bilangan yang dimaksud muncul. Jika bilangan ditemukan, frekuensinya akan ditampilkan, dan jika tidak ditemukan, program akan memberi tahu bahwa bilangan tersebut tidak ada dalam array.


   ## SOAL 3
Program di atas adalah program yang mencatat dan menampilkan hasil pertandingan antara dua klub. Program ini mendefinisikan tipe data Klub yang memiliki dua atribut: nama (nama klub) dan skor (jumlah skor yang dicetak klub).
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt' dan 'math').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Tipe data Klub, yang mendefinisikan struktur untuk menyimpan nama dan skor klub.
      - Variabel klubA dan klubB, yang masing-masing menyimpan informasi nama dan skor dari kedua klub.
      - Loop yang digunakan untuk mencatat hasil pertandingan secara berulang hingga salah satu klub memasukkan skor negatif.
      - Array hasil untuk menyimpan hasil pertandingan setiap ronde, yang akan ditampilkan di akhir program.
      
   ## Code Explanation
   ```go
   type Klub struct {
	nama string
	skor int
   }
   ```
   Kode di atas adalah deklarasi tipe data Klub. Tipe data Klub adalah sebuah struct, yang digunakan untuk mengelompokkan beberapa nilai yang berhubungan. 
  
   ```go
   var klubA, klubB Klub
   ```
   Kode var klubA, klubB Klub adalah deklarasi dua variabel, klubA dan klubB, dengan tipe data Klub yang telah didefinisikan sebelumnya. Kedua variabel ini akan digunakan untuk menyimpan informasi tentang dua klub yang berbeda, seperti nama dan skor mereka. 

   ```go
   var nomorPertandingan int
   ```
   Kode var nomorPertandingan int adalah deklarasi variabel nomorPertandingan dengan tipe data int (bilangan bulat). Variabel ini digunakan untuk menyimpan nomor atau urutan pertandingan dalam program.

   ```go
   var hasil []string
   ```
   Kode var hasil []string adalah deklarasi variabel hasil yang bertipe slice of string ([]string). Slice ini digunakan untuk menyimpan hasil-hasil pertandingan dalam bentuk string. 

   ```go
   fmt.Print("Klub A : ")
   fmt.Scan(&klubA.nama)

   fmt.Print("Klub B : ")
   fmt.Scan(&klubB.nama)
   ```
   Kode di atas meminta pengguna untuk memasukkan nama dua klub yang akan bertanding. Nama klub A dan klub B dimasukkan melalui input dan disimpan dalam atribut nama dari variabel klubA dan klubB.

   ```go
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
   ```
   Kode di atas menjalankan sebuah loop yang mencatat hasil pertandingan antara dua klub. Setiap kali pertandingan dimulai, nomor pertandingan (nomorPertandingan) akan bertambah satu. Program kemudian meminta input skor untuk klub A dan klub B. Jika salah satu klub memasukkan skor negatif, loop akan berhenti. Setelah itu, program membandingkan skor kedua klub: jika klub A menang, nama klub A ditambahkan ke dalam slice hasil, jika klub B menang, nama klub B yang ditambahkan, dan jika keduanya seri, "Draw" yang ditambahkan. Hasil pertandingan tersebut disimpan dalam slice hasil dengan format yang sesuai.

   ```go
   fmt.Println("\nHasil pertandingan:")
	for i := 0; i < len(hasil); i++ {
		fmt.Println(hasil[i])
   }
   ```
   Kode di atas digunakan untuk menampilkan hasil pertandingan setelah loop selesai. Pertama, fmt.Println("\nHasil pertandingan:") menampilkan judul untuk hasil pertandingan. Kemudian, dengan perulangan for, program akan menampilkan setiap elemen dari slice hasil satu per satu menggunakan fmt.Println(hasil[i]). Setiap elemen berisi informasi tentang hasil pertandingan, seperti pemenang atau jika pertandingan berakhir dengan hasil seri.

   ```go
   fmt.Println("\nPertandingan selesai.")
   ```
   Kode fmt.Println("\nPertandingan selesai.") digunakan untuk menampilkan pesan yang menunjukkan bahwa semua pertandingan telah selesai. Pesan ini akan muncul setelah hasil semua pertandingan ditampilkan, menandakan bahwa proses pencatatan hasil telah selesai dan program telah berakhir.


## SOAL 4
Program di atas adalah program yang memanipulasi array tipe rune untuk memeriksa apakah sebuah kata merupakan palindrom (kata yang sama jika dibaca dari depan dan belakang).
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt' dan 'math').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Konstanta dan Tipe Data
	      	1. const NMAX int = 127 mendefinisikan batas maksimum panjang array, yaitu 127 elemen.
		2. type tabel [NMAX]rune mendefinisikan tipe data tabel sebagai array yang berisi elemen bertipe rune (untuk menyimpan karakter Unicode).
      - Fungsi isiArray, Fungsi ini meminta input berupa kata dari pengguna, kemudian memasukkan karakter-karakter dari kata tersebut ke dalam array tabel satu per satu hingga mencapai batas maksimum NMAX atau sampai semua karakter dimasukkan.
      - Fungsi cetakArray, Fungsi ini mencetak array tabel dengan mengubah setiap elemen rune menjadi string menggunakan string(t[i]).
      - Fungsi balikanArray, Fungsi ini membalikkan urutan elemen-elemen dalam array dan mengembalikannya sebagai array baru bertipe tabel.
      - Fungsi cekPalindrom, Fungsi ini memeriksa apakah kata dalam array adalah palindrom dengan membandingkan elemen-elemen dari kedua ujung array menuju ke tengah. Jika ada perbedaan, fungsi akan mengembalikan false; jika tidak ada perbedaan, mengembalikan true.
      
   ## Code Explanation
   ```go
   const NMAX int = 127 
   ```
   Kode di atas mendeklarasikan sebuah konstanta dengan nama NMAX yang bertipe int dan memiliki nilai 127. Konstanta ini digunakan untuk menentukan ukuran maksimum array dalam program.
  
   ```go
   type tabel [NMAX]rune
   ```
   mendefinisikan tipe data baru dengan nama tabel, yang merupakan array dengan panjang tetap sebesar 127 elemen, dan setiap elemen bertipe rune. Tipe rune digunakan untuk menyimpan karakter Unicode, sehingga array tabel dapat menyimpan hingga 127 karakter (baik huruf maupun simbol Unicode).

   ```go
   var tab tabel 
   var m int 
   ```
   Kode var tab tabel mendeklarasikan variabel tab bertipe array tabel, yang dapat menyimpan hingga 127 karakter bertipe rune. Sementara itu, var m int mendeklarasikan variabel m bertipe int, yang digunakan untuk melacak jumlah elemen yang telah dimasukkan ke dalam array tab.

   ```go
   func isiArray(t *tabel, n *int) {
	var input string
	fmt.Print("Masukkan kata : ")
	fmt.Scanln(&input)

	for i := 0; i < len(input); i++ {
		if *n >= NMAX {
			break
		}
		t[*n] = rune(input[i]) 
		(*n)++
	}
   }
   ```
   Fungsi isiArray menerima dua parameter, t yang merupakan pointer ke array tabel dan n yang merupakan pointer ke integer untuk melacak jumlah elemen dalam array. Fungsi ini meminta pengguna untuk memasukkan sebuah kata menggunakan fmt.Scanln(&input). Kemudian, karakter-karakter dari kata tersebut dimasukkan ke dalam array t satu per satu. Setiap karakter dikonversi menjadi tipe rune sebelum dimasukkan ke dalam array, dan jumlah elemen dalam array (n) ditambah 1 setelah setiap karakter ditambahkan. Fungsi ini berhenti jika jumlah elemen mencapai batas maksimum yang ditentukan oleh NMAX.

   ```go
   func cetakArray(t tabel, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += string(t[i]) 
	}
	return result
   }
   ```
   Kode di atas adalah fungsi cetakArray yang menerima dua parameter, t bertipe array tabel dan n bertipe int, yang menunjukkan jumlah elemen dalam array. Fungsi ini menggabungkan elemen-elemen array t yang berjumlah n menjadi sebuah string. Setiap elemen array (yang bertipe rune) dikonversi menjadi string menggunakan string(t[i]) dan kemudian ditambahkan ke variabel result. Setelah selesai, fungsi ini mengembalikan string yang berisi gabungan karakter-karakter tersebut.

   ```go
   func balikanArray(t tabel, n int) tabel {
	var reversed tabel
	for i := 0; i < n; i++ {
		reversed[i] = t[n-i-1] 
	}
	return reversed
   }
   ```
   Fungsi balikanArray menerima dua parameter, t yang merupakan array bertipe tabel dan n yang menunjukkan jumlah elemen dalam array. Fungsi ini membalikkan urutan elemen-elemen dalam array t dan menyimpannya dalam array baru reversed. Proses pembalikan dilakukan dengan menyalin elemen-elemen array t mulai dari indeks terakhir (yaitu n-i-1) ke array reversed. Setelah selesai, fungsi ini mengembalikan array reversed yang berisi elemen-elemen array yang telah dibalik.

   ```go
   func cekPalindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false 
		}
	}
	return true 
   }
   ```
   Fungsi cekPalindrom memeriksa apakah array t adalah palindrom dengan membandingkan elemen dari kedua ujung array. Jika ada elemen yang tidak cocok, fungsi mengembalikan false, jika cocok semua, mengembalikan true.

   ```go
   var tab tabel  
   var m int    
   ```
   Kode var tab tabel mendeklarasikan variabel tab bertipe array tabel, yang dapat menyimpan hingga 127 karakter bertipe rune. Sementara itu, var m int mendeklarasikan variabel m bertipe int, yang digunakan untuk melacak jumlah elemen yang telah dimasukkan ke dalam array tab.

   ```go
   isiArray(&tab, &m)   
   ```
   Perintah 'isiArray(&tab, &m)' memanggil fungsi isiArray, mengirimkan alamat variabel tab dan m untuk mengisi array dan memperbarui jumlah elemen.

   ```go
   tabBalik := balikanArray(tab, m)   
   ```
   Perintah 'tabBalik := balikanArray(tab, m)' memanggil fungsi balikanArray untuk membalikkan urutan elemen dalam array tab sebanyak m elemen, dan menyimpannya dalam variabel tabBalik.

   ```go
   fmt.Print("Array asli : ")
   fmt.Println(cetakArray(tab, m))     
   ```
   Perintah 'fmt.Print("Array asli : ")' mencetak teks "Array asli : ", kemudian 'fmt.Println(cetakArray(tab, m))' memanggil fungsi cetakArray untuk mengubah array tab yang berisi m elemen menjadi sebuah string, lalu mencetak string tersebut.

   ```go
   fmt.Print("Array setelah dibalik : ")
   fmt.Println(cetakArray(tabBalik, m))     
   ```
   Perintah 'fmt.Print("Array setelah dibalik : ")' mencetak teks "Array setelah dibalik : ", kemudian 'fmt.Println(cetakArray(tabBalik, m))' memanggil fungsi cetakArray untuk mengubah array tabBalik yang berisi hasil pembalikan array tab menjadi sebuah string, lalu mencetak string tersebut.

   ```go
   if cekPalindrom(tab, m) {
		fmt.Println("Palindrom ? true")
   } else {
		fmt.Println("Palindrom ? false")
   }  
   ```
   Perintah 'if cekPalindrom(tab, m)' memeriksa apakah array tab dengan m elemen merupakan palindrom menggunakan fungsi cekPalindrom. Jika hasilnya true, maka program mencetak "Palindrom ? true". Jika hasilnya false, program mencetak "Palindrom ? false".

   

   
    









   

