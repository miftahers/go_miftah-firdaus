# Intro RESTful API

## API (Application Programming Interface)

API adalah seperangkat fungsi dan prosedur yang memungkinkan pembuatan aplikasi yang mengakses fitur atau data sistem operasi, aplikasi, atau layanan lainnya.

Hubungan antara Client dan Server ada pada 2 proses yaitu `Request` dan `Response`.

## Frontend & Backend Integration

Frontend mengirim request/upload -> backend menerima request dan mengelola request -> backend mengirim response ke frontend.

## Backend & Backend Integration

contoh: Service to Service

## What is REST?

REST  : REpresentational State Transfer

REST ITU
1. Menggunakan HTTP Protocol
2. Memiliki Endpoint, contoh: https://www.instagram.com/api/users
3. HTTP REQUEST METHOD
   1. GET: Minta data
   2. POST: Upload data
   3. PUT: Edit data
   4. DELETE: Hapus data
   5. HEAD
   6. OPTION
   7. PATCH
4. Request & Response format
   1. JSON
   2. XML
   3. SOAP

## JSON (Javascript Object Notation)
``` JSON
Contoh JSON:

  {
    "key": "value"        <- string
    "umur": 17,           <- int
    "single": true,       <- bool
    "hobi": [             <- slice
      "belajar",
      "renang"
    ]
    "alamat": {           <- JSON in JSON
      "rumah": "malang",
      "no": 4,
      "rt": "03",
      "rw": "11"
    }
  }
```

## HTTP Response Code

HTTP Response code berguna untuk mengidentifikasi apakah request itu berhasil atau tidak.

1. Response code 200: OK
2. Response code 201: Created (Berhasil PUT/POST)
3. Response code 400: Bad Request (Terjadi kegagalan saat POST/PUT. Contoh nya jika kita tidak lolos pada proses validasi)
4. Response code 404: Not Found (Resource yang diminta tidak ada)
5. Response code 401: Unauthorized (Mengakses tanpa autentikasi/tidak memiliki izin. contoh nya saat kita akses tapi belum login)
6. Response code 405: Method Not Allowed (Method yang digunakan tidak sesuai dengan endpoint)
7. Response code 500: Internal Server Error

## API TOOLS

Contoh tools testing API:
- Katalon
- Apache JMeter
- SoapUI
- Postman
- dll.

### POSTMAN

Postman adalah http client untuk testing web services.

## REST API Deisgn Best Practice

### User Nouns Instead of Verbs
Karena perintah dari REST API sudah kata kerja jadi gausah pake katakerja pada saat membuat endpoint

### Naming using Plural Nouns
Disarankan untuk menggunakan kata jamak dibandingkan kata tunggal
```
// BENAR:
GET /cars/123
POST /cars
GET /cars

// SALAH
GET /car/123
POST /car
GET /car
```

### Use Resource Nesting to show relations or hierarchy

```
// Contoh

/users <- menunjukan daftar user
/users/123 <- menunjuk user yang spesifik
/users/123/orders <- menunjuk daftar order yang dilakukan user tertentu
/users/123/orders/0001 <- menunjuk order yang spesifik yang dilakukan oleh user tertentu

```

### Contoh Format Response JSON

#### Contoh pertama
``` JSON
{
  "code"      : 200,
  "message"   : "Success get data",
  "status"    : "success",
  "data"      : {
    ...
  }
}
```

#### Contoh kedua
``` JSON
{
  "meta"            : {
    "error_type"    : "OAuthException",
    "code"          : 400,
    "error_message" : "..."
  },
  "data"            : {
    ...
  },
  "pagination"      : {
    "next_url"      : "...",
    "next_max_id"   : "13872296"
  }
}
```

### Filtering, Sorting, Paging, and Field Selection

#### Filtering
```
GET /users?country=USA
GET /users?creaton_date=2019-11-11

```
#### Sorting
```
GET /users?sort=birthdate_date:asc
GET /users?sort=birthdate_date:desc

```
#### Paging
```
GET /users?limit=100
GET /users?offset=2

```

#### Field Selection
```
GET /users/123?fields=username,email <- for one spesific user
GET /users?fields=username,email <- for a full list of users

```

### Handle Trailing Slashes Gracefully
```
(1) GET /cars
(2) GET /cars/
```
API Kita harus bisa menghandle endpoint keduanya

### Versioning

```
// Contoh API:
https://us6.api.mailchimp.com/3.0/ (major + minor version indication)
https://api.stripe.com/v1 (major version indication only)
https://developer.github.com/v3/ (major version indication only)

```

