package usecases

import (
	"github.com/vidu171/clean-architecture-go/domain"
)

type BookInteractor struct {
	BookRepository domain.BookRepository
	Logger         Logger
}

func NewBookInteractor(repository domain.BookRepository, Logger Logger) BookInteractor {
	return BookInteractor{repository, Logger}
}

func (interactor *BookInteractor) CreateBook(book domain.Book) error {
	err := interactor.BookRepository.SaveBook(book)
	if err != nil {
		interactor.Logger.Log(err.Error())
		return err
	}
	return nil
}

func (interactor *BookInteractor) FindAll() ([]*domain.Book, error) {
	results, err := interactor.BookRepository.FindAll()
	if err != nil {
		interactor.Logger.Log(err.Error())
		return nil, err
	}
	return results, nil
}
