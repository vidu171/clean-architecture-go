package usecases

import (
	"log"

	"github.com/vidu171/clean-architecture-go/domain"
)

type BookInteractor struct {
	BookRepository domain.BookRepository
}

func NewBookInteractor(repository domain.BookRepository) BookInteractor {
	return BookInteractor{repository}
}

func (interactor *BookInteractor) CreateBook(book domain.Book) error {
	err := interactor.BookRepository.SaveBook(book)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (interactor *BookInteractor) FindAll() ([]*domain.Book, error) {
	results, err := interactor.BookRepository.FindAll()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return results, nil
}
