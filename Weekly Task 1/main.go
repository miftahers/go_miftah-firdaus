package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Book struct {
	Id       string
	Title    string
	Price    int
	Category string
}

var BookTitle []string
var Books []Book

func main() {
	var input int

	for input != 5 {
		fmt.Println("== Book Management ==")
		fmt.Println("1. Get All Book")
		fmt.Println("2. Create a Book")
		fmt.Println("3. Update a Book")
		fmt.Println("4. Delete a Book")
		fmt.Println("5. Exit")
		fmt.Print("Choose your menu: ")
		if _, err := fmt.Scan(&input); err != nil {
			fmt.Println(err.Error())
		}

		switch input {
		case 1:
			GetBooks()
		case 2:
			if err := CreateBook(); err != nil {
				fmt.Println(err.Error())
			}
		case 3:
			if err := UpdateBook(); err != nil {
				fmt.Println(err.Error())
			}
		case 4:
			if err := DeleteBook(); err != nil {
				fmt.Println(err.Error())
			}
		case 5:
			fmt.Println("Thanks for using this program!")
			time.Sleep(3 * time.Second)
		default:
			fmt.Print("Invalid input, try again next time.\n")
		}
	}
}

func GetBooks() {

	fmt.Println()
	fmt.Println("All Books")

	if len(Books) == 0 {
		fmt.Println("No Book added, Add Book first!")
		fmt.Println()
		time.Sleep(2 * time.Second)
	} else {
		for _, v := range BookTitle {
			for i := range Books {
				if Books[i].Title == v {
					fmt.Println("==")
					fmt.Printf("ID\t\t: %v\n", Books[i].Id)
					fmt.Printf("Title\t\t: %v\n", Books[i].Title)
					fmt.Printf("Price\t\t: %v\n", Books[i].Price)
					fmt.Printf("Category\t: %v\n", Books[i].Category)
					fmt.Println("==")
				}
				time.Sleep(10 * time.Millisecond)
			}
		}
		time.Sleep(2 * time.Second)
	}
}

func CreateBook() error {
	var id, title, category string
	var price int
	var book Book

	id = CreateUUID()
	book.Id = id

	fmt.Print("enter title\t\t: ")
	_, err := fmt.Scan(&title)
	if err != nil {
		return err
	}
	book.Title = title

	fmt.Print("enter price\t\t: ")
	_, err = fmt.Scan(&price)
	if err != nil {
		return err
	}
	book.Price = price

	fmt.Print("enter category\t\t: ")
	_, err = fmt.Scan(&category)
	if err != nil {
		return err
	}
	book.Category = category

	BookTitle = append(BookTitle, title)
	sort.Strings(BookTitle)
	Books = append(Books, book)
	fmt.Println("Book added!")

	return nil
}

func UpdateBook() error {
	var id, title, category string
	var price, index int
	var err error

	fmt.Print("enter id\t: ")
	if _, err = fmt.Scan(&id); err != nil {
		return err
	}
	for i := range Books {
		if Books[i].Id == id {
			index = i
			fmt.Println("ID found!")
			break
		} else {
			return errors.New("ID Not Found :(")
		}
	}
	fmt.Print("enter title\t: ")
	if _, err = fmt.Scan(&title); err != nil {
		return err
	}

	fmt.Print("enter price\t: ")
	if _, err = fmt.Scan(&price); err != nil {
		return err
	}

	fmt.Print("enter category\t: ")
	if _, err = fmt.Scan(&category); err != nil {
		return err
	}

	Books[index].Title = title
	Books[index].Price = price
	Books[index].Category = category
	BookTitle[index] = title
	sort.Strings(BookTitle)
	fmt.Println("Book updated!")

	return nil
}

func DeleteBook() error {
	var id string
	var index int
	fmt.Print("enter id		: ")
	_, err := fmt.Scan(&id)
	if err != nil {
		return err
	}

	for i := range Books {
		if Books[i].Id == id {
			index = i
			fmt.Println("ID Found")
			break
		} else {
			return errors.New("ID not found")
		}
	}

	for i := range BookTitle {
		if BookTitle[i] == Books[index].Title {

			if len(BookTitle) == 1 {
				BookTitle = nil
			} else if len(BookTitle) == 2 {
				if i == 0 {
					BookTitle = BookTitle[i+1:]
				} else {
					BookTitle = BookTitle[:i]
				}
			} else {
				if i != len(BookTitle)-1 {
					BookTitle = append(BookTitle[:i], BookTitle[i+1:]...)
				} else {
					BookTitle = BookTitle[:i]
				}
			}

			if len(BookTitle) > 1 {
				sort.Strings(BookTitle)
			}
		}
	}
	if len(Books) == 1 {
		Books = nil
	} else if len(Books) == 2 {
		if index == 0 {
			Books = Books[index+1:]
		} else {
			Books = Books[:index]
		}
	} else {
		if index != len(Books)-1 {
			Books = append(Books[:index], Books[index+1:]...)
		} else {
			Books = Books[:index]
		}
	}

	return nil
}

func CreateUUID() string {
	uuidWithHyphen := uuid.New()
	result := strings.Replace(uuidWithHyphen.String(), "", "", -1)
	return result
}
