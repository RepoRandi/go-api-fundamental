package main

import (
	"fmt"
	"log"
	"web_api_fundamental/book"
	"web_api_fundamental/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/web-api-fundamental?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db connection error!")
	}

	fmt.Println("Connected to database!")

	db.AutoMigrate(book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/api/v1")

	v1.GET("/", bookHandler.GetBooksHandler)

	v1.GET("/book/:id", bookHandler.GetBookHandler)

	v1.GET("/query", bookHandler.QueryHandler)

	v1.POST("/books", bookHandler.CreateBooksHandler)

	v1.PUT("/book/:id", bookHandler.UpdateBookHandler)

	v1.DELETE("/book/:id", bookHandler.DeleteBookHandler)

	router.Run(":8888")
}
