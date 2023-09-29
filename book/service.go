package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	FindByTitle(title string) (Book, error)
	Create(bookRequest *BookRequest) (Book, error)
	Update(ID int, bookRequest *BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) FindAll() ([]Book, error) {
	return s.repo.FindAll()
}

func (s *service) FindByID(id int) (Book, error) {
	return s.repo.FindByID(id)
}

func (s *service) FindByTitle(title string) (Book, error) {
	return s.repo.FindByTitle(title)
}

func (s *service) Create(bookRequest *BookRequest) (Book, error) {

	var book Book

	book.Title = bookRequest.Title
	book.Description = bookRequest.Description
	book.Price = bookRequest.Price
	book.Rating = bookRequest.Rating
	book.Discount = bookRequest.Discount

	return s.repo.Create(&book)
}

func (s *service) Update(ID int, bookRequest *BookRequest) (Book, error) {

	book, err := s.repo.FindByID(ID)

	if err != nil {
		return book, err
	}

	book.Title = bookRequest.Title
	book.Description = bookRequest.Description
	book.Price = bookRequest.Price
	book.Rating = bookRequest.Rating
	book.Discount = bookRequest.Discount

	return s.repo.Update(&book)
}

func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repo.FindByID(ID)

	if err != nil {
		return book, err
	}

	return s.repo.Delete(book)
}
