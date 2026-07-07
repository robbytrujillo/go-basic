// Analogi method
func (r ReceiverType) namaMethod() {
    // isi fungsi
}

// method sederhana
type Mahasiswa struct {
    Nama string
    Umur int
}

func (m Mahasiswa) Perkenalan() {
    fmt.Printf("Halo, nama saya %s dan umur saya %d\n", m.Nama, m.Umur)
}

// penggunaan
func main() {
    mhs := Mahasiswa{Nama: "Adit", Umur: 21}
    mhs.Perkenalan()
}

// method dengan return value
func (m Mahasiswa) TahunLahir(tahunSekarang int) int {
    return tahunSekarang - m.Umur
}

// Value Receiver (copy)
func (m Mahasiswa) UbahNamaBaru(nama string) {
    m.Nama = nama
}

// Pointer Receiver (langsung ubah)
func (m *Mahasiswa) UbahNama(nama string) {
    m.Nama = nama
}

func main() {
    mhs := Mahasiswa{Nama: "Rudi", Umur: 22}
    mhs.UbahNama("Santi")
    fmt.Println(mhs.Nama) // Output: Santi
}

