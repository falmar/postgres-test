// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "time"

type ticket struct {
	ID       int64     `json:"id" sql:"id"`
	Title    string    `json:"title" sql:"title"`
	Username string    `json:"username" sql:"username"`
	CreateAt time.Time `json:"created_at" sql:"created_at"`
}
