package main
import (
	"awesomeProject/calcBooks"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/calcBooks/books", calcBooks.BooksHandleFunc)
	http.HandleFunc("/calcBooks/books/", calcBooks.BookHandleFunc)
	http.HandleFunc("/calcBooks/hello", calcBooks.HelloHandleFunc)
	calcBooks.AlgebricMinus(9,1)
	http.ListenAndServe(":8080",nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) ==0 {
		port = "8080"
	}
	return ":" + port
}

func index (w http.ResponseWriter, r * http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"Hello from GO server")

}
