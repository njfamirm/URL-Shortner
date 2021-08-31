package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/amir-mhmd-najafi/URL-Shortner/database/app"
	"github.com/amir-mhmd-najafi/URL-Shortner/database/databaseconfig"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = databaseconfig.ConnectToDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("end init")
}

func main() {
	
	// static file
	fs := http.FileServer(http.Dir("../template/statistic"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", redirect)
	http.HandleFunc("/shortened", shortened)
	err := http.ListenAndServe(":5500", nil)
	if err != nil {
		panic(err)
	}
}

// home page => input for link => shortened link
func redirect(w http.ResponseWriter, r *http.Request) {
	app.Redirect(w, r, DB)
}

// show counter link and other data
func shortened(w http.ResponseWriter, r *http.Request) {
	app.Shortner(w, r, DB)
}
