package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"backend/models"
)

type LoginHandler struct {
	UserStore *models.UserStorage
}

func NewLoginHandler(userStore *models.UserStorage) *LoginHandler {
	return &LoginHandler{userStore}
}

func (h LoginHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
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
	pass, err := h.UserStore.Find(reqBody.Username)

	if err != nil || reqBody.Password != pass {
		if errors.Is(err, models.UserNotFoundErr) || reqBody.Password != pass {
			resBody.Success = false
			resBody.Message = "incorrect credentials"
		} else {
			InternalServerErrorHandler(writer)
			return
		}
	} else {
		resBody.Success = true
		resBody.Message = "login successful"
	}

	jsonBytes, err := json.Marshal(resBody)

	if err != nil {
		InternalServerErrorHandler(writer)
		return
	}

	if resBody.Success {
		cookie := http.Cookie{}
		cookie.Name = "loggedIn"
		cookie.Value = "true"
		cookie.Expires = time.Now().Add(24 * time.Hour)
		cookie.Secure = false
		cookie.HttpOnly = true

		http.SetCookie(writer, &cookie)
	}

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	if _, err = writer.Write(jsonBytes); err != nil {
		InternalServerErrorHandler(writer)
		return
	}
}
