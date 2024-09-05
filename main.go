package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "qwerty"
	dbname   = "test"
)

const (
	createTableBook = `CREATE TABLE book (
		book_id serial primary key,
		title VARCHAR(50),
		author VARCHAR(30),
		price DECIMAL(8, 2),
		amount INT
	);`

	insertTableBook = `INSERT INTO book (title, author, price, amount) 
		VALUES ('Мастер и Маргарита', 'Булгаков М.А.', '670.99', '3');`

	selectTableBook = `SELECT book_id, title, author, price, amount FROM book;`
)

type Book struct {
	BookID int
	Title  string
	Author string
	Price  float64
	Amount int
}

func main() {
	needCreateTable := flag.Bool("n", false, "if need create table")
	flag.Parse()

	ctx := context.Background()

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	if *needCreateTable {
		if _, err = db.Exec(createTableBook); err != nil {
			log.Fatal(err)
		}
	}

	if _, err = db.Exec(insertTableBook); err != nil {
		log.Fatal(err)
	}

	var books []Book

	rows, err := db.QueryContext(ctx, selectTableBook)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var book Book

		if err = rows.Scan(&book.BookID, &book.Title, &book.Author, &book.Price, &book.Amount); err != nil {
			log.Fatal(err)
		}

		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// For one row
	// if err = db.QueryRowContext(ctx, selectTableBook).
	// 	Scan(&book.BookID, &book.Title, &book.Author, &book.Price, &book.Amount); err != nil {
	// 	log.Fatal(err)
	// }

	log.Println(books)
}
