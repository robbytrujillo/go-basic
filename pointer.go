var angka int = 10
var ptr *int = &angka

fmt.Println("Nilai angka:", angka)
fmt.Println("Alamat angka:", &angka)
fmt.Println("Pointer ptr:", ptr)

// hasil
fmt.Println("Nilai melalui pointer:", *ptr)

// Pointer dan Fungsi
// Tanpa pointer
func ubah(x int) {
    x = 100
}

func main() {
    a := 10
    ubah(a)
    fmt.Println(a) // Tetap 10
}

// Dengan pointer
func ubah(x *int) {
    *x = 100
}

func main() {
    a := 10
    ubah(&a)
    fmt.Println(a) // Jadi 100
}

// Pointer Default
angka := new(int)
*angka = 50
fmt.Println(*angka) // Output: 50


