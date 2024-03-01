package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"backend/models"
)

type RegisterHandler struct {
	UserStore *models.UserStorage
}

func NewRegisterHandler(userStore *models.UserStorage) *RegisterHandler {
	return &RegisterHandler{userStore}
}

func (h RegisterHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		ForbiddenHandler(writer)
		return
	}

	var reqBody models.LoginRegisterRequest

	if err := json.NewDecoder(request.Body).Decode(&reqBody); err != nil {
		InternalServerErrorHandler(writer)
		return
	}

	var resBody models.UserActionResponse

	if err := h.UserStore.Add(reqBody.Username, reqBody.Password); err != nil {
		if !errors.Is(err, models.UserAlreadyExistsErr) {
			InternalServerErrorHandler(writer)
			return
		}

		resBody.Success = false
		resBody.Message = err.Error()

	} else {
		resBody.Success = true
		resBody.Message = "user created"
	}

	jsonBytes, err := json.Marshal(resBody)

	if err != nil {
		InternalServerErrorHandler(writer)
		return
	}

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	if _, err = writer.Write(jsonBytes); err != nil {
		InternalServerErrorHandler(writer)
		return
	}
}
