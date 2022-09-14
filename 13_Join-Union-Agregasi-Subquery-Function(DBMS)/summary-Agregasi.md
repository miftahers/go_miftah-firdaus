# AGREGASI

Daftar agregasi:
- MIN()
- MAX()
- SUM()
- AVG()
- COUNT()
- HAVING
  
  #  MIN()

  Function MIN() berguna untuk mengambil nilai terkecil dari daftar nilai yang ada pada tabel. Perlu dicatat bahwa fungsi `DISTINCT tidak bekerja pada function MIN()`. 

  `CONTOH PENGGUNAAN MIN()`
  Untuk mendapatkan informasi karyawan yang memiliki gaji terendah, Anda menggunakan subquery berikut:

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
  
  OUTPUT:
  ![](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-MIN-with-Subquery-example.png)


  `SQL MIN() dengan contoh GROUP BY`

    SELECT
      department_id,
      MIN(salary)
    FROM
      employees
    GROUP BY
      department_id;

  OUTPUT:
  ![](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-MIN-with-GROUP-BY-example.png)


  # MAX()

  Function MAX() berguna untuk mengambil nilai terkecil dari daftar nilai yang ada pada tabel. Perlu dicatat bahwa fungsi `DISTINCT tidak bekerja pada function MAX()`. 

  `CONTOH PENGGUNAAN MAX()`

  Untuk mendapatkan karyawan yang memiliki gaji tertinggi, Anda menggunakan subquery sebagai berikut:

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

  OUTPUT:
  ![](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-MAX-highest-salary.png)


  `SQL MAX dengan contoh GROUP BY`

  Sebagai contoh, kita dapat menggunakan fungsi MAX untuk mencari gaji tertinggi karyawan di setiap departemen sebagai berikut:

    SELECT
      department_id,
      MAX(salary)
    FROM
      employees
    GROUP BY
      department_id;

  OUTPUT:
  ![](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-MAX-highest-salary-per-department.png)\


  # SUM()

  Fungsi SQL SUM adalah fungsi agregat yang mengembalikan jumlah semua atau nilai yang berbeda. Kita dapat menerapkan fungsi SUM ke kolom numerik saja.

  Berikut ini mengilustrasikan sintaks fungsi SUM.

    SUM([ALL|DISTINCT] expression)

  `CONTOH PENGGUNAAN SUM()`

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

  
  # AVG() / AVERAGE FUNCTION

  Fungsi SQL AVG adalah fungsi agregat yang menghitung nilai rata-rata satu set. Berikut ini mengilustrasikan sintaks fungsi SQL AVG:

    AVG([ALL|DISTINCT] expression)

  `Contoh penggunaan AVG()`
  Untuk menghitung gaji rata-rata semua karyawan, Anda menerapkan fungsi AVG pada kolom gaji sebagai berikut:

    SELECT 
      AVG(salary)
  FROM
      employees;

  OUTPUT:
  ![](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-AVG-example.png)

  `SQL AVG dengan contoh klausa GROUP BY`
  Untuk menghitung nilai rata-rata grup, kami menggunakan fungsi AVG dengan klausa GROUP BY . Misalnya, pernyataan berikut mengembalikan departemen dan gaji rata-rata karyawan setiap departemen.

    SELECT
      department_id,
      AVG(salary)
    FROM
      employees
    GROUP BY
      department_id;

  OUTPUT:
  ![](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-AVG-GROUP-BY-example.png)

  
  # COUNT()

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

    SELECT
    department_id,
    COUNT(*)
  FROM
    employees
  GROUP BY
    department_id;

  OUTPUT:
  ![](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-COUNT-with-GROUP-BY-example.png)


  # HAVING

  APA PERBEDAAN HAVING DAN WHERE?

  HAVING BOLEH MENGGUNAKAN OPERATOR LAIN SELAIN '='

  CONTOH WHERE:

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

  CONTOH HAVING:

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
  