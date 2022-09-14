# UNION

Operator UNION menggabungkan kumpulan hasil dari dua atau lebih pernyataan SELECT ke dalam kumpulan hasil tunggal.

`UNION Syntax`

    SELECT 
        column1, column2
    FROM
        table1 
    UNION [ALL]
    SELECT 
        column3, column4
    FROM
        table2;

Untuk menggunakan operator UNION, Anda menulis pernyataan SELECT individu dan menggabungkannya dengan kata kunci UNION.

Kolom yang dikembalikan oleh pernyataan SELECT harus memiliki tipe data, ukuran, dan urutan yang sama atau dapat dikonversi.

Sistem database memproses kueri dengan mengeksekusi dua pernyataan SELECT terlebih dahulu. Kemudian, ini menggabungkan dua set hasil individu menjadi satu dan menghilangkan baris duplikat. Untuk menghilangkan baris duplikat, sistem database mengurutkan hasil gabungan yang ditetapkan oleh setiap kolom dan memindainya untuk menemukan baris yang cocok yang terletak bersebelahan.

Untuk mempertahankan baris duplikat di set hasil, Anda menggunakan operator UNION ALL.

Misalkan, kita memiliki dua himpunan hasil A(1,2) dan B(2,3). Gambar berikut mengilustrasikan A UNION B:

!["Ilustrasi A UNION B"](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-UNION.png)

Dan gambar berikut mengilustrasikan A UNION ALL B

!["Ilustrasi A UNION ALL B"](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-UNION-ALL.png)


`Sistem database melakukan langkah-langkah berikut:`

- Pertama, jalankan setiap pernyataan SELECT satu per satu.
- Kedua, gabungkan kumpulan hasil dan hapus baris duplikat untuk membuat kumpulan hasil gabungan.
- Ketiga, urutkan hasil gabungan yang ditetapkan oleh kolom yang ditentukan dalam klausa ORDER BY.
  
Dalam prakteknya, kita sering menggunakan operator UNION untuk menggabungkan data dari tabel yang berbeda. Lihat tabel karyawan dan tanggungan berikut:

!["tabel employees dan dependents"](https://www.sqltutorial.org/wp-content/uploads/2016/03/employees_dependents_tables.png)

Pernyataan berikut menggunakan operator UNION untuk menggabungkan nama depan dan nama belakang karyawan dan tanggungan.

    SELECT
      first_name,
      last_name
    FROM
      employees
    UNION
    SELECT
      first_name,
      last_name
    FROM
      dependents
    ORDER BY
      last_name;

!["Output UNION pada tabel employees dan dependents"](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-UNION-practical-example.png)

