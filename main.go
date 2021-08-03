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
	helloHandler := handlers.NewHello(l)
	goodByeHandler := handlers.NewGoodBye(l)
	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/goodbye", goodByeHandler)

	fmt.Println("Server running...")
	server := http.Server{
		Addr : ":9000",
		Handler: serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	//handle system interruptions
	signChan := make(chan os.Signal)
	signal.Notify(signChan, os.Interrupt)
	signal.Notify(signChan, os.Kill)

	sig := <- signChan
	l.Println("Receiving terminate, graceful shutdown...", sig)

	timeOutContext, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	server.Shutdown(timeOutContext)
}





