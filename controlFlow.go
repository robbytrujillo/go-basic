// Struktur Kondisional
// if, else if, else
nilai := 85

if nilai >= 90 {
    fmt.Println("Nilai A")
} else if nilai >= 80 {
    fmt.Println("Nilai B")
} else {
    fmt.Println("Nilai C atau di bawahnya")
}

// Struktur Pemilihan
// switch
hari := "Senin"

switch hari {
case "Senin":
    fmt.Println("Awal minggu")
case "Sabtu", "Minggu":
    fmt.Println("Weekend!")
default:
    fmt.Println("Hari biasa")
}

// Struktur Perulangan
// standar (miri for di C/C++)
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// tanpa kondisi (infinite loop)
for {
    fmt.Println("Terus jalan...")
    break // keluar dari loop
}

// dengan (range)
// digunakan untuk melakukan iterasi pada array, slice, map, dan string
angka := []int{1, 2, 3, 4}

for index, value := range angka {
    fmt.Println(index, value)
}

// kalau cuman butuh hasilnya saja
for _, value := range angka {
    fmt.Println(value)
}

