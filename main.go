package main

import (
	//"errors"
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	
)

//Now lets create data strcuture for the book library


type book struct{
	ID       string  `json:"id"`
    Title    string   `json:"title"`
    Author   string    `json:"author"`
	Quauntity int      `json:"quantity"`
	DOC       doc
}

//Create a struct for the date of book creation

type doc struct{
	ID       string   `json:"id"`
    Day      int   `json:"day"`
    Month    int    `json:"month"`
	Year     int   `json:"year"`

}

/// now lets write a function to create books

func createBook(c *gin.Context){


	var newbook book
	if err :=c.BindJSON(&newbook); err !=nil{

		return
	}
	books = append(books,newbook)
	c.IndentedJSON(http.StatusCreated,newbook)

}

// creating a slice of book struct that will be used to store the book records

var books = []book{

	{ID: "1", Title: "Gold Truck  for the lord", Author: "Truck wills", Quauntity: 80},
	{ID: "2", Title: "clement Songs", Author: "Clement Doe",Quauntity: 30},
	{ID: "4", Title: "Dragon of prraise", Author: "John Wein", Quauntity: 50},
	{ID: "5", Title: "The golden angels", Author: "Shakes Wein", Quauntity: 10},
	{ID: "6", Title: "The songs of jubilee", Author: "John john", Quauntity: 20},
}
// now rite a fun ction that collects all the books

func getBooks(c *gin.Context){
c.IndentedJSON(http.StatusOK,books)

}

func addBook(c *gin.Context){

	var newbk book
	if err := c.BindJSON(&newbk); err!=nil{

		return
	}
	books = append(books,newbk)
	c.IndentedJSON(http.StatusCreated,newbk)
}

func main(){


	router := gin.Default() 
	router.GET("/books",getBooks)
	router.POST("/books",createBook)
	//router.POST("/books",addBook)
	router.Run("localhost:8080")
}