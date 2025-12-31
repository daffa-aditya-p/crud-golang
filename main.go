package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Donation struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Amount int    `json:"amount"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/donasigolang_db")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/donation", donationHandler)
	http.HandleFunc("/donation/", donationByIDHandler)

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// =====================
// HANDLERS
// =====================
func donationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
                w.Header().Set("Content-Type", "application/json")

		rows, _ := db.Query("SELECT id, name, email, amount FROM donations")
		var data []Donation

		for rows.Next() {
			var d Donation
			rows.Scan(&d.ID, &d.Name, &d.Email, &d.Amount)
			data = append(data, d)
		}
		json.NewEncoder(w).Encode(data)

	case "POST":
		var d Donation
		json.NewDecoder(r.Body).Decode(&d)
		db.Exec(
			"INSERT INTO donations (name, email, amount) VALUES (?, ?, ?)",
			d.Name, d.Email, d.Amount,
		)
		w.Write([]byte("created"))

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func donationByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/donation/"):]
	id, _ := strconv.Atoi(idStr)

	switch r.Method {

	case "GET":
		var d Donation
		db.QueryRow(
			"SELECT id, name, email, amount FROM donations WHERE id=?",
			id,
		).Scan(&d.ID, &d.Name, &d.Email, &d.Amount)

		json.NewEncoder(w).Encode(d)

	case "DELETE":
		db.Exec("DELETE FROM donations WHERE id=?", id)
		w.Write([]byte("deleted"))

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
