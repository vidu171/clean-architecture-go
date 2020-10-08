package usecases

import (
	"github.com/vidu171/clean-architecture-go/domain"
)

type AuthorInteractor struct {
	AuthorRepository domain.AuthorRepository
	Logger             Logger
}

func NewAuthorInteractor(AuthorRepository domain.AuthorRepository, Logger Logger) AuthorInteractor {
	return AuthorInteractor{AuthorRepository, Logger}
}

func (interactor *AuthorInteractor) CreateAuthor(author domain.Author) error {
	err := interactor.AuthorRepository.SaveAuthor(author)
	if err != nil {
		interactor.Logger.Log(err.Error())
		return err
	}
	return nil
}
