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
    conn, err := pgx.Connect(context.Background(), "postgres://chris:@localhost:5432/figures") // local db because
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
    file, err := os.Open("../extract/test.jsonl")
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
        err = insertFigure(figure, conn)
        if err != nil {
            log.Fatal(err)
        }
    }

}

func insertFigure(figure Figure, conn *pgx.Conn) error {
    _, err := conn.Exec(context.Background(), `
        `)
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
