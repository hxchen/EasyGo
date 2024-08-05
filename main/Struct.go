package main

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books
	var Book2 Books

	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.book_id = 6495407

	Book2.title = "Telecom Billing"
	Book2.author = "Zara Ali"
	Book2.subject = "Telecom Billing Tutorial"
	Book2.book_id = 6495700

	printBook(Book1)
	printBook(Book2)
}

func printBook(book Books) {
	println("Book title: ", book.title)
	println("Book author: ", book.author)
	println("Book subject: ", book.subject)
	println("Book book_id: ", book.book_id)
}
