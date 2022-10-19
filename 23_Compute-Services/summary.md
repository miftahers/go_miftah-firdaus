# Compute services

System dan software deployment: Deployment adalah kegiatan yang bertujuan untuk  menyebarkan aplikasi/produk yang telah dikerjakan oleh para pengembang seringkali untuk mengubah dari status sementara ke permanen. penyebarannya dapat melalui beragam cara tergantung dari jenis aplikasinya, aplikasi web & api ke server sedangkan aplikasi mobile ke playstore/appstore.

Strategi deployment:
- Big-bang deployment / replace deployment strategy
  - Kelebihan:
    - Mudah diimplementasikan. cara klasik, tinggal replace.
    - Perubahan kepada sistem langsung 100% secara instan
  - Kekurangan:
    - Terlalu beresiko, rata-rata downtime cukup lama
- Rollout deployment strategy
  - Kelebihan:
    - Lebih aman dan less downtime dari versi sebelumnya
  - Kekurangan:
    - Akan ada 2 versi aplikasi yang berjalan secara barengan sampai semua server terdeploy, dan bisa membuat bingung.
    - Karena sifatnya perlahan satu-persatu, untuk deployment dan rollback lebihlama dari yang bigbang. Karena prosesnya perlahan-lahan sampai semua server terkena efeknya.
    - Tidak ada control request. Server yang baru ter-deploy dengan aplikasi versi baru, langsung mendapat request yang sama banyaknya dengan server yang lain.
- Blue/Green deployment strategy
  - Kelebihan:
    - Perubahan sangat cepat, sekali switch service langsung berubah 100%
    - Tidak ada issue beda versi pada service seperti yang terjadi pada rollout development
  - Kekurangan:
    - Resource yang dibutuhkan lebih banyak. Karena untuk setiap deployment kita harus menyediakan service yang serupa environtments dengan yang sedang berjalan di production.
    - Testing harus benar-benar sangat diprioritaskan sebelum di switch, aplikasi harus kita pastikan aman dari request yang tiba-tiba banyak.
- Canary Deployment strategy
  - Kelebihan:
    - Cukup aman
    - Mudah untuk rollback jika terjadi bug/error, tanpa berimbas ke semua user
  - Kekurangan:
    - Untuk mencapai 100% cukup lama dibanding dengan blue/green deploymen. Dengan blue/green deployment, aplikasi langsung 100% terdeploy keseluruh user.

Simple Dev-Ops Cycle

Dev(Plan -> Code -> Build -> Test) -:> Ops(Release -> Deploy -> Operate -> Monitor)


