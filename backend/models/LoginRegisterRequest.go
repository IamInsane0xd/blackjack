package models

type LoginRegisterRequest struct {
	Username string "json:\"username\""
	Password string "json:\"password\""
}
