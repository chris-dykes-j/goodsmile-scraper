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
	_, err := conn.Exec(context.Background(), `
        INSERT INTO nendoroid (nendoroid_number) VALUES ($1) RETURNING id;

        INSERT INTO nendoroid_name (nendoroid_id, language_code, text) VALUES (id, 'en', $2);
        INSERT INTO nendoroid_name (nendoroid_id, language_code, text) VALUES (id, 'ja', $3);
        INSERT INTO nendoroid_name (nendoroid_id, language_code, text) VALUES (id, 'zh', $4);

        INSERT INTO nendoroid_description (nendoroid_id, language_code, text) VALUES (id, 'en', $5);
        INSERT INTO nendoroid_description (nendoroid_id, language_code, text) VALUES (id, 'ja', $6);
        INSERT INTO nendoroid_description (nendoroid_id, language_code, text) VALUES (id, 'zh', $7);

        INSERT INTO nendoroid_details (nendoroid_id, language_code, details) VALUES (id, 'en', $8);
        INSERT INTO nendoroid_details (nendoroid_id, language_code, details) VALUES (id, 'ja', $9);
        INSERT INTO nendoroid_details (nendoroid_id, language_code, details) VALUES (id, 'zh', $10);

        INSERT INTO nendoroid_link (nendoroid_id, language_code, text) VALUES (id, 'en', $11);
        INSERT INTO nendoroid_link (nendoroid_id, language_code, text) VALUES (id, 'ja', $12);
        INSERT INTO nendoroid_link (nendoroid_id, language_code, text) VALUES (id, 'zh', $13);

        INSERT INTO nendoroid_blog_link (nendoroid_id, language_code, text) VALUES (id, 'en', $14);
        INSERT INTO nendoroid_blog_link (nendoroid_id, language_code, text) VALUES (id, 'ja', $15);
        INSERT INTO nendoroid_blog_link (nendoroid_id, language_code, text) VALUES (id, 'zh', $16);
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
