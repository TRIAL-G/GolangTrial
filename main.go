//connecting to server 
package main

import ("fmt"
		"log"
		"net/http"

)

func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path := "hello"{
		http,Error(w, "404 not ")
	}

}

func main{
	fileserver := http.Fileserver(http.dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err !null {
		log.Fatal(err)
	}
}
