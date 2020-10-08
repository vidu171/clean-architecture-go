package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vidu171/clean-architecture-go/domain"

	"github.com/vidu171/clean-architecture-go/usecases"
)

type AuthorController struct {
	authorInteractor usecases.AuthorInteractor
}

func NewAuthorController(authorInteractor usecases.AuthorInteractor) *AuthorController {
	return &AuthorController{authorInteractor}
}

func (controller *AuthorController) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var author domain.Author
	err := json.NewDecoder(req.Body).Decode(&author)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}
	err2 := controller.authorInteractor.CreateAuthor(author)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
}
