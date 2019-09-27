package main

import (
	"fmt"
	"net/http"
)

type helloInterface interface {
	methodWithHelloMessage() string
}

type printHelloConsole string

type hello struct {
	args1, args2 string
}

func (p printHelloConsole) methodWithHelloMessage() string {
	return string("hello world " + p)
}

func (m hello) methodWithHelloMessage() string {
	return m.args1 +" "+m.args2
}

//Call this method from any package, with a value which satisfies helloInterface
func PrintHello(h helloInterface) string{
	return h.methodWithHelloMessage()
}

func main() {

	http.HandleFunc("/",index)
	http.ListenAndServe(":8080",nil)

}

func index (w http.ResponseWriter, r * http.Request) {

	helloVar := hello{args1: "Hello", args2: "World!!!!!!!"}
	p := printHelloConsole("Hello variable initialized")
	fmt.Println(p.methodWithHelloMessage())
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,PrintHello(helloVar))

}
