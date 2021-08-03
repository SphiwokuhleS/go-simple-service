package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request){
		h.l.Println("Hello There!")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Bad stuff", http.StatusBadRequest)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Printf("Hello %s\n", d)
	}