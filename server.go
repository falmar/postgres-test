// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

var postgres *sql.DB

func main() {
	var err error

	postgres, err = sql.Open("postgres", "")

	postgres.Ping()

	if err != nil {
		log.Fatal(err)
	}

	r := httprouter.New()

	setRoutes(r)

	http.ListenAndServe(":8080", r)
}
