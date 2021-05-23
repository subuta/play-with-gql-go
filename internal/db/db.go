// +build !js,!wasm

// Loaded on non-WASM.
package db

import (
	"database/sql"
	"github.com/k0kubun/pp"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Run () {
	// ----- Try [mattn/go-sqlite3: sqlite3 driver for go using database/sql](https://github.com/mattn/go-sqlite3)
	// Won't work with WASM build.
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select 1 + 3;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var count int
		rows.Scan(&count)
		pp.Println(count)
	}
}