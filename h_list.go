// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"
)

func listHandler(w http.ResponseWriter, r *http.Request) {
	stmt, err := postgres.Prepare(`SELECT t.id,
    t.title,
    COALESCE(t.username, '') as username,
    t.created_at
    FROM servidesk.ticket t`)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	defer rows.Close()

	var tickets []*ticket

	for rows.Next() {
		tck := &ticket{}

		err = rows.Scan(&tck.ID, &tck.Title, &tck.Username, &tck.CreateAt)

		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		tickets = append(tickets, tck)
	}

	response, err := json.Marshal(tickets)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(response)
}
