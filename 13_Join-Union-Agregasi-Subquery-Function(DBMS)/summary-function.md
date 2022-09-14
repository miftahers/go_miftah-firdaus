# FUNCTION

Function adalah kumpulan statement/pernyataan yang akan mengembalikan/ me-return nilai balik pada pemanggilnya.

Contoh function untuk mengembalikan jumlah data dari tweets per user

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

`DETERMINISTIC / NOT DETERMINISTIC` adalah untuk menentukan yang bisa menggunakan function ini adalah user pembuatnya saja (deterministic) atau user siapa saja (not deterministic)

Cara menggunakan function

    SELECT sf_count_tweet_peruser(2);

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

    CREATE TRIGGER trigger_name [BEFORE|AFTER] event
    ON table_name trigger_type
    BEGIN
      -- trigger_logic
    END;

Mari kita periksa sintaks lebih detail:

- Pertama, tentukan nama TRIGGER setelah CREATE TRIGGERklausa.
- Selanjutnya, gunakan salah satu BEFOREatau AFTERkata kunci untuk menentukan kapan TRIGGER harus muncul sebagai respons terhadap peristiwa tertentu, misalnya, INSERT, UPDATE, atau DELETE.
- Kemudian, tentukan nama tabel tempat TRIGGER diikat.
- Setelah itu, tentukan jenis TRIGGER menggunakan salah satu FOR EACH ROW atau FOR EACH STATEMENT. Kami akan membahas lebih lanjut tentang ini di bagian berikutnya.
- Akhirnya, letakkan logika TRIGGER di BEGIN ... END blok.

Selain menggunakan kode di BEGIN ENDblok, Anda dapat menjalankan prosedur tersimpan sebagai berikut:

    CREATE TRIGGER trigger_name 
    [BEFORE|AFTER] event
    ON table_name trigger_type
    EXECUTE stored_procedure_name;

  # making trigger
  `The following before_update_salary trigger logs the salary changes to the salary_changes table.`

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


  # Modify triggers

    CREATE OR REPLACE TRIGGER trigger_name 
    [BEFORE|AFTER] event
    ON table_name trigger_type
    BEGIN
      -- trigger_logic
    END;

  # Delete triggers
  To delete a trigger, you use the DROP TRIGGER statement as follows:

    DROP TRIGGER [IF EXISTS] trigger_name;

  Code language: SQL (Structured Query Language) (sql)
  `IF EXISTS` memungkinkan Anda untuk menghapus TRIGGER jika TRIGGERnya ada. Jika TRIGGER tidak ada, maka pernyataan tidak melakukan apa-apa. Namun, jika Anda tidak memiliki IF EXISTSopsi, sistem database mungkin mengeluarkan kesalahan jika Anda mencoba untuk menghapus TRIGGER yang tidak ada.

Sekali lagi, jika Anda menjatuhkan tabel , semua TRIGGER yang terkait dengan tabel juga akan dihapus. Pernyataan berikut menghapus before_update_salary TRIGGER:

    DROP TRIGGER IF EXISTS before_update_salary;




