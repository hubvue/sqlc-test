package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"tutorial.sqlc.dev/app/tutorial"
)

func run() error {
	ctx := context.Background()

	db, err := sql.Open("mysql", "root:xxx@tcp(49.232.212.40:3306)/sqlctest")
	if err != nil {
		return err
	}

	queries := tutorial.New(db)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Println(authors)

	// crate an author
	result, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of the C Programming Language an The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}
	insertedAuthorId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(insertedAuthorId)

	// get the author we just inserted
	featchedAuthor, err := queries.GetAuthor(ctx, insertedAuthorId)
	if err != nil {
		return err
	}
	log.Println(reflect.DeepEqual(insertedAuthorId, featchedAuthor.ID))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
