// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"context"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func ticketIDFromCtx(ctx context.Context) (int64, error) {
	if ps, ok := ctx.Value("params").(httprouter.Params); ok {
		ticketID, err := strconv.ParseInt(ps.ByName("id"), 10, 64)

		return ticketID, err
	}

	return 0, nil
}
