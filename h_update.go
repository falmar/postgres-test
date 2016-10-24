// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func updateHandler(w http.ResponseWriter, r *http.Request) {
	tck := &ticket{}

	ticketID, err := ticketIDFromCtx(r.Context())

	if ticketID == 0 || err != nil {
		http.Error(w, "No identifier specified\n", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(tck)

	stmt, err := postgres.Prepare("UPDATE servidesk.ticket SET title = $2, username = $3, created_at = $4 WHERE id = $1")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(ticketID, tck.Title, tck.Username, tck.CreateAt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

}
