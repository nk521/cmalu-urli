package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

var db *sql.DB

// Link struct to store the original link and its shortened version
type Link struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
	Expiry      int64  `json:"expiry"`
}

// Connect to the MySQL database
func init() {
	var err error
	var connstr string = fmt.Sprintf("mason:p5g41tmlx@tcp(%s:%s)/cmalu_urli", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	db, err = sql.Open("mysql", connstr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}
}

// shortenLink function to shorten the given link
func shortenLink(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var link Link
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid request body")
		fmt.Fprint(w, err)
		return
	}

	link.Expiry = time.Now().Unix() + 24*60*60

	// make sure that you always find a unique short link
	for {
		link.ShortURL = generateShortURL()

		query := fmt.Sprintf("INSERT INTO links (original_url, short_url, expiry) VALUES ('%s', '%s', %d)", link.OriginalURL, link.ShortURL, link.Expiry)
		_, err = db.Exec(query)

		if err != nil {
			if strings.Contains(err.Error(), "1062") {
				continue
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			break
		}
	}

	// output the short link
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "c.raju.dev/c/"+link.ShortURL)
}

// redirect if short link is valid
func redirectLink(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var link Link
	short_url := ps.ByName("short_url")
	query := fmt.Sprintf("SELECT original_url, short_url, expiry FROM links WHERE short_url = '%s'", short_url)
	err := db.QueryRow(query).Scan(&link.OriginalURL, &link.ShortURL, &link.Expiry)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if link.Expiry < time.Now().Unix() {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	redirect_url := link.OriginalURL
	http.Redirect(w, r, redirect_url, http.StatusMovedPermanently)
}

// generateShortURL function to generate a unique shortened link
func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyz"
	word := ""
	for i := 0; i < 5; i++ {
		word += string(letters[rand.Intn(26)])
	}
	return word
}

func main() {
	router := httprouter.New()
	router.POST("/cmalu", shortenLink)
	router.GET("/c/:short_url", redirectLink)

	log.Println("Server starting on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
