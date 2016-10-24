// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "time"

type ticket struct {
	ID       int64     `sql:"id"`
	Title    string    `sql:"title"`
	Username string    `sql:"username"`
	CreateAt time.Time `sql:"created_at"`
}

func scanTicket() *ticket {
	tck := &ticket{}

	return tck
}
