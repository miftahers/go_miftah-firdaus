# Unit Testing

`How can u consider yourself to be a professional if you do not know that all of your code works?`

`how can u know all your code works if you don't test it every time u make a change?`

`how can you test it every time you make a change if you don't have automated unit tests with every high coverage?`

Outline:
- Testing definition
- Purpose
- Level
- Framework
- Structure
- Runner
- Mocking
- Coverage
- Step

## Software testing
Adalah proses untuk menganalisa sebuah sistem atau software untuk mendeteksi perbedaan antara kondisi fitur sistem saat ini dengan fitur yang diinginkan oleh stakeholder.
## Purpose of Testing
- Preventing `Regression`
  
  Regression adalah kondisi dimana software yang sudah berjalan dengan baik menjadi salah karena ada perubahan fitur.

- Confidence Level in `Refactoring`
  
  Refactoring adalah proses mengubah sistem atau code tanpa mengubah fungsionalitasnya.

- Improve `Code Design`
- Code `Documentation`
- Scaling the `Team`

## Level of Testing

1. UI Test
   UI test adalah test yang dilakukan untuk keseluruhan fitur bersama dengan User Interface, atau bisa dibilang test untuk mengetahui apakah aplikasi berjalan dengan baik dan sinkron dengan UI nya.
  
2. Integration Test
   Testing terhadap spesific modul tertentu. Contoh: testing terhadap API yang kita buat. Apakah API berjalan dengan baik? Apakah request dan response nya sesuai?

3. Unit Test
   Testing terhadap unit terkecil dari aplikasi yang kita buat. Unit testing diantaranya adalah test terhadap fungsi, method dan mengecek semua logic di code kita.

## Framework
Framework menyediakan alat, dan kebutuhan struktural untuk testing secara effisien.


Go framework to test
https://github.com/stretchr/testify

List framework other languages
https://en.wikipedia.org/wiki/List_of_unit_testing_frameworks#cite_note-216

## Structure
2 Ussual pattern:
1. `Centralize` your test file inside test folder
2. save test file `together` with production file

Test File (Collection of Test Suites) -> Test Suites (Collection of Test Cases) => (1) Test Fixtures(Setup & Teardown) & (2) Test Cases (Input, Process, Output)

## Runner
Runner adalah aplikasi yang di desain untuk menjalankan testing itu sendiri. 

Fitur/tool yang biasanya ada di runner:
- Tool yang menjalankan test file
- Terdapat beberapa tool yang memiliki fitur watch mode. Untuk melakukan test secara otomatis ketika kita melakukan perubahan.
- (Tips) Pilih runner yang paling cepat

## Mocking
Mocking adalah sesuatu yang dibuat sebagai imitasi/tiruan, dalam konteks testing contohnya adalah ...

Test case need to be `INDEPENDENT`

Need to create `*MOCK OBJECT`

* : mock object adalah object palsu/tiruan yang bertingkah/berprilaku layaknya objek asli

## DO NOT TEST
1. 3rd Party Modules
2. Database
3. 3rd Party API or other external system
4. Object class that have been tested/will be tested in other place

## Coverage
Testing coverage adalah alat ukur untuk menunjukan source code program yang sudah di execute ketika kita menjalankan sebuah test.

Coders need to `make sure` whether they have created enough tests.

Coverage tool menganalisis application code ketika test sedang dijalankan.

### Coverage report
- Function
- Statement
- Branch
- Lines

### Report format
- CLI
- XML
- HTML
- LCOV

Use format tool like sonarqube (static code analyzer) able to read it


## Simple Steps to Test

### Create a New TEST FILE in Go
1. Name ending library_test.go (e.g. user_test.go)
2. Location file:
   1. same folder, same package
   2. same folder different package

Create package file
```go
# lib/calculate/calculate.go
package calculate

func Addition(number_a, number_b int) int {
	result := number_a + number_b
	return result
}

func Substract(number_a, number_b int) int {
	result := number_a - number_b
	return result
}

func Mult(number_a, number_b int) int {
	result := number_a * number_b
	return result
}

func Div(number_a, number_b int) int {
	result := number_a / number_b
	return result
}

```

### Write a Test Function
1. Name: Test & Capitalised Word
2. Should have the signature func TestXxxx(t *testing.T)

```go
package calculate

import (
	"testing"
)

func TestAddition(t *testing.T) {
	if Addition(1, 2) != 3 {
		t.Error("Expected 1 (+) 2 to equal 3")
	}
	if Addition(-1, -2) != -3 {
		t.Error("Expected -1 (+) -2 to equal -3")
	}
}

```

Understand what u just made!
- More than one test function!
- Follow the rule testing in Go

Run Testing:
```sh
go test ./lib/calculate -cover
```


Run Testing with Report Coverage:
```sh
go test ./lib/calculate -coverprofile=cover.out && go tool cover -html=cover.out
```


## Test Scenario && Test Case

### What is test scenario ?

Adalah gambaran umum terhadap apa yang kita akan test. Test scenario is like a high-Level Test Case

#### Test scenario example:
- Check the create new account functionality
- Check the login functionality


### What is test case?

Adalah kumpulan langkah-langkah uji positif dan negatif dalam sebuah test scenario. Test case a set of pre-conditions, steps, expected result, status and actual result.

#### Test case example:

Test scenario: Check the login Functionality

Test case 1: Enter `valid email` and `valid password`
Test case 2: Enter `valid email` and `invalid password`
Test case 3: Enter `invalid email` and `valid password`
Test case 4: Enter `invalid email` and `invalid password`

Test scenario: `What to be tested`
Test case: `How to be tested`