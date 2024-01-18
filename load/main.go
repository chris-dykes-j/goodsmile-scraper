package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	// Connect to database
	conn, err := pgx.Connect(context.Background(), "postgres://chris:@localhost:5432/figures")
	if err != nil {
		fmt.Println("Damnit")
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	// Run init script
	init, err := os.ReadFile("init.sql")
	if err != nil {
		log.Fatal(err)
	}

	sql := string(init)
	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}

	// Read jsonl file, line by line
	file, err := os.Open("../scrape/test.jsonl")
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	for decoder.More() {
		var figure Figure
		if err := decoder.Decode(&figure); err != nil {
			log.Fatal(err)
		}

		// Create insert statements
		err = insertNendoroid(figure, conn)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func insertNendoroid(figure Figure, conn *pgx.Conn) error {
	init, err := os.ReadFile("insert.sql")
	if err != nil {
		log.Fatal(err)
	}

	sql := string(init)
	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Exec(context.Background(), `
        SELECT insert_nendo(
            $1::int,
            $2::varchar, $3::varchar, $4::varchar,
            $5::text, $6::text, $7::text,
            $8::jsonb, $9::jsonb, $10::jsonb,
            $11::text, $12::text, $13::text,
            $14::text, $15::text, $16::text);
        `,
		figure.English.ItemNumber,
		figure.English.Name, figure.Japanese.Name, figure.Chinese.Name,
		figure.English.Description, figure.Japanese.Description, figure.Chinese.Description,
		figure.English.Details, figure.Japanese.Details, figure.Chinese.Details,
		figure.English.ItemLink, figure.Japanese.ItemLink, figure.Chinese.ItemLink,
		figure.English.BlogLink, figure.Japanese.BlogLink, figure.Chinese.BlogLink)
	if err != nil {
		return err
	}
	return nil
}

// Input Data Structs
type Figure struct {
	English  Data `json:"en"`
	Japanese Data `json:"ja"`
	Chinese  Data `json:"zh"`
}

type Data struct {
	ItemNumber  string    `json:"itemNumber"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ItemLink    string    `json:"itemLink"`
	BlogLink    string    `json:"blogLink"`
	Details     []Details `json:"details"`
}

type Details struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
