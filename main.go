package main

import "github.com/gofiber/fiber/v2"

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

	books = append(books, Book{ID: 1, Title: "Coder", Author: "New"})
	books = append(books, Book{ID: 2, Title: "Sleep", Author: "New"})

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Post("/upload",uploadFile)

	app.Listen(":8080")
}

func uploadFile(c *fiber.Ctx) error{
	// input data from image key ---> image:
	file , err := c.FormFile("image")

	if err != nil{
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// save under folder
	err = c.SaveFile(file,"./uploads/" + file.Filename)

	if err != nil{
		return 	c.Status(fiber.StatusInternalServerError).SendString(err.Error())

	}
	return c.SendString("File Uploaded !")

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
