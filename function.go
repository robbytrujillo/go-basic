// cara membuat function (fungsi)
// func sapa() {
// 	fmt.Println("Halo dari fungsi!")
// }
// untuk menjalankan:
// sapa()

// fungsi dengan parameter
func halo(nama string) {
	fmt.Println("Halo", nama)
}
// Pemanggilan:
// halo("Golang")

// fungsi dapat punya parameter
func tambah(a int, b int) {
    fmt.Println("Hasil:", a + b)
}
// atau disingkat
func tambah(a, b int) {
    fmt.Println("Hasil:", a + b)
}

// fungsi dengan return value (fungsi bisa mengembalikan nilai menggunakan return)
func kali(a, b int) int {
    return a * b
}

// cara penggunaan
hasil := kali(3, 4)
fmt.Println(hasil) // Output: 12

// Multiple Return Value
func hitung(a, b int) (int, int) {
    return a + b, a * b
}
// pemanggilan
jumlah, perkalian := hitung(2, 5)
fmt.Println("Jumlah:", jumlah)
fmt.Println("Perkalian:", perkalian)

// Fungsi dengan Return Bernama
func bagi(a, b int) (hasil int) {
    hasil = a / b
    return
}

// Fungsi di dalam Fungsi (Function as Value)
f :func(nama string) {
	fmt.Println("Hai", nama)
}
f("Golang")
