# Intro RESTful API

## API (Application Programming Interface)

API adalah seperangkat fungsi dan prosedur yang memungkinkan pembuatan aplikasi yang mengakses fitur atau data sistem operasi, aplikasi, atau layanan lainnya.

Hubungan antara Client dan Server ada pada 2 proses yaitu `Request` dan `Response`.

## Frontend & Backend Integration

Frontend mengirim request/upload -> backend menerima request dan mengelola request -> backend mengirim response ke frontend.

## Backend & Backend Integration

contoh: Service to Service

## What is REST?

REST  : REpresentational State Transfer

REST ITU
1. Menggunakan HTTP Protocol
2. Memiliki Endpoint, contoh: https://www.instagram.com/api/users
3. HTTP REQUEST METHOD
   1. GET: Minta data
   2. POST: Upload data
   3. PUT: Edit data
   4. DELETE: Hapus data
   5. HEAD
   6. OPTION
   7. PATCH
4. Request & Response format
   1. JSON
   2. XML
   3. SOAP

## JSON (Javascript Object Notation)
``` JSON
Contoh JSON:
{
  "key": "value"        <- string
  "umur": 17,           <- int
  "single": true,       <- bool
  "hobi": [             <- slice
    "belajar",
    "renang"
  ]
  "alamat": {           <- JSON in JSON
    "rumah": "malang",
    "no": 4,
    "rt": "03",
    "rw": "11"
  }
}
```

## HTTP Response Code

HTTP Response code berguna untuk mengidentifikasi apakah request itu berhasil atau tidak.

1. Response code 200: OK
2. Response code 201: Created (Berhasil PUT/POST)
3. Response code 400: Bad Request (Terjadi kegagalan saat POST/PUT. Contoh nya jika kita tidak lolos pada proses validasi)
4. Response code 404: Not Found (Resource yang diminta tidak ada)
5. Response code 401: Unauthorized (Mengakses tanpa autentikasi/tidak memiliki izin. contoh nya saat kita akses tapi belum login)
6. Response code 405: Method Not Allowed (Method yang digunakan tidak sesuai dengan endpoint)
7. Response code 500: Internal Server Error

## API TOOLS

Contoh tools testing API:
- Katalon
- Apache JMeter
- SoapUI
- Postman
- dll.

### POSTMAN

Postman adalah http client untuk testing web services.

## REST API Deisgn Best Practice

### User Nouns Instead of Verbs
Karena perintah dari REST API sudah kata kerja jadi gausah pake katakerja pada saat membuat endpoint

### Naming using Plural Nouns
Disarankan untuk menggunakan kata jamak dibandingkan kata tunggal
```
// BENAR:
GET /cars/123
POST /cars
GET /cars

// SALAH
GET /car/123
POST /car
GET /car
```

### Use Resource Nesting to show relations or hierarchy

```
// Contoh

/users <- menunjukan daftar user
/users/123 <- menunjuk user yang spesifik
/users/123/orders <- menunjuk daftar order yang dilakukan user tertentu
/users/123/orders/0001 <- menunjuk order yang spesifik yang dilakukan oleh user tertentu

```

### Contoh Format Response JSON

#### Contoh pertama
``` JSON
{
  "code"      : 200,
  "message"   : "Success get data",
  "status"    : "success",
  "data"      : {
    ...
  }
}
```

#### Contoh kedua
``` JSON
{
  "meta"            : {
    "error_type"    : "OAuthException",
    "code"          : 400,
    "error_message" : "..."
  },
  "data"            : {
    ...
  },
  "pagination"      : {
    "next_url"      : "...",
    "next_max_id"   : "13872296"
  }
}
```

### Filtering, Sorting, Paging, and Field Selection

#### Filtering
```
GET /users?country=USA
GET /users?creaton_date=2019-11-11

```
#### Sorting
```
GET /users?sort=birthdate_date:asc
GET /users?sort=birthdate_date:desc

```
#### Paging
```
GET /users?limit=100
GET /users?offset=2

```

#### Field Selection
```
GET /users/123?fields=username,email <- for one spesific user
GET /users?fields=username,email <- for a full list of users

```

### Handle Trailing Slashes Gracefully
```
(1) GET /cars
(2) GET /cars/
```
API Kita harus bisa menghandle endpoint keduanya

### Versioning

```
// Contoh API:
https://us6.api.mailchimp.com/3.0/ (major + minor version indication)
https://api.stripe.com/v1 (major version indication only)
https://developer.github.com/v3/ (major version indication only)

```

# JSON data

JSON atau Javascript Object Notation adalah notasi standar yang umum digunakan untuk komunikasi data dalam web. JSON merupakan subset dari javascript.

Go menyediakan package encoding/json yang berisikan banyak fungsi untuk kebutuhan operasi json.

## Decode JSON Ke Variabel Object Struct
Di Go, data json dituliskan sebagai string. Dengan menggunakan json.Unmarshal, json string bisa dikonversi menjadi bentuk object, entah itu dalam bentuk map[string]interface{} ataupun object struct.

Program berikut ini adalah contoh cara decoding json ke bentuk object. Pertama import package yang dibutuhkan, lalu siapkan struct User.
```go
package main

import "encoding/json"
import "fmt"

type User struct {
  FullName string `json:"Name"`
  Age int
}
```
Struct user ini nantinya digunakan untuk membuat variabel baru penampung hasil decode json string. Proses decode sendiri dilakukan lewat fungsi json.Unmarshal(), dengan json string tersebut dimasukan ke statement fungsi tersebut.

``` go
func main() {
  var jsonString = `{"Name": "john doe", "Age": 27}`

  var jsonData = []byte(jsonString)

  var data User
  
  var err = json.Unmarshal(jsonData, &data)
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Println("user :", data.FullName)
  fmt.Println("age  :", data.Age)
}
```
Fungsi unmarshal hanya menerima data json dalam bentuk []byte, maka dari itu data json string pada kode di atas di-casting terlebih dahulu ke tipe []byte sebelum dipergunakan pada fungsi unmarshal.

Juga, perlu diperhatikan, argument ke-2 fungsi unmarshal harus diisi dengan pointer dari object yang nantinya akan menampung hasilnya.

![](https://blog.sanbercode.com/wp-content/uploads/2021/09/materi-go-hari13-1.png)

Jika kita perhatikan lagi, pada struct User, salah satu property-nya yaitu FullName memiliki tag json:"Name". Tag tersebut digunakan untuk mapping informasi json ke property yang bersangkutan.

Data json yang akan diparsing memiliki 2 property yaitu Name dan Age. Kebetulan penulisan Age pada data json dan pada struktur struct adalah sama, berbeda dengan Name yang tidak ada pada struct.

Dengan menambahkan tag json, maka property FullName struct akan secara cerdas menampung data json property Name.

Pada kasus decoding data json string ke variabel object struct, semua level akses property struct penampung harus publik.

## Decode JSON Ke map[string]interface{} & interface{}
Tak hanya ke object cetakan struct, target decoding data json juga bisa berupa variabel bertipe map[string]interface{}.
```go
var data1 map[string]interface{}
json.Unmarshal(jsonData, &data1)

fmt.Println("user :", data1["Name"])
fmt.Println("age  :", data1["Age"])
```
Variabel bertipe interface{} juga bisa digunakan untuk menampung hasil decode. Dengan catatan pada pengaksesan nilai property, harus dilakukan casting terlebih dahulu ke map[string]interface{}.
```go
var data2 interface{}
json.Unmarshal(jsonData, &data2)

var decodedData = data2.(map[string]interface{})
fmt.Println("user :", decodedData["Name"])
fmt.Println("age  :", decodedData["Age"])
```
## Decode Array JSON Ke Array Object
Decode data dari array json ke slice/array object masih sama, siapkan saja variabel penampung hasil decode dengan tipe slice struct. Contohnya bisa dilihat pada kode berikut.
```go
var jsonString = `[
    {"Name": "john doe", "Age": 27},
    {"Name": "doe john", "Age": 32}
]`

var data []User

var err = json.Unmarshal([]byte(jsonString), &data)
if err != nil {
    fmt.Println(err.Error())
    return
}

fmt.Println("user 1:", data[0].FullName)
fmt.Println("user 2:", data[1].FullName)
```
## Encode object Ke JSON String

Setelah sebelumnya dijelaskan beberapa cara decode data dari json string ke object, sekarang kita akan belajar cara encode data object ke bentuk json string.

Fungsi `json.Marshal digunakan untuk encoding data ke json string`. Sumber data bisa berupa variabel object cetakan struct, map[string]interface{}, atau slice.

Pada contoh berikut, data slice struct dikonversi ke dalam bentuk json string. Hasil konversi berupa []byte, casting terlebih dahulu ke tipe string agar bisa ditampilkan bentuk json string-nya.
```go
var object = []User{{"john doe", 27}, {"doe john", 32}}
var jsonData, err = json.Marshal(object)
if err != nil {
    fmt.Println(err.Error())
    return
}

var jsonString = string(jsonData)
fmt.Println(jsonString)
```
Output:
![](https://blog.sanbercode.com/wp-content/uploads/2021/09/materi-go-hari13-2.png)


 

# Web Server

Go menyediakan package net/http, berisi berbagai macam fitur untuk keperluan pembuatan aplikasi berbasis web. Termasuk di dalamnya web server, routing, templating, dan lainnya.

Go memiliki web server sendiri, dan web server tersebut berada di dalam Go, tidak seperti bahasa lain yang servernya terpisah dan perlu diinstal sendiri (seperti PHP yang memerlukan Apache, .NET yang memerlukan IIS).

## Membuat Aplikasi Web Sederhana
Package net/http memiliki banyak sekali fungsi yang bisa dimanfaatkan. Di bagian ini kita akan mempelajari beberapa fungsi penting seperti routing dan start server.

Program di bawah ini merupakan contoh sederhana untuk memunculkan text di web ketika url tertentu diakses.
```go
package main

import "fmt"
import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "apa kabar!")
}

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "halo!")
  })
   
  http.HandleFunc("/index", index)
  
  fmt.Println("starting web server at http://localhost:8080/")

  http.ListenAndServe(":8080", nil)
}
```
Jalankan program tersebut.

![""](https://blog.sanbercode.com/wp-content/uploads/2021/09/materi-go-hari13-3.png)


Jika muncul dialog Do you want the application “main” to accept incoming network connections? atau sejenis, pilih allow. Setelah itu, buka url http://localhost/ dan http://localhost/index lewat browser.



Fungsi `http.HandleFunc()` digunakan untuk routing aplikasi web. Maksud dari routing adalah penentuan aksi ketika url tertentu diakses oleh user.

Pada kode di atas 2 rute didaftarkan, yaitu / dan /index. Aksi dari rute / adalah menampilkan text "halo" di halaman website. Sedangkan /index menampilkan text "apa kabar!".

!["contoh"](https://blog.sanbercode.com/wp-content/uploads/2021/09/materi-go-hari13-4.png)

Fungsi `http.HandleFunc()` memiliki 2 buah parameter yang harus diisi. Parameter pertama adalah rute yang diinginkan. Parameter kedua adalah callback atau aksi ketika rute tersebut diakses. Callback tersebut bertipe fungsi func(w http.ResponseWriter, r *http.Request).

Pada pendaftaran rute /index, callback-nya adalah fungsi index(), hal seperti ini diperbolehkan asalkan tipe dari fungsi tersebut sesuai.

Fungsi `http.listenAndServe()` digunakan untuk menghidupkan server sekaligus menjalankan aplikasi menggunakan server tersebut. Di Go, 1 web aplikasi adalah 1 buah server berbeda.

Pada contoh di atas, server dijalankan pada port 8080.

Perlu diingat, setiap ada perubahan pada file .go, go run harus dipanggil lagi.

Untuk menghentikan web server, tekan CTRL+C pada terminal atau CMD, dimana pengeksekusian aplikasi berlangsung.



# Web Service API Server

Web Service API adalah sebuah web yang menerima request dari client dan menghasilkan response, biasa berupa JSON/XML.

## Membuat API dengan Method GET
Method GET paling sering digunakan untuk menampilkan sebuah data json. Maka saya akan berikan contoh pertama dengan method GET.

Method GET dapat di akses melalui browser atau tools api lainnya. Selain itu juga dapat di ambil dari website lainnya, dengan catatan harus mengaktifkan cors.
```go
package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
)

type Movie struct {
  ID    int    `json:"id"`
  Title string `json:"title"`
  Year  int    `json:"year"`
}

func Movies() []Movie {
  movs := []Movie{
    {1, "Spider-Man", 2002},
    {2, "John Wick", 2014},
    {3, "Avengers", 2018},
    {4, "Logan", 2017},
  }
  return movs
}

// GetMovies
func getMovies(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
    movies := Movies()
    dataMovies, err := json.Marshal(movies)
    
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(dataMovies)
    return
  }
  http.Error(w, "ERROR....", http.StatusNotFound)
}

func main() {
  http.HandleFunc("/movies", getMovies)
  fmt.Println("server running at http://localhost:8080")

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
}
```
Setelah web server sudah berjalan, web service yang telah dibuat perlu untuk di-tes. Salah satu tools yang digunakan untuk mengecek web service API adalah postman

Hasilnya :
![](https://blog.sanbercode.com/wp-content/uploads/2021/09/materi-go-hari14-1.png)


Perhatikan kode di atas, terdapat http server yang mengalankan fungsi GetMovies(). Dimana di dalam fungsi tersebut akan memeriksa dahulu method yang di akses, apabila method itu GET maka akan di teruskan ke baris kode berikutnya.

Kode baris berikutnya memanggil method Movies(), dimana nilai balik dari fungsi tersebut adalah data array dari struct Movie. Setelah itu maka data tersebut di ubah format ke dalam tipe byte dengan sintaks marshal.

Untuk menampilkan ke halaman browser menggunakan kode w.Write(dataMovie). Agar Format JSON harus menambahkan sintaks w.Header().Set("Content-Type", "application/json").

Sedangkan untuk sintaks w.WriteHeader(http.StatusOK) digunakan untuk memeberikan kode status header. Jika belum paham mengenai kode status http bisa membaca di https://restfulapi.net/http-status-codes/.

## Membuat API dengan Method POST
Kalau menggunakan Method POST ini tidak bisa langsung di akses di browser, tapi dapat menggunakan form HTML atau Postman sebagai tools. Di karenakan kita tidak membuat website namun membuat API Web Service dengan Golang maka untuk Uji Coba saya menggunakan `Postman`.

Sebelum ke kode program ada hal yang perlu kita pahami tentang input data post. Secara umum dan sering di pakai dapat menggunakan 2 macam cara untuk menerima data yaitu dengan form dan teks json, kedua hal ini biasa di sebut `content type`. Cara menerima datanya pun berbeda.

```go
// PostMovie
func PostMovie(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var Mov Movie
  if r.Method == "POST" {
    if r.Header.Get("Content-Type") == "application/json" {
      // parse dari json
      decodeJSON := json.NewDecoder(r.Body)
      if err := decodeJSON.Decode(&Mov); err != nil {
        log.Fatal(err)
      }
    } else {
      // parse dari form
      getID := r.PostFormValue("id")
      id, _ := strconv.Atoi(getID)
      title := r.PostFormValue("title")
      getYear := r.PostFormValue("year")
      year, _ := strconv.Atoi(getYear)
      Mov = Movie{
         ID:    id,
         Title: title,
         Year:  year,
      }
    }

    dataMovie, _ := json.Marshal(Mov) // to byte
    w.Write(dataMovie)                // cetak di browser
    return
  }

  http.Error(w, "NOT FOUND", http.StatusNotFound)
  return
}

// pada function main

http.HandleFunc("/post_movie", PostMovie)
```


Percobaan pada kode diatas dilakukan 2 kali yaitu dengan menggunakan raw data json dan form-data.


Uji coba dengan JSON

![](https://blog.sanbercode.com/wp-content/uploads/2021/09/materi-go-hari14-2.png)


Uji coba dengan form-data

![](https://blog.sanbercode.com/wp-content/uploads/2021/09/materi-go-hari14-3.png)


Routing "/post_movie" mengakses fungsi dengan nama PostMovie. Sama seperti hal nya dengan method GET untuk memeriksa method yang di akses, namun di ganti dengan method POST.

Untuk menangkap tipe konten yang di inputkan dapat menggunakan kode `r.Header.Get(“Content-Type”) `. Jika tipe tersebut json untuk menerima data harus menggunakan tipe data maupun struct di tandai dengan kode `decodeJSON := json.NewDecoder(r.Body)` dan decodeJSON.Decode(&Mov) agar masuk struct dengan nama Mahasiswa.

Jika data tersebut bukan json maka untuk mengambil data nya dapat menggunakan sintaks PostFormValue dan di dalam nya di tulis nama field nya. Data kembalian terhadap fungsi itu yaitu string, apabila ingin mengubah ke bentuk tipe data integer dapat menggunakan manipulasi string yaitu Atoi.