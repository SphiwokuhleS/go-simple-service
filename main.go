package main

import (
	"awesomeProject/handlers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main(){
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	productHandler := handlers.NewProducts(l)
	serveMux := http.NewServeMux()
	serveMux.Handle("/", productHandler)

	server := http.Server{
		Addr : ":9000",
		Handler: serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		fmt.Println("Server running...")
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	//handle system interruptions
	signChan := make(chan os.Signal)
	signal.Notify(signChan, os.Interrupt)
	signal.Notify(signChan, os.Kill)

	l.Println("Receiving terminate, graceful shutdown...", <- signChan)

	timeOutContext, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	server.Shutdown(timeOutContext)
}





