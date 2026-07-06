// Contoh Interface
type BisaBerjalan interface {
    Jalan()
}

// dua struct
type Manusia struct {
    Nama string
}

type Robot struct {
    Kode string
}

// masing-masing punya method Jalan()
func (m Manusia) Jalan() {
    fmt.Println(m.Nama, "sedang berjalan")
}

func (r Robot) Jalan() {
    fmt.Println("Robot", r.Kode, "berjalan otomatis")
}

// Menggunakan interface
func Gerakkan(b BisaBerjalan) {
    b.Jalan()
}

func main() {
    m := Manusia{Nama: "Andi"}
    r := Robot{Kode: "RX-01"}

    Gerakkan(m)
    Gerakkan(r)
}

// Output
// Andi sedang berjalan
// Robot RX-01 berjalan otomatis

// Interface Kosong (interface{})
func CetakData(data interface{}) {
    fmt.Println(data)
}
