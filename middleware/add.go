package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DatasetText struct {
	Text   string
	IsSpam bool
}

func addConnection() (*sql.DB, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db, err := sql.Open("postgres", os.Getenv("POSGRESQL_URL"))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AddText(res http.ResponseWriter, req *http.Request) {
	db, err := addConnection()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var dataset DatasetText
	err = json.NewDecoder(req.Body).Decode(&dataset)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	var text = dataset.Text
	var isSpam = dataset.IsSpam
	var tableQuery = ""
	fmt.Println(isSpam)
	if isSpam {
		tableQuery = "INSERT INTO spam (sentence) VALUES ($1)"
	} else {
		tableQuery = "INSERT INTO nonspam (sentence) VALUES ($1)"
	}

	qr, err := db.Exec(tableQuery, text)
	if err != nil {
		fmt.Println("error di query")
		log.Fatal(err)
	}
	fmt.Println("Query result: ", qr)
}

func DetectText(res http.ResponseWriter, req *http.Request) {

}
