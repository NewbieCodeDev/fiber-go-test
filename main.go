package main

import ("github.com/gofiber/fiber/v2")


// fiber is framework like express.js

func main() {
	app := fiber.New()

	app.Get("/hello",func(c *fiber.Ctx) error{
		return c.SendString("hello World!")
	})

	app.Listen(":8080")
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

