package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
Drop Table artists;
CREATE TABLE artists (
    id text,
    name text,
    description text
);
`

func main() {
	db, err := sqlx.Open("sqlite3", "./test.db")
	if err != nil {
		fmt.Println(err.Error())
	}

	var artists []Artist
	artists = append(artists, Artist{Description: gofakeit.Sentence(5), Id: gofakeit.UUID(), Name: gofakeit.Name()})
	artists = append(artists, Artist{Description: gofakeit.Sentence(6), Id: gofakeit.UUID(), Name: gofakeit.Name()})
	artists = append(artists, Artist{Description: gofakeit.Sentence(4), Id: gofakeit.UUID(), Name: gofakeit.Name()})
	artists = append(artists, Artist{Description: gofakeit.Sentence(12), Id: gofakeit.UUID(), Name: gofakeit.Name()})

	db.MustExec(schema)

	tx := db.MustBegin()
	for _, u := range artists {
		fmt.Printf("%#v\n", &u)
		_, err := tx.NamedExec("INSERT into artists (id,name,description) values(:id,:name,:description)", u)
		checkErr(err)
	}
	tx.Commit()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Artist struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
