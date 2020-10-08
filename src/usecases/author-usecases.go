package usecases

import (
	"log"

	"github.com/vidu171/clean-architecture-go/domain"
)

type AuthorInteractor struct {
	AuthorRepository domain.AuthorRepository
}

func NewAuthorInteractor(AuthorRepository domain.AuthorRepository) AuthorInteractor {
	return AuthorInteractor{AuthorRepository}
}

func (interactor *AuthorInteractor) CreateAuthor(author domain.Author) error {
	err := interactor.AuthorRepository.SaveAuthor(author)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
