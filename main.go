package main

import (
	"log"
	handlers "microservice/Handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags) //defines a new logger object

	ph := handlers.NewProducts(l) //product handler
	smux := mux.NewRouter()

	getRouter := smux.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := smux.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := smux.Methods("POST").Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	deleteRouter := smux.Methods("DELETE").Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", ph.DeleteProducts)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	//we are using a custom server with tunes settings and for Read write timeout
	server := &http.Server{
		Addr:         ":8080",          //configure bind address
		Handler:      smux,             //default handler
		ReadTimeout:  time.Second * 1,  //max time to read request from client
		WriteTimeout: time.Second * 1,  //max time to write response to the client
		IdleTimeout:  time.Second * 30, //max time for connections using TCP keep-Alive
	}
	//we run our server is a seperate go routine so as it dosent block the main.go code
	go func() {

		l.Println("Server started on PORT: 8080")
		err := server.ListenAndServe()
		if err != nil {
			l.Printf("Error Starting server %s", err)
			os.Exit(1)

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
