package main

import ("github.com/gofiber/fiber/v2"
		"strconv")


// fiber is framework like express.js
/* concept -- > declare struct for collect data , req.body to struct*/

// Book struct to hold book data
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	app := fiber.New()

	books = append(books, Book{ID:1, Title: "Coder", Author:"New"})
	books = append(books, Book{ID:2, Title: "Sleep", Author:"New"})

	app.Get("/books",getBooks)
	app.Get("/books/:id",getBook)
	app.Post("/books",createBook)
	app.Put("/books/:id",updateBook)
	app.Delete("/books/:id",deleteBook)
	


	app.Listen(":8080")
}

func getBooks(c *fiber.Ctx) error{
		
		return c.JSON(books)
	}

func getBook(c *fiber.Ctx) error{
		bookId , err := c.ParamsInt("id",0)
		
		if err != nil{
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		for _, book := range books{
			if book.ID == bookId{
				return c.JSON(book)
			}
		}

		return c.SendStatus(fiber.StatusNotFound)
		
		
		
		
	}

func createBook(c *fiber.Ctx) error{
	book := new(Book)
	if err := c.BodyParser(book) ; err != nil{
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}   //แปลงข้อมูลจาก body req --> map to struct book 
	books = append(books, *book)
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error{
	id , err := strconv.Atoi(c.Params("id"))
	
	if err != nil{
		return c.SendStatus(fiber.StatusBadRequest)
	}

	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err != nil{
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	
	
	for i, book := range books{
		if book.ID == id {
			books[i].Title = bookUpdate.Title
			books[i].Author = bookUpdate.Author
			return c.JSON(books)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
	
}

func deleteBook(c *fiber.Ctx) error{
	id , err := strconv.Atoi(c.Params("id"))
	 if err != nil{
		return c.SendStatus(fiber.StatusBadRequest)
	 }

	 for i, book := range books{
		if book.ID == id {
			books = append(books[:i],books[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	 }
	 return c.SendStatus(fiber.StatusBadRequest)

}



/* code before use Fiber
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello World!")
}

func main() {

	
	http.HandleFunc("/hello",helloHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
*/

