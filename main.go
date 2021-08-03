package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		log.Println("Server running...")
		d, err:= ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Bad stuff", http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Printf("Data: %s\n", d)
		fmt.Fprintf(w, "Hello %s\n", d)
	})

	http.ListenAndServe(":9000", nil)
}





