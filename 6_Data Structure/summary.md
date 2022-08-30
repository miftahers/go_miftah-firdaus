# Data structure (Big-O-Time, Array, Slice, Map dan Function)


# Algoritma dan Big-O-Time

Dalam pembuatan sebuah algoritma atau langkah untuk melakukan perhitungan komputer perlu diperhatikan efektivitas dan efisiensinya. Efektivitas dan efisiensi yang dimaksud adalah penggunaan waktu pemrosesan (Time Complexity) dan ruang pada memori (Space Complexity). Adapun Big-O-Time adalah suatu metode penggambaran operation/operasi terhadap variabel yang terlibat di dalamnya.
  
  # Time Complexity
  
  `constant time - O(1)`
  
    // Contoh
    func dominant (n int) int {
      var result int = 0
      result = result + 10
      return result
    }

  `Linear time - O(n)`

    // Contoh
    func linear (n int, a []int) int {
      for i := 0;  i < n; i++ {
        if a[i] == 0 {
          return 0
        }
      }
      return 1
    }

  `Linear time - O(n+m)`

    // Contoh
    func linear2 (n int, m int) int {
      var result int = 0
      for i := 0; i < n; i++ {
        result++
      }
      for j := 0; j < m; j++ {
        result++
      }
      return result
    }

  `Logarithmic time - O(log(n))`

    // Contoh
    func logarithmic (n int) int {
      var result int = 0;
      for n > 1 {
        n /= 2
        result++
      }
      return result
    }
    // log(32) = 5

  `Quadratic time - O(n^2)`

    // Contoh
    func quadratic(n int) int {
      var result int = 0
      for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
          result++
        }
      }
    return result
    }

  Urutan dari yang paling cepat ke paling lambat:
  1) O(1) - algotima konstan
  2) O(log n) - algoritma logaritmik
  3) O(n) -  algoritma linear
  4) O(n log n) - algoritma n log n
  5) O(n^2) - algoritma kuadratik
  6) O(n^3) - algoritma kubik
  7) O(2^n) - algoritma eksponensial
  8) O(n!) - algoritma faktorial

  `Time limit`
  
  Rata-rata kecepatan komputasi jaman sekarang adalah 10^8 operasi per kurang dari satu detik. Batas waktu untuk tes secara daring biasanya adalah 1 sampai dengan 10 detik.
  Berikut adalah perkiraan jumlah data dan ekspektasi dari time complexity yang mesti digunakan:
  - n <= 1.000.000, yang mesti digunakan adalah O(n) atau O(n log n)
  - n <= 10.000, yang mesti digunakan adalah O(n^2)
  - n <= 500, yang mesti digunakan adalah O(n^3)
  *list diatas hanyalah perkiraan, time complexity sebaiknya disesuaikan dengan tugas yang harus dipenuhi.
  *Proses harus diproses kurang dari 2 detik untuk memberikan kenyamanan psikologis untuk pengguna

  # space complexity

  Pada space complexity kita bergantung pada variabel dan tipe data yang digunakan, sebisa mungkin kita menggunakan sedikit variabel untuk mengurangi penggunaan memori.
  
  `constant O(1)`
  
  Variabel dideklarasikan jumlahnya konstan sehingga hanya menggunakan sejumlah alamat memori dan tidak bertambah jumlahnya karena bersipat hanya mengubah nilai tapi tidak menambah data.
  
    `Contoh`
    
      func penjumlahan(a int, b int) (c int) {
        c = a + b
      }
      // Variable yang digunakan hanya a, b, dan c sehingga hanya menggunakan tiga data saja.

  `linear O(n)`

  variabel yang dideklarasikan berupa array,slice atau map yang jumlah datanya bisa banyak.

    `contoh`

      func addFoods (foodName string, bag []string) []string {
        bag = append(foods, foodName)
        return bag
      }
      // Data yang ditambahkan bisa banyak sehingga menggunakan banyak memori

  Jadi, untuk membuat penggunaan penyimpanan/memori lebih efisien gunakan array/slice/map hanya pada kondisi terburuk saja.

# Array

Array adalah `data struktur` yang digunakan untuk menyimpan lebih dari satu data pada satu variabel dengan `tipe data sejenis` dan harus `dideklarasikan kapasitasnya`.

`cara mendeklarasikan array`
    
    // Single Dimensional Array
    
    // 1
    var bilanganGanjil [5]int

    // cara menginput nilainya:
    bilanganGanjil[0] = 1
    bilanganGanjil[1] = 3
    
    // 2
    var bilanganGanjil [5]int {1, 3, 5, 7, 9}
    
    //Multi Dimensional Array

    array := [3][3]string {
      {"Kacang polong", "Gandum", "Padi"},
      {"Kedondong", "Mangga", "Pepaya"},
      {"Kopi", "Jus Buah", "Susu kedelai"},
    }

`cara menampilkan data array`
    
    array := [3]int{1, 2, 3}
    
    // looping - len()
    for i := 0; i < len(array); i++ {
      fmt.Printf("index [%d] => %d", i, array[i])
    }

    // looping - range
    for index, element := range array {
      fmt.Printf("index [%d] => %d", index, element)
    }

# Slice

Slice adalah `struktur data` yang digunakan untuk menyimpan lebih dari satu data pada satu variabel dengan `tipe data sejenis` dan `kapasitas nya tidak perlu dideklarasikan`. 

Slice `bukan`lah `array yang dinamis` tetapi slice adalah `tipe yang menggunakan referensi.` Jadi, jika kita menyalin nilai dari array untuk dijadikan slice maka apabila nilai array nya dirubah maka nilai dari slice yang telah dibuat dari array tersebut akan ikut berubah.

Slice memiliki panjang(`len`) dan kapasitas(`cap`). Panjang/length merupakan banyaknya nilai/value pada slice, sedangkan kapasitas/Capacity adalah banyaknya ruang pada slice untuk diisi dengan data.

`cara mendeklarasikan slice`

    // 1
    var even_number []int
    // 2
    var odd_number []int{1, 3, 5, 7, 9}
    // 3
    numbers := []int{1, 2, 3, 4, 5}
    // 4 - Make
    // var <variable> = make([]<data_type>, length, capacity)
    // or
    // <variable> := make([]<data_type>, length)
    var primes = make([]int, 5, 10)

`Append dan Copy`

    // Append - Berfungsi untuk menambah data ke dalam slice
    var fruits []string{"Melon", "Coco", "Strawberry"}
    fmt.Println(fruits)
    fruits = append(fruits, "Apple", "Banana")
    fmt.Println(fruits)

    // copy - berfungsi untuk menyalin data dari slice/array ke dalam sebuah slice
    // <variable> := copy(<copied_slice variable>, <slice variable>)
    copied_fruits := make([]string, 2)
    copy := copy(copied_slice, fruits)
    fmt.Println("Copied Slice: ", copied_slice)
    fmt.Println("Total copied number of slice: ", copy)

# Map

Map adalah `struktur data` yang digunakan untuk menyimpan lebih dari satu data pada sebuah variabel yang memiliki `key` dan `value` yang keduanya saling berkaitan dimana `key` itu memiliki sifat unik. 

Contoh relasi antara `key` dan `value`
  
  `map`

    // 1
    rataNilai := make(map[string]int)
    rataNilai["Rian"] = 80
    rataNilai["Udin"] = 79
    rataNilai["Samsudin"] = 80

    for key, value := range rataNilai {
      fmt.Printf("Rataan nilai %s = %d", key, value)
    }


    // 2
    harga := map[string]int {
      "Siomay": 5000,
      "Batagor": 10000,
      "Baso Ikan": 5000,
    }
    
    fmt.Println("=== DAFTAR HARGA WARUNG BATOK ===")
    for key, value := range harga {
      fmt.Printf("%s = %d", key, value)
    }

*Key `Rian` memiliki nilai/value `80` dan key `"Rian" tidak dapat memiliki 2 nilai`, nilai/value 80 bisa dimiliki oleh siapa saja dan `tidak boleh ada key "Rian" lagi pada map tersebut` karena sifatnya unik.


# Function
Function adalah sekumpulan kode yang dikelompokan dengan nama tertentu. Function diperlukan untuk membuat kodingan yang clean, tidy dan modular (Bisa dicabut-pasang).

`Simple function`
    
    // Without Parameter
    func sayHello() {
      fmt.Print("Hello, ")
    }

    // With Parameter
    func greeting(hour int) {
      if hour < 12 {
        fmt.Println("Selamat Pagi!")
      } else if hour < 18 {
        fmt.Println("Selamat Sore!")
      } else{
        fmt.Println("Selamat Malam!")
      }
    }

    func main() {
      //Without Parameter
      sayHello()
      // With Parameter
      greeting(11)
    }

`Function with Return Value (Fungsi dengan nilai kembalian)`

    // Single return value
    func calculateSquare(side int) int {
      return side*side
    }
    // Multiple return value
    func calculateCircle(diameter float64)(keliling float64, luas float64) {
      keliling := math.Pi * diameter
      luas := math.Pi * math.Pow(diameter/2,2)
      return
    }

    func main(){
      // Luas Persegi
      side:= 5
      luasPersegi := calculateSquare(side)
      fmt.Println("Keliling persegi = ", luasPersegi)

      // Keliling & Luas Lingkaran
      var diameter float64 = 10
      kelilingLingkaran, luasLingkaran := calculateCircle(diameter)
      fmt.Printf("Keliling lingkaran = %.2f dan luas lingkaran = %.2f", kelilingLingkaran, luasLingkaran)
    }
