# METHOD

Method adalah `fungsi yang menempel pada type (bisa struct atau tipe data lainnya).` Method bisa `diakses lewat variabel objek.`

Keunggulan method dibanding fungsi biasa adalah memiliki akses ke property struct hingga level private (level akses nantinya akan dibahas lebih detail pada bab selanjutnya). Dan juga, dengan menggunakan method sebuah proses bisa di-enkapsulasi dengan baik.

`Penerapan Method`
Cara menerapkan method sedikit berbeda dibanding penggunaan fungsi. Ketika deklarasi, ditentukan juga siapa pemilik method tersebut. Contohnya bisa dilihat pada kode berikut:

    package main

    import "fmt"
    import "strings"

    type student struct {
      name  string
      grade int
    }

    func (s student) sayHello() {
      fmt.Println("halo", s.name)
    }

Cara deklarasi method sama seperti fungsi, hanya saja perlu ditambahkan deklarasi variabel object di sela-sela keyword func 
dan nama fungsi. Struct yang digunakan akan menjadi pemilik method.

func (s student) sayHello() maksudnya adalah fungsi sayHello dideklarasikan sebagai method milik struct student. Pada contoh di atas struct student memiliki dua buah method, yaitu sayHello() 

Contoh pemanfaatan method bisa dilihat pada kode berikut.

    package main

    import "fmt"
    import "strings"

    type student struct {
      name  string
      grade int
    }

    func (s student) sayHello() {
      fmt.Println("halo", s.name)
    }

    func main() {
        var john = student{"john wick", 21}
        john.sayHello()
    }