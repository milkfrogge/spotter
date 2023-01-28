/*
 * Copyright (c)  Nikita Cherkasov, 2023.
 * Spotter Project
 */

package handler

import (
	"SpotterBackend/src/internal/handler"
	"SpotterBackend/src/internal/user/model"
	"SpotterBackend/src/internal/user/service"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type userHandler struct {
	log     *logrus.Logger
	service *service.UserService
}

func NewUserHandler(logger *logrus.Logger, userService *service.UserService) handler.Handler {
	return &userHandler{
		log:     logger,
		service: userService,
	}
}

func (h *userHandler) Register(router *httprouter.Router) {
	h.log.Println("Register routes for user")
	router.POST("/auth/signup/", h.createByEmail)
}

func (h *userHandler) createByEmail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.log.Println("in createEmail")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Errorf("Error while read body of Request: %s", err)
		w.WriteHeader(500)
		return
	}
	dto := model.CreateByEmailDTO{}
	err = json.Unmarshal(body, &dto)
	if err != nil {
		h.log.Errorf("Error while unmarshall: %s", err)
		w.WriteHeader(500)
		return
	}
	var userId int64
	userId, err = h.service.CreateUserByEmail(dto)
	if err != nil {
		h.log.Errorf("Error while creating user: %s", err)
		w.WriteHeader(500)
		return
	}
	respJSON, err := json.Marshal(map[string]int64{"id": userId})
	if err != nil {
		h.log.Errorf("Error while marshall: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Write(respJSON)
}
