# Advance Function

# Variadic function
Digunakan ketika parameter yang di harapkan berubah-ubah jumlahnya. Cara nya adalah dengan mengubah parameter function dengan slice.
    
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

# Anonymous Function == Literal Function
Adalah function yang tidak memiliki nama. Berguna ketika ingin membuat inline function.
  
  `1`
    
    func main(){

      func(){
        fmt.Println("Hello World")
        fmt.Println("Alterra Academy")
      }()

      fmt.Println("Hello World Bagian 2")
      fmt.Println("Alterra Academy Bagian 2")
    }

  `2 set to variable`

    func main(){

      bagianAwal := func(){
        fmt.Println("Hello World")
        fmt.Println("Alterra Academy")
      }

      fmt.Println("Hello World Bagian 2")
      fmt.Println("Alterra Academy Bagian 2")

      bagianAwal()
    }

  `3 with parameter`

    func main(){
      
      bagianKe := func(urutan string){
        fmt.Println("Hello World " + "bagian " + urutan)
        fmt.Println("Alterra Academy " + "bagian " + urutan)
      }

      fmt.Println("Hello World Bagian 2")
      fmt.Println("Alterra Academy Bagian 2")

      bagianKe("akhir")
    }

  `4 Closure`

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

# Defer Function
adalah function yang akan dieksekusi ketika parent function nya mengembalikan nilai. Multiple return juga bisa digunakan dan dijalankan sebagai stack(LIFO), satu per satu.
  `defer - simple function`

    func main(){
      defer fmt.Print(" 1 ")
      defer fmt.Print(" 2 ")
      fmt.Print(" 3 ")
      // Output: 3  2  1
    }

  `defer anonymous function`

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

defer biasanya digunakan saat open dan close database

    func main(){
        
      defer func() {
        fmt.Println(" Close DB ")
      }()
      fmt.Println(" Open DB ")
      
      // Jadi tidak akan lupa untuk close DB dan dijalankan setelah seluruh perintah dalam function selesai
    }