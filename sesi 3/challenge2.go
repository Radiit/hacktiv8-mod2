package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
)

type ItemBook struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Genre  string `json:"genre"`
	Author string `json:"author"`
}

var book []ItemBook
var counter = 1
var db *sql.DB

func main() {
	var err error
	gin.SetMode(gin.ReleaseMode)
	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=pswd dbname=hacktiv8 sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	router.GET("book", GetAllBook)
	router.GET("book/:id", GetBookById)
	router.POST("book", AddBook)
	router.PUT("book/:id", UpdateBook)
	router.DELETE("book/:id", DeleteBook)
	router.Run()
}

func GetAllBook(c *gin.Context) {
	query := "select * from Book"
	rows, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mapBook := make([]ItemBook, 0)

	for rows.Next() {
		var books ItemBook
		err = rows.Scan(&books.Name, &books.Genre, &books.Author, &books.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		mapBook = append(mapBook, books)
	}

	c.JSON(http.StatusOK, mapBook)
}

func AddBook(c *gin.Context) {
	var newBook ItemBook
	err := c.BindJSON(&newBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	query := "insert into Book(Name, Genre, Author, ID) values($1, $2, $3, $4) returning *"
	row := db.QueryRow(query, newBook.Name, newBook.Genre, newBook.Author, newBook.ID)
	err = row.Scan(&newBook.Name, &newBook.Genre, &newBook.Author, &newBook.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, newBook)
}

func UpdateBook(c *gin.Context) {
	idString := c.Param("id")
	var newBook ItemBook

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = c.BindJSON(&newBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	query := "update book set Name=$1, Genre=$2, Author=$3 where id=$4 returning id"
	row := db.QueryRow(query, newBook.Name, newBook.Genre, newBook.Author, id)

	err = row.Scan(&newBook.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, newBook)
}

func GetBookById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var bookTarget ItemBook
	query := "select * from book where id=$1"

	row := db.QueryRow(query, id)
	err = row.Scan(&bookTarget.Name, &bookTarget.Genre, &bookTarget.Author, &bookTarget.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, bookTarget)
}

func DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var deletedBook ItemBook
	query := "delete from book where id=$1 returning *"

	row := db.QueryRow(query, id)

	err = row.Scan(&deletedBook.Name, &deletedBook.Genre, &deletedBook.Author, &deletedBook.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, deletedBook)
}
