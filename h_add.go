// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func addHandler(w http.ResponseWriter, r *http.Request) {
	tck := &ticket{}

	err := json.NewDecoder(r.Body).Decode(tck)

	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	stmt, err := postgres.Prepare(`
    INSERT INTO servidesk.ticket
    (title, username, created_at) VALUES ($1, $2, $3)
    RETURNING id;`,
	)

	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(tck.Title, tck.Username, tck.CreateAt).Scan(&tck.ID)

	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	response, err := json.Marshal(tck)

	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(response)

}
