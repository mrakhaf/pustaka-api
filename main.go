package main

import (
	"github.com/mrakhaf/pustaka-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", handler.RootHandler)
	router.GET("/books/:id", handler.BooksHandler)
	router.GET("/query", handler.QueryHandler)
	router.POST("/uploadbooks", handler.UploadBooksHandler)

	router.Run()
}
