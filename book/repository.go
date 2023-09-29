package book

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	FindByTitle(title string) (Book, error)
	Create(book *Book) (Book, error)
	Update(book *Book) (Book, error)
	Delete(book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (repo *repository) FindAll() ([]Book, error) {
	var books []Book
	err := repo.db.Find(&books).Error
	return books, err
}

func (repo *repository) FindByID(id int) (Book, error) {
	var book Book
	err := repo.db.First(&book, id).Error
	return book, err
}

func (repo *repository) FindByTitle(title string) (Book, error) {
	var book Book
	err := repo.db.Where("title LIKE  ? ", title).Find(&book).Error
	return book, err
}

func (repo *repository) Create(book *Book) (Book, error) {
	err := repo.db.Create(book).Error
	return *book, err
}

func (repo *repository) Update(book *Book) (Book, error) {
	err := repo.db.Save(book).Error
	return *book, err
}

func (repo *repository) Delete(book Book) (Book, error) {
	err := repo.db.Delete(&book).Error
	return book, err
}
