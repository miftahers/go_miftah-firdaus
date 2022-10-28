# Recursive-NumberTheory-Sorting-Searching

## Recursive
`Pengantar Rekursif`
Rescursion adalah keadaan dimana sebuah fungsi (func) memanggil dirinya sendiri untuk menyelsaikan sebuah masalah.
Jika permasalahan yang dikerjakan sederhana maka fungsi yang rekursif akan bisa mengeluarkan jawaban yang lebih dan jika permasalahan yang dikerjakan terlalu besar maka fungsi tersebut akan melakukan rekursi akan memperkecil cakupan masalah.

`Keuntungan menggunakan rekursif`
- Banyak permasalahan akan mudah dilakukan jika menggunakan fungsi yang rekursif dan juga kode yang ditulis akan lebih rapi
- Dengan rekursif akan lebih mudah melihat dan mendesain langkah penyelsaian suatu masalah

`Hal yang perlu diperhatikan ketika menggunakan rekursif`
Apa yang menjadi solusi paling sederhana untuk kasus tersebut?
Jika tidak dapat menyelsaikan permasalahan dengan satu kali rekursi maka carilah rekursif yang memiliki kemiripan dengan masalah yang dihadapi dan manfaatkan untuk menyelsaikan masalah kita.

`Contoh penggunaan rekursif`
```go
// Mencari Faktorial
func faktorial(num int) int {
  // (check) Jika nilai sama dengan 1 maka faktorialnya 1
  if num == 1 {
    return 1
  } else {
    // (multiplication) Jika nilai bukan sama dengan 1 maka #kalikan dengan (num-1) dengan terlebih dahulu cek apakah num == 1
    return num * factorial(num-1)
  }
}
```

## NumberTheory
Number Theory adalah cabang ilmu matematika yang meneliti integer.

`Contoh NumberTheory`
```go
// Faktorial
func faktorial(num int) int {
  if num == 1 {
    return 1
  } else {
    return num * factorial(num-1)
  }
}

// Bilangan Prima
func isPrime(number int) bool {
if number == 1 {
  return false
}
for i := 2; i*i < number; i++ {
  if number%i == 0 {
    return false
  }
}
return true
}

// Faktor Persekutuan Terbesar (Greatest Common Divisor)
func gcd(a, b int) int {
  if(b == 0) {
    return a
  }
  return gcd(b, a % b)
}

// Kelipatan Persekutuan Terkecil (Least Common Multiple)
func lcm(a,b int) int {
  return a* (b / gcd(a, b))
}
```
## Searching

`Linear Search O(n)`
```go
func linearSearch(arr []string, x int) int {
  for i := 0; i < len(arr); i++ {
    if arr[i] == x {
      return i
    }
  }
  // define return -1 as not found in main()
  return -1
}
```
`Golang Builtins - sort.SearchInts()`
```go
elements := []ints{1,2,3,4,5,6,7}
value := 4
findIndex := sort.SearchInts(elements, value)
if value == elements[findIndex] {
  fmt.Println("Nilai ditemukan pada elements")
} else {
  fmt.Println("Nilai tidak ditemukan pada elements)
}
```
# Sorting
Sorting atau mengurutkan adalah proses menyusun data menjadi terurut dengan ketentuan tertentu.

`Selection sort - O(n^2)`
Cari element dengan nilai terkecil dan tukar dengan element pertama dari array. Kemudian cari element dengan nilai terkecil kedua dan tukar dengan element kedua dari array dan ulangi seperti itu.
```go
func selectionSort(elements []int) []int {
  n := len(elements)
  for k := 0; k < n; k++ {
    minimal := k
    for j := k + 1; j < n; j++ {
      if elements[j] < elements[minimal] {
        minimal = j
      }
    }
    elements[k], elements[minimal] = elements[minimal], elements[j]
  }
  return elements
}
```
`Counting sort - O(n+k)`
Hitung elements di array count, lakukan perulangan melewati array count secara increment.
```go
func countingSort(elements []int, k int) []int {
  count := make([]int, k + 1)
  for i := 0; i < len(elements); i++{
    count[elements[i]]++
  }
  counter := 0
  for i := 0; i < k +1; i++ {
    for j := 0; j < count[i]; j++ {
      elements[counter] = i
      counter++
    }
  }
  return elements
}
```
`Golang Builtins - sort.Strings()`
```go
func main() {
  // sort strings
  strs := []strings{c, a, b}
  sort.Strings(strs)
  fmt.Println("Sorted Strings: ", strs)

  // sort Ints
  ints := []int {3, 4, 6, 1, 3, 3}
  sort.Ints(ints)
  fmt.Println("Sorted Ints: ", ints)

  // Check if string/ints already sorted
  s := sort.IntsAreSorted(ints)
  fmt.Println("Sorted: ", s)
}
```