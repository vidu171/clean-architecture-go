package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vidu171/clean-architecture-go/infrastructure/db"

	"github.com/vidu171/clean-architecture-go/usecases"

	"github.com/vidu171/clean-architecture-go/interface/controllers"

	"github.com/vidu171/clean-architecture-go/infrastructure/router"
	"github.com/vidu171/clean-architecture-go/interface/repository"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
	dbHandler  db.DBHandler
)

func getBookController() controllers.BookController {
	bookRepo := repository.NewBookRepo(dbHandler)
	bookInteractor := usecases.NewBookInteractor(bookRepo)
	bookController := controllers.NewBookController(bookInteractor)
	return *bookController
}

func getAuthorController() controllers.AuthorController {
	authorRepo := repository.NewAuthorRepo(dbHandler)
	authorInteractor := usecases.NewAuthorInteractor(authorRepo)
	authorController := controllers.NewAuthorController(authorInteractor)
	return *authorController
}

func main() {

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "App is up and running..")
	})
	var err error
	dbHandler, err = db.NewDBHandler("mongodb://localhost:27017", "bookstore")
	if err != nil {
		log.Println("Unable to connect to the DataBase")
		return
	}
	bookController := getBookController()
	authorController := getAuthorController()
	httpRouter.POST("/book/add", bookController.Add)
	httpRouter.GET("/book", bookController.FindAll)
	httpRouter.POST("/author/add", authorController.Add)
	httpRouter.SERVE(":8000")
}
