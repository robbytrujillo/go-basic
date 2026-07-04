// Array
var angka [3]int

angka[0] = 10
angka[1] = 20
angka[2] = 30

// atau

angka := [3]int{10, 20, 30}

fmt.Println(angka[1]) // Output: 20

// Slice
angka := []int{10, 20, 30}
fmt.Println(angka[0]) // Output: 10

// menambahkan elemen append
angka = append(angka, 40)
fmt.Println(angka) // [10 20 30 40]

// Membuat slice dan array
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4] // Ambil elemen ke-1 sampai sebelum ke-4
fmt.Println(slice) // [2 3 4]

// Map
nilai := map[string]int{
    "Andi": 90,
    "Budi": 85,
}

fmt.Println(nilai["Andi"]) // Output: 90

// menambah data
nilai["Citra"] = 88

// Menghapus data
delete(nilai, "Budi")

// cek apakah key ada?
val, ok := nilai["Dina"]

if ok {
    fmt.Println("Ada:", val)
} else {
    fmt.Println("Tidak ditemukan")
}


