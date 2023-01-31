/*
 * Copyright (c)  Nikita Cherkasov, 2023.
 * Spotter Project
 */

package handler

import (
	"SpotterBackend/src/internal/constants"
	"SpotterBackend/src/internal/handler"
	"SpotterBackend/src/internal/user/model"
	"SpotterBackend/src/internal/user/service"
	"SpotterBackend/src/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
)

const endpointCreateByEmail = "/auth/signup/email/"
const endpointCreateByPhone = "/auth/signup/phone/"
const endpointAboutUser = "/user/:id"

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
	router.POST(endpointCreateByEmail, h.createByEmail)
	router.POST(endpointCreateByPhone, h.createByPhone)
	router.GET(endpointAboutUser, h.AboutUser)
}

func (h *userHandler) createByEmail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := io.ReadAll(r.Body)
	h.log.Errorf("%s", string(body))
	if err != nil {
		h.log.Errorf("Error while read body of Request: %s", err)
		utils.WriteResponseError(w, errors.New(constants.InternalServerError))
		return
	}
	dto := model.CreateByEmailDTO{}
	err = json.Unmarshal(body, &dto)
	if err != nil {
		h.log.Errorf("Error while unmarshall: %s", err)
		utils.WriteResponseError(w, errors.New(constants.InternalServerError))
		return
	}
	var userId int64
	userId, err = h.service.CreateUserByEmail(dto)
	if err != nil {
		utils.WriteResponseError(w, err)
		return
	}
	respJSON, err := json.Marshal(map[string]int64{"id": userId})
	if err != nil {
		h.log.Errorf("Error while marshall: %s", err)
		utils.WriteResponseError(w, errors.New(constants.InternalServerError))
		return
	}
	w.Write(respJSON)
}

func (h *userHandler) createByPhone(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Errorf("Error while read body of Request: %s", err)
		utils.WriteResponseError(w, errors.New(constants.InternalServerError))
		return
	}
	dto := model.CreateByPhoneNumberDTO{}
	err = json.Unmarshal(body, &dto)
	if err != nil {
		h.log.Errorf("Error while unmarshall: %s", err)
		utils.WriteResponseError(w, errors.New(constants.InternalServerError))
		return
	}
	var userId int64
	userId, err = h.service.CreateUserByPhone(dto)
	if err != nil {
		utils.WriteResponseError(w, err)
		return
	}
	respJSON, err := json.Marshal(map[string]int64{"id": userId})
	if err != nil {
		h.log.Errorf("Error while marshall: %s", err)
		utils.WriteResponseError(w, errors.New(constants.InternalServerError))
		return
	}
	w.Write(respJSON)
}
func (h *userHandler) AboutUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO: remove _ (handle error)
	id, _ := strconv.Atoi(ps.ByName("id"))
	fmt.Println(id)
	user, err := h.service.AboutUser(id)
	fmt.Println(user)
	if err != nil {
		utils.WriteResponseError(w, err)
		return
	}
	respJSON, err := json.Marshal(map[string]int64{"id": 0})
	if err != nil {
		h.log.Errorf("Error while marshall: %s", err)
		utils.WriteResponseError(w, errors.New(constants.InternalServerError))
		return
	}
	w.Write(respJSON)
}
