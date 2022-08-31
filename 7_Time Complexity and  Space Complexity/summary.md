
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
  3) O(sqrt(n))
  4) O(n) -  algoritma linear
  5) O(n log n) - algoritma n log n
  6) O(n^2) - algoritma kuadratik
  7) O(n^3) - algoritma kubik
  8) O(2^n) - algoritma eksponensial
  9) O(n!) - algoritma faktorial

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
