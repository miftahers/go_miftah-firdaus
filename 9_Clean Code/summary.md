# CLEAN CODE

# Apa itu "Clean Code"?
Clean Code adalah istilah untuk kode yang mudah dibaca, difahami dan diubah oleh programmer.

    "Working code isn't necessary good code. Your code also need to be easy to read understand, and modify"
    ~ Brandon Gregory

    "Any fool can write code that a computer can understand. Good programmers write code that humans can understand."
    ~ Martin Fowler

# Kenapa "Clean Code"?
- Work Collaboration
- Feature Development
- Faster Development

# Karakteristik "Clean Code"

    "There are only two hard things in Computer Science: cache invalidation and naming things"
    ~ Marten Fowler

`1. Mudah difahami - Penamaan mudah difahami`

  `Contoh salah`

    let b = 125.0;
    let data = [2,3,5,7]

    const locations = ["Austin", "New York", "San Francisco"];
    locations.forEach((1)=> {
      doStuff();
      doSomeOtherStuff();
      // .
      // ..
      // ...
      // Sebentar, 1 itu untuk apa lagi?
      dispatch(1);
    });

  `Contoh Benar`

    let userBalance = 125.0;
    let primeList = [2,3,4,7];

    const locations = ['Austin', 'New York', 'San Francisco'];
    locations.forEach((location) => {
    doStuff();
    doSomeOtherStuff();
    // ...
    // ...
    // ...
    dispatch(location);
    });

`2. Mudah dieja dan dicari`

  `Contoh Salah`

    const yyyymmdstr = moment().format('YYYY/MM/DD');
    let fName = "Ismawul Umam";
    let dvdr = 3;
    // Nah kan, 86400000 itu untuk apa?
    setTimeout(blastOff, 86400000);

  `Contoh Benar`

    const currentDate = moment().format('YYYY/MM/DD');
    let fullName = "Iswanul Umam";
    let divider = 3;
    // Tulis const dalam huruf kapital dan secara global
    const MILLISECONDS_IN_A_DAY = 86400000;
    setTimeout(blastOff, MILLISECONDS_IN_A_DAY);

`3. Singkat namun mendeskripsikan konteks`

  `Contoh Salah`

    function inv (user) {}

    const address = 'One Infinite Loop, Cupertino 95014';
    const cityZipCodeRegex = /^[^,\\]+[,\\\s]+(.+?)\s*(\d{5})?$/; 
    saveCityZipCode(address.match(cityZipCodeRegex)[1], address.match(cityZipCodeRegex)[2]);

  `Contoh Benar`

    function inviteUSer(emailAddress){}

    const address = 'One Infinite Loop, Cupertino 95014';
    const cityZipCodeRegex = /^[^,\\]+[,\\\s]+(.+?)\s*(\d{5})?$/; 
    const [, city, zipCode] = address.match(cityZipCodeRegex) || [];
    saveCityZipCode(city, zipCode);

`4. Konsisten`

  `Contoh Salah`

    getUserId();
    getClientName();
    getCUstomerRecord();

    const DAY_IN_WEEK = 7;
    const daysInMonth = 30;

    const songs = ['Back In Black'. 'Stairway to Heaven', 'Hey Jude'];
    const Artist = ['ACDC', 'Led Zeppelin', 'The Beatles'];

    function eraseDatabase(){}
    function restore_database(){}

    class animal{}
    class Alpaca{}

  `Contoh Benar`

    getUserId();
    getClientName();
    getCUstomerRecord();

    const DAY_IN_WEEK = 7;
    const DAY_IN_MONTH = 30;

    const songs = ['Back In Black'. 'Stairway to Heaven', 'Hey Jude'];
    const Artist = ['ACDC', 'Led Zeppelin', 'The Beatles'];

    function eraseDatabase(){}
    function restore_database(){}

    class Animal{}
    class Alpaca{}

`5. Hindari penambahan konteks yang tak perlu`

  `Contoh Salah`

    let fullNameString;

    const Car = {
      carmake: 'Honda';
      carModel: 'Accord';
      carColor: 'Blue';
    };

    function paintCar(car) {
      car.carColor = 'Red';
    }

  `Contoh Benar`

    let fullName;

    const Car = {
      make: 'Honda';
      model: 'Accord';
      color: 'Blue';
    };

    function paintCar(car) {
      car.color = 'Red';
    }

`6. Komentar`

  `Contoh Salah`

  `#1`

    function changeVocals (sentence) {
      //meletakan variable
      let result = '';
      //menyimpan seluruh huruf vokal
      const vowels = ['a','i','u','e','o','A','I','U','E','O'];
      //looping seluruh sentence
      for (let char of sentence) {
        if (vowels.indexOf(char) !== -1) {
          //ganti nilai char vokal menjadi char baru (bergeser satu huruf)
          let newcar = char.carCodeAt(0)+1;
          result += String.fromCharCode(newChar);
        } else {
          //tambah result dengan char
          result += char;
        }
      }
      //return result
      return result
    }

  `#2`

    doStuff();
    // doOtherStuff();
    // doSomeOtherStuff();
    // doSoMuchStuff();

  `Contoh Benar`

  `#1`

    function changeVocals (sentence) {
      let result = '';
      const vowels =  ['a','i','u','e','o','A','I','U','E','O'];

      for (let char of sentence) {
        //  jika ditemukan karakter vokal, maka direplace dengan karakter baru (charCode + 1)
        if vowels.indexOf(char) != -1 {
          let newChar = char.charCodeAt(0) + 1;
          result += String.fromCharCode(newChar);
        } else {
          result += char;
        }
      }
      return result
    }

  `#2`

    // Good
    doStuff();

`7. Good Function`

  `Contoh terlalu banyak argumen`

    //terlalu banyak argumen
    function createMenu(title, body, buttonText, cancellable) {}

  ``Contoh Benar`

    function createMenu({title, body, buttonText, cancellable}) {}

    createMenu({
      title: 'Foo', 
      body: 'Bar',
      buttonText: 'Baz',
      cancellable: true
    });

  `Contoh Salah`

    // hindari efek samping
    // variabel global direferensi oleh beberapa fungsi
    // jika kita mempunyai fungsi lain yang menghubungkan nama variabel init, sekarang akan menjadi array dan dapat merusak ini semua.
    let name = 'Ryan McDermott';
    function splitIntoFirstAndLastName() {
      name = name.split(' ');
    }
    splitIntoFirstAndLastName();
    console.log(name); // ['Ryan', 'McDermott'];

  `Contoh Benar`

    function splitIntoFirstAndLastName(){
      return name.split();
    }
    const name = 'Ryan McDermott';
    const newName = splitIntoFirstAndLastName(name);
    console.log(name); // 'Ryan McDermott';
    console.log(newName); // ['Ryan', 'McDermott'];

  `Contoh komentar yang bagus pada fungsi`

    /**
    * Function sendEmail()
    * 
    * Getting data from api dialect
    * @param {object} propertyNotif -  Property object from notification
    * @param {object} propertyQuery - Property object data from Query
    * @return Object
    * /

    function sendEmail(propertyNotif, propertyQuery) {
      // some of code here
    }

`8. Gunakan Konvensi`
  `Contoh "style guide"`
  - airbnb: https://github.com/airbnb/javascript
  - google: https://google.github.io/styleguide/pyguide.html

`9. Formatting`

  `formatting tips`
  - lebar baris code 80-120 karakter
  - satu class 300 - 500 baris.
  - baris code yang berhubungan saling berdekatan
  - dekatkan fungsi dengan pemanggilnya
  - deklarasi variabel berdekatan dengan penggunanya
  - perhatikan indentasi
  - menggunakan <i>prettier</i> atau <i>formatter</i>

# "Clean Code" Principle
  # KISS (keep it so simple)
  Hindari membuat fungsi yang dibuat untuk `melakukan A`, sekaligus `memodifikasi B`, `mengecek fungsi C`, dst.
  - Fungsi atau class harus kecil.
  - fungsi dibuat untuk melakuka satu tugas saja
  - jangan gunakan terlalu banyak argumen pada fungsi
  - harus diperhatikan untuk mencapai kondisi yang seimbang, kecil dan jumlahnya minimal.
  # DRY(Don't Repeat Yourself)
  Duplikasi code terjadi karena sering copy-paste. Untuk menghindari duplikasi code buatlah fungsi yang dapat digunakan secara berulang-ulang.
  # Refactoring
  refactoring adalah proses restrukturisasi kode yang dibuat, dengan cara mengubah struktur internal tanpa mengubah perilaku eksternal. Prinsip KISS dan DRY bisa dicapai dengan cara refactor.
  
  `teknik refactoring`

  - Membuat sebuah abstraksi
  - memecah kode dengan fungsi /class
  - perbaiki penamaan dan lokasi kode
  - deteksi kode yang memiliki duplikasi

