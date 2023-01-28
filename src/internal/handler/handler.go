/*
 * Copyright (c)  Nikita Cherkasov, 2023.
 * Spotter Project
 */

package handler

import "github.com/julienschmidt/httprouter"

type Handler interface {
	Register(router *httprouter.Router)
}
