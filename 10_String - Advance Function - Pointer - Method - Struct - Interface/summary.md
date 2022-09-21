# STRING

## Menghitung panjang string
Berapa Panjang string A?
```go
func main(){
  // len strings
  sentence := "Hello"
  lenSentence := len(sentence)
  fmt.Println(lenSentence)
}
```
## Compare strings
Apakah string A sama dengan string B?
```go
func main(){
  // compare strings
  str1 := "abc"
  str2 := "abd"
  fmt.Println(str1 == str2)
}
```
## Contain strings
Apakah string A mengandung string B?
```go
func main(){
  // Contains
  const (
    str = "something"
    substr = "some"
  )
  res := strings.Contains(str, substr)
  fmt.Println(res)
}
```
## Catch string pake index
Bagaimana mengambil kata "dog" dari "watchdog"?
```go
func main(){
  str := "watchdog"
  //index 01234567
  res := str[5:len(str)]
  fmt.Println(res)
}
```
## Replace strings`
Bagaimana cara mengganti "katak" menjadi "kotak"?
```go
func main(){
  str := "katak"
  kotok := strings.Replace(str, "a", "o", -1) // Mengubah "katak" menjadi "kotok" karena parameter terakhir -1 jadi seluruh huruf/kata yang ditemukan dirubah
  kotak := strings.Replace(str, "a", "o", 1) //  Mengubah "katak" menjadi "kotak" karena parameter terakhir nya 1 jadi hanya 1 huruf pertama yang ditemukan
  fmt.Println("%s\n%s\n", kotok, kotak)
}
```

## Insert string
Bagaimana cara mengubah bh menjadi buah?
```go
func main(){
  str := "bh"
  // i=>  01
  res := str[:1] + "ua" + str[1:]
  fmt.Println(res)
}
```


# Advance Function

## Variadic function
Digunakan ketika parameter yang di harapkan berubah-ubah jumlahnya. Cara nya adalah dengan mengubah parameter function dengan slice.
```go
// function
func sum(angka ...int){
  res := 0
  for _, v := range angka {
    res += v
  }
  return res
}

// Main
func main() {
  fmt.Println(sum(1,2,3,4,5))
  fmt.Println(sum(1,3,5,2,3,5,6,7,7,4))
}
```
## Anonymous Function == Literal Function
Adalah function yang tidak memiliki nama. Berguna ketika ingin membuat inline function.
  
### 1
```go
func main(){

  func(){
    fmt.Println("Hello World")
    fmt.Println("Alterra Academy")
  }()

  fmt.Println("Hello World Bagian 2")
  fmt.Println("Alterra Academy Bagian 2")
}
```
### 2 set to variable
```go
func main(){

  bagianAwal := func(){
    fmt.Println("Hello World")
    fmt.Println("Alterra Academy")
  }

  fmt.Println("Hello World Bagian 2")
  fmt.Println("Alterra Academy Bagian 2")

  bagianAwal()
}
```
### 3 with parameter
```go
func main(){
  
  bagianKe := func(urutan string){
    fmt.Println("Hello World " + "bagian " + urutan)
    fmt.Println("Alterra Academy " + "bagian " + urutan)
  }

  fmt.Println("Hello World Bagian 2")
  fmt.Println("Alterra Academy Bagian 2")

  bagianKe("akhir")
}
```
### 4 Closure
```go
func NewCounter() func() int {
  count := 0
  return func() int {
    count += 1
    return count
  }
}

func main() {
  counter := NewCounter()
  fmt.Print(counter() + " ")
  fmt.Print(counter() + " ")
  fmt.Print(counter() + " ")
  // OUTPUT: 1 2 3
}
```
## Defer Function
adalah function yang akan dieksekusi ketika parent function nya mengembalikan nilai. Multiple return juga bisa digunakan dan dijalankan sebagai stack(LIFO), satu per satu.
### defer - simple function`
```go
func main(){
  defer fmt.Print(" 1 ")
  defer fmt.Print(" 2 ")
  fmt.Print(" 3 ")
  // Output: 3  2  1
}
```
`defer anonymous function`
```go
func main(){
  
  defer func() {
    fmt.Print(" 1 ")
  }()
  defer func() {
    fmt.Print(" 2 ")
  }()
  fmt.Print(" 3 ")

  // OUTPUT: 3  2  1

}
```
defer biasanya digunakan saat open dan close database
```go
func main(){
    
  defer func() {
    fmt.Println(" Close DB ")
  }()
  fmt.Println(" Open DB ")
  
  // Jadi tidak akan lupa untuk close DB dan dijalankan setelah seluruh perintah dalam function selesai
}
```


# POINTER

Pointer adalah reference atau alamat memori. Variabel pointer berarti variabel yang berisi alamat memori suatu nilai. Sebagai contoh sebuah variabel bertipe integer memiliki nilai 4, maka yang dimaksud pointer adalah alamat memori dimana nilai 4 disimpan, bukan nilai 4 itu sendiri.

Variabel-variabel yang memiliki reference atau alamat pointer yang sama, saling berhubungan satu sama lain dan nilainya pasti sama. Ketika ada perubahan nilai, maka akan memberikan efek kepada variabel lain (yang referensi-nya sama) yaitu nilainya ikut berubah.

## 1. Penerapan Pointer
Variabel bertipe pointer ditandai dengan adanya tanda asterisk (*) tepat sebelum penulisan tipe data ketika deklarasi.
```go
var number *int
var name *string
```
Nilai default variabel pointer adalah nil (kosong). Variabel pointer tidak bisa menampung nilai yang bukan pointer, dan sebaliknya variabel biasa tidak bisa menampung nilai pointer.

Ada dua hal penting yang perlu diketahui mengenai pointer:

Variabel biasa bisa diambil nilai pointernya, caranya dengan menambahkan tanda ampersand (&) tepat sebelum nama variabel. Metode ini disebut dengan referencing.
Dan sebaliknya, nilai asli variabel pointer juga bisa diambil, dengan cara menambahkan tanda asterisk (*) tepat sebelum nama variabel. Metode ini disebut dengan dereferencing.
```go
var numberA int = 4
var numberB *int = &numberA
fmt.Println("numberA (value)   :", numberA) // 4
fmt.Println("numberA (address) :", &numberA) // 0xc20800a220
fmt.Println("numberB (value)   :", *numberB) // 4
fmt.Println("numberB (address) :", numberB) // 0xc20800a220
```
Variabel numberB dideklarasikan bertipe pointer int dengan nilai awal adalah referensi variabel numberA (bisa dilihat pada kode &numberA). Dengan ini, variabel numberA dan numberB menampung data dengan referensi alamat memori yang sama.

Variabel pointer jika di-print akan menghasilkan string alamat memori (dalam notasi heksadesimal), contohnya seperti numberB yang diprint menghasilkan 0xc20800a220.

Nilai asli sebuah variabel pointer bisa didapatkan dengan cara di-dereference terlebih dahulu (bisa dilihat pada kode *numberB).


## 2. Efek Perubahan Nilai Pointer
Ketika salah satu variabel pointer di ubah nilainya, sedang ada variabel lain yang memiliki referensi memori yang sama, maka nilai variabel lain tersebut juga akan berubah.
```go
var numberA int = 4
var numberB *int = &numberA

fmt.Println("numberA (value)   :", numberA)
fmt.Println("numberA (address) :", &numberA)
fmt.Println("numberB (value)   :", *numberB)
fmt.Println("numberB (address) :", numberB)

fmt.Println("")

numberA = 5

fmt.Println("numberA (value)   :", numberA)
fmt.Println("numberA (address) :", &numberA)
fmt.Println("numberB (value)   :", *numberB)
fmt.Println("numberB (address) :", numberB)
```
Variabel numberA dan numberB memiliki referensi memori yang sama. Perubahan pada salah satu nilai variabel tersebut akan memberikan efek pada variabel lainnya. Pada contoh di atas, numberA nilainya di ubah menjadi 5. membuat nilai asli variabel numberB ikut berubah menjadi 5.

## 3. Parameter Pointer
Parameter bisa juga didesain sebagai pointer. Cara penerapannya kurang lebih sama, dengan cara mendeklarasikan parameter sebagai pointer.
```go
package main

import "fmt"

func main() {
  var number = 4
  fmt.Println("before :", number) // 4
  
  change(&number, 10)
  fmt.Println("after  :", number) // 10
}

func change(original *int, value int) {
  *original = value
}
```
Fungsi change() memiliki 2 parameter, yaitu original yang tipenya adalah pointer int, dan value yang bertipe int. Di dalam fungsi tersebut nilai asli parameter pointer original diubah.
Fungsi change() kemudian diimplementasikan di main. Variabel number yang nilai awalnya adalah 4 diambil referensi-nya lalu digunakan sebagai parameter pada pemanggilan fungsi change().

Nilai variabel number berubah menjadi 10 karena perubahan yang terjadi di dalam fungsi change adalah pada variabel pointer.


# METHOD

Method adalah `fungsi yang menempel pada type (bisa struct atau tipe data lainnya).` Method bisa `diakses lewat variabel objek.`

Keunggulan method dibanding fungsi biasa adalah memiliki akses ke property struct hingga level private (level akses nantinya akan dibahas lebih detail pada bab selanjutnya). Dan juga, dengan menggunakan method sebuah proses bisa di-enkapsulasi dengan baik.

## Penerapan Method
Cara menerapkan method sedikit berbeda dibanding penggunaan fungsi. Ketika deklarasi, ditentukan juga siapa pemilik method tersebut. Contohnya bisa dilihat pada kode berikut:
```go
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
```
Cara deklarasi method sama seperti fungsi, hanya saja perlu ditambahkan deklarasi variabel object di sela-sela keyword func 
dan nama fungsi. Struct yang digunakan akan menjadi pemilik method.

func (s student) sayHello() maksudnya adalah fungsi sayHello dideklarasikan sebagai method milik struct student. Pada contoh di atas struct student memiliki dua buah method, yaitu sayHello() 

Contoh pemanfaatan method bisa dilihat pada kode berikut.
```go
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
```


# Struct

Go tidak memiliki class yang ada di bahasa-bahasa strict OOP lain. Tapi Go memiliki tipe data struktur yang disebut dengan Struct.

Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu. Property dalam struct, tipe datanya bisa bervariasi. Mirip seperti map, hanya saja key-nya sudah didefinisikan di awal, dan tipe data tiap itemnya bisa berbeda.

## 1. Deklarasi Struct
Keyword type digunakan untuk deklarasi struct. Di bawah ini merupakan contoh cara penggunaannya.
```go
type student struct {
    name string
    grade int
}
```
Struct student dideklarasikan memiliki 2 property, yaitu name dan grade.

## 2. Penerapan Struct
berbeda dengan map, struct tidak langsung bisa di gunakan begitu saja, struct bisa di gunakan dengan membuat data atau biasa disebut object berdasarkan dari struct yang sudah di buat. berikut contohnya:
```go
func main() {
  var john student
  john.name = "john doe"
  john.grade = 2

  fmt.Println("name  :", john.name)
  fmt.Println("grade :", john.grade)
}
```
seperti yang terlihat pada contoh diatas bahwa cara membuat object dari sebuah struct sama seperti kita mendeklarasikan variabel, di mulai dari nama variabel dan diikuti dengan tipe data, pada contoh diatas tipe datanya adalah student yang didapatkan dari type student struct.

## 3. Struct Literals
 Sebelumnya kita telah membuat object dari struct, namun sebenarnya ada banyak cara yang bisa kita gunakan untuk membuat data dari struct, berikut ini contohnya:
```go
// cara pertama
var john = student{}
john.name = "wick"
john.grade = 2

// cara kedua tetapi isinya harus berurutan
var doe = student{"doe", 2}

// cara ketiga dengan nama property tetapi tidak harus berurutan
var jack = student{name: "jack", grade: 2}

fmt.Println("student 1 :", john.name)
fmt.Println("student 2 :", doe.name)
fmt.Println("student 3 :", jack.name)
```
## 4. Embedded Struct
Embedded struct adalah mekanisme untuk menempelkan sebuah struct sebagai properti struct lain. Agar lebih mudah dipahami, mari kita bahas kode berikut.
```go
package main

import "fmt"

type person struct {
  name string
  age  int
}

type student struct {
  grade int
  person
}

func main() {
  // contoh 1
  var john = student{}
  john.name = "john"
  john.age = 21
  john.grade = 2

  fmt.Println("name  :", john.name)
  fmt.Println("age   :", john.age)
  fmt.Println("age   :", john.person.age)
  fmt.Println("grade :", john.grade)
  
  // contoh 2
  var doeData = person{name: "doe", age: 21}
  var doe = student{person: doeData, grade: 2}

  fmt.Println("name  :", doe.name)
  fmt.Println("age   :", doe.age)
  fmt.Println("grade :", doe.grade)
}
```
dapat terlihat diatas bahwa object john di inisialisasi berdasarkan object student dan property yang sebelumnya terdapat di object person dapat di gunakan jika struct student embed struct dari person

### 5. Anonymous Struct
Anonymous struct adalah struct yang tidak dideklarasikan di awal sebagai tipe data baru, melainkan langsung ketika pembuatan object. Teknik ini cukup efisien untuk pembuatan object yang struct-nya hanya dipakai sekali.
```go
package main

import "fmt"

type person struct {
  name string
  age  int
}

func main() {
  var john = struct {
    person
    grade int
  }{}
  john.person = person{"wick", 21}
  john.grade = 2

  fmt.Println("name  :", john.person.name)
  fmt.Println("age   :", john.person.age)
  fmt.Println("grade :", john.grade)
}
```
Pada kode di atas, variabel john langsung diisi objek anonymous struct yang memiliki property grade, dan property person yang merupakan embedded struct.

Salah satu aturan yang perlu diingat dalam pembuatan anonymous struct adalah, deklarasi harus diikuti dengan inisialisasi. Bisa dilihat pada john setelah deklarasi struktur struct, terdapat kurung kurawal untuk inisialisasi objek. Meskipun nilai tidak diisikan di awal, kurung kurawal tetap harus ditulis.
```go
// anonymous struct tanpa pengisian property
var john = struct {
person
grade int
}{}

// anonymous struct dengan pengisian property
var doe = struct {
person
grade int
}{
person: person{"wick", 21},
grade:  2,
}
```
## 6. Nested struct
Nested struct adalah anonymous struct yang di-embed ke sebuah struct. Deklarasinya langsung didalam struct peng-embed. Contoh:
```go
type student struct {
  person struct {
    name string
    age  int
  }
  grade   int
}
```


# Interface

Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), yang dibungkus dengan nama tertentu.

Interface merupakan tipe data. Nilai object bertipe interface zero value-nya adalah nil. Interface mulai bisa digunakan jika sudah ada isinya, yaitu object konkret yang memiliki definisi method minimal sama dengan yang ada di interface-nya.

## 1. Penerapan Interface
Yang pertama perlu dilakukan untuk menerapkan interface adalah menyiapkan interface beserta definisi method nya. Keyword type dan interface digunakan untuk pendefinisian interface.
```go
package main

import "fmt"
import "math"

type hitung interface {
  luas() float64
  keliling() float64
}
```
Pada kode di atas, interface hitung memiliki 2 definisi method, luas() dan keliling(). Interface ini nantinya digunakan sebagai tipe data pada variabel, dimana variabel tersebut akan menampung object bangun datar hasil dari struct yang akan kita buat.

Dengan memanfaatkan interface hitung, perhitungan luas dan keliling bangun datar bisa dilakukan, tanpa perlu tahu jenis bangun datarnya sendiri itu apa.

Siapkan struct bangun datar lingkaran, struct ini memiliki method yang beberapa diantaranya terdefinisi di interface hitung.
```go
type lingkaran struct {
  diameter float64
}

func (l lingkaran) jariJari() float64 {
  return l.diameter / 2
}

func (l lingkaran) luas() float64 {
  return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
  return math.Pi * l.diameter
}
```
Struct lingkaran di atas memiliki tiga method, jariJari(), luas(), dan keliling().

Selanjutnya, siapkan struct bangun datar persegi.
```go
type persegi struct {
  sisi float64
}

func (p persegi) luas() float64 {
  return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
  return p.sisi * 4
}
```
Perbedaan struct persegi dengan lingkaran terletak pada method jariJari(). Struct persegi tidak memiliki method tersebut. Tetapi meski demikian, variabel object hasil cetakan 2 struct ini akan tetap bisa ditampung oleh variabel cetakan interface hitung, karena dua method yang ter-definisi di interface tersebut juga ada pada struct persegi dan lingkaran, yaitu luas() dan keliling().

Buat implementasi perhitungan di main.
```go
func main() {
  var bangunDatar hitung
  
  bangunDatar = persegi{10.0}
  fmt.Println("===== persegi")
  fmt.Println("luas      :", bangunDatar.luas())
  fmt.Println("keliling  :", bangunDatar.keliling())
  
  bangunDatar = lingkaran{14.0}
  fmt.Println("===== lingkaran")
  fmt.Println("luas      :", bangunDatar.luas())
  fmt.Println("keliling  :", bangunDatar.keliling())
  fmt.Println("jari-jari :", bangunDatar.(lingkaran).jariJari())
}
```
Perhatikan kode di atas. Variabel object bangunDatar bertipe interface hitung. Variabel tersebut digunakan untuk menampung object konkrit buatan struct lingkaran dan persegi.

Dari variabel tersebut, method luas() dan keliling() diakses. Secara otomatis Golang akan mengarahkan pemanggilan method pada interface ke method asli milik struct yang bersangkutan.

Method jariJari() pada struct lingkaran tidak akan bisa diakses karena tidak terdefinisi dalam interface hitung. Pengaksesannya dengan paksa akan menyebabkan error.

Untuk mengakses method yang tidak ter-definisi di interface, variabel-nya harus di-casting terlebih dahulu ke tipe asli variabel konkritnya (pada kasus ini tipenya lingkaran), setelahnya method akan bisa diakses.

Cara casting object interface sedikit unik, yaitu dengan menuliskan nama tipe tujuan dalam kurung, ditempatkan setelah nama interface dengan menggunakan notasi titik (seperti cara mengakses property, hanya saja ada tanda kurung nya). Contohnya bisa dilihat di kode berikut. Statement bangunDatar.(lingkaran) adalah contoh casting pada object interface.
```go
var bangunDatar hitung = lingkaran{14.0}
var bangunLingkaran lingkaran = bangunDatar.(lingkaran)
```
bangunLingkaran.jariJari()
Perlu diketahui juga, jika ada interface yang menampung object konkrit dimana struct-nya tidak memiliki salah satu method yang terdefinisi di interface, error juga akan muncul. Intinya kembali ke aturan awal, variabel interface hanya bisa menampung object yang minimal memiliki semua method yang terdefinisi di interface-nya.

## 2. Embedded Interface
Interface bisa di-embed ke interface lain, sama seperti struct. Cara penerapannya juga sama, cukup dengan menuliskan nama interface yang ingin di-embed ke dalam interface tujuan.

Pada contoh berikut, disiapkan interface bernama hitung2d dan hitung3d. Kedua interface tersebut kemudian di-embed ke interface baru bernama hitung.

```go
package main

import "fmt"
import "math"

type hitung2d interface {
  luas() float64
  keliling() float64
}

type hitung3d interface {
  volume() float64
}

type hitung interface {
  hitung2d
  hitung3d
}
```
Interface hitung2d berisikan method untuk kalkulasi luas dan keliling, sedang hitung3d berisikan method untuk mencari volume bidang. Kedua interface tersebut diturunkan di interface hitung, menjadikannya memiliki kemampuan untuk menghitung luas, keliling, dan volume.

Next, siapkan struct baru bernama kubus yang memiliki method luas(), keliling(), dan volume().
```go
type kubus struct {
  sisi float64
}

func (k kubus) volume() float64 {
  return math.Pow(k.sisi, 3)
}

func (k kubus) luas() float64 {
  return math.Pow(k.sisi, 2) * 6
}
func (k kubus) keliling() float64 {
  return k.sisi * 12
}
```
object hasil cetakan struct kubus di atas, nantinya akan ditampung oleh object cetakan interface hitung yang isinya merupakan gabungan interface hitung2d dan hitung3d.

Terakhir, buat implementasi-nya di main.
```go 
func main() {
  var bangunRuang hitung = kubus{4}

  fmt.Println("===== kubus")
  fmt.Println("luas      :", bangunRuang.luas())
  fmt.Println("keliling  :", bangunRuang.keliling())
  fmt.Println("volume    :", bangunRuang.volume())
}
 ```
Bisa dilihat di kode di atas, lewat interface hitung, method luas, keliling, dan volume bisa di akses.



 

## Interface Kosong

Interface kosong atau empty interface yang dinotasikan dengan interface{} merupakan tipe data yang sangat spesial. Variabel bertipe ini bisa menampung segala jenis data, bahkan array, pointer, apapun. Tipe data dengan konsep ini biasa disebut dengan dynamic typing.

### 1. Penggunaan interface{}
interface{} merupakan tipe data, sehingga cara penggunaannya sama seperti pada tipe data lainnya, hanya saja nilai yang diisikan bisa apa saja. Contoh:
```go
package main

import "fmt"

func main() {
  var secret interface{}
  
  secret = "ethan hunt"
  fmt.Println(secret)

  secret = []string{"apple", "manggo", "banana"}
  fmt.Println(secret)

  secret = 12.4
  fmt.Println(secret)
}
```
Keyword interface seperti yang kita tau, digunakan untuk pembuatan interface. Tetapi ketika ditambahkan kurung kurawal ({}) di belakang-nya (menjadi interface{}), maka kegunaannya akan berubah, yaitu sebagai tipe data.



### 2. Casting Variabel Interface Kosong
Variabel bertipe interface{} bisa ditampilkan ke layar sebagai string dengan memanfaatkan fungsi print, seperti fmt.Println(). Tapi perlu diketahui bahwa nilai yang dimunculkan tersebut bukanlah nilai asli, melainkan bentuk string dari nilai aslinya.

Hal ini penting diketahui, karena untuk melakukan operasi yang membutuhkan nilai asli pada variabel yang bertipe interface{}, diperlukan casting ke tipe aslinya. Contoh seperti pada kode berikut.
```go
package main

import "fmt"
import "strings"

func main() {
  var secret interface{}

  secret = 2
  var number = secret.(int) * 10
  fmt.Println(secret, "multiplied by 10 is :", number)

  secret = []string{"apple", "manggo", "banana"}
  var fruits = strings.Join(secret.([]string), ", ")
  fmt.Println(fruits, "is my favorite fruits")
}
```
Pertama, variabel secret menampung nilai bertipe numerik. Ada kebutuhan untuk mengalikan nilai yang ditampung variabel tersebut dengan angka 10. Maka perlu dilakukan casting ke tipe aslinya, yaitu int, setelahnya barulah nilai bisa dioperasikan, yaitu secret.(int) * 10.

Pada contoh kedua, secret berisikan array string. Kita memerlukan string tersebut untuk digabungkan dengan pemisah tanda koma. Maka perlu di-casting ke []string terlebih dahulu sebelum bisa digunakan di strings.Join(), contohnya pada strings.Join(secret.([]string), ", ").



Teknik casting pada interface disebut dengan type assertions.

### 3. Casting Variabel Interface Kosong Ke object Pointer
Variabel interface{} bisa menyimpan data apa saja, termasuk data object, pointer, ataupun gabungan keduanya. Di bawah ini merupakan contoh penerapan interface untuk menampung data object pointer.
```go
package main

import (
  "fmt"
)

type person struct {
  name string
  age int
}

func main(){
  var secret interface{} = &person{name: "wick", age: 27}
  var name = secret.(*person).name
  fmt.Println(name)
}
```
Variabel secret dideklarasikan bertipe interface{} menampung referensi object cetakan struct person. Cara casting dari interface{} ke struct pointer adalah dengan menuliskan nama struct-nya dan ditambahkan tanda asterisk `(*)` di awal, contohnya seperti secret.`(*person)`. Setelah itu barulah nilai asli bisa diakses.

# PANIC, RECOVER, ERROR

# Panic

Panic function adalah function yang bisa kita gunakan untuk menghentikan program. Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan. Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi. berikut ini contoh penggunaannya:
```go
package main

import (
  "fmt"
)

func endApp(){
  fmt.Println("End App")
}

func runApp(error bool){
  defer endApp()
  if error{
    panic("ERROR")
  }
}

func main(){
  runApp(true)
}
```
pada kode diatas terlihat bahwa didalam function runApp terdapat kode panic("ERROR") ini berarti program akan di berhentikan dengan panic dan messagenya adalah "ERROR".


# Recover

Recover adalah function yang bisa kita gunakan untuk menangkap data panic. Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan. berikut contoh penggunaannya:
```go
package main

import (
  "fmt"
)

func endApp(){
  fmt.Println("End App")
  message := recover()
  fmt.Println("Terjadi Error", message)
}

func runApp(error bool){
  defer endApp()
  if error{
    panic("ERROR")
  }
}

func main(){
  runApp(true)
}
```
pada kode tersebut jika dibandingkan dengan contoh di panic, terdapat perbedaan, yaitu pada function endApp terdapat variabel message yang berisi recover(), recover() tersebut berfungsi untuk menangkap pesan error yang di keluarkan oleh panic.

 

# Error

Error merupakan sebuah tipe. Error memiliki 1 buah property berupa method Error(), method ini mengembalikan detail pesan error dalam string. Error termasuk tipe yang isinya bisa nil. cara menuliskan tipe data error cukup dengan menulis error (huruf kecil semua).

Di Go, banyak sekali fungsi yang mengembalikan nilai balik lebih dari satu. Biasanya, salah satu kembalian adalah bertipe error. Contohnya seperti pada fungsi strconv.Atoi(). Fungsi tersebut digunakan untuk konversi data string menjadi numerik. Fungsi ini mengembalikan 2 nilai balik. Nilai balik pertama adalah hasil konversi, dan nilai balik kedua adalah error. Ketika konversi berjalan mulus, nilai balik kedua akan bernilai nil. Sedangkan ketika konversi gagal, penyebabnya bisa langsung diketahui dari error yang dikembalikan.

Dibawah ini merupakan contoh program sederhana untuk deteksi inputan dari user, apakah numerik atau bukan. Dari sini kita akan belajar mengenai pemanfaatan error.
```go
package main

import (
  "fmt"
  "strconv"
)

func main() {
  var input string
  fmt.Print("Type some number: ")
  fmt.Scanln(&input)
  
  number, err := strconv.Atoi(input)
  
  if err == nil {
    fmt.Println(number, "is number")
  } else {
    fmt.Println(input, "is not number")
    fmt.Println(err.Error())
  }
}
```
Jalankan program, maka muncul tulisan "Type some number: ". Ketik angka bebas, jika sudah maka enter. Statement fmt.Scanln(&input) dipergunakan untuk men-capture inputan yang diketik oleh user sebelum dia menekan enter, lalu menyimpannya sebagai string ke variabel input.

Selanjutnya variabel tersebut dikonversi ke tipe numerik menggunakan strconv.Atoi(). Fungsi tersebut mengembalikan 2 data, ditampung oleh number dan err. Data pertama (number) berisi hasil konversi. Dan data kedua err, berisi informasi errornya (jika memang terjadi error ketika proses konversi).

Setelah itu dilakukan pengecekkan, ketika tidak ada error, number ditampilkan. Dan jika ada error, input ditampilkan beserta pesan errornya.

Pesan error bisa didapat dari method Error() milik tipe error.

## Membuat Custom Error

Selain memanfaatkan error hasil kembalian suatu fungsi internal yang tersedia, kita juga bisa membuat objek error sendiri dengan menggunakan fungsi errors.New() (harus import package errors terlebih dahulu).

berikut contoh penggunaanya:
```go
package main

import (
  "fmt"
  "errors"
)

func pembagian(nilai uint, pembagi uint)(float64, error){
  if pembagi == 0 {
    return 0.0, errors.New("Maaf pembagi tidak boleh NOL")
  }else{
    return float64(nilai/pembagi), nil
  }
}

func main(){
  hasil, err := pembagian(8,4)
  if err == nil{
    fmt.Println("Hasil", hasil)
  }else{
    fmt.Println("Error", err.Error())
  }
}
```
pada kode tersebut terdapat function pembagian dimana terdapat pengecekan pembagi sama dengan nol atau tidak jika pembaginya nol makan akan dikembalikan nilai errornya.

