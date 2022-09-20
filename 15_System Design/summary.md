# SYSTEM DESIGN
list:
  1. Diagram
  2. Distributed system
  3. Job queue and microservices
  4. 
# Diagram

Diagram yang sering digunakan untuk mendesain sebuah system adalah Flow Chart, Use Case, Entitiy Relationship Diagram, dan High Level Architecture.

1. Flow Chart
[Detail Flow Chart](lucidchart.com/pages/what-is-a-flowchart-tutorial 'Klik disini untuk detail tentang Flow Chart')

2. Use Case Diagram
[Detail Use Case Diagram](https://www.lucidchart.com/pages/uml-use-case-diagram 'Klik disini untuk detail tentang Use Case Diagram')

3. Entity Relationship Diagram
[Detail Entity Relationship Diagram](https://www.lucidchart.com/pages/er-diagrams 'Klik disini untuk detail tentang Entity Relationship Diagram')

4. High Level Architeture Diagram
[Detail High Level Architecture Diagram](https://www.lucidchart.com/pages/architecture-diagram 'Klik disini untuk detail tentang High Level Architecture Diagram')


# Distributed System

System Design Basics

ada yang perlu diperhatikan ketika mendesain system yang besar:
1. What are the different architectural pieces that can be used?
2. How do these pieces work with each other?
3. How can we best utilize these pieces: what are the right tradeoffs?

Familiarizing these concept would greatly benefit in understanding distributed system concepts.

Karakteristik dari distributed system:
1. Scalability

      `Scaling`

      Vertical Scale    : Increase Ability of system

      Horizontal Scale  : Duplikasi Server atau Tambah Server

      Jika kita masih memiliki kesempatan untuk meningkatkan kemampuan dari sistem yang kita bangun, gunakan vertical scale, tidak ada perubahan pada code kita.

      Jika kita tidak bisa memprediksi atau ada kemungkinan mengamalami peningkatan penggunaan layanan kita yang sangat cepat dan besar, gunakan horizontal scale atau tambahlah server, ada kemungkinan perlu merubah code.

2. Reliability
  Artinya system dari layanan kita harus bisa tetap berjalan ketika terjadi kegagalan seperti server down. Salah satu cara supaya bisa reliability adalah dengan memiliki multiple server.
3. Availability
  artinya system dari layanan kita harus selalu sedia walaupun terjadi kegagalan seperti server down. Sama seperti reliability yang berarti kita mesti memanfaatkan multiple server supaya siap ketika salah satu server down.
4. Efficiency
  Dua standar yang digunakan untuk mengukur efisiensi adalah latency dan bandwith. Semakin cepat latensi dan semakin besar bandwith maka efisiensi nya semakin bagus.
5. Serviceability & Manageability
  artinya sistem harus bisa diatur dan diperbaiki dengan mudah. Salah satu cara untuk membuat sistem yang serviceability dan manageability adalah dengan mengimplementasikan clean code.


# Job Queue and Microservices

# Load Balancer
  Load balancer adalah bagian dari sistem yang harus dibuat dimana berfungsi untuk mengatur dimana data akan dikelola selanjutnya. contoh ada 3 server untuk backend, nah data akan dimasukan dulu dari web server ke load balancer kemudian load balancer akan menentukan ke backend server yang mana yang siap untuk memproses data tersebut.

# Monolithic
  Monolithic berarti kita membuat sebuah system dalam sebuah satu kesatuan dimana server menghandle seluruh fitur/request dari client.
  keunggulan dari monolithic adalah kita hemat biaya server karena server yang digunakan hanyalah satu kesatuan.
  Kekurangan dari monolithic adalah kita akan kewalahan jika salah satu fitur rusak maka kita mesti menghentikan/maintenance server secara keseluruhan dan Seluruh layanan akan terganggu.

# Microservices
  microservices berarti kita membuat setiap fitur atau beberapa fitur menjadi bagian-bagian yang terpisah dan di handle oleh server yang berbeda sehingga ketika salah satu fitur rusak, fitur-fitur yang lain bisa tetap berjalan seperti biasa. antar fitur ini berkomunikasi menggunakan networks sebagai contoh berkomunikasi via RestAPI.
  Keunggulan dari microservices ini adalah ketika salah satu fitur rusak maka layanan yang lain akan tetap bisa berjalan, bahasa pemrograman bisa beragam dan disesuaikan untuk mengefisiensikan tiap fitur.
  Kekurangan dari microservices ini adalah kita akan menggunakan lebih banyak resources untuk handle bagian-bagian dari fitur yang berbeda sehingga cost yang diperlukan juga lebih tinggi dari rataan metode monolithic.



# SQL & NoSQL

SQL atau Relational Database adalah database yang memiliki relasi dan dibuat secara terstruktur.

NoSQL atau No relational Database adalah database yang tidak memiliki relasi.

Manfaat Relational DB
1. Dirancang untuk segala keperluan
2. SQL memiliki standar yang jelas
3. memiliki banyak tool (Administrasi, reporting, framework)

4 Prinsip Relational Database
1. Atomicity: Transaksi terjadi semua atau tidak sama sekali
2. Consistency: Data tertulis merupakan data valid yang ditentukan berdasarkan aturan tertentu
3. Isolation: pada saat terjadi request yang bersamaan(Concurrent), memastikan bahwa transaksi dieksekusi seperti dijalankan secara sekuensial.
4. Durability: jaminan bahwa transaksi yang telah tersimpan, tetap tersimpan.


Not Only SQL

No SQL is new way of thinking about a database.

DBMS yang menyediakan mekanisme yang lebih fleksibel dibandingkan model RDBMS (Sifat ACID).

No SQL digunakan untuk:
1. Menghindari pada sifat transaksi ACID
2. Menghindari Kompleksitas SQL
3. Menghindari Design Schema di depan
4. Menghindari Transactions (Ditangani oleh aplikasi)

Kelebihan No SQL:
- Schema less
- fast development
- etc.

Kapan No SQL digunakan?
ketika membutuhkan skema yang fleksibel, ACID tidak diperlukan, data logging (JSON), data sementara (cache), etc.

Kapan tidak diperlukan?
Data finansial (karena butuh tracking dana), data transaksi, bussiness critical, ACID - Compliant.


Apa itu Scema less?
- Traditional RDMBS
  - Tidak bisa menambah data yang tidak sesuai skema
  - perlu menambah data NULL pada item yang tidak memiliki data
  - memiliki tipe data yang strict
  - tidak dapat menambah beberapa item data pada sebuah field
- No SQL DBMS
  - tidak memiliki skema ketika menambahkan data
  - aplikasi menangani proses validasi tipe data
  - mendukung proses aggregasi dokumen pada item

Contoh NoSQL DMBS:
- redis -> key: value
- cassandra -> column - family
- neo4j -> graph
- mongoDB -> document based
- etc.

