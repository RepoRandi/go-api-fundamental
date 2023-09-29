package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"web_api_fundamental/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetBooksHandler(c *gin.Context) {

	books, err := h.bookService.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := convertToBookResponse(b)

		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success FindAll Books",
	})

	c.JSON(http.StatusOK, gin.H{
		"books": booksResponse,
	})

}

func (h *bookHandler) GetBookHandler(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println(err)
	}

	b, err := h.bookService.FindByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	bookResponse := convertToBookResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"status": "Success FindByID Book",
	})

	c.JSON(http.StatusOK, gin.H{
		"book": bookResponse,
	})
}

func (h *bookHandler) QueryHandler(c *gin.Context) {
	title := c.Query("title")

	book, err := h.bookService.FindByTitle(title)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success FindByTitle Books",
		"title":  book.Title,
	})
}

func (h *bookHandler) CreateBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {

		errorMessages := []string{}

		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field: %s, Condition: %s", err.Field(), err.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})

		return
	}

	book, err := h.bookService.Create(&bookRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	bookResponse := convertToBookResponse(book)

	c.JSON(http.StatusOK, gin.H{
		"status": "Success Create Book",
	})

	c.JSON(http.StatusOK, gin.H{
		"book": bookResponse,
	})
}

func (h *bookHandler) UpdateBookHandler(c *gin.Context) {

	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {

		errorMessages := []string{}

		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field: %s, Condition: %s", err.Field(), err.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})

		return
	}

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println(err)
	}

	book, err := h.bookService.Update(id, &bookRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	bookResponse := convertToBookResponse(book)

	c.JSON(http.StatusOK, gin.H{
		"status": "Success Update Book",
	})

	c.JSON(http.StatusOK, gin.H{
		"book": bookResponse,
	})
}

func (h *bookHandler) DeleteBookHandler(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println(err)
	}

	book, err := h.bookService.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	bookResponse := convertToBookResponse(book)

	c.JSON(http.StatusOK, gin.H{
		"status": "Success Delete Book",
	})

	c.JSON(http.StatusOK, gin.H{
		"book": bookResponse,
	})
}

func convertToBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}
}
