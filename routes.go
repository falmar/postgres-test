// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func wrapper(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)

		h.ServeHTTP(w, r.WithContext(ctx))
	}
}

func setRoutes(r *httprouter.Router) {

	r.GET("/tickets", wrapper(http.HandlerFunc(listHandler)))
	r.POST("/tickets", wrapper(http.HandlerFunc(addHandler)))
	r.GET("/tickets/:id", wrapper(http.HandlerFunc(getHandler)))
	r.PUT("/tickets/:id", wrapper(http.HandlerFunc(updateHandler)))
	r.DELETE("/tickets/:id", wrapper(http.HandlerFunc(deleteHandler)))

}
