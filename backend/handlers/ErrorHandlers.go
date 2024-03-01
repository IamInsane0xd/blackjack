package handlers

import "net/http"

func ForbiddenHandler(writer http.ResponseWriter) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/text")
	writer.WriteHeader(http.StatusForbidden)
	_, _ = writer.Write([]byte("403 forbidden"))
}

func InternalServerErrorHandler(writer http.ResponseWriter) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/text")
	writer.WriteHeader(http.StatusInternalServerError)
	_, _ = writer.Write([]byte("500 internal server error"))
}
