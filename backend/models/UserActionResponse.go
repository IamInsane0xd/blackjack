package models

type UserActionResponse struct {
	Success bool   "json:\"success\""
	Message string "json:\"message\""
}
