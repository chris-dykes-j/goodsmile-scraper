package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
    // Connect to database
    conn, err := pgx.Connect(context.Background(), "connectionstring")
    if err != nil {
        log.Fatal(err)
    }

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
    figure := Figure{}

    // Create insert statements
    err = insertFigure(figure, conn)
    if err != nil {
        log.Fatal(err)
    }

    // Commit transaction
}

func insertFigure(figure Figure, conn *pgx.Conn) error {
    sql := createInsertScript(figure)
    _, err := conn.Exec(context.Background(), sql)
    if err != nil {
        log.Fatal(err)
    }
    return nil
}

func createInsertScript(figfigure Figure) string {
    return ""
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
