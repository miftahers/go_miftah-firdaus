# STRING

  `Menghitung panjang string`
  Berapa Panjang string A?

    func main(){
      // len strings
      sentence := "Hello"
      lenSentence := len(sentence)
      fmt.Println(lenSentence)
    }

  `Compare strings`
  Apakah string A sama dengan string B?

    func main(){
      // compare strings
      str1 := "abc"
      str2 := "abd"
      fmt.Println(str1 == str2)
    }

  `Contain strings`
  Apakah string A mengandung string B?

    func main(){
      // Contains
      const (
        str = "something"
        substr = "some"
      )
      res := strings.Contains(str, substr)
      fmt.Println(res)
    }
  
  `Catch string pake index`
  Bagaimana mengambil kata "dog" dari "watchdog"?

    func main(){
      str := "watchdog"
      //index 01234567
      res := str[5:len(str)]
      fmt.Println(res)
    }
  
  `Replace strings`
  Bagaimana cara mengganti "katak" menjadi "kotak"?

    func main(){
      str := "katak"
      kotok := strings.Replace(str, "a", "o", -1) // Mengubah "katak" menjadi "kotok" karena parameter terakhir -1 jadi seluruh huruf/kata yang ditemukan dirubah
      kotak := strings.Replace(str, "a", "o", 1) //  Mengubah "katak" menjadi "kotak" karena parameter terakhir nya 1 jadi hanya 1 huruf pertama yang ditemukan
      fmt.Println("%s\n%s\n", kotok, kotak)
    }

  `Insert string`
  Bagaimana cara mengubah bh menjadi buah?

    func main(){
      str := "bh"
      // i=>  01
      res := str[:1] + "ua" + str[1:]
      fmt.Println(res)
    }
