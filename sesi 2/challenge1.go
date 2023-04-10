package main

import (
	"github.com/gin-gonic/gin"
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

func main() {
	router := gin.Default()
	router.GET("book", GetAllBook)
	router.GET("book/:id", GetBookById)
	router.POST("book", AddBook)
	router.PUT("book/:id", UpdateBook)
	router.DELETE("book/:id", DeleteBook)
	router.Run()
}

func GetAllBook(c *gin.Context) {
	c.JSON(http.StatusOK, book)
}

func AddBook(c *gin.Context) {
	var newBook ItemBook
	err := c.BindJSON(&newBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	newBook.ID = counter
	counter++
	book = append(book, newBook)
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
	book[id] = newBook
	c.JSON(http.StatusOK, newBook)
}

func GetBookById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	for _, v := range book {
		if v.ID == id {
			c.JSON(http.StatusOK, v)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
}

func DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	for i, v := range book {
		if v.ID == id {
			id = i
			break
		}
	}
	book = append(book[:id], book[id+1:]...)
	c.JSON(http.StatusOK, gin.H{"message": "successs"})
}

//package main
//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"strconv"
//)
//
//type Album struct {
//	ID     int     `json:"id"`
//	Title  string  `json:"title"`
//	Artist string  `json:"artist"`
//	Price  float64 `json:"price"`
//}
//
//var mapAlbum = make(map[int]Album, 0)
//var counter = 1
//
//func main() {
//	router := gin.Default()
//	router.GET("/albums", getAlbums)
//	router.GET("/albums/:id", GetAlbumsById)
//	router.POST("/albums", postAlbums)
//	router.DELETE("/albums/:id", deleteAlbums)
//	router.PUT("/albums/:id", putAlbums)
//	router.Run("localhost:8000")
//}
//
//func getAlbums(c *gin.Context) {
//	albums := make([]Album, 0)
//	for _, v := range mapAlbum {
//		albums = append(albums, v)
//	}
//	c.JSON(http.StatusOK, albums)
//}
//
//func postAlbums(c *gin.Context) {
//	var newAlbum Album
//	err := c.ShouldBindJSON(&newAlbum)
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
//		return
//	}
//	newAlbum.ID = counter
//	mapAlbum[counter] = newAlbum
//	counter++
//
//	c.JSON(http.StatusOK, newAlbum)
//}
//
//func deleteAlbums(c *gin.Context) {
//	idString := c.Param("id")
//
//	id, err := strconv.Atoi(idString)
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
//			"error": err,
//		})
//		return
//	}
//	v, found := mapAlbum[id]
//	//fmt.Println(mapAlbum[id])
//	if !found {
//		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
//			"error": "Not Found",
//		})
//		return
//	}
//	fmt.Println(v)
//	delete(mapAlbum, id)
//	c.JSON(http.StatusOK, v)
//}
//
//func putAlbums(c *gin.Context) {
//	idString := c.Param("id")
//
//	id, err := strconv.Atoi(idString)
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
//			"error": err,
//		})
//		return
//	}
//	_, found := mapAlbum[id]
//	if !found {
//		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
//			"error": "Not Found",
//		})
//		return
//	}
//	var updatedAlbum Album
//
//	err = c.ShouldBindJSON(&updatedAlbum)
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
//		return
//	}
//	mapAlbum[id] = updatedAlbum
//	c.JSON(http.StatusOK, updatedAlbum)
//}
//
//func GetAlbumsById(c *gin.Context) {
//	target := c.Param("id")
//
//	id, err := strconv.Atoi(target)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, err)
//	}
//	data, found := mapAlbum[id]
//	if !found {
//		c.JSON(http.StatusNotFound, gin.H{
//			"error": "data ga nemu",
//		})
//	}
//	c.JSON(http.StatusOK, data)
//}
