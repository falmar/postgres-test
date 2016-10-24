// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

var postgres *sql.DB

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	conString := `
		host=` + os.Getenv("DB_HOST") + `
		user=` + os.Getenv("DB_USER") + `
		dbname=` + os.Getenv("DB_NAME") + `
		password=` + os.Getenv("DB_PASSWORD") + `
		sslmode=` + os.Getenv("DB_SSLMODE")

	postgres, err = sql.Open("postgres", conString)

	postgres.Ping()

	if err != nil {
		log.Fatal(err)
	}

	r := httprouter.New()

	setRoutes(r)

	http.ListenAndServe(":8080", r)
}
