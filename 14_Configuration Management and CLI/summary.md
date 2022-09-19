# UNIX Shell

  # Directory

  Directory command:
    - pwd
    - ls
    - mkdir
    - cd
    - rm
    - cp
    - mv
    - ln
    
      biner:
      
      000 = 0
      001 = 1
      010 = 2
      011 = 3
      100 = 4
      101 = 5
      110 = 6
      111 = 7

  Deskripsi:
  - Perintah `pwd` atau "print working directory" menampilkan direktori saat ini tempat Anda berada.
  - Perintah direktori `ls` or "list" sama dengan "dir" perintah di Windows. Perintah ini mencantumkan file dan direktori di direktori saat ini atau jalur alternatif apa pun yang ditentukan.
    - ls -al / ls -a
      Digunakan untuk menampilkan list file/folder beserta hak aksesnya. Cara mengubah hak akses tersebut dengan mengetikan `chmod 777 <nama_file/direktori>` maka kode hak aksesnya akan menjadi `-rwxrwxrwx`. Adapun angka setelah `chmod` itu akan dikonversi menjadi biner. Jadi jika ingin mengubah hak akses jadi `--wxrwx--x` itu dengan command `chmod 371 <nama_file>`.
  - Membuat direktori di Linux adalah melalui perintah `mkdir` atau make directory . Untuk membuat direktori baru bernama MyAwesomeLinuxDir di direktori home Anda (ditandai dengan jalur khusus ~) `mkdir ~/MyAwesomeLinuxDir`
  - Beroperasi di Linux sering berarti mengubah direktori, dilakukan melalui `cd` atau mengubah perintah direktori. Untuk kembali ke direktori sebelumnya -> `cd ..`
  -  gunakan `rm`untuk menghapus direktori dan isinya. Contoh-> `rm MyAwesomeLinuxDir`
  -  Salin file di Linux dengan `cp`. Perintah tidak hanya menyalin direktori dan file di Linux, tetapi juga atribut file dan membuat tautan simbolik. Untuk salin direktori + isinya gunakan `cp -r MyAwesomeLinuxDir/ <destinasi_folder>`.
  -  Untuk memindahkan direktori dan file di Linux, gunakan `mv`.
  -  `find` untuk mencari file / folder -> `sudo find / -type d -name "var"`

  # Files

  Files command:
    - create: touch
    - view: head, cat, tail, less
    - editor: vim, nano
    - permission: chown, chmod
    - different: diff

  `CREATE`

  Touch digunakan untuk membuat file baru yang kosong -> `touch file_baru.txt`. Jika ingin buat banyak file dalam sekali command gunakan `touch nama_file{01..05}`.

  `VIEW`

  - `Head`: The head command, as the name implies, print the top N number of data of the given input.
  - Cat (concatenate) command is very frequently used in Linux. It reads data from the file and gives their content as output. It helps us to create, view, concatenate files. 
    - To view One files `cat filename`, 
  - To view Multiple files `car file1 file2`
  - Create a file `cat > newfile`

  # Network

  Network commands:
  - ping
  - ssh
  - netstat
  - nmap
  - ipp addr (ifconfig)
  - wget
  - curl
  - telnet
  - netcat

  # Utility
  
  Utility command:
  - man
  - env
  - echo
  - date
  - which
  - watch
  - sudo
  - history
  - grep
  - locate
  
# SHELL SCRIPT

- SHELL IS BRIDGE BETWEEN THE USER AND KERNEL
- SHELL SCRIPT IS A PROGRAMMING LANGUAGE COMPILED BASED ON SHELL COMMANDS

  # create using echo

  `create new file`
  echo 'content' > file_name
  `echo 'hello world' > hello.sh`

  `add to file`
  echo 'content' >> file_name

  # run file
  
  use `./hello.sh`

  # write using nano

  use `nano file_name`
  `nano hello.sh`

  # write variables
  
  example:

  1) create a file using touch
    
    `touch new-file.sh`
    
  2) edit file using nano
    
    `nano new-file.sh`
  
  3) write this on nano

    #!bin/sh

    #text
    test = "Linux"
    echo "Ayo belajar " $test
    test = "Hal Baru"
    echo "Ayo belajar " $test

    #penjumlahan
    angka1 = 20
    angka2 = 15
    sum=$((angka1+angka2))
    echo maka $angka1 +$angka2 = $sum

    #dapatkan input
    echo "masukan namamu: "
    read nama
    echo Hai $nama Salam Kenal

  4) save file nano using `ctrl+O` kemudian keluar dari nano dengan `ctrl+x`

  5) pastikan hak akses file bisa di eksekusi. dan jalankan!


