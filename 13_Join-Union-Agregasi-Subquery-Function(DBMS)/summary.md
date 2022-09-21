# JOIN

## INNER JOIN

Misalkan, Anda memiliki dua tabel: A dan B.

Tabel A memiliki empat baris: (1,2,3,4) dan tabel B memiliki empat baris: (3,4,5,6)

Ketika tabel A bergabung dengan tabel B menggunakan inner join, Anda memiliki himpunan hasil (3,4) yang merupakan perpotongan tabel A dan tabel B.

Lihat gambar berikut.
!["Inner Join"](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-INNER-JOIN.png)

Untuk setiap baris dalam tabel A, klausa gabungan bagian dalam menemukan baris yang cocok dalam tabel B. Jika sebuah baris cocok, itu termasuk dalam kumpulan hasil akhir.

Misalkan kolom dalam tabel A dan B adalah adan b. Pernyataan berikut mengilustrasikan klausa inner join:
```SQL
  SELECT a
  FROM A
  INNER JOIN B ON b = a;
```
Klausa INNER JOIN muncul setelah klausa FROM. Kondisi untuk mencocokkan antara tabel A dan tabel B ditentukan setelah kata kunci ON. Kondisi ini disebut kondisi join yaitu, B.n = A.n

Klausa INNER JOIN dapat menggabungkan tiga tabel atau lebih selama mereka memiliki hubungan, biasanya hubungan kunci asing.

Misalnya, pernyataan berikut menggambarkan cara menggabungkan 3 tabel: A, B, dan C:
```SQL
  SELECT
  A.n
  FROM A
  INNER JOIN B ON B.n = A.n
  INNER JOIN C ON C.n = A.n;
```

`Contoh SQL INNER JOIN`

1) Menggunakan SQL INNER JOIN untuk menggabungkan dua tabel
Kami akan menggunakan employees dan departments tabel dari database sampel untuk menunjukkan cara kerja INNER JOIN.

!["tabel employees dan departments"](https://www.sqltutorial.org/wp-content/uploads/2016/03/emp_dept_tables.png)

Setiap karyawan milik satu dan hanya satu departemen sementara setiap departemen dapat memiliki lebih dari satu karyawan. Hubungan antara departmentsdan employeesadalah satu-ke-banyak.

Kolom department_iddalam employeestabel adalah kolom kunci asing yang menghubungkan employeeske departmentstabel.

Untuk mendapatkan informasi departemen id 1,2, dan 3, Anda menggunakan pernyataan berikut.
```SQL
  SELECT
    department_id,
    department_name
  FROM
    departments
  WHERE
    department_id IN (1, 2, 3);
```
!["output informasi departemen id 1,2, dan 3"](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-INNER-JOIN-departments-table.png)

Perhatikan bahwa kita menggunakan operator IN dalam klausa WHERE untuk mendapatkan baris dengan department_id 1, 2 dan 3.

Untuk mendapatkan informasi karyawan yang bekerja di departemen id 1, 2 dan 3, Anda menggunakan query berikut:
```SQL
  SELECT
    first_name,
    last_name,
    department_id
  FROM
    employees
  WHERE
    department_id IN (1, 2, 3)
  ORDER BY
    department_id;
```
!["output informasi karyawan yang bekerja di departemen id 1, 2 dan 3"](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-INNER-JOIN-employees-table.png)

Untuk menggabungkan data dari dua tabel ini, Anda menggunakan klausa INNER JOIN sebagai kueri berikut:
```SQL
  SELECT 
      first_name,
      last_name,
      employees.department_id,
      departments.department_id,
      department_name
  FROM
      employees
          INNER JOIN
      departments ON departments.department_id = employees.department_id
  WHERE
      employees.department_id IN (1 , 2, 3);
```
!["output gabungan data dari dua tabel di atas"](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-INNER-JOIN-example.png)

Untuk setiap baris dalam employeestabel, pernyataan memeriksa apakah nilai department_idkolom sama dengan nilai department_idkolom dalam departmentstabel. Jika kondisi

Jika kondisi employees.department_id = departments.department_idterpenuhi, baris gabungan yang menyertakan data dari baris di keduanya employeesdan departmentstabel disertakan dalam kumpulan hasil.

Perhatikan bahwa keduanya employeesdan departmentstabel memiliki nama kolom yang sama department_id, oleh karena itu kita harus mengkualifikasi department_idkolom tersebut menggunakan sintaks table_name.column_name.


`Contoh tabel SQL INNER JOIN 3`

Setiap karyawan memegang satu pekerjaan sementara pekerjaan dapat dipegang oleh banyak karyawan. Hubungan antara jobstabel dan employeestabel adalah satu-ke-banyak.

Diagram database berikut menggambarkan hubungan antara employees, departments dan jobs:

!["diagram hubungan antara employees, departments dan jobs"](https://www.sqltutorial.org/wp-content/uploads/2016/03/emp_dept_jobs_tables.png)

Kueri berikut menggunakan klausa INNER JOIN untuk menggabungkan 3 tabel: karyawan, departemen, dan pekerjaan untuk mendapatkan nama depan, nama belakang, jabatan, dan nama departemen karyawan yang bekerja di departemen id 1, 2, dan 3.
```SQL
  SELECT
    first_name,
    last_name,
    job_title,
    department_name
  FROM
    employees e
  INNER JOIN departments d ON d.department_id = e.department_id
  INNER JOIN jobs j ON j.job_id = e.job_id
  WHERE
    e.department_id IN (1, 2, 3);
```
!["output INNER JOIN untuk menggabungkan 3 tabel"](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-INNER-JOIN-3-tables-example.png)

## LEFT JOIN

!["LEFT JOIN"](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-LEFT-JOIN.png)

Dalam SQL, kami menggunakan sintaks berikut untuk menggabungkan tabel A dengan tabel B.
```SQL
  SELECT
    A.n
  FROM
    A
  LEFT JOIN 
    B 
  ON 
    B.n = A.n;
```
Klausa LEFT JOIN muncul setelah klausa FROM. Kondisi yang mengikuti kata kunci ON disebut kondisi joinB.n = A.n


`Contoh SQL LEFT JOIN`

!["tabel countries dan locations"](https://www.sqltutorial.org/wp-content/uploads/2016/03/countries_locations_tables.png)

Setiap lokasi milik satu dan hanya satu negara sementara setiap negara dapat memiliki nol atau lebih lokasi. Hubungan antara tabel negara dan lokasi adalah satu-ke-banyak.

Kolom country_id di tabel lokasi adalah kunci asing yang menautkan ke kolom country_id di tabel negara.

Untuk menanyakan nama negara AS, Inggris, dan Cina, Anda menggunakan pernyataan berikut.
```SQL
  SELECT
    country_id,
    country_name
  FROM
    countries
  WHERE
    country_id IN ('US', 'UK', 'CN');
```
!["output LEFT JOIN"](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-LEFT-JOIN-countries-data.png)

## RIGHT JOIN

`RIGHT JOIN Syntax`
```SQL
  SELECT 
    column_name(s)
  FROM 
    table1
  RIGHT JOIN 
    table2
  ON 
    table1.column_name = table2.column_name;
```
!["visualization RIGHT JOIN / RIGHT OUTER JOIN"](https://www.w3schools.com/Sql/img_rightjoin.gif)

# UNION

Operator UNION menggabungkan kumpulan hasil dari dua atau lebih pernyataan SELECT ke dalam kumpulan hasil tunggal.

`UNION Syntax`
```SQL
    SELECT 
        column1, column2
    FROM
        table1 
    UNION [ALL]
    SELECT 
        column3, column4
    FROM
        table2;
```
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
```SQL
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
```
!["Output UNION pada tabel employees dan dependents"](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-UNION-practical-example.png)

# AGREGASI

Daftar agregasi:
- MIN()
- MAX()
- SUM()
- AVG()
- COUNT()
- HAVING

##  MIN()

Function MIN() berguna untuk mengambil nilai terkecil dari daftar nilai yang ada pada tabel. Perlu dicatat bahwa fungsi `DISTINCT tidak bekerja pada function MIN()`. 

`CONTOH PENGGUNAAN MIN()`
Untuk mendapatkan informasi karyawan yang memiliki gaji terendah, Anda menggunakan subquery berikut:
```SQL
  SELECT
    employee_id,
    first_name,
    last_name,
    salary
  FROM
    employees
  WHERE
    salary = (
      SELECT
        MIN(salary)
      FROM
        employees
    );
```
OUTPUT:
![](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-MIN-with-Subquery-example.png)


`SQL MIN() dengan contoh GROUP BY`
```SQL
  SELECT
    department_id,
    MIN(salary)
  FROM
    employees
  GROUP BY
    department_id;
```
OUTPUT:
![](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-MIN-with-GROUP-BY-example.png)


## MAX()

Function MAX() berguna untuk mengambil nilai terkecil dari daftar nilai yang ada pada tabel. Perlu dicatat bahwa fungsi `DISTINCT tidak bekerja pada function MAX()`. 

`CONTOH PENGGUNAAN MAX()`

Untuk mendapatkan karyawan yang memiliki gaji tertinggi, Anda menggunakan subquery sebagai berikut:
```SQL
  SELECT
    employee_id,
    first_name,
    last_name,
    salary
  FROM
    employees
  WHERE
    salary = (
      SELECT
        MAX(salary)
      FROM
        employees
    );
```
OUTPUT:
![](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-MAX-highest-salary.png)


`SQL MAX dengan contoh GROUP BY`

Sebagai contoh, kita dapat menggunakan fungsi MAX untuk mencari gaji tertinggi karyawan di setiap departemen sebagai berikut:
```SQL
  SELECT
    department_id,
    MAX(salary)
  FROM
    employees
  GROUP BY
    department_id;
```
OUTPUT:
![](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-MAX-highest-salary-per-department.png)\


## SUM()

Fungsi SQL SUM adalah fungsi agregat yang mengembalikan jumlah semua atau nilai yang berbeda. Kita dapat menerapkan fungsi SUM ke kolom numerik saja.

Berikut ini mengilustrasikan sintaks fungsi SUM.

  SUM([ALL|DISTINCT] expression)

`CONTOH PENGGUNAAN SUM()`
```SQL
  SELECT
    e.department_id,
    department_name,
    SUM(salary)
  FROM
    employees e
  INNER JOIN departments d ON d.department_id = e.department_id
  GROUP BY
    e.department_id
  ORDER BY
    SUM(salary) DESC;
```

## AVG() / AVERAGE FUNCTION

Fungsi SQL AVG adalah fungsi agregat yang menghitung nilai rata-rata satu set. Berikut ini mengilustrasikan sintaks fungsi SQL AVG:

  AVG([ALL|DISTINCT] expression)

`Contoh penggunaan AVG()`
Untuk menghitung gaji rata-rata semua karyawan, Anda menerapkan fungsi AVG pada kolom gaji sebagai berikut:
```SQL
  SELECT 
    AVG(salary)
FROM
    employees;
```
OUTPUT:
![](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-AVG-example.png)

`SQL AVG dengan contoh klausa GROUP BY`
Untuk menghitung nilai rata-rata grup, kami menggunakan fungsi AVG dengan klausa GROUP BY . Misalnya, pernyataan berikut mengembalikan departemen dan gaji rata-rata karyawan setiap departemen.
```SQL
  SELECT
    department_id,
    AVG(salary)
  FROM
    employees
  GROUP BY
    department_id;
```
OUTPUT:
![](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-AVG-GROUP-BY-example.png)


## COUNT()

Fungsi SQL COUNTadalah fungsi agregat yang mengembalikan jumlah baris yang dikembalikan oleh kueri. Anda dapat menggunakan COUNT() dalam pernyataan SELECT untuk mendapatkan jumlah karyawan, jumlah karyawan di setiap departemen, jumlah karyawan yang memegang pekerjaan tertentu, dll.

Berikut ini mengilustrasikan sintaks dari COUNT() SQL:

  COUNT([ALL | DISTINCT] expression);

Hasil COUNT() tergantung pada argumen yang Anda berikan padanya.
- Kata `ALL` akan menyertakan nilai duplikat dalam hasil. Misalnya, jika Anda memiliki grup (1, 2, 3, 3, 4, 4) dan menerapkan COUNT(), hasilnya adalah 6. Secara default, COUNT() menggunakan `ALL`.
- Kata `DISTINCT` hanya menghitung nilai unik. Misalnya, COUNT() mengembalikan 4 jika Anda menerapkannya ke grup (1, 2, 3, 3, 4, 4).
- `expression` adalah kolom tabel tempat Anda ingin menghitung nilainya.

Bentuk lain dari COUNTfungsi yang menerima tanda bintang (*) sebagai argumen adalah sebagai berikut:

  COUNT(*)

Fungsi COUNT(*)mengembalikan jumlah baris dalam tabel dalam kueri. Ini menghitung baris duplikat dan baris yang berisi nilai nol.


`Contoh fungsi SQL COUNT`

Contoh berikut menggunakan COUNTfungsi dengan klausa GROUP BY untuk mencari jumlah karyawan untuk setiap departemen:
```SQL
  SELECT
  department_id,
  COUNT(*)
FROM
  employees
GROUP BY
  department_id;
```
OUTPUT:
![](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-COUNT-with-GROUP-BY-example.png)


## HAVING

APA PERBEDAAN HAVING DAN WHERE?

HAVING BOLEH MENGGUNAKAN OPERATOR LAIN SELAIN '='

CONTOH WHERE:
```SQL
  SELECT 
      manager_id,
      first_name,
      last_name,
      COUNT(employee_id) direct_reports
  FROM
      employees
  WHERE
      manager_id IS NOT NULL
  GROUP BY manager_id;
```
CONTOH HAVING:
```SQL
  SELECT 
      manager_id,
      first_name,
      last_name,
      COUNT(employee_id) direct_reports
  FROM
      employees
  WHERE
      manager_id IS NOT NULL
  GROUP BY manager_id
  HAVING direct_reports >= 5;
```
`SQL HAVING with SUM function example`
```
  SELECT 
      department_id, SUM(salary)
  FROM
      employees
  GROUP BY department_id
  HAVING SUM(salary) BETWEEN 20000 AND 30000
  ORDER BY SUM(salary);
```
`SQL HAVING with MIN function example`
```SQL
  SELECT
    e.department_id,
    department_name,
    MIN(salary)
  FROM
    employees e
  INNER JOIN departments d ON d.department_id = e.department_id
  GROUP BY
    e.department_id
  HAVING
    MIN(salary) >= 10000
  ORDER BY
    MIN(salary);
```
`SQL HAVING clause with AVG function example`

```SQL
  SELECT
    e.department_id,
    department_name,
    ROUND(AVG(salary), 2)
  FROM
    employees e
  INNER JOIN departments d ON d.department_id = e.department_id
  GROUP BY
    e.department_id
  HAVING
    AVG(salary) BETWEEN 5000
  AND 7000
  ORDER BY
    AVG(salary);

```

# SUBQUERY

Subquery dapat digunakan bersama `SELECT`, `INSERT`, `UPDATE`, dan `DELETE` bersama operator seperti `=`, `<`, `>`, `>=`, `<=`, `IN`, `BETWEEN`, dan lainnya.

PERATURAN SUBQUERY:
- Harus tertutup dalam tanda kurung
- Sebuah subquery hanya dapat memiliki satu kolom pada klausa `SELECT`, kecuali beberapa kolom yang di query utama untuk subquery untuk membandingkan kolom yang dipilih.
- Subqueries yang kembali lebih dari satu baris hanya dapat digunakan dengan beberapa value operator, seperti operator `IN`
- Daftar `SELECT` tidak bisa menyertakan referensi ke nilai-nilai yang mengevaluasi ke BLOB, ARRAY, CLOB, atau NCLOB
- Sebuah subquery tidak dapat segera tertutup dalam fungsi set

Contoh:

Tampilkan data tabel users yang jumlah total favourite_count per user lebih dari 5 pada tabel tweets
```SQL
  SELECT * FROM users
  WHERE 
    id
  IN
    (
      SELECT user_id FROM tweets
      GROUP BY user_id
      HAVING SUM(favourite_count) > 5
    );
```
Tampilkan data tabel users yang user_id nya ada pada tabel tweets
```SQL
  SELECT * FROM users
  WHERE
    id
  IN
    (
      SELECT user_id FROM tweets
      GROUP BY user_id
    );
```

# FUNCTION

Function adalah kumpulan statement/pernyataan yang akan mengembalikan/ me-return nilai balik pada pemanggilnya.

Contoh function untuk mengembalikan jumlah data dari tweets per user
```SQL
  DELIMITER //
  CREATE FUNCTION sf_count_tweet_peruser
  (user_id_p INT) RETURN INT DETERMINISTIC
  BEGIN
  DECLARE total INT;
  SELECT COUNT(*) INTO total FROM tweets
  WHERE user_id = user_id_p AND TYPE = 'tweets';
  RETURN total;
  END //
  DELIMITER ;
```
`DETERMINISTIC / NOT DETERMINISTIC` adalah untuk menentukan yang bisa menggunakan function ini adalah user pembuatnya saja (deterministic) atau user siapa saja (not deterministic)

Cara menggunakan function
```SQL
  SELECT sf_count_tweet_peruser(2);
```
*Akan menampilkan jumlah tweets user_id 2


# TRIGGER

TRIGGER adalah sepotong kode yang dieksekusi secara otomatis sebagai respons terhadap peristiwa tertentu yang terjadi pada tabel dalam database.

TRIGGER selalu dikaitkan dengan tabel tertentu. Jika tabel dihapus , semua TRIGGER terkait juga dihapus secara otomatis.

TRIGGER dipanggil sebelum atau setelah kejadian berikut:

- INSERT – ketika baris baru dimasukkan
- UPDATE – ketika baris yang ada diperbarui
- DELETE – ketika sebuah baris dihapus.

Saat Anda mengeluarkan  INSERT, UPDATE, atau DELETE pernyataan, sistem manajemen basis data relasional (RDBMS) TRIGGER TRIGGER yang sesuai.

Dalam beberapa RDMBS, TRIGGER juga dipanggil dalam hasil eksekusi pernyataan yang memanggil pernyataan INSERT, UPDATE, atau DELETE. Misalnya, MySQL memiliki LOAD DATA INFILE , yang membaca baris dari file teks dan menyisipkan ke dalam tabel dengan kecepatan sangat tinggi, memanggil BEFORE INSERT dan AFTER INSERT TRIGGER.

Di sisi lain, pernyataan dapat menghapus baris dalam tabel tetapi tidak memanggil TRIGGER terkait. Misalnya, pernyataan TRUNCATE TABLE menghapus semua baris dalam tabel tetapi tidak memanggil BEFORE DELETE dan AFTER DELETE TRIGGER.

`Trigger creation statement syntax`
```SQL
  CREATE TRIGGER trigger_name [BEFORE|AFTER] event
  ON table_name trigger_type
  BEGIN
    -- trigger_logic
  END;
```
Mari kita periksa sintaks lebih detail:

- Pertama, tentukan nama TRIGGER setelah CREATE TRIGGERklausa.
- Selanjutnya, gunakan salah satu BEFOREatau AFTERkata kunci untuk menentukan kapan TRIGGER harus muncul sebagai respons terhadap peristiwa tertentu, misalnya, INSERT, UPDATE, atau DELETE.
- Kemudian, tentukan nama tabel tempat TRIGGER diikat.
- Setelah itu, tentukan jenis TRIGGER menggunakan salah satu FOR EACH ROW atau FOR EACH STATEMENT. Kami akan membahas lebih lanjut tentang ini di bagian berikutnya.
- Akhirnya, letakkan logika TRIGGER di BEGIN ... END blok.

Selain menggunakan kode di BEGIN ENDblok, Anda dapat menjalankan prosedur tersimpan sebagai berikut:
``` SQL
  CREATE TRIGGER trigger_name 
  [BEFORE|AFTER] event
  ON table_name trigger_type
  EXECUTE stored_procedure_name;
```
## making trigger
`The following before_update_salary trigger logs the salary changes to the salary_changes table.`
``` SQL
  DELIMITER //

  CREATE TRIGGER before_update_salary
  BEFORE UPDATE ON employees
  FOR EACH ROW
  BEGIN
    IF NEW.salary <> OLD.salary THEN
    INSERT INTO salary_changes(employee_id,old_salary,new_salary)
          VALUES(NEW.employee_id,OLD.salary,NEW.salary);
      END IF;
  END; //

  DELIMITER ;
```

## Modify triggers

  CREATE OR REPLACE TRIGGER trigger_name 
  [BEFORE|AFTER] event
  ON table_name trigger_type
  BEGIN
    -- trigger_logic
  END;

# Delete triggers
To delete a trigger, you use the DROP TRIGGER statement as follows:
``` SQL
  DROP TRIGGER [IF EXISTS] trigger_name;
```
Code language: SQL (Structured Query Language) (sql)
`IF EXISTS` memungkinkan Anda untuk menghapus TRIGGER jika TRIGGERnya ada. Jika TRIGGER tidak ada, maka pernyataan tidak melakukan apa-apa. Namun, jika Anda tidak memiliki IF EXISTSopsi, sistem database mungkin mengeluarkan kesalahan jika Anda mencoba untuk menghapus TRIGGER yang tidak ada.

Sekali lagi, jika Anda menjatuhkan tabel , semua TRIGGER yang terkait dengan tabel juga akan dihapus. Pernyataan berikut menghapus before_update_salary TRIGGER:
``` SQL
  DROP TRIGGER IF EXISTS before_update_salary;

```


