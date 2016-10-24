// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	var ps httprouter.Params
	var ticketID int64
	tck := &ticket{}

	ps, ok := r.Context().Value("params").(httprouter.Params)

	if !ok {
		w.Write([]byte("No identifier specified\n"))
		return
	}

	ticketID, _ = strconv.ParseInt(ps.ByName("id"), 10, 64)

	if ticketID == 0 {
		w.Write([]byte("No identifier specified\n"))
		return
	}

	stmt, _ := postgres.Prepare("SELECT id, title, username, created_at FROM servidesk.ticket WHERE id = $1")

	defer stmt.Close()

	stmt.QueryRow(ticketID).Scan(&tck.ID, &tck.Title, &tck.Username, &tck.CreateAt)

	if tck.ID <= 0 {
		w.WriteHeader(404)
		return
	}

	response, _ := json.Marshal(tck)

	w.Header().Set("Content-Type", "application/json")

	w.Write(response)
}
