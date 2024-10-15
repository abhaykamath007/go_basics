package main

import (
	"bufio"
	"fmt"
	"os"
)

type Book struct {
	Title  string
	Author string
	Year   int
}

func addBook(library []Book, book Book) []Book {
	library = append(library, book)
	return library
}

func deleteBook(library []Book, title string) ([]Book, bool) {
	for i, book := range library {
		if book.Title == title {
			return append(library[:i], library[i+1:]...), true
		}
	}
	return library, false
}

func searchBook(library []Book, title string) {
	for _, book := range library {
		if book.Title == title {
			fmt.Println("Found:", book.Title, "by", book.Author, "in the year", book.Year)
			return // Exit the function if the book is found
		}
	}
	fmt.Println("Book not found.")
}

func viewBooks(library []Book) {
	if len(library) == 0 {
		fmt.Println("No books in the library.")
		return
	}
	fmt.Println("Books in the Library:")
	for i, book := range library {
		fmt.Printf("%d: %s by %s (%d)\n", i+1, book.Title, book.Author, book.Year)
	}
}

func main() {

	var library []Book
	scanner := bufio.NewScanner(os.Stdin)

	library = addBook(library, Book{"The Hobbit", "J.R.R. Tolkien", 1937})
	library = addBook(library, Book{"1984", "George Orwell", 1949})
	library = addBook(library, Book{"To Kill a Mockingbird", "Harper Lee", 1960})
	library = addBook(library, Book{"Harry Potter", "JK Rowling", 1990})

	for {
		fmt.Println("\nLibrary Management System:")
		fmt.Println("1. View Books")
		fmt.Println("2. Add Book")
		fmt.Println("3. Delete Book")
		fmt.Println("4. Search Book")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			viewBooks(library)

		case 2:
			var title, author string
			var year int
			fmt.Print("Enter book title: ")
			scanner.Scan()
			title = scanner.Text()
			fmt.Print("Enter author: ")
			scanner.Scan()
			author = scanner.Text()
			fmt.Print("Enter publication year: ")
			fmt.Scanln(&year)
			library = addBook(library, Book{title, author, year})
			fmt.Println("Book added successfully!")

		case 3:
			fmt.Print("Enter the title of the book to delete: ")
			var title string
			scanner.Scan()
			title = scanner.Text()
			var res bool
			library, res = deleteBook(library, title)
			if res {
				fmt.Println("Book deleted successfully!")
			} else {
				fmt.Println("Book not found.")
			}

		case 4:
			fmt.Print("Enter the title of the book to search: ")
			var titleToSearch string
			scanner.Scan()
			titleToSearch = scanner.Text()
			searchBook(library, titleToSearch)

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")

		}
	}
}
