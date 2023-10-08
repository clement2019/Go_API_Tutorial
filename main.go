package main

import (
	//"errors"
	//"fmt"
	"errors"
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
// now write a fun ction that collects all the books

func getBooks(c *gin.Context){

c.IndentedJSON(http.StatusOK,books)

}

//Now lets return books

func returnbook(c *gin.Context){

	id,ok :=c.GetQuery("id")
	if !ok{

		c.IndentedJSON(http.StatusBadRequest,gin.H{"message": "Book ID not found"})
		return
	}
	book,err :=getbookByID(id)
	if err !=nil{

		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book Not Found!"})
		return
	}
	if book.Quauntity ==0{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message": "Book not available"})
    }
    book.Quauntity +=1
	c.IndentedJSON(http.StatusOK,book)
}



//Now lets look at checkoking out books

func checkout(c *gin.Context){
id,ok :=c.GetQuery("id")

	if !ok{

		c.IndentedJSON(http.StatusBadRequest,gin.H{"message": "Book ID not found"})
		return
	}
	book,err :=getbookByID(id)
	if err !=nil{

		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book Not Found!"})
		return
	}
	if book.Quauntity ==0{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message": "Book not available"})

	}
	book.Quauntity -=1
	c.IndentedJSON(http.StatusOK,book)

}



func bookById(c *gin.Context){

	id := c.Param("id")
	book,err := getbookByID(id)

		if err != nil{
			c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book Not Found!"})
			return
		}
		c.IndentedJSON(http.StatusOK,book)
	}


func getbookByID(id string) (*book,error){
//form lopp 
 for i,b :=range books{
  if b.ID==id{

	return &books[i],nil
  }

 }
 return nil,errors.New("error book ID not found")

}



func main(){


	router := gin.Default() 
	router.GET("/books",getBooks)
	router.POST("/books",createBook)
	router.GET("/books/:id",bookById)
	router.PATCH("/checkout",checkout)
	router.PATCH("/return",returnbook)
	
	router.Run("localhost:8080")
}