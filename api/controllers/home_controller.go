package controllers

import (
	"net/http"

	"github.com/mazeko/go-blog/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome")
}
