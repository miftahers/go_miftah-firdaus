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

    SELECT * FROM users
    WHERE 
      id
    IN
      (
        SELECT user_id FROM tweets
        GROUP BY user_id
        HAVING SUM(favourite_count) > 5
      );

Tampilkan data tabel users yang user_id nya ada pada tabel tweets

    SELECT * FROM users
    WHERE
      id
    IN
      (
        SELECT user_id FROM tweets
        GROUP BY user_id
      );

