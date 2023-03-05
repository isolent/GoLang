package lecture

import (
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request){
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request){
	fmt.Printf("got / hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main(){
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	fmt.Println("Server is listening")
	err := http.ListenAndServe(":8181", nil)
	if err != nil{
		return
	}
}