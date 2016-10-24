// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"
)

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	ticketID, err := ticketIDFromCtx(r.Context())

	if ticketID == 0 || err != nil {
		http.Error(w, "No identifier specified\n", http.StatusNotFound)
		return
	}

	stmt, err := postgres.Prepare("DELETE FROM servidesk.ticket WHERE id = $1")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(ticketID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

}
