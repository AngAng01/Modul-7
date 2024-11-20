package main

import "fmt"

const NMAX int = 127 
type tabel [NMAX]rune
	var tab tabel 
	var m int 

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

func cetakArray(t tabel, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += string(t[i]) 
	}
	return result
}

func balikanArray(t tabel, n int) tabel {
	var reversed tabel
	for i := 0; i < n; i++ {
		reversed[i] = t[n-i-1] 
	}
	return reversed
}

func cekPalindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false 
		}
	}
	return true 
}

func main() {
	var tab tabel  
	var m int      
	
	isiArray(&tab, &m)

	tabBalik := balikanArray(tab, m)

	fmt.Print("Array asli : ")
	fmt.Println(cetakArray(tab, m))  

	fmt.Print("Array setelah dibalik : ")
	fmt.Println(cetakArray(tabBalik, m))  

	if cekPalindrom(tab, m) {
		fmt.Println("Palindrom ? true")
	} else {
		fmt.Println("Palindrom ? false")
	}
}
