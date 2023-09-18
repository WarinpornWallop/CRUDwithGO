package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//รูปแบบของข้อมูลหนังสือ
type book struct {
  ID string `json:"id"`
  Name string `json:"name"`
  Author string  `json:"author"`
  Price  float64 `json:"price"`
}

// รายการของหนังสือที่มีทั้งหมด
var books= []book{
    {
      ID: "1", 
      Name: "Harry Potter", 
      Author: "J.K. Rowling", 
      Price: 15.9,
    },
    {
      ID: "2", 
      Name: "One Piece", 
      Author: "Oda Eiichirō", 
      Price: 2.99,
     },
    { 
      ID: "3", 
      Name: "demon slayer", 
      Author: "koyoharu gotouge", 
      Price: 2.99,
    },
}
func main() {
    router := gin.Default()
	//สร้าง route GET/books สำหรับเรียกฟังก์ชั่น getBooksในฟังก์ชั่น main
    router.GET("/books", getBooks)
	//สร้าง route GET /book/:idสำหรับเรียกฟังก์ชั่น getBookByID ในฟังก์ชั่น main
	router.GET("/book/:id", getBookByID)
	//สร้าง route POST/books สำหรับเรียกฟังก์ชั่น addBookในฟังก์ชั่น main
	router.POST("/books", addBook)
	//สร้าง route PUT/book/:id สำหรับเรียกฟังก์ชั่นที่มาจัดการแก้ไขหนังสือ
	router.PUT("/book/:id", updateBook)
	//สร้าง route DELETE/book/:id สำหรับเรียกฟังก์ชั่นที่มาจัดการลบรายการหนังสือ
	router.DELETE("/book/:id", deleteBookByID)
    router.Run("localhost:8080")
}
//gin.Context เป็นส่วนสำคัญที่สุดของ gin มีรายละเอียดของ request, validates, จัดรูปแบบเป็น JSON เป็นต้น
// books =  ตัวถัดไปคือ slice ของรายการหนังสือทั้งหมดที่จะตอบกลับไป
func getBooks(c *gin.Context) {
    c.JSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) { 
  paramID := c.Param("id")
  for _, book := range books {
   //ทำการ loop หาหนังสือที่ตรงกับ id ที่ส่งมา
    if book.ID == paramID { 
      c.JSON(http.StatusOK, book)
      return
    }
  }
 c.JSON(http.StatusNotFound, "data not found")
}

func addBook(c *gin.Context) {
    var newBook book

    // เรียก BindJSON เพื่อผูก JSON ที่รับมากับ newBook
    if err := c.BindJSON(&newBook); err != nil {
        return
    }

    // เพิ่มรายการหนังสือเล่มใหม่เข้าไปใน slice
    books = append(books, newBook)
    c.JSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
  var editBook book  
  if err := c.BindJSON(&editBook); err != nil {
    return
  }
  
  paramID := c.Param("id")
  for i := 0; i <= len(books)-1; i++ {
    //loop หา id หนังสือตาม id และทำการแก้ไขรายการหนังสือนั่น
    if books[i].ID == paramID {
      books[i].Name = editBook.Name
      books[i].Author = editBook.Author
      books[i].Price = editBook.Price
      
      c.JSON(http.StatusOK, books[i])
      return
    }
 }
 //หาก id ที่ส่งมาไม่ตรงกับ id รายการหนังสือจะส่ง status code 404 พร้อม ข้อความ “data not found”
 c.JSON(http.StatusNotFound, "data not found")
}

func deleteBookByID(c *gin.Context) {
	paramID := c.Param("id")
	for i := 0; i <= len(books)-1; i++ {
		if books[i].ID == paramID {
			// ทำการลบรายการหนังสือจาก id ที่ส่งมา 
			//โดยตัดข้อมูลแล้วเอามาต่อโดยไม่มีรายการหนังสือที่เราส่งมา  
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, "delete success")
			return
		}
	}
	c.JSON(http.StatusNotFound, "data not found")
}