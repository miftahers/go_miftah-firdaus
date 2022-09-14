# JOIN

  # INNER JOIN

  Misalkan, Anda memiliki dua tabel: A dan B.

  Tabel A memiliki empat baris: (1,2,3,4) dan tabel B memiliki empat baris: (3,4,5,6)

  Ketika tabel A bergabung dengan tabel B menggunakan inner join, Anda memiliki himpunan hasil (3,4) yang merupakan perpotongan tabel A dan tabel B.

  Lihat gambar berikut.
  !["Inner Join"](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-INNER-JOIN.png)

  Untuk setiap baris dalam tabel A, klausa gabungan bagian dalam menemukan baris yang cocok dalam tabel B. Jika sebuah baris cocok, itu termasuk dalam kumpulan hasil akhir.

  Misalkan kolom dalam tabel A dan B adalah adan b. Pernyataan berikut mengilustrasikan klausa inner join:

    SELECT a
    FROM A
    INNER JOIN B ON b = a;

  Klausa INNER JOIN muncul setelah klausa FROM. Kondisi untuk mencocokkan antara tabel A dan tabel B ditentukan setelah kata kunci ON. Kondisi ini disebut kondisi join yaitu, B.n = A.n

  Klausa INNER JOIN dapat menggabungkan tiga tabel atau lebih selama mereka memiliki hubungan, biasanya hubungan kunci asing.

  Misalnya, pernyataan berikut menggambarkan cara menggabungkan 3 tabel: A, B, dan C:

    SELECT
    A.n
    FROM A
    INNER JOIN B ON B.n = A.n
    INNER JOIN C ON C.n = A.n;


  `Contoh SQL INNER JOIN`

  1) Menggunakan SQL INNER JOIN untuk menggabungkan dua tabel
  Kami akan menggunakan employees dan departments tabel dari database sampel untuk menunjukkan cara kerja INNER JOIN.

  !["tabel employees dan departments"](https://www.sqltutorial.org/wp-content/uploads/2016/03/emp_dept_tables.png)

  Setiap karyawan milik satu dan hanya satu departemen sementara setiap departemen dapat memiliki lebih dari satu karyawan. Hubungan antara departmentsdan employeesadalah satu-ke-banyak.

  Kolom department_iddalam employeestabel adalah kolom kunci asing yang menghubungkan employeeske departmentstabel.

  Untuk mendapatkan informasi departemen id 1,2, dan 3, Anda menggunakan pernyataan berikut.

    SELECT
      department_id,
      department_name
    FROM
      departments
    WHERE
      department_id IN (1, 2, 3);

  !["output informasi departemen id 1,2, dan 3"](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-INNER-JOIN-departments-table.png)

  Perhatikan bahwa kita menggunakan operator IN dalam klausa WHERE untuk mendapatkan baris dengan department_id 1, 2 dan 3.

  Untuk mendapatkan informasi karyawan yang bekerja di departemen id 1, 2 dan 3, Anda menggunakan query berikut:

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

  !["output informasi karyawan yang bekerja di departemen id 1, 2 dan 3"](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-INNER-JOIN-employees-table.png)

  Untuk menggabungkan data dari dua tabel ini, Anda menggunakan klausa INNER JOIN sebagai kueri berikut:

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

  !["output gabungan data dari dua tabel di atas"](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-INNER-JOIN-example.png)

  Untuk setiap baris dalam employeestabel, pernyataan memeriksa apakah nilai department_idkolom sama dengan nilai department_idkolom dalam departmentstabel. Jika kondisi

  Jika kondisi employees.department_id = departments.department_idterpenuhi, baris gabungan yang menyertakan data dari baris di keduanya employeesdan departmentstabel disertakan dalam kumpulan hasil.

  Perhatikan bahwa keduanya employeesdan departmentstabel memiliki nama kolom yang sama department_id, oleh karena itu kita harus mengkualifikasi department_idkolom tersebut menggunakan sintaks table_name.column_name.


  `Contoh tabel SQL INNER JOIN 3`

  Setiap karyawan memegang satu pekerjaan sementara pekerjaan dapat dipegang oleh banyak karyawan. Hubungan antara jobstabel dan employeestabel adalah satu-ke-banyak.

  Diagram database berikut menggambarkan hubungan antara employees, departments dan jobs:

  !["diagram hubungan antara employees, departments dan jobs"](https://www.sqltutorial.org/wp-content/uploads/2016/03/emp_dept_jobs_tables.png)

  Kueri berikut menggunakan klausa INNER JOIN untuk menggabungkan 3 tabel: karyawan, departemen, dan pekerjaan untuk mendapatkan nama depan, nama belakang, jabatan, dan nama departemen karyawan yang bekerja di departemen id 1, 2, dan 3.

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

  !["output INNER JOIN untuk menggabungkan 3 tabel"](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-INNER-JOIN-3-tables-example.png)

  # LEFT JOIN

  !["LEFT JOIN"](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-LEFT-JOIN.png)

  Dalam SQL, kami menggunakan sintaks berikut untuk menggabungkan tabel A dengan tabel B.

    SELECT
      A.n
    FROM
      A
    LEFT JOIN 
      B 
    ON 
      B.n = A.n;

  Klausa LEFT JOIN muncul setelah klausa FROM. Kondisi yang mengikuti kata kunci ON disebut kondisi joinB.n = A.n


  `Contoh SQL LEFT JOIN`

  !["tabel countries dan locations"](https://www.sqltutorial.org/wp-content/uploads/2016/03/countries_locations_tables.png)

  Setiap lokasi milik satu dan hanya satu negara sementara setiap negara dapat memiliki nol atau lebih lokasi. Hubungan antara tabel negara dan lokasi adalah satu-ke-banyak.

  Kolom country_id di tabel lokasi adalah kunci asing yang menautkan ke kolom country_id di tabel negara.

  Untuk menanyakan nama negara AS, Inggris, dan Cina, Anda menggunakan pernyataan berikut.

    SELECT
      country_id,
      country_name
    FROM
      countries
    WHERE
      country_id IN ('US', 'UK', 'CN');

  !["output LEFT JOIN"](https://www.sqltutorial.org/wp-content/uploads/2016/03/SQL-LEFT-JOIN-countries-data.png)

  # RIGHT JOIN

  `RIGHT JOIN Syntax`

    SELECT 
      column_name(s)
    FROM 
      table1
    RIGHT JOIN 
      table2
    ON 
      table1.column_name = table2.column_name;

  !["visualization RIGHT JOIN / RIGHT OUTER JOIN"](https://www.w3schools.com/Sql/img_rightjoin.gif)

  