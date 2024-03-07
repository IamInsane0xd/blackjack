package handlers

import (
	"encoding/json"
	"net/http"

	"backend/models"
)

type ForgotPasswordHandler struct {
	userStore *models.UserStorage
}

func NewForgotPasswordHandler(userStore *models.UserStorage) *ForgotPasswordHandler {
	return &ForgotPasswordHandler{userStore}
}

func (h ForgotPasswordHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		ForbiddenHandler(writer)
		return
	}

	var reqBody models.ForgotPasswordRequest

	if err := json.NewDecoder(request.Body).Decode(&reqBody); err != nil {
		InternalServerErrorHandler(writer)
		return
	}

	pass, err := h.userStore.Find(reqBody.Username)
	var resBody models.UserActionResponse

	if err != nil {
		resBody.Success = false
		resBody.Message = err.Error()
	} else {
		resBody.Success = true
		resBody.Message = pass
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
