package main

import (
	"log"
	handlers "microservice/Handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/net/context"
)

const message = "Hello world"

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)
	mux := http.NewServeMux()
	mux.Handle("/", hh)
	mux.Handle("/goodbye", gh)

	//we are using a custom server with tunes settings and for Read write timeout
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  time.Second * 1,
		WriteTimeout: time.Second * 1,
		IdleTimeout:  time.Second * 30,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	//we make a signal channel to notify whenever we interrupt or shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("\n Recieved terminate, graceful shutdown", sig)

	//shutdown the server gracefully
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
