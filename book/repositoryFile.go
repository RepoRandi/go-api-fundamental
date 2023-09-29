package book

import "fmt"

type fileRepository struct {
}

func NewFileRepository() *fileRepository {
	return &fileRepository{}
}

func (repo *fileRepository) FindAll() ([]Book, error) {

	var book []Book

	fmt.Println("FindAll")

	return book, nil
}

func (repo *fileRepository) FindByID(id int) (Book, error) {

	var book Book

	fmt.Println("FindByID")

	return book, nil
}

func (repo *fileRepository) Create(book *Book) (Book, error) {

	fmt.Println("Create")

	return *book, nil
}
