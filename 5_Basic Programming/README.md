# Introduction Golang
  Golang dikembangkan `google`, golang sendiri merupakan bahasa pemrograman yang `easy to build`, `simple`, `reliable`, dan `efficient software`.
  Golang bagus digunakan untuk membuat lower-level program yang menyediakan layanan untuk sistem-sistem yang lain, atau disebut `system programming`.
  Penggunaan golang biasanya untuk:
  1. Aplikasi: E-commerce, Music Player, Social Media Apps
  2. System Programs: APIs, Game Engines, CLI apps

  Alasan harus pakai golang:
  - Simple
  - dikompile mirip bahasa pemrograman `statically typed` seperti C, dan `dynamically typed` seperti Javascript
  - lightweight syntax
  - ada concurrency bawaan yang bisa digunakana untuk computing skala besar dan perangkat keras <i>multicore</i>
  - open source
  - digunakan perusahaan besar

  Perusahaan yang menggunakan go: `Uber`, `Google`, `Medium`, `Grab`, `Tokopedia`, `Alterra`, dsb.

## Variables and Types

  Variables:
  - Booleans
  - Numeric
    - int/uint (int8,int16,int32,int64,rune, dsb)
    - Float (Float32, Float64)
    - Complex (Complex64, Complex128)
  - String
  
  `Variable Declaration`
``` go
// var <variable_name> <data_type>
  var age int
  var month string
// var <variable_name> <data_type> = <value>
  var age int = 20
  var prima bool = true
  var name string = "Miftah"
// var <list_variable_name> <data_type>
  var name, address string
// var <list_variable_name> <data_type> = <value>
  var name, city string = "Miftah", "Sumedang"
// <variable_name> := <value>
  name := "Miftah"
```
  `Single Constant`
```go
const pi float64 = 3.14
```
  `multple Constant`
```go
  const (
    pi float64 = 3.14
    pi2
    age = 20
  )
```
## Operator Expression

`Operator`
```go
a := b + c //Summation
a := b - c //Substraction
a := b * c //Multiplication
a := b / c //Distribution
a := b % c //Modulus
```
`String Operation`
```go
helloWorld := "Hello" + " " + "world"
fmt.Println("Hallo" + " " + "Dunia")
```

## Branching

`If-else`
```go
if a>2 {
fmt.Println(a, " is more than 2")
} else if a>5 {
fmt.Println(a, " is more than 2 and 5")
} else {
fmt.Println(a, " is below 2")
}
```
`Switch`
```go
a := 0
switch a {
case 0:
  fmt.Println(a, "is false")
  fallthrough
case 1:
  fmt.Println(a, "is true")
default:
  fmt.Println(a, "is not bool num")
}
```
karena golang otomatis melakukan break pada switch-case maka diperlukan fallthrough untuk mengecek kondisi selanjutnya.

## Looping

`LOOPING BIASA`
```go
for i := 0; i < 5; i++ {
fmt.Println(i)
}
```
`LOOPING dengan Continue dan Break`
```go
for i := 0; i <= 4; i++ {
  if i > 0 {
    fmt.Println ("i lebih dari 0")
    continue
  }
  if i > 3 {
    fmt.Println ("i lebih dari 3")
    break
  }
}
```
`LOOPING OVER STRING`
```go
sentence := "kata"
// 1 - len
for i := 0; i < len(sentence); i++ {
  fmt.Printf(string(sentence[i]) + " ")
}  

// 2 - range
for pos, char := range sentence {
  fmt.Printf("character %c start at byte position %d", char, pos)
}
```