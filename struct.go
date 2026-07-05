// Struct adalah kumpulan dari beberapa field(varable) yang dikelompokkan dalam satu tipe data
// contoh tipe data mahasiswa
// type Mahasiswa struct {
//     Nama  string
//     Umur  int
//     Aktif bool
// }

// cara menggunakan
// Membuat variable dan struct
var mhs1 Mahasiswa
mhs1.Nama = "Budi"
mhs1.Umur = 20
mhs1.Aktif = true

fmt.Println(mhs1)

// struct literal (langsung isi)
mhs2 := Mahasiswa{
    Nama:  "Ani",
    Umur:  22,
    Aktif: false,
}

// akses field
fmt.Println("Nama:", mhs2.Nama)
fmt.Println("Umur:", mhs2.Umur)

// Struct di dalam Struct
// struct dapat menampung struct lain
type Alamat struct {
    Kota   string
    Jalan  string
}

type Mahasiswa struct {
    Nama   string
    Umur   int
    Alamat Alamat
}

// Penggunaan
mhs := Mahasiswa{
    Nama: "Citra",
    Umur: 21,
    Alamat: Alamat{
        Kota:  "Bandung",
        Jalan: "Jl. Merdeka",
    },
}

fmt.Println(mhs.Alamat.Kota) // Output: Bandung

// Pointer ke struct
func ubahUmur(m *Mahasiswa) {
    m.Umur = 25
}

func main() {
    m := Mahasiswa{Nama: "Doni", Umur: 20}
    ubahUmur(&m)
    fmt.Println(m.Umur) // Output: 25
}

