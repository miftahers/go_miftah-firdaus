package main

import (
	"fmt"
	"github.com/google/uuid"
	"sort"
	"strings"
)

type Book struct {
	Id       string
	Title    string
	Price    int
	Category string
}

var (
	Books     []Book
	BookTitle []string
	input     int
)

func main() {
	for input != 5 {
		input = menu()
		switch input {
		case 1:
			getAllBook()
		case 2:
			CreateBook()
		case 3:
			UpdateBook()
		case 4:
			DeleteBook()
		default:
			fmt.Print("Invalid input, try again.\n")
		}

	}
}

func menu() int {
	var input int
	fmt.Println("== Book Management ==")
	fmt.Println("1. Get All Book")
	fmt.Println("2. Create a Book")
	fmt.Println("3. Update a Book")
	fmt.Println("4. Delete a Book")
	fmt.Println("5. Exit")
	fmt.Print("Choose your menu: ")
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Println("Invalid input!")
	}
	return input
}
func getAllBook() {
	fmt.Println("All Books")

	if len(Books) == 0 {
		fmt.Println("No book added, Add book first!")
	} else {
		for _, v := range BookTitle {
			for i := range Books {
				if Books[i].Title == v {
					fmt.Println("==")
					fmt.Printf("ID\t\t: %v\n", Books[i].Id)
					fmt.Printf("Title\t\t: %v\n", Books[i].Title)
					fmt.Printf("Price\t\t: %v\n", Books[i].Price)
					fmt.Printf("Category\t\t: %v\n", Books[i].Category)
					fmt.Println("==")
				}
			}
		}
	}
}

func CreateBook() {
	var id, title, category string
	var price int
	var err error

	fmt.Println()
	id = CreateUUID()
	fmt.Print("enter title\t\t: ")
	_, err = fmt.Scanf("%s", &title)
	fmt.Print("enter price\t\t: ")
	_, err = fmt.Scanf("%d", &price)
	fmt.Print("enter category\t: ")
	_, err = fmt.Scanf("%s", &category)

	if err != nil {
		fmt.Println("invalid input")
	} else {
		BookTitle = append(BookTitle, title)
		sort.Strings(BookTitle)
		Books = append(Books, Book{
			Id:       id,
			Title:    title,
			Price:    price,
			Category: category,
		})
		fmt.Println("Book added!")
	}
}

func UpdateBook() {
	var id, title, category string
	var price int
	var err error

	fmt.Print("enter id\t: ")
	_, err = fmt.Scanf("%s", &id)
	fmt.Print("enter title\t: ")
	_, err = fmt.Scanf("%s", &title)
	fmt.Print("enter price\t: ")
	_, err = fmt.Scanf("%d", &price)
	fmt.Print("enter category\t: ")
	_, err = fmt.Scanf("%s", &category)

	if err != nil {
		fmt.Println("Invalid input!")
	} else {
		for i := range Books {
			if Books[i].Id == id {
				Books[i].Title = title
				Books[i].Price = price
				Books[i].Category = category
				sort.Strings(BookTitle)
				fmt.Println("Book updated!")
			} else {
				fmt.Println("ID not found!")
			}
		}

	}
}

func DeleteBook() {
	var id string
	fmt.Print("enter id		: ")
	_, err := fmt.Scanf("%s", &id)
	if err != nil {
		fmt.Println("Invalid input")
	} else {
		for i := range Books {
			if Books[i].Id == id {
				Books = append(Books[:i], Books[i+1:]...)
				BookTitle = append(BookTitle[:i], BookTitle[i+1:]...)
				sort.Strings(BookTitle)
			}
		}
	}
}

func CreateUUID() string {
	uuidWithHyphen := uuid.New()
	result := strings.Replace(uuidWithHyphen.String(), "", "", -1)
	return result
}
